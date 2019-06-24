package main

import (
	"fmt"
	"os"

	"github.com/op/go-logging"
)

func initializeLogger() *logging.Logger {

	var format = logging.MustStringFormatter(
		fmt.Sprintf(
			"%v %v",
			"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.5s}",
			"%{id:03x}%{color:reset} %{message}",
		),
	)

	var backend = logging.NewBackendFormatter(
		logging.NewLogBackend(os.Stdout, "", 0),
		format)

	var level logging.Level

	if verbose {
		level = logging.DEBUG
	} else {
		level = logging.INFO
	}

	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(level, "")

	logging.SetBackend(backendLeveled)

	return logging.MustGetLogger("logger")

}
