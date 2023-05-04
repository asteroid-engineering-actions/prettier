package main

import (
	"bytes"
	"github.com/sethvargo/go-githubactions"
	"testing"
)

func TestLoadActionInput(t *testing.T) {
	actionLogOutput := bytes.NewBuffer(nil)
	envMap := map[string]string{
		"INPUT_NAME": "AJ",
	}
	getEnv := func(key string) string {
		return envMap[key]
	}

	action := githubactions.New(
		githubactions.WithGetenv(getEnv),
		githubactions.WithWriter(actionLogOutput),
	)

	got := loadActionInput(action)
	want := actionInput{
		Name: "AJ",
	}

	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
}
