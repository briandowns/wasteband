package command

import (
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"

	"github.com/briandowns/wasteband/config"
	"github.com/briandowns/wasteband/utils"
	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// Show holds in the passed in configuration
type Show struct {
	config *config.Configuration
}

// NewShow creates a new instance of Delete
func NewShow(conf *config.Configuration) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Show{
			config: conf,
		}, nil
	}
}

// Run shows a given resource.
func (s *Show) Run(args []string) int {
	if flags.Has("help", args) || len(args) < 1 {
		fmt.Print(s.Help())
		return 1
	}

	// make sure that the command is a valid one
	if !utils.InSlice(args[0], allowedResources["show"]) {
		fmt.Print("ERROR: invalid option for show\n\n")
		fmt.Print(s.Help())
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
	case "cluster-health":
		result, err := utils.ESConn(d.config.Endpoint).Health()
		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 0, '\t', 0)

		v := reflect.ValueOf(result)

		fmt.Fprint(w, "\n")

		// iterate through the values of the struct and write to the tabwriter
		for i := 0; i < v.NumField(); i++ {
			// colorize the value of Status field based on it's content
			switch v.Field(i).Interface() {
			case "green":
				fmt.Fprintf(w, "%s\t%v\n", v.Type().Field(i).Name, green(v.Field(i).Interface().(string)))
				continue
			case "yellow":
				fmt.Fprintf(w, "%s\t%v\n", v.Type().Field(i).Name, yellow(v.Field(i).Interface().(string)))
				continue
			case "red":
				fmt.Fprintf(w, "%s\t%v\n", v.Type().Field(i).Name, red(v.Field(i).Interface().(string)))
				continue
			}
			fmt.Fprintf(w, "%s\t%v\n", v.Type().Field(i).Name, v.Field(i).Interface())
		}

		fmt.Fprintf(w, "\n")
		w.Flush() // dump out of writing to the tabwriter / os.Stdout
	}
	return 1
}

// Help provides full help inforamation for the subcommand
func (s *Show) Help() string { return "" }

// Synopsis provides a brief description of the command
func (s *Show) Synopsis() string {
	return "Show an Elasticsearch resource"
}
