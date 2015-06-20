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

func (l *Create) Run(args []string) int { return 0 }
func (l *Create) Help() string          { return "" }
func (l *Create) Synopsis() string      { return "Create an Elasticsearch resource" }
