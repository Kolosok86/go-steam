package steam

import (
	"github.com/kolosok86/go-steam/protocol/steamlang"
	"github.com/kolosok86/go-steam/steamid"
)

type LoggedOnEvent struct {
	Result                    steamlang.EResult
	ExtendedResult            steamlang.EResult
	OutOfGameSecsPerHeartbeat int32
	InGameSecsPerHeartbeat    int32
	PublicIp                  uint32
	ServerTime                uint32
	AccountFlags              steamlang.EAccountFlags
	ClientSteamId             steamid.SteamId `json:",string"`
	EmailDomain               string
	CellId                    uint32
	CellIdPingThreshold       uint32
	Steam2Ticket              []byte
	UsePics                   bool
	WebApiUserNonce           string
	IpCountryCode             string
	VanityUrl                 string
	NumLoginFailuresToMigrate int32
	NumDisconnectsToMigrate   int32
}

type LogOnFailedEvent struct {
	Result steamlang.EResult
}

type LoggedOffEvent struct {
	Result steamlang.EResult
}

type AccountInfoEvent struct {
	PersonaName          string
	Country              string
	CountAuthedComputers int32
	AccountFlags         steamlang.EAccountFlags
	FacebookId           uint64 `json:",string"`
	FacebookName         string
}
