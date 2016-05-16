package main

import (
	"gopkg.in/yaml.v2"
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
