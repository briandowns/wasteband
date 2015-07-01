package command

import (
	"fmt"

	"github.com/briandowns/wasteband/config"
	"github.com/briandowns/wasteband/utils"
	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// Optimize holds in the passed in configuration
type Optimize struct {
	config *config.Configuration
}

// NewOptimize creates a new instance of Delete
func NewOptimize(conf *config.Configuration) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Optimize{
			config: conf,
		}, nil
	}
}

// Run optimizes a given resource.
func (o *Optimize) Run(args []string) int {
	if flags.Has("help", args) || len(args) < 1 {
		fmt.Print(o.Help())
		return 1
	}

	// process the subcommand and it's options
	switch args[1] {
	case "index":
		result, err := utils.ESConn(o.config.Endpoint).OptimizeIndices(nil, args[2])
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			return 1
		}
		fmt.Printf("Result: %b, ShardStatus: %v\n", result.Ok)

	default:
		fmt.Printf("ERROR: invalid option for optimize\n")
	}

	return 1
}

// Help provides full help inforamation for the subcommand
func (o *Optimize) Help() string {
	return `Usage: wasteband optimize <argument>

  Optimize an index

Options:

  index       Optimize an index.  Takes the index name as an argument.
  
`
}

// Synopsis provides a brief description of the command
func (o *Optimize) Synopsis() string {
	return "Optimize an Elasticsearch index"
}
