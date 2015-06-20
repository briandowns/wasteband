package command

import (
	"fmt"

	"github.com/briandowns/wasteband/config"
	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

type Show struct {
	config *config.Configuration
}

func NewShow(conf *config.Configuration) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Show{
			config: conf,
		}, nil
	}
}

func (s *Show) Run(args []string) int {
	if flags.Has("help", args) {
		fmt.Print(s.Help())
		return 1
	}
	return 1
}

// Help provides full help inforamation for the subcommand
func (s *Show) Help() string { return "" }

// Synopsis provides a brief description of the command
func (s *Show) Synopsis() string {
	return "Show an Elasticsearch resource"
}
