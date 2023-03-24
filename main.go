package main

import (
	"flag"
	"fmt"
	"github.com/asteroid-engineering-actions/prettier/version"
	"os"
	"runtime"

	"github.com/sethvargo/go-githubactions"

	"github.com/asteroid-engineering-actions/prettier/prettieraction"
)

func main() {
	flag.Parse()

	if version.VersionFlag {
		version.PrintVersion(os.Stdout)
	}

	fmt.Printf("Asteroid Engineering: Prettier Action\nVersion: %s\non: %s_%s\n", version.String(), runtime.GOOS, runtime.GOARCH)

	action := githubactions.New()

	actionInputs := prettieraction.LoadActionInput(action)

	prettieraction.Hello(os.Stdout, actionInputs.Name)
}
