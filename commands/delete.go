package commands

import (
	"github.com/mitchellh/cli"
)

type Delete struct{}

func NewDelete() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Delete{}, nil
	}
}

func (l *Delete) Run(args []string) int { return 0 }
func (l *Delete) Help() string          { return "" }
func (l *Delete) Synopsis() string      { return "" }
