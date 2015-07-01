package main

import (
	"fmt"
	"log"
	"os"

	"github.com/briandowns/wasteband/command"
	"github.com/briandowns/wasteband/config"
	"github.com/mitchellh/cli"
)

// wastebandVersion holds the current version of wasteband
const (
	wastebandVersion = "0.1"
	wastebandName    = "wasteband"
)

var conf *config.Configuration
var configFile = "config.json"

// Commands is the mapping of all the available wasteband commands.
var Commands map[string]cli.CommandFactory

func main() {
	if retval, err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(retval)
	}
}

func run() (int, error) {
	conf, err := config.Load(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	c := &cli.CLI{
		Name:     wastebandName,
		Version:  wastebandVersion,
		Args:     os.Args[1:],
		HelpFunc: cli.BasicHelpFunc(wastebandName),
		Commands: map[string]cli.CommandFactory{
			"show":     command.NewShow(conf),
			"create":   command.NewCreate(conf),
			"delete":   command.NewDelete(conf),
			"optimize": command.NewOptimize(conf),
			"search":   command.NewDelete(conf),
			"set":      command.NewSet(conf),
			"version":  command.NewVersion(wastebandVersion),
		},
	}

	retval, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1, err
	}

	return retval, nil
}
