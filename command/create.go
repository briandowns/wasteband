package command

import (
	"fmt"

	"github.com/briandowns/wasteband/config"
	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

type Create struct {
	config *config.Configuration
}

func NewCreate(conf *config.Configuration) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Create{
			config: conf,
		}, nil
	}
}

func (c *Create) Run(args []string) int {
	if flags.Has("help", args) {
		fmt.Print(c.Help())
		return 1
	}
	return 1
}

// Help provides full help inforamation for the subcommand
func (c *Create) Help() string {
	return `Usage: wasteband create index <jSON> 

  Create an index, etc.
  
`
}

// Synopsis provides a brief description of the command
func (c *Create) Synopsis() string {
	return "Create an Elasticsearch resource"
}
