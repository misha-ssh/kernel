# Misha ssh - kernel

[![Go Report Card](https://goreportcard.com/badge/github.com/misha-ssh/kernel)](https://goreportcard.com/report/github.com/misha-ssh/kernel)
[![Go Docs](https://godoc.org/github.com/misha-ssh/kernel?status.svg)](https://godoc.org/github.com/misha-ssh/kernel)
[![Release](https://img.shields.io/github/release/misha-ssh/kernel?status.svg)](https://github.com/misha-ssh/kernel/releases)
[![Action Lint](https://github.com/misha-ssh/kernel/actions/workflows/lint.yml/badge.svg)](https://github.com/misha-ssh/kernel)
[![Action Tests](https://github.com/misha-ssh/kernel/actions/workflows/tests.yml/badge.svg)](https://github.com/misha-ssh/kernel)
[![Action Coverage](https://github.com/misha-ssh/kernel/actions/workflows/coverage.yml/badge.svg)](https://github.com/misha-ssh/kernel)

This package acts as the core for an ssh client written in go

Documentation package - [link](https://pkg.go.dev/github.com/misha-ssh/kernel)

Made using data from packages:
* [crypto](https://pkg.go.dev/golang.org/x/crypto)
* [go-keyring](http://github.com/zalando/go-keyring)
* [term](https://pkg.go.dev/golang.org/x/term)

## 📝 Features

- **Multi-Model:** choose from a wide range of LLMs or add your own via OpenAI- or Anthropic-compatible APIs
- **Flexible:** switch LLMs mid-session while preserving context
- **Session-Based:** maintain multiple work sessions and contexts per project
- **LSP-Enhanced:** Crush uses LSPs for additional context, just like you do
- **Extensible:** add capabilities via MCPs (`http`, `stdio`, and `sse`)
- **Works Everywhere:** first-class support in every terminal on macOS, Linux, Windows (PowerShell and WSL), FreeBSD, OpenBSD, and NetBSD

## Install

install this package in your repository

```bash
go get github.com/misha-ssh/kernel
```

## Examples

You will be provided with a list of commands that you can use in your projects

The code with the commands will be on the way - [link](./examples/command)

### Connect

The command to connect to the server

```go
package main

import (
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

// main for success connect start make command: up-ssh
// kernel.Connect It will perform the connection by type and connection properties
func main() {
	connection := &connect.Connect{
		Alias:     "test",
		Login:     "root",
		Password:  "password",
		Address:   "localhost",
		Type:      connect.TypeSSH,
		CreatedAt: "",
		UpdatedAt: "",
		SshOptions: &connect.SshOptions{
			Port:       22,
			PrivateKey: "",
		},
	}

	err := kernel.Connect(connection)
	if err != nil {
		panic(err)
	}
}
```

Usage example:


### Create

### Update

### Delete

### List


### Init Connection

descriptions fields

### Run ssh server

example make command for start

### Description variable

App dir - 

Name files -

Type Logging - 


## 🧪 Testing

You can run the command for testing after the step with local installation

Run Lint and Analyze code(phpstan/rector/phpcs):
```bash
make lint
```

Run Unit tests:
```bash
make test
```

Run mutation tests:
```bash
make test-mutation
```

Run test coverage:
```bash
make test-coverage
```

Run test coverage:
```bash
make test-coverage
```
## 🤝 Feedback

We appreciate your support and look forward to making our product even better with your help!

[@Denis Korbakov](https://github.com/deniskorbakov)

---

📝 Generated from [deniskorbakov/skeleton-php-docker](https://github.com/deniskorbakov/skeleton-php-docker)
