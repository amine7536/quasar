package quasar

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/amine7536/quasar/conf"
	"github.com/amine7536/quasar/event"
	api "github.com/osrg/gobgp/api"
	gobgpConfig "github.com/osrg/gobgp/config"
	gobgp "github.com/osrg/gobgp/server"
)

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
			EbgpMultihop: gobgpConfig.EbgpMultihop{
				Config: gobgpConfig.EbgpMultihopConfig{
					Enabled:     true,
					MultihopTtl: 80,
				},
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

					// Parsers
					bgpevent := event.Event{}
					err := event.Parse(&bgpevent, path)
					if err != nil {
						logger.Info(err)
					}

					// Outputs
					for name, out := range conf.MapOutputs {
						logger.Debugf("output=%s", name)
						go func(o conf.OutputHandler, e event.Event) {
							logger.Debugf("event=%+v", e)
							if err := o.Send(e); err != nil {
								logger.Errorf("output failed: %v\n", err)
							}
						}(out, bgpevent)

					}

				}
			}
		}
	}

}
