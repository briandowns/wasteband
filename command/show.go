package command

import (
	"fmt"
	"reflect"

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

	// process the subcommand and it's options
	switch args[0] {
	case "config":
		s.showConfig()
	case "indices":
		s.showIndices()
	case "cluster-health":
		s.showClusterHealth()
	case "cluster-stats":
		//
	default:
		fmt.Println("ERROR: invalid option for show\n")
	}

	return 1
}

// Help provides full help inforamation for the subcommand
func (s *Show) Help() string {
	return `Usage: wasteband show <option> <arguments> 

  Show a resource

Options:

  config            Display the current wasteband configuration
  indexes           Display all indexes
  cluster-health    Display the health of the cluster
  cluster-state     Display the state of the cluster
  cluster-stats     Display the stats from the cluster
  
`
}

// Synopsis provides a brief description of the command
func (s *Show) Synopsis() string {
	return "Show an Elasticsearch resource"
}

// showConfig outputs the current running configuration
func (s *Show) showConfig() {
	fmt.Print("\nwasteband config:\n")
	w := utils.NewTabWriter()

	v := reflect.ValueOf(*s.config)

	fmt.Fprint(w, "\n")

	// iterate through the values of the struct and write to the tabwriter
	for i := 0; i < v.NumField(); i++ {
		fmt.Fprintf(w, "%s\t%v\n", v.Type().Field(i).Name, v.Field(i).Interface())
	}

	fmt.Fprintf(w, "\n")
	w.Flush()
}

// showIndices outputs data associated with all indices
func (s *Show) showIndices() {
	// most of this functionaity will have to wait on PR 209 in the elastigo
	// project to be merged.
	result := utils.ESConn(s.config.Endpoint).GetCatIndexInfo("")

	w := utils.NewTabWriter()

	fmt.Fprintf(w, "\nHealth\tName\tShards\tReplicas\tDocuments\tSize")
	fmt.Fprintf(w, "\n----------\t----------\t----------\t----------\t----------\t----------\n")

	for _, i := range result {
		fmt.Fprintf(w, "%s\t%s\t%d\t%d\t%d\t%d\n",
			i.Health, i.Name, i.Shards, i.Replicas, i.Docs.Deleted, i.Store.Size)
	}

	fmt.Fprintf(w, "\n")
	w.Flush()
}

// showClusterHealth outputs the health of the cluster
func (s *Show) showClusterHealth() int {
	result, err := utils.ESConn(s.config.Endpoint).Health()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return 1
	}

	fmt.Printf("\nCluster Health: %s\n", s.config.Name)
	w := utils.NewTabWriter()

	v := reflect.ValueOf(result)

	fmt.Fprint(w, "\n")

	// iterate through the values of the struct and write to the tabwriter
	for i := 0; i < v.NumField(); i++ {
		// colorize the value of Status field based on it's content
		switch v.Field(i).Interface() {
		case "green":
			fmt.Fprintf(w, "%s\t%v\n",
				v.Type().Field(i).Name, green(v.Field(i).Interface().(string)))
			continue
		case "yellow":
			fmt.Fprintf(w, "%s\t%v\n",
				v.Type().Field(i).Name, yellow(v.Field(i).Interface().(string)))
			continue
		case "red":
			fmt.Fprintf(w, "%s\t%v\n",
				v.Type().Field(i).Name, red(v.Field(i).Interface().(string)))
			continue
		}
		fmt.Fprintf(w, "%s\t%v\n",
			v.Type().Field(i).Name, v.Field(i).Interface())
	}

	fmt.Fprintf(w, "\n")
	w.Flush() // dump out of writing to the tabwriter / os.Stdout
	return 1
}
