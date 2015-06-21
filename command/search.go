package command

import (
	"fmt"

	"github.com/briandowns/wasteband/config"
	"github.com/briandowns/wasteband/utils"
	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// Search holds in the passed in configuration
type Search struct {
	config *config.Configuration
}

// NewSearch creates a new instance of Delete
func NewSearch(conf *config.Configuration) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Search{
			config: conf,
		}, nil
	}
}

// Run searches for a given resource.
func (s *Search) Run(args []string) int {
	if flags.Has("help", args) || len(args) < 1 {
		fmt.Print(s.Help())
		return 1
	}

	// make sure that the command is a valid one
	if !utils.InSlice(args[0], allowedResources["search"]) {
		fmt.Print("ERROR: invalid option for show\n\n")
		fmt.Print(s.Help())
		return 1
	}

	// process the subcommand and it's options
	switch args[1] {
	case "index":
		_, err := utils.ESConn(d.config.Endpoint).Search()
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			return 1
		}
	}

	return 1
}

// Help provides full help inforamation for the subcommand
func (s *Search) Help() string { return "" }

// Synopsis provides a brief description of the command
func (s *Search) Synopsis() string {
	return "Search an Elasticsearch resource"
}
