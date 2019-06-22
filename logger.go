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

	logging.SetBackend(backend)

	return logging.MustGetLogger("logger")

}
