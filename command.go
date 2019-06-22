package main

import (
	"flag"
	"fmt"
	"os"
)

type Command struct {
	Run              func(cmd *Command, args []string)
	Name             string
	Args             string
	ShortDescription string
	LongDescription  string
	Flags            flag.FlagSet
}

func (c *Command) PrintUsage() {

	str := ""
	str += fmt.Sprintf("%s\n\n", c.LongDescription)
	str += fmt.Sprintf("Usage:\n\tghtp %s %s\n", c.Name, c.Args)

	fmt.Printf("%v\n\n", str)
	c.Flags.PrintDefaults()
	fmt.Print("\n")

	os.Exit(1)

}
