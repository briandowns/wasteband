package main

import "github.com/mattbaird/elastigo/lib"

// esConn connects to the given Elasticsearch cluster
func esConn(server string) *elastigo.Conn {
	c := elastigo.NewConn()
	c.Domain = server
	c.Port = "9200"
	return c
}
