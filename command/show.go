package command

import (
	"github.com/briandowns/wasteband/config"
	"github.com/mitchellh/cli"
)

type Show struct {
	config *config.Configuration
}

func NewShow(conf *config.Configuration) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Show{
			config: conf,
		}, nil
	}
}

func (s *Show) Run(args []string) int { return 0 }
func (s *Show) Help() string          { return "" }
func (s *Show) Synopsis() string      { return "Show an Elasticsearch resource" }
