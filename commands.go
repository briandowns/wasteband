package main

import (
	"github.com/briandowns/wasteband/commands"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available Consul commands.
var Commands map[string]cli.CommandFactory

func init() {
	Commands = map[string]cli.CommandFactory{
		"show":    commands.NewShow(conf),
		"create":  commands.NewCreate(conf),
		"delete":  commands.NewDelete(conf),
		"set":     commands.NewSet(conf),
		"version": commands.NewVersion(WastebandVersion),
	}
}
