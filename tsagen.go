package main

import (
	"os"

	logger "github.com/apsdehal/go-logger"
	"github.com/sparkymat/tsagen/commands"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var log, _ = logger.New("tsagen", 1, os.Stdout)

var (
	app = kingpin.New("tsagen", "Utility for generating Typescript React/Redux apps, and pages for them")

	initApp      = app.Command("init", "Initializes a new TSApp with the specified name")
	initAppName  = initApp.Arg("name", "Name of the new app").Required().String()
	initAppForce = initApp.Flag("force", "Overwrite existing app folder").Short('f').Bool()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Register user
	case initApp.FullCommand():
		returnCode := commands.InitialiseApp(*initAppName, *initAppForce)
		os.Exit(returnCode)
	}
}
