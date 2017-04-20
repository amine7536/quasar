package main

import (
	"log"

	"github.com/amine7536/quasar/cmd"
	stackimpact "github.com/stackimpact/stackimpact-go"
)

const (
	// Version : app version
	Version = "0.1.0"
	// ProgramName : app name
	ProgramName = "Quasar"
)

func main() {

	agent := stackimpact.NewAgent()
	agent.Start(stackimpact.Options{
		AgentKey: "66cf1969d58a40cdac6da1455a6dbffe8611efbb",
		AppName:  ProgramName,
	})

	if err := cmd.NewRootCmd(Version, ProgramName).Execute(); err != nil {
		log.Fatal(err)
	}
}
