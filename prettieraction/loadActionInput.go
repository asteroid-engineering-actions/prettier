package prettieraction

import "github.com/sethvargo/go-githubactions"

type ActionInput struct {
	Name string
}

func LoadActionInput(a *githubactions.Action) ActionInput {
	actionInput := ActionInput{}

	actionInput.Name = a.GetInput("name")

	return actionInput
}
