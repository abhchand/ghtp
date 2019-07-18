package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Github struct {
		Organization string `yaml:"organization"`
		Repository   string `yaml:"repository"`
		AuthToken    string `yaml:"auth_token"`
	} `yaml:"github"`

	TargetProcess struct {
		Domain    string `yaml:"domain"`
		AuthToken string `yaml:"auth_token"`
	} `yaml:"target_process"`

	SyncRules []SyncRule `yaml:"sync"`
}

type SyncRule struct {
	IfHas   string `yaml:"if_has"`
	ThenSet string `yaml:"then_set"`
}

func readConfigFile(filepath string) Config {

	// Read file contents

	contents, err := ioutil.ReadFile(filepath)
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
