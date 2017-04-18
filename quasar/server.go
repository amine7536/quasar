package quasar

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/amine7536/quasar/conf"
	"github.com/amine7536/quasar/utils"
	api "github.com/osrg/gobgp/api"
	gobgpConfig "github.com/osrg/gobgp/config"
	gobgp "github.com/osrg/gobgp/server"
)

type event struct {
	time     time.Time
	network  string
	nexthop  net.IP
	withdraw bool
	neighbor neighbor
}

type neighbor struct {
	address string
	asn     uint32
	name    []string
}

// Start the App
func Start(config *conf.Config, logger *logrus.Entry) {

	syssigChan := make(chan os.Signal, 1)
	signal.Notify(syssigChan, syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(syssigChan, os.Kill)

	s := gobgp.NewBgpServer()
	go s.Serve()

	// start grpc api server. this is not mandatory
	// but you will be able to use `gobgp` cmd with this.
	logger.Infof("Starting gRPC API=%t", config.API)
	if config.API {
		g := api.NewGrpcServer(s, ":50051")
		go g.Serve()
	}

	// global configuration
	global := &gobgpConfig.Global{
		Config: gobgpConfig.GlobalConfig{
			As:       config.Asn,
			RouterId: config.RouterID,
			Port:     -1, // gobgp won't listen on tcp:179
		},
	}

	if err := s.Start(global); err != nil {
		logger.Fatal(err)
	}

	for _, v := range config.Neighbors {
		// neighbor configuration
		n := &gobgpConfig.Neighbor{
			Config: gobgpConfig.NeighborConfig{
				NeighborAddress: v.Address,
				PeerAs:          v.Asn,
			},
		}

		if err := s.AddNeighbor(n); err != nil {
			logger.Fatal(err)
		}
	}

	// monitor new routes
	w := s.Watch(gobgp.WatchBestPath(false))

mainLoop:
	for {
		select {
		case <-syssigChan:
			logger.Info("Got SIGTERM, bye !")
			// Break mainLoop
			s.Shutdown()
			break mainLoop
		case ev := <-w.Event():
			switch msg := ev.(type) {
			case *gobgp.WatchEventBestPath:
				for _, path := range msg.PathList {

					neighborName, _ := utils.ResolveName(path.GetSource().Address)

					e := event{
						time: path.GetTimestamp(),
						neighbor: neighbor{
							address: path.GetSource().Address.String(),
							asn:     path.GetSource().AS,
							name:    neighborName,
						},
						withdraw: path.IsWithdraw,
						nexthop:  path.GetNexthop(),
						network:  path.GetNlri().String(),
					}

					fmt.Printf("%+v\n", e)
					fmt.Printf("%+v\n", path)

				}
			}
		}
	}

}
