package event

import (
	"net"
	"time"
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
