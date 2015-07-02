package config

import (
	"encoding/json"
	"os"
)

// Configuration contains the ccpbot configuration
type Configuration struct {
	Name     string `json:"cluster_name"`
	Endpoint string `json:"cluster_endpoint"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// HasEndpoint checks if an endpoint has been set
func (c *Configuration) HasEndpoint() bool {
	if c.Endpoint != "" {
		return true
	}
	return false
}

// Load builds a config obj
func Load(cf string) (*Configuration, error) {
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
