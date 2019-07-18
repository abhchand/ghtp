package main

import (
	"io/ioutil"
	"os"
	"testing"
)

var configDir string

func TestReadConfigFile(t *testing.T) {

	//
	// Setup
	//

	configDir = "/tmp/ghtp"
	configFile = configDir + "/config.yml"

	body := `
github:
  organization: abhchand
  repository: ghtp
  auth_token: abcde
target_process:
  domain: acme
  auth_token: fghij

sync:
  - if_has: bug
    then_set: Development
  - if_has: Code Review
    then_set: In Review
`

	t.Run("Success", func(t *testing.T) {

		err := setupConfigFile(body)
		if err != nil {
			t.Error("Error setting up test fixture file. Got", err)
		}

		actual := readConfigFile(configFile)

		if actual.Github.Organization != "abhchand" {
			t.Error("expected", "abhchand", "got", actual)
		}

	})

}

func setupConfigFile(body string) error {

	var err error

	// Create directory if it does not exist
	_, err = os.Stat(configDir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(configDir, 0775)

		if err != nil {
			log.Debugf("Error creating directory %v", configDir)
			return err
		}
	}

	// Overwrite existing test file
	err = ioutil.WriteFile(configFile, []byte(body), 0755)
	if err != nil {
		return err
	}

	return nil

}
