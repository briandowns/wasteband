package main

import (
	"encoding/json"
	"os"
)

// cluster represents an Elasticsearch cluster
type cluster struct {
	name     string `json:"cluster_name"`
	endpoint string `json:"cluster_endpoint"`
	username string `json:"username"`
	password string `json:"password"`
}

// configuration contains the ccpbot configuration
type configuration struct {
	clusters []cluster
}

// getConfig builds a config obj
func getConfig(cf string) (*configuration, error) {
	confFile, err := os.Open(cf)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(confFile)
	conf := &configuration{}
	err = decoder.Decode(conf)

	if err != nil {
		return nil, err
	}

	return conf, nil
}
