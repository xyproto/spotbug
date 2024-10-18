package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/xyproto/ollamaclient/v2"
	"github.com/xyproto/usermodel"
	"github.com/xyproto/wordwrap"
)

// spotBugs uses Ollama and the given model to try to find bugs in one or more source code files
// prompt is the start of the multimodal prompt: the instructions which will be followed by the source code
// outputFile is an (optional) filename to write the resulting description to
// writeWidth is the width that the returned or written string should be wrapped to, if it is >0
// filenames is a list of input source code files
// The result is returned as a string.
func spotBugs(prompt, model, outputFile string, wrapWidth int, filenames []string, verbose bool) (string, error) {
	if prompt == "" {
		return "", errors.New("the given prompt can not be empty")
	}

	if wrapWidth == -1 {
		wrapWidth = getTerminalWidth()
	}

	if len(filenames) < 1 {
		return "", fmt.Errorf("no source code filenames provided")
	}

	var sourceCode []string
	for _, filename := range filenames {
		logVerbose(verbose, "[%s] Reading... ", filename)
		data, err := os.ReadFile(filename)
		if err == nil { // success
			sourceCode = append(sourceCode, string(data))
			logVerbose(verbose, "OK\n")
		} else {
			logVerbose(verbose, "FAILED: "+err.Error()+"\n")
		}
	}

	if len(sourceCode) == 0 {
		return "", fmt.Errorf("no source code to analyze")
	}
	if model == "" {
		model = usermodel.GetCodeModel() // get the llm-manager defined model for the "vision" task, perhaps "llava"
	}

	oc := ollamaclient.New()
	oc.ModelName = model
	oc.Verbose = verbose

	if err := oc.PullIfNeeded(true); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to pull model: %v\n", err)
		fmt.Fprintln(os.Stderr, "Ollama must be up and running")
		os.Exit(1)
	}

	oc.SetReproducible()

	promptAndSourceCode := append([]string{prompt}, sourceCode...)

	logVerbose(verbose, "[%s] Analyzing...\n", oc.ModelName)
	output, err := oc.GetOutput(promptAndSourceCode...)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}
	logVerbose(verbose, "[%s] Analysis complete.\n", oc.ModelName)

	if output == "" {
		return "", errors.New("generated output is empty")
	}

	if wrapWidth > 0 {
		lines, err := wordwrap.WordWrap(output, wrapWidth)
		if err == nil { // success
			output = strings.Join(lines, "\n")
		}
	}

	if outputFile != "" {
		err := os.WriteFile(outputFile, []byte(output), 0o644)
		if err != nil {
			return "", fmt.Errorf("error writing to file: %v", err)
		}
		return "", nil
	}

	return output, nil
}
