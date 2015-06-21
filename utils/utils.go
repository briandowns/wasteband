package utils

import "github.com/mattbaird/elastigo/lib"

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
