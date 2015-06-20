package command

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

func (s *Set) Run(args []string) int { return 0 }
func (s *Set) Help() string          { return "" }

// Synopsis provides a brief description of the command
func (s *Set) Synopsis() string {
	return "Set a cluster config parameter"
}
