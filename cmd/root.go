package cmd

import (
	"log"

	"github.com/amine7536/quasar/conf"
	"github.com/amine7536/quasar/quasar"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "quasar",
	Short: "BGP Black Hole",
	Run:   run,
}

var version string
var progName string

// NewRootCmd will setup and return the root command
func NewRootCmd(v string, p string) *cobra.Command {
	// Set Version and ProgramName
	version = v
	progName = p

	rootCmd.PersistentFlags().StringP("config", "c", "", "Config file to use")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	config, err := conf.LoadConfig(cmd)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}

	logger, err := conf.ConfigureLogging(&config.Logs)
	if err != nil {
		log.Fatal("Failed to configure logging: " + err.Error())
	}

	logger.Infof("Starting with config: %+v", config)

	// Start the Application
	quasar.Start(config, logger)

}
