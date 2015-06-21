// package command contains all command logic to succesfully run wasteband
package command

// allowedResources holds the allowable subcommand options
var allowedResources = map[string][]string{
	"create": {"index", "document"},
	"delete": {"index", "document"},
	"set":    {""},
	"show":   {"indexes", "cluster-health", "cluster-state", "cluster-stats"},
	"search": {""},
}
