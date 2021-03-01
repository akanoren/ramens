package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {
	out, err := exec.Command("go", "run", "cmd/ramen/main.go", "examples/hello.ramen").Output()
	if err != nil {
		t.Fatal(err)
	}

	// ignore line feeds or carridge returns
	ostr := strings.TrimSpace(string(out))
	if ostr != "Hello, World!" {
		t.Fatalf("%q must be equal to \"Hello, World!\"", string(out))
	}
}
