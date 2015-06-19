package commands

import (
	"github.com/mitchellh/cli"
)

type Create struct{}

func NewCreate() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Create{}, nil
	}
}

func (l *Create) Run(args []string) int { return 0 }
func (l *Create) Help() string          { return "" }
func (l *Create) Synopsis() string      { return "" }
