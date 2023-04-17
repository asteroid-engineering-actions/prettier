package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/sethvargo/go-githubactions"

	"github.com/asteroid-engineering-actions/prettier/go_action/version"
)

func main() {
	flag.Parse()

	if version.VersionFlag {
		version.PrintVersion(os.Stdout)
	}

	fmt.Printf("Asteroid Engineering: Prettier Action\nVersion: %s\non: %s_%s\n\n", version.String(), runtime.GOOS, runtime.GOARCH)

	action := githubactions.New()

	actionInputs := loadActionInput(action)

	hello(os.Stdout, actionInputs.Name)
}

func start(handler handlerFunc) {
	aEvent := &actionEvent{
		getEnv: os.Getenv,
	}

	exitCode, err := handler(aEvent)

	if err != nil {
		os.Exit(1)
	}

	os.Exit(exitCode)
}
