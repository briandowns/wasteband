package commands

import (
	"github.com/briandowns/wasteband/config"
	"github.com/mitchellh/cli"
)

type Set struct {
	config *config.Configuration
}

func NewSet(conf *config.Configuration) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Set{
			config: conf,
		}, nil
	}
}

func (l *Set) Run(args []string) int { return 0 }
func (l *Set) Help() string          { return "" }
func (l *Set) Synopsis() string      { return "" }
