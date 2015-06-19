package commands

import (
	"github.com/briandowns/wasteband/config"
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

func (l *Delete) Run(args []string) int { return 0 }
func (l *Delete) Help() string          { return "" }
func (l *Delete) Synopsis() string      { return "" }
