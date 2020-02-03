package main

import (
	"flag"
	"fmt"
)

const version = "0.3"

var flagsForVersion = defineFlagsForVersion()

var cmdVersion = &Command{
	Name:             "version",
	ShortDescription: "Display version",
	LongDescription:  "Displays the version of ghtp",
	Run:              runVersion,
	Flags:            flagsForVersion,
}

func defineFlagsForVersion() flag.FlagSet {

	flagSet := *flag.NewFlagSet("version", flag.ExitOnError)

	return flagSet

}

func runVersion(cmd *Command, args []string) {
	fmt.Printf("ghtp (%s)\n", version)
}
