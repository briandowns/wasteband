package main

import (
	"github.com/briandowns/wasteband/commands"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available Consul commands.
var Commands map[string]cli.CommandFactory

func init() {
	Commands = map[string]cli.CommandFactory{
		"show":    commands.NewShow(),
		"create":  commands.NewCreate(),
		"delete":  commands.NewDelete(),
		"set":     commands.NewSet(),
		"version": commands.NewVersion(WastebandVersion),
	}
}
