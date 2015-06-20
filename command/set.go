package command

import (
	"fmt"

	"github.com/briandowns/wasteband/config"
	"github.com/fatih/flags"
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

func (s *Set) Run(args []string) int {
	if flags.Has("help", args) {
		fmt.Print(s.Help())
		return 1
	}
	return 1
}

func (s *Set) Help() string { return "" }

// Synopsis provides a brief description of the command
func (s *Set) Synopsis() string {
	return "Set a cluster config parameter"
}
