// package command contains all command logic to succesfully run wasteband
package command

import (
	"github.com/fatih/color"
)

var (
	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
)

// allowedResources holds the allowable subcommand options
var allowedResources = map[string][]string{
	"create": {"index", "document"},
	"delete": {"index", "document"},
	"set":    {""},
	"show":   {"config", "indexes", "cluster-health", "cluster-state", "cluster-stats"},
	"search": {"index"},
}
