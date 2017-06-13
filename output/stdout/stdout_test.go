package outputstdout

import (
	"net"
	"testing"
	"time"

	"github.com/amine7536/quasar/event"
)

func Test_main(t *testing.T) {

	// Init Output
	stdOut := Output{
		Mode: "json",
	}

	// Mock Event
	var bgpevent event.Event
	bgpevent.Time = time.Now()
	bgpevent.Neighbor = event.Neighbor{
		Address: "10.2.2.2",
		Asn:     65000,
		Name:    []string{"router1.test"},
	}
	bgpevent.Withdraw = true
	bgpevent.Nexthop = event.Nexthop{
		Net:  "172.16.0.254",
		Name: []string{"router2.test"},
	}

	_, networkNet, _ := net.ParseCIDR("192.168.1.36/32")
	bgpevent.Network = event.Network{
		Net:  networkNet.String(),
		Name: []string{"service.test"},
	}

	if err := stdOut.Send(bgpevent); err != nil {
		t.Fatalf("Methon Send on Output failed with : %s\n", err.Error())
	}

}
