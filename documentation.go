package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

var usageTemplate = `
ghtp is a tool that provides integration between Github and TargetProcess

Usage:
	ghtp <command> [arguments]

Available commands:
{{range .}}
    {{.Name | printf "%-8s"}}` + "\t\t" + `{{.ShortDescription}}{{end}}

Use "ghtp help <command>" for more information about a command.
`

func printHelp(args []string) {

	if len(args) == 0 {
		printUsage(os.Stdout)
		return
	}

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: ghtp help <command>\n\n")
		fmt.Fprintf(os.Stderr, "Too many arguments given.\n")
		os.Exit(1)
	}

	for _, cmd := range commands {
		if cmd.Name == args[0] {
			cmd.PrintUsage()
			return
		}
	}

}

func printUsageError() {
	printUsage(os.Stderr);
	os.Exit(1)
}

func printUsage(w io.Writer) {
	renderTemplate(w, usageTemplate, commands)
}

func renderTemplate(w io.Writer, text string, data interface{}) {

	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace})

	template.Must(t.Parse(strings.TrimSpace(text) + "\n\n"))

	err := t.Execute(w, data)
	if err != nil {
		panic(err)
	}

}
