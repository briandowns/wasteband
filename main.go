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
const wastebandVersion = "0.1"

var conf *config.Configuration
var configFile = "config.json"

// Commands is the mapping of all the available wasteband commands.
var Commands map[string]cli.CommandFactory

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	conf, err := config.GetConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	c := &cli.CLI{
		Name:     "wasteband",
		Version:  wastebandVersion,
		Args:     os.Args[1:],
		HelpFunc: cli.BasicHelpFunc("wasteband"),
		Commands: map[string]cli.CommandFactory{
			"show":    command.NewShow(conf),
			"create":  command.NewCreate(conf),
			"delete":  command.NewDelete(conf),
			"set":     command.NewSet(conf),
			"version": command.NewVersion(wastebandVersion),
		},
	}

	_, err = c.Run() // ignoring the return value for now...
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return err
	}

	return nil
}
