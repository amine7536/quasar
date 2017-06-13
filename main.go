package main

import (
	"log"

	"github.com/amine7536/quasar/cmd"
)

const (
	// Version : app version
	Version = "0.3.7"
	// ProgramName : app name
	ProgramName = "Quasar"
)

func main() {

	if err := cmd.NewRootCmd(Version, ProgramName).Execute(); err != nil {
		log.Fatal(err)
	}
}
