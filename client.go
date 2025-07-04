package steam

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"net"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kolosok86/go-steam/cryptoutil"
	"github.com/kolosok86/go-steam/netutil"
	"github.com/kolosok86/go-steam/protocol"
	"github.com/kolosok86/go-steam/protocol/protobuf"
	"github.com/kolosok86/go-steam/protocol/steamlang"
	"github.com/kolosok86/go-steam/steamid"
	"golang.org/x/net/proxy"
	"google.golang.org/protobuf/proto"
)

// Represents a client to the Steam network.
// Always poll events from the channel returned by Events() or receiving messages will stop.
// All access, unless otherwise noted, should be threadsafe.
//
// When a FatalErrorEvent is emitted, the connection is automatically closed. The same client can be used to reconnect.
// Other errors don't have any effect.
type Client struct {
	// these need to be 64 bit aligned for sync/atomic on 32bit
	sessionId    int32
	_            uint32
	steamId      uint64
	currentJobId uint64

	Auth           *Auth
	Social         *Social
	Web            *Web
	Notifications  *Notifications
	GC             *GameCoordinator
	LocalIpAddress *net.TCPAddr
	Proxy          proxy.Dialer

	events        chan interface{}
	handlers      []PacketHandler
	handlersMutex sync.RWMutex

	jobHandlersMutex sync.RWMutex
	jobHandlers      map[protocol.JobId]chan *protocol.Packet
	tempSessionKey   []byte

	ConnectionTimeout time.Duration

	mutex          sync.RWMutex // guarding conn and writeChan
	conn           connection
	writeChan      chan protocol.IMsg
	writeBuf       *bytes.Buffer
	heartbeat      chan struct{}
	heartbeatMutex sync.Mutex // guarding heartbeat channel
}

type PacketHandler interface {
	HandlePacket(*protocol.Packet)
}

func NewClient() *Client {
	client := &Client{
		events:      make(chan interface{}, 3),
		writeBuf:    new(bytes.Buffer),
		jobHandlers: make(map[protocol.JobId]chan *protocol.Packet),
	}

	client.Auth = &Auth{client: client}
	client.RegisterPacketHandler(client.Auth)

	client.Social = newSocial(client)
	client.RegisterPacketHandler(client.Social)

	client.Web = &Web{client: client}
	client.RegisterPacketHandler(client.Web)

	client.Notifications = newNotifications(client)
	client.RegisterPacketHandler(client.Notifications)

	client.GC = newGC(client)
	client.RegisterPacketHandler(client.GC)

	return client
}

// Get the event channel. By convention all events are pointers, except for errors.
// It is never closed.
func (c *Client) Events() <-chan interface{} {
	return c.events
}

// SetIpAddress sets the local IP address to use for the connection.
func (c *Client) SetIpAddress(ipAddress string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	addr, err := net.ResolveIPAddr("ip", ipAddress)
	if err != nil {
		return fmt.Errorf("invalid IP address format: %s", ipAddress)
	}

	c.LocalIpAddress = &net.TCPAddr{
		IP: addr.IP,
	}

	return nil
}

// SetProxy sets the proxy to use for the connection.
func (c *Client) SetProxy(address string) error {
	url, err := url.Parse(address)
	if err != nil {
		return fmt.Errorf("invalid proxy format: %s", address)
	}

	proxy, err := proxy.FromURL(url, nil)
	if err != nil {
		return fmt.Errorf("invalid proxy format: %s", address)
	}

	c.Proxy = proxy
	return nil
}

func (c *Client) Emit(event interface{}) {
	c.events <- event
}

// Emits a FatalErrorEvent formatted with fmt.Errorf and disconnects.
func (c *Client) Fatalf(format string, a ...interface{}) {
	c.Emit(FatalErrorEvent(fmt.Errorf(format, a...)))
	c.Disconnect()
}

// Emits an error formatted with fmt.Errorf.
func (c *Client) Errorf(format string, a ...interface{}) {
	c.Emit(fmt.Errorf(format, a...))
}

// Registers a PacketHandler that receives all incoming packets.
func (c *Client) RegisterPacketHandler(handler PacketHandler) {
	c.handlersMutex.Lock()
	defer c.handlersMutex.Unlock()
	c.handlers = append(c.handlers, handler)
}

func (c *Client) GetNextJobId() protocol.JobId {
	return protocol.JobId(atomic.AddUint64(&c.currentJobId, 1))
}

func (c *Client) SteamId() steamid.SteamId {
	return steamid.SteamId(atomic.LoadUint64(&c.steamId))
}

func (c *Client) SessionId() int32 {
	return atomic.LoadInt32(&c.sessionId)
}

func (c *Client) Connected() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.conn != nil
}

// Connects to a random Steam server and returns its address.
// If this client is already connected, it is disconnected first.
// This method tries to use an address from the Steam Directory and falls
// back to the built-in server list if the Steam Directory can't be reached.
// If you want to connect to a specific server, use `ConnectTo`.
func (c *Client) Connect() (*netutil.PortAddr, error) {
	var server *netutil.PortAddr

	// try to initialize the directory cache
	if !steamDirectoryCache.IsInitialized() {
		_ = steamDirectoryCache.Initialize()
	}
	if steamDirectoryCache.IsInitialized() {
		server = steamDirectoryCache.GetRandomCM()
	} else {
		server = GetRandomCM()
	}

	err := c.ConnectTo(server)
	return server, err
}

// Connects to a specific server.
// You may want to use one of the `GetRandom*CM()` functions in this package.
// If this client is already connected, it is disconnected first.
func (c *Client) ConnectTo(addr *netutil.PortAddr) error {
	return c.ConnectToBind(addr, c.LocalIpAddress)
}

// Connects to a specific server, and binds to a specified local IP
// If this client is already connected, it is disconnected first.
func (c *Client) ConnectToBind(addr *netutil.PortAddr, local *net.TCPAddr) error {
	c.Disconnect()

	conn, err := dialTCP(local, addr.ToTCPAddr(), c.Proxy)
	if err != nil {
		c.Fatalf("Connect failed: %v", err)
		return err
	}
	c.conn = conn
	c.writeChan = make(chan protocol.IMsg, 5)

	go c.readLoop()
	go c.writeLoop()

	return nil
}

func (c *Client) Disconnect() {
	// stop heartbeat before locking the main mutex
	c.heartbeatMutex.Lock()
	if c.heartbeat != nil {
		close(c.heartbeat)
		c.heartbeat = nil
	}
	c.heartbeatMutex.Unlock()

	// clear all pending job handlers on disconnect
	c.jobHandlersMutex.Lock()
	for jobID, ch := range c.jobHandlers {
		close(ch)
		delete(c.jobHandlers, jobID)
	}
	c.jobHandlersMutex.Unlock()

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.conn == nil {
		return
	}

	c.conn.Close()
	c.conn = nil

	close(c.writeChan)

	c.Emit(&DisconnectedEvent{})
}

// Adds a message to the send queue. Modifications to the given message after
// writing are not allowed (possible race conditions).
//
// Writes to this client when not connected are ignored.
func (c *Client) Write(msg protocol.IMsg) {
	if cm, ok := msg.(protocol.IClientMsg); ok {
		cm.SetSessionId(c.SessionId())
		cm.SetSteamId(c.SteamId())
	}
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if c.conn == nil {
		return
	}
	c.writeChan <- msg
}

// WriteUnified sends a unified message and returns a packet with a 10 second timeout
func (c *Client) WriteUnified(service string, body proto.Message) (*protocol.Packet, error) {
	msg := protocol.NewClientMsgProtobuf(steamlang.EMsg_ServiceMethodCallFromClient, body)

	msg.Header.Proto.TargetJobName = &service
	jobID := c.GetNextJobId()
	msg.SetSourceJobId(jobID)

	ch := make(chan *protocol.Packet, 1)

	c.jobHandlersMutex.Lock()
	c.jobHandlers[protocol.JobId(jobID)] = ch
	c.jobHandlersMutex.Unlock()

	c.Write(msg)

	select {
	case packet := <-ch:
		// check if handler still exists before deleting
		c.jobHandlersMutex.Lock()
		delete(c.jobHandlers, protocol.JobId(jobID))
		c.jobHandlersMutex.Unlock()
		return packet, nil
	case <-time.After(time.Second * 10):
		c.jobHandlersMutex.Lock()
		delete(c.jobHandlers, protocol.JobId(jobID))
		c.jobHandlersMutex.Unlock()
		return nil, fmt.Errorf("timeout: no response received within 10 seconds")
	}
}

func (c *Client) readLoop() {
	for {
		// This *should* be atomic on most platforms, but the Go spec doesn't guarantee it
		c.mutex.RLock()
		conn := c.conn
		c.mutex.RUnlock()
		if conn == nil {
			return
		}
		packet, err := conn.Read()

		if err != nil {
			c.Fatalf("Error reading from the connection: %v", err)
			return
		}
		c.handlePacket(packet)
	}
}

func (c *Client) writeLoop() {
	for {
		c.mutex.RLock()
		conn := c.conn
		c.mutex.RUnlock()
		if conn == nil {
			return
		}

		msg, ok := <-c.writeChan
		if !ok {
			return
		}

		err := msg.Serialize(c.writeBuf)
		if err != nil {
			c.writeBuf.Reset()
			c.Fatalf("Error serializing message %v: %v", msg, err)
			return
		}

		err = conn.Write(c.writeBuf.Bytes())

		c.writeBuf.Reset()

		if err != nil {
			c.Fatalf("Error writing message %v: %v", msg, err)
			return
		}
	}
}

func (c *Client) heartbeatLoop(seconds time.Duration) {
	c.heartbeatMutex.Lock()
	// stop previous heartbeat if it exists
	if c.heartbeat != nil {
		close(c.heartbeat)
	}

	c.heartbeat = make(chan struct{})
	heartbeatChan := c.heartbeat
	c.heartbeatMutex.Unlock()

	ticker := time.NewTicker(seconds * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientHeartBeat, new(protobuf.CMsgClientHeartBeat)))
		case <-heartbeatChan:
			return
		}
	}
}

func (c *Client) handlePacket(packet *protocol.Packet) {
	switch packet.EMsg {
	case steamlang.EMsg_ChannelEncryptRequest:
		c.handleChannelEncryptRequest(packet)
	case steamlang.EMsg_ChannelEncryptResult:
		c.handleChannelEncryptResult(packet)
	case steamlang.EMsg_Multi:
		c.handleMulti(packet)
	case steamlang.EMsg_ServiceMethodResponse:
		c.handleServiceMethodResponse(packet)
	}

	c.handlersMutex.RLock()
	defer c.handlersMutex.RUnlock()
	for _, handler := range c.handlers {
		handler.HandlePacket(packet)
	}
}

func (c *Client) handleServiceMethodResponse(packet *protocol.Packet) {
	c.jobHandlersMutex.Lock()
	defer c.jobHandlersMutex.Unlock()

	if ch, ok := c.jobHandlers[packet.TargetJobId]; ok {
		// Use non-blocking send to avoid deadlock
		select {
		case ch <- packet:
			// Successfully sent packet
		default:
			// Channel is blocked or closed, just remove handler
		}

		delete(c.jobHandlers, packet.TargetJobId)
	}
}

func (c *Client) handleChannelEncryptRequest(packet *protocol.Packet) {
	body := steamlang.NewMsgChannelEncryptRequest()
	packet.ReadMsg(body)

	if body.Universe != steamlang.EUniverse_Public {
		c.Fatalf("Invalid universe %v!", body.Universe)
	}

	c.tempSessionKey = make([]byte, 32)
	rand.Read(c.tempSessionKey)
	encryptedKey := cryptoutil.RSAEncrypt(GetPublicKey(steamlang.EUniverse_Public), c.tempSessionKey)

	payload := new(bytes.Buffer)
	payload.Write(encryptedKey)
	binary.Write(payload, binary.LittleEndian, crc32.ChecksumIEEE(encryptedKey))
	payload.WriteByte(0)
	payload.WriteByte(0)
	payload.WriteByte(0)
	payload.WriteByte(0)

	c.Write(protocol.NewMsg(steamlang.NewMsgChannelEncryptResponse(), payload.Bytes()))
}

func (c *Client) handleChannelEncryptResult(packet *protocol.Packet) {
	body := steamlang.NewMsgChannelEncryptResult()
	packet.ReadMsg(body)

	if body.Result != steamlang.EResult_OK {
		c.Fatalf("Encryption failed: %v", body.Result)
		return
	}
	c.conn.SetEncryptionKey(c.tempSessionKey)
	c.tempSessionKey = nil

	c.Emit(&ConnectedEvent{})
}

func (c *Client) handleMulti(packet *protocol.Packet) {
	body := new(protobuf.CMsgMulti)
	packet.ReadProtoMsg(body)

	payload := body.GetMessageBody()

	if body.GetSizeUnzipped() > 0 {
		r, err := gzip.NewReader(bytes.NewReader(payload))
		if err != nil {
			c.Errorf("handleMulti: Error while decompressing: %v", err)
			return
		}

		payload, err = io.ReadAll(r)
		if err != nil {
			c.Errorf("handleMulti: Error while decompressing: %v", err)
			return
		}
	}

	pr := bytes.NewReader(payload)
	for pr.Len() > 0 {
		var length uint32
		binary.Read(pr, binary.LittleEndian, &length)
		packetData := make([]byte, length)
		pr.Read(packetData)
		p, err := protocol.NewPacket(packetData)
		if err != nil {
			c.Errorf("Error reading packet in Multi msg %v: %v", packet, err)
			continue
		}
		c.handlePacket(p)
	}
}
