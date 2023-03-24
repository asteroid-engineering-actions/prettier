package prettieraction

import (
	"bytes"
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	inputName := "AJ"
	outputData := bytes.Buffer{}

	Hello(&outputData, inputName)

	got := outputData.String()
	want := fmt.Sprintf("Hello %s!", inputName)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
