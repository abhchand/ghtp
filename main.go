package main

import (
	"fmt"
	"flag"
	"os"
)

var log = initializeLogger()
var verbose	bool

var commands = []*Command{
	cmdVersion,
	cmdSync,
}


func main() {

	args := parseArgs();

	if len(args) < 1 {
		printUsageError()
	}

	if args[0] == "help" {
		printHelp(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name == args[0] {

			cmd.Flags.Parse(args[1:])
			cmdArgs := cmd.Flags.Args()

			cmd.Run(cmd, cmdArgs)
			return

		}
	}

	fmt.Fprintf(os.Stderr, "ghtp: unknown command %q\n", args[0])
	fmt.Fprintf(os.Stderr, "Run 'ghtp help' for usage.\n")
	os.Exit(2)

}

func parseArgs() []string {

	flag.Parse()
	return flag.Args()

}
