package main

import (
	"flag"
)

var flagsForSync = defineFlagsForSync()
var (
	githubAuthToken        string
	targetProcessAuthToken string
)

var cmdSync = &Command{
	Name:             "sync",
	Args:             "[-g] [-t] [-v]",
	ShortDescription: "Update TargetProcess state to match Github PR state",
	LongDescription:  "Update TargetProcess state to match Github PR state",
	Run:              runSync,
	Flags:            flagsForSync,
}

func defineFlagsForSync() flag.FlagSet {

	flagSet := *flag.NewFlagSet("version", flag.ExitOnError)

	flagSet.StringVar(&githubAuthToken, "g", "", "Github Auth Token (Required)")
	flagSet.StringVar(&targetProcessAuthToken, "t", "", "Target Process Token (Required)")
	flagSet.BoolVar(&verbose, "v", false, "Enable verbose output")

	return flagSet

}

func runSync(cmd *Command, args []string) {
	validateOptions()
}
