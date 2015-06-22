package utils

import (
	"os"
	"text/tabwriter"

	"github.com/mattbaird/elastigo/lib"
)

// NewTabWriter setups up a new and configured tabwriter
func NewTabWriter() *tabwriter.Writer {
	t := new(tabwriter.Writer)
	t.Init(os.Stdout, 0, 8, 1, '\t', 0)
	return t
}

// InSlice checks if the given string is in the given slice
func InSlice(item string, dr []string) bool {
	for _, i := range dr {
		if item == i {
			return true
		}
	}
	return false
}

// ESConn connects to the given Elasticsearch cluster
func ESConn(server string) *elastigo.Conn {
	c := elastigo.NewConn()
	c.Domain = server
	c.Port = "9200"
	return c
}
