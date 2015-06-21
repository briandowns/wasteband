package command

import (
	"fmt"

	"github.com/briandowns/wasteband/config"
	"github.com/briandowns/wasteband/utils"
	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// Delete holds in the passed in configuration
type Delete struct {
	config *config.Configuration
}

// NewDelete creates a new instance of Delete
func NewDelete(conf *config.Configuration) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Delete{
			config: conf,
		}, nil
	}
}

// Delete removes a given resource.
func (d *Delete) Run(args []string) int {
	if flags.Has("help", args) || len(args) < 1 {
		fmt.Print(d.Help())
		return 1
	}

	// make sure that the command is a valid one
	if !utils.InSlice(args[0], allowedResources["delete"]) {
		fmt.Print("ERROR: invalid option for delete\n\n")
		fmt.Print(d.Help())
		return 1
	}

	// process the subcommand and it's options
	switch args[1] {
	case "index":
		_, err := utils.ESConn(d.config.Endpoint).DeleteIndex(args[2])
		if err != nil {
			fmt.Printf("%v\n", err.Error())
			return 1
		}
	case "document":
		_, err := utils.ESConn(d.config.Endpoint).Delete(args[2], args[3], args[4], nil)
		if err != nil {
			fmt.Printf("%v\n", err)
			return 1
		}
	}

	return 1
}

// Help provides full help inforamation for the subcommand
func (d *Delete) Help() string {
	return `Usage: wasteband delete <option> <arguments> 

  Delete a resource

Options:

  index       Delete an index.  Takes the index name as an argument.
  document    Delete a document. Takes the index, type, and id as arguments.
  
`
}

// Synopsis provides a brief description of the command
func (d *Delete) Synopsis() string {
	return "Delete an Elasticsearch resource"
}
