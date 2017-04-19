package event

import (
	"net"
	"time"

	"github.com/amine7536/quasar/utils"
	gobgpTable "github.com/osrg/gobgp/table"
)

type Event struct {
	Time     time.Time
	Network  Nilr
	Nexthop  net.IP
	Withdraw bool
	Neighbor Neighbor
}

type Nilr struct {
	Net  string
	Name []string
}

type Neighbor struct {
	Address string
	Asn     uint32
	Name    []string
}

// Parse Event
func Parse(bgpevent *Event, path *gobgpTable.Path) error {
	// Try to resolve DNS Names
	neighborName, _ := utils.ResolveName(path.GetSource().Address.String())
	nlirName, _ := utils.ResolveNilrName(path.GetNlri().String())

	// Update Event
	bgpevent.Time = path.GetTimestamp()
	bgpevent.Neighbor = Neighbor{
		Address: path.GetSource().Address.String(),
		Asn:     path.GetSource().AS,
		Name:    neighborName,
	}

	bgpevent.Withdraw = path.IsWithdraw
	bgpevent.Nexthop = path.GetNexthop()
	bgpevent.Network = Nilr{
		Net:  path.GetNlri().String(),
		Name: nlirName,
	}

	return nil
}
