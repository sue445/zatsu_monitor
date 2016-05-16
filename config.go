package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config map[string](map[string]string)

func LoadConfigFromData(yamlData string) (Config, error) {
	c := Config{}

	err := yaml.Unmarshal([]byte(yamlData), &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func LoadConfigFromFile(yamlFile string) (Config, error) {
	buf, err := ioutil.ReadFile(yamlFile)

	if err != nil {
		return nil, err
	}

	return LoadConfigFromData(string(buf))
}
