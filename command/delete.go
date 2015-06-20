package command

import (
	"fmt"

	"github.com/briandowns/wasteband/config"
	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

type Delete struct {
	config *config.Configuration
}

func NewDelete(conf *config.Configuration) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Delete{
			config: conf,
		}, nil
	}
}

func (d *Delete) Run(args []string) int {
	if flags.Has("help", args) {
		fmt.Print(d.Help())
		return 1
	}
	return 1
}

// Help provides full help inforamation for the subcommand
func (d *Delete) Help() string { return "" }

// Synopsis provides a brief description of the command
func (d *Delete) Synopsis() string {
	return "Delete an Elasticsearch resource"
}
