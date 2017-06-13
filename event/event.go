package event

import (
	"fmt"
	"net"
	"time"

	"github.com/amine7536/quasar/utils"
	gobgpTable "github.com/osrg/gobgp/table"
)

// Event struct
type Event struct {
	Time        time.Time `json:"time"`
	Network     Network   `json:"network"`
	Nexthop     Nexthop   `json:"nexthop"`
	Withdraw    bool      `json:"withdraw"`
	Neighbor    Neighbor  `json:"neighbor"`
	Communities []string  `json:"communities"`
}

// Network struct
type Network struct {
	Net  string   `json:"net"`
	Name []string `json:"name"`
}

// Nexthop struct
type Nexthop struct {
	Net  string   `json:"address"`
	Name []string `json:"name"`
}

// Neighbor struct
type Neighbor struct {
	Address string   `json:"address"`
	Asn     uint32   `json:"asn"`
	Name    []string `json:"name"`
}

// Parse Event
func Parse(bgpevent *Event, path *gobgpTable.Path) error {
	// Update Event
	bgpevent.Time = path.GetTimestamp()

	// Neighbor info
	neighborNet := path.GetSource().Address.String()
	neighborName, _ := utils.ResolveName(neighborNet)

	bgpevent.Neighbor = Neighbor{
		Address: neighborNet,
		Asn:     path.GetSource().AS,
		Name:    neighborName,
	}

	// Is withdraw
	bgpevent.Withdraw = path.IsWithdraw

	// Nexthop info
	nexthopNet := path.GetNexthop().String()
	if nexthopNet == "<nil>" {
		nexthopNet = ""
	}
	nexthopName, _ := utils.ResolveName(nexthopNet)

	bgpevent.Nexthop = Nexthop{
		Net:  nexthopNet,
		Name: nexthopName,
	}

	// Network info
	_, networkNet, _ := net.ParseCIDR(path.GetNlri().String())
	networkName, _ := utils.ResolveNilrName(networkNet.String())

	bgpevent.Network = Network{
		Net:  networkNet.String(),
		Name: networkName,
	}

	// Communities
	for _, comm := range path.GetCommunities() {
		bgpevent.Communities = append(bgpevent.Communities, communityToString(comm))
	}

	return nil
}

func communityToString(comm uint32) string {
	return fmt.Sprintf("%d:%d", comm>>16, comm&0x0000ffff)
}
