package command

import (
	"fmt"

	"github.com/briandowns/wasteband/config"
	"github.com/briandowns/wasteband/utils"
	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// Show holds in the passed in configuration
type Show struct {
	config *config.Configuration
}

// NewShow creates a new instance of Delete
func NewShow(conf *config.Configuration) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Show{
			config: conf,
		}, nil
	}
}

// Run shows a given resource.
func (s *Show) Run(args []string) int {
	if flags.Has("help", args) || len(args) < 1 {
		fmt.Print(s.Help())
		return 1
	}

	// make sure that the command is a valid one
	if !utils.InSlice(args[0], allowedResources["show"]) {
		fmt.Print("ERROR: invalid option for show\n\n")
		fmt.Print(s.Help())
		return 1
	}

	// process the subcommand and it's options
	switch args[1] {
	case "index":
		_, err := utils.ESConn(d.config.Endpoint).DeleteIndex(args[2])
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			return 1
		}
	}

	return 1
}

// Help provides full help inforamation for the subcommand
func (s *Show) Help() string { return "" }

// Synopsis provides a brief description of the command
func (s *Show) Synopsis() string {
	return "Show an Elasticsearch resource"
}
