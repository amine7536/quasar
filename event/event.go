package event

import (
	"time"

	"github.com/amine7536/quasar/utils"
	gobgpTable "github.com/osrg/gobgp/table"
)

// Event struct
type Event struct {
	Time     time.Time `json:"time"`
	Network  Nlri      `json:"network"`
	Nexthop  Nlri      `json:"nexthop"`
	Withdraw bool      `json:"withdraw"`
	Neighbor Neighbor  `json:"neighbor"`
}

// Nlri struct
type Nlri struct {
	Net  string   `json:"net"`
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
	// Try to resolve DNS Names
	neighborName, _ := utils.ResolveName(path.GetSource().Address.String())
	nlirName, _ := utils.ResolveNilrName(path.GetNlri().String())
	nexthopName, _ := utils.ResolveName(path.GetNexthop().String())

	// Update Event
	bgpevent.Time = path.GetTimestamp()
	bgpevent.Neighbor = Neighbor{
		Address: path.GetSource().Address.String(),
		Asn:     path.GetSource().AS,
		Name:    neighborName,
	}
	bgpevent.Withdraw = path.IsWithdraw
	bgpevent.Nexthop = Nlri{
		Net:  path.GetNexthop().String(),
		Name: nexthopName,
	}
	bgpevent.Network = Nlri{
		Net:  path.GetNlri().String(),
		Name: nlirName,
	}

	return nil
}
