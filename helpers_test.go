package main

import (
	"os"
	"testing"
)

func TestAbsolutePath(t *testing.T) {

	mockWorkingDir := "/tmp/ghtp"
	currWorkingDir, _ := os.Getwd()

	// Create directory if it does not exist
	_, err := os.Stat(mockWorkingDir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(mockWorkingDir, 0775)

		if err != nil {
			t.Error("Error creating directory " + mockWorkingDir)
		}
	}

	// Change to working directory, ensuring we change back when done
	os.Chdir(mockWorkingDir)
	defer os.Chdir(currWorkingDir)

	t.Run("Success", func(t *testing.T) {

		testCases := map[string]string{
			"config.yml":           "/private/tmp/ghtp/config.yml",
			"./config.yml":         "/private/tmp/ghtp/config.yml",
			"../config.yml":        "/private/tmp/config.yml",
			"/tmp/ghtp/config.yml": "/tmp/ghtp/config.yml",
		}

		for input, expected := range testCases {

			actual := absolutePath(input)

			if actual != expected {
				t.Error("expected", expected, "got", actual)
			}
		}

	})

}
