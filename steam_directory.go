package steam

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/kolosok86/go-steam/netutil"
)

const (
	steamDirectoryURL = "https://api.steampowered.com/ISteamDirectory/GetCMList/v1/?cellId=0"
	httpTimeout       = 15 * time.Second
)

var steamDirectoryCache = &steamDirectory{}

type steamDirectory struct {
	mu      sync.RWMutex
	servers []string
}

type cmListResult struct {
	ServerList []string `json:"serverlist"`
	Result     uint32   `json:"result"`
	Message    string   `json:"message"`
}

type cmResponse struct {
	Response cmListResult `json:"response"`
}

// Initialize fetches the CM server list from the Steam Directory.
// The HTTP request is performed without holding the lock to avoid
// blocking concurrent reads for the duration of the request.
func (sd *steamDirectory) Initialize() error {
	servers, err := fetchCMServers()
	if err != nil {
		return err
	}

	sd.mu.Lock()
	sd.servers = servers
	sd.mu.Unlock()
	return nil
}

func fetchCMServers() ([]string, error) {
	client := &http.Client{Timeout: httpTimeout}
	resp, err := client.Get(steamDirectoryURL)
	if err != nil {
		return nil, fmt.Errorf("steam directory: %w", err)
	}
	defer resp.Body.Close()

	var r cmResponse
	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("steam directory: decode response: %w", err)
	}
	if r.Response.Result != 1 {
		return nil, fmt.Errorf("steam directory: result %d: %s", r.Response.Result, r.Response.Message)
	}
	if len(r.Response.ServerList) == 0 {
		return nil, fmt.Errorf("steam directory: empty server list")
	}
	return r.Response.ServerList, nil
}

func (sd *steamDirectory) GetRandomCM() *netutil.PortAddr {
	sd.mu.RLock()
	servers := sd.servers
	sd.mu.RUnlock()

	if len(servers) == 0 {
		return nil
	}
	return netutil.ParsePortAddr(servers[rand.Intn(len(servers))])
}

func (sd *steamDirectory) IsInitialized() bool {
	sd.mu.RLock()
	defer sd.mu.RUnlock()
	return len(sd.servers) > 0
}
