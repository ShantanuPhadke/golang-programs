package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Shantanu")
	actualOutput := buffer.String()
	expectedOutput := "Hello, Shantanu"

	if actualOutput != expectedOutput {
		t.Errorf("Actual string printed out %q but expected %q", actualOutput, expectedOutput)
	}
}
