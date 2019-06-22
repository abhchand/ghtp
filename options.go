package main

import (
	"os"
)

func validateOptions() {

	if githubAuthToken == "" {
		log.Fatal("Missing Github Auth Token. Please specify with -g")
		os.Exit(1)
	}

	if targetProcessAuthToken == "" {
		log.Fatal("Missing TargetProcess Auth Token. Please specify with -t")
		os.Exit(1)
	}

}
