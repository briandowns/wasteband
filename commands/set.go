package commands

import (
	"github.com/mitchellh/cli"
)

type Set struct{}

func NewSet() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Set{}, nil
	}
}

func (l *Set) Run(args []string) int { return 0 }
func (l *Set) Help() string          { return "" }
func (l *Set) Synopsis() string      { return "" }
