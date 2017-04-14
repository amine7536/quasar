package cmd

import (
	"log"

	"github.com/amine7536/quasar/config"
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
	conf, err := config.LoadConfig(cmd)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}

	logger, err := config.ConfigureLogging(&conf.LogConfig)
	if err != nil {
		log.Fatal("Failed to configure logging: " + err.Error())
	}

	logger.Infof("Starting with config: %+v", conf)

	// Start the Application
	//app.Start(conf, logger)

}
