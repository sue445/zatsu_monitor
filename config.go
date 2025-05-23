package main

import (
	"github.com/cockroachdb/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config represents config file
type Config map[string](map[string]string)

// LoadConfigFromData load config from yaml data
func LoadConfigFromData(yamlData string) (Config, error) {
	c := Config{}

	err := yaml.Unmarshal([]byte(yamlData), &c)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}

// LoadConfigFromFile load config from yaml file
func LoadConfigFromFile(yamlFile string) (Config, error) {
	buf, err := ioutil.ReadFile(yamlFile)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return LoadConfigFromData(string(buf))
}
