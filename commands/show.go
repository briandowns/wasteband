package commands

import (
	"github.com/mitchellh/cli"
)

type Show struct{}

func NewShow() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Show{}, nil
	}
}

func (s *Show) Run(args []string) int { return 0 }
func (s *Show) Help() string          { return "" }
func (s *Show) Synopsis() string      { return "" }
