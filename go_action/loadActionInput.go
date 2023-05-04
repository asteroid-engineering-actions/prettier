package main

import "github.com/sethvargo/go-githubactions"

type actionInput struct {
	Name string
}

func loadActionInput(a *githubactions.Action) actionInput {
	actionInput := actionInput{}

	actionInput.Name = a.GetInput("name")

	return actionInput
}
