package main

import (
	"os"
)

func validateOptions() {

	if githubAuthToken == "" {
		log.Fatal("Missing Github Auth Token. Please specify with -gt")
		os.Exit(1)
	}

	if githubOrganization == "" {
		log.Fatal("Missing Github Organization. Please specify with -gh-org")
		os.Exit(1)
	}

	if githubRepository == "" {
		log.Fatal("Missing Github Repository. Please specify with -gh-repo")
		os.Exit(1)
	}

	if targetProcessAuthToken == "" {
		log.Fatal("Missing TargetProcess Auth Token. Please specify with -tt")
		os.Exit(1)
	}

}
