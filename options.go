package main

import (
	"os"
)

func validateOptions() {

	if githubAuthToken == "" {
		log.Fatal("Missing Github Auth Token. Please specify with -gt")
		os.Exit(1)
	}

	if configFile == "" {
		log.Fatal("Missing Config File. Please specify with -config-file")
		os.Exit(1)
	}

	if targetProcessAuthToken == "" {
		log.Fatal("Missing TargetProcess Auth Token. Please specify with -tt")
		os.Exit(1)
	}

}
