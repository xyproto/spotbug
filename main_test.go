package main

import (
	"fmt"
	"strings"
	"testing"
)

const (
	promptHeader = ""
	outputFile   = ""
	wrapWidth    = 0
	verbose      = true
)

func TestFoundNoBugs(t *testing.T) {
	// Define the input parameters
	filenames := []string{"main.go", "spotbugs.go"}
	// Call the spotBug function
	output, err := spotBugs(defaultPrompt, "", "", wrapWidth, filenames, verbose)
	if err != nil {
		t.Fatalf("spotBug failed: %v", err)
	}

	fmt.Println(output)

	// Check if the output contains the word "puppy"
	if !strings.Contains(output, "puppy") {
		t.Errorf("Expected output to contain 'puppy', but it did not. Output: %s", output)
	}
}
