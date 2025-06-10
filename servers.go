package steam

import (
	"math/rand"
	"time"

	"github.com/kolosok86/go-steam/netutil"
)

// CMServers contains a list of worldwide servers
//
// curl -s "https://api.steampowered.com/ISteamDirectory/GetCMListForConnect/v1/?key=$STEAM_API_KEY&format=json&celllid=1&cmtype=netfilter" | jq -r '.response.serverlist[] | .endpoint' | awk 'BEGIN {print "var CMServers = []string{"} {print "  \""$1"\","} END {print "}"}'
var CMServers = []string{
	"155.133.248.38:27017",
	"155.133.248.39:27017",
	"155.133.226.76:27017",
	"162.254.197.40:27017",
	"162.254.197.54:27017",
	"162.254.197.38:27017",
	"155.133.226.74:27017",
	"185.25.182.20:27017",
	"185.25.182.52:27017",
	"155.133.226.78:27017",
	"162.254.197.39:27017",
	"155.133.226.75:27017",
	"155.133.252.54:27017",
	"155.133.252.40:27017",
	"162.254.198.104:27017",
	"155.133.252.39:27017",
	"162.254.198.46:27017",
	"162.254.198.44:27017",
	"146.66.155.54:27017",
	"162.254.196.68:27017",
	"162.254.196.84:27017",
	"146.66.155.38:27017",
	"162.254.196.83:27017",
	"162.254.196.67:27017",
}

// GetRandomCM returns a random server from a built-in IP list.
//
// Prefer Client.Connect(), which uses IPs from the Steam Directory,
// which is always more up-to-date.
func GetRandomCM() *netutil.PortAddr {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	addr := netutil.ParsePortAddr(CMServers[rng.Int31n(int32(len(CMServers)))])
	if addr == nil {
		panic("invalid address in CMServers slice")
	}
	return addr
}
