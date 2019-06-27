package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Github struct {
		Organization string `yaml:"organization"`
		Repository   string `yaml:"repository"`
	} `yaml:"github"`

	SyncRules []SyncRule `yaml:"sync"`
}

type SyncRule struct {
	if_has   string `yaml:"if_has"`
	then_set string `yaml:"then_set"`
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
