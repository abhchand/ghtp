package main

import (
	"flag"
	"os"
)

var flagsForSync = defineFlagsForSync()
var (
	githubAuthToken        string
	githubOrganization     string
	githubRepository       string
	configFile             string
	targetProcessAuthToken string
)

var cmdSync = &Command{
	Name:             "sync",
	Args:             "[-config-file] [-gh] [-gh-org] [-gh-repo] [-tt] [-v]",
	ShortDescription: "Update TargetProcess state to match Github PR state",
	LongDescription:  "Update TargetProcess state to match Github PR state",
	Run:              runSync,
	Flags:            flagsForSync,
}

func defineFlagsForSync() flag.FlagSet {

	flagSet := *flag.NewFlagSet("version", flag.ExitOnError)

	flagSet.StringVar(
		&githubAuthToken, "gt", "", "Github Auth Token (Required)")
	flagSet.StringVar(
		&githubOrganization, "gh-org", "", "Github Org (Required)")
	flagSet.StringVar(
		&githubRepository, "gh-repo", "", "Github Repository (Required)")

	flagSet.StringVar(
		&configFile, "config-file", "", "Config File (Required)")

	flagSet.StringVar(
		&targetProcessAuthToken, "tt", "", "Target Process Token (Required)")

	flagSet.BoolVar(
		&verbose, "v", false, "Enable verbose output")

	return flagSet

}

func runSync(cmd *Command, args []string) {

	validateOptions()

	// Find eligible pull requests

	prs := findEligiblePullRequests()
	log.Infof("Found %v eligible pull request(s)", len(prs))
	log.Debug(prs)

	if len(prs) == 0 {
		log.Info("Exiting")
		os.Exit(0)
	}

	// Parse Config File

	config := readConfigFile()
	log.Debugf("Config: %v", config)

	// Set the appropriate TP state for each Pull Request

	for _, pr := range prs {

		targetProcessAssignable := findTargetProcessAssignableById(pr.targetProcessAssignableId())

		currentState := targetProcessAssignable.getCurrentEntityState()

		nextStateName := pr.expectedTargetProcessNextStateName(config.SyncRules)
		nextState := targetProcessAssignable.findNextStateByName(nextStateName)

		if len(nextStateName) == 0 {
			log.Infof("[%v] No next state could be determined from rule set",
				pr.toString())
			continue
		}

		if currentState.Name == nextStateName {
			log.Infof("[%v] Already has state: %v ✅", pr.toString(), currentState.toString())
			continue
		}

		if nextState.Id == 0 {
			log.Errorf("[%v] Invalid state: %v", pr.toString(), nextState.toString())
			continue
		}

		log.Infof("[%v] current state: %v, next state: %v",
			pr.toString(),
			currentState.toString(),
			nextState.toString())

		updateTargetProcessEntityState(targetProcessAssignable, nextState)

	}

}
