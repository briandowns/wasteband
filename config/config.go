package config

import (
	"encoding/json"
	"os"
)

// Cluster represents an Elasticsearch cluster
type Cluster struct {
	Name     string `json:"cluster_name"`
	Endpoint string `json:"cluster_endpoint"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Configuration contains the ccpbot configuration
type Configuration struct {
	Clusters []Cluster
}

// GetConfig builds a config obj
func GetConfig(cf string) (*Configuration, error) {
	confFile, err := os.Open(cf)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(confFile)
	conf := &Configuration{}
	err = decoder.Decode(conf)

	if err != nil {
		return nil, err
	}

	return conf, nil
}
