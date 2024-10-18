package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

const versionString = "SpotBug 1.0.0"

const defaultPrompt = "You are an expert software developer with a PhD in Computer Science. Say \"ALL IS WELL\" if you can not spot any bugs in the following code, or \"LGTM\" if you are even just a bit unsure."

func main() {
	var (
		outputFile   string
		promptHeader string
		wrapWidth    int
		showVersion  bool
		verbose      bool
		model        string
	)

	pflag.BoolVarP(&verbose, "verbose", "V", false, "verbose output")
	pflag.StringVarP(&model, "model", "m", "", "Specify the Ollama model to use")
	pflag.StringVarP(&promptHeader, "prompt", "p", defaultPrompt, "Provide a custom prompt header")
	pflag.StringVarP(&outputFile, "output", "o", "", "Specify an output file")
	pflag.IntVarP(&wrapWidth, "wrap", "w", 0, "Word wrap at specified width. Use '-1' for terminal width")
	pflag.BoolVarP(&showVersion, "version", "v", false, "display version")

	pflag.Parse()

	if showVersion {
		fmt.Println(versionString)
		return
	}

	filenames := pflag.Args()

	output, err := spotBugs(promptHeader, model, outputFile, wrapWidth, filenames, verbose)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if output != "" {
		fmt.Println(output)
	}
}
