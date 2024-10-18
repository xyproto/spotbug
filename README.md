# SpotBug

Use LLMs that are running locally (or on a server defined in `OLLAMA_HOST`) to find bugs in the given source code files.

This repository contains a command line utility that can be used to find bugs in source code with Ollama.

### NOTE: This utility is a bit experimental, and a work in progress.

## Requirements

### Run-time requirements

* Ollama (the service must be up and running, and there must be enough memory and CPU and/or GPU available to be able to use the user-configured LLM model for vision tasks (for example the [`llava`](https://ollama.com/library/llava) model).
* [`llm-manager`](https://github.com/xyproto/llm-manager) can be used to configure which model to use for the `vision` task.

### Build-time requirements

* Go 1.22 or later

## Installation

    go install github.com/xyproto/spotbug@latest

The executable ends up in `~/go/bin` unless Go has been configured to place it somewhere else.

## Example use

```sh
spotbug main.go
```

> LGTM

## General info

* Version: 1.0.0
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
