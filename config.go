package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	SyncRules []SyncRule `yaml:"sync"`
}

type SyncRule struct {
	IfHas   string `yaml:"IfHas"`
	ThenSet string `yaml:"ThenSet"`
}

func readConfigFile() Config {

	// Read file contents

	contents, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	// Parse contents

	var config Config
	err = yaml.Unmarshal(contents, &config)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return config

}
