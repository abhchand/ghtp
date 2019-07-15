package main

import (
	"flag"
	"os"
)

var flagsForSync = defineFlagsForSync()
var (
	configFile string

	githubOrganization string
	githubRepository   string
	githubAuthToken    string

	targetProcessDomain    string
	targetProcessAuthToken string
)

var cmdSync = &Command{
	Name:             "sync",
	Args:             "[-config-file] [-v]",
	ShortDescription: "Update TargetProcess state to match Github PR state",
	LongDescription:  "Update TargetProcess state to match Github PR state",
	Run:              runSync,
	Flags:            flagsForSync,
}

func defineFlagsForSync() flag.FlagSet {

	flagSet := *flag.NewFlagSet("version", flag.ExitOnError)

	flagSet.StringVar(
		&configFile, "config-file", "config.yml", "YML Config File of options")

	flagSet.BoolVar(
		&verbose, "v", false, "Enable verbose output")

	return flagSet

}

func runSync(cmd *Command, args []string) {

	// Validate command line options

	validateSyncArguments()

	// Parse, validate, and load Config File

	config := readConfigFile()
	log.Debugf("Config: %v", config)

	validateSyncConfigFile(config)
	loadSyncOptions(config)

	// Find eligible pull requests

	prs := findEligiblePullRequests()
	log.Infof("Found %v eligible pull request(s)", len(prs))
	log.Debugf("%v", prs)

	if len(prs) == 0 {
		log.Info("Exiting")
		os.Exit(0)
	}

	// Set the appropriate TP state for each Pull Request

	for _, pr := range prs {
		synchronize(pr, config)
	}

}

// Synchronizes TargetProcess state to match Github state for a given Pull
// Request. It applies the `sync` rules from the config file to determine what
// TargetProcess state to set for a given set of Github labels.
//
// It also performs type checking against TargetProcess to ensure that the
// desired state is a valid workflow state for that Assignable to move to.
//
//   - If no action can be deterined from the rule set, do nothing
//   - If a TP Assignable already has the desired state, do nothing
//   - If the desired TargetProcess state is invalid, log an error and continue
//   - If none of the above apply, attempt to update TargetProcess to the desired
//     state
//
func synchronize(pr PullRequest, config Config) {

	targetProcessAssignable := findTargetProcessAssignableById(pr.targetProcessAssignableId())

	currentState := targetProcessAssignable.getCurrentEntityState()

	nextStateName := pr.expectedTargetProcessNextStateName(config.SyncRules)
	nextState := targetProcessAssignable.findNextStateByName(nextStateName)

	if len(nextStateName) == 0 {
		log.Infof("[%v] No next state could be determined from rule set",
			pr.toString())
		return
	}

	if currentState.Name == nextStateName {
		log.Infof("[%v] Already has state: %v ✅", pr.toString(), currentState.toString())
		return
	}

	if nextState.Id == 0 {
		log.Errorf("[%v] Invalid state: %v", pr.toString(), nextState.toString())
		return
	}

	log.Infof("[%v] current state: %v, next state: %v",
		pr.toString(),
		currentState.toString(),
		nextState.toString())

	updateTargetProcessEntityState(targetProcessAssignable, nextState)

}

func loadSyncOptions(config Config) {

	githubOrganization = config.Github.Organization
	githubRepository = config.Github.Repository

	targetProcessDomain = config.TargetProcess.Domain
	targetProcessAuthToken = config.TargetProcess.AuthToken

	var token string

	if token = os.Getenv("GITHUB_AUTH_TOKEN"); token != "" {
		githubAuthToken = token
	} else {
		githubAuthToken = config.Github.AuthToken
	}

	if token = os.Getenv("TARGET_PROCESS_AUTH_TOKEN"); token != "" {
		targetProcessAuthToken = token
	} else {
		targetProcessAuthToken = config.TargetProcess.AuthToken
	}

}

func validateSyncArguments() {

	if configFile == "" {
		log.Fatal("Missing Config File. Please specify with -config-file")
		os.Exit(1)
	}

}

func validateSyncConfigFile(config Config) {

	if config.Github.Organization == "" {
		log.Fatal("Missing Github Organization. Please specify github.organization in config file")
		os.Exit(1)
	}

	if config.Github.Repository == "" {
		log.Fatal("Missing Github Repository. Please specify github.repository in config file")
		os.Exit(1)
	}

	if os.Getenv("GITHUB_AUTH_TOKEN") == "" && config.Github.AuthToken == "" {
		log.Fatal("Missing Github AuthToken. Please specify github.auth_token in config file or $GITHUB_AUTH_TOKEN")
		os.Exit(1)
	}

	if config.TargetProcess.Domain == "" {
		log.Fatal("Missing TargetProcess Domain. Please specify target_process.domain in config file")
		os.Exit(1)
	}

	if os.Getenv("TARGET_PROCESS_AUTH_TOKEN") == "" && config.TargetProcess.AuthToken == "" {
		log.Fatal("Missing TargetProcess AuthToken. Please specify target_process.auth_token in config file or $TARGET_PROCESS_AUTH_TOKEN")
		os.Exit(1)
	}

}
