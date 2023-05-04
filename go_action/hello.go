package main

import (
	"fmt"
	"io"
)

func hello(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello %s!", name)
}
