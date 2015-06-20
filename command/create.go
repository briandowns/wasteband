package command

import (
	"github.com/briandowns/wasteband/config"
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

func (c *Create) Run(args []string) int { return 0 }
func (c *Create) Help() string          { return "" }

// Synopsis provides a brief description of the command
func (c *Create) Synopsis() string {
	return "Create an Elasticsearch resource"
}
