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
- **Works Everywhere:** first-class support in every terminal on macOS, Linux, Windows (PowerShell and WSL), FreeBSD,
  OpenBSD, and NetBSD

## Install

install this package in your repository

```bash
go get github.com/misha-ssh/kernel
```

## Examples

You will be provided with a list of commands that you can use in your projects

The code with the commands will be on the way - [link](./examples/command)

### Connect

The command to connect to the remote server

```go
package main

import (
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

func main() {
	connection := &connect.Connect{...}

	err := kernel.Connect(connection)
	if err != nil {
		panic(err)
	}
}
```

### Create

The command to create a connection

this command saves the connection to a file and goes through the dependency initialization cycle

```go
package main

import (
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

func main() {
	connection := &connect.Connect{...}

	err := kernel.Create(connection)
	if err != nil {
		panic(err)
	}
}
```

### Update

The command to update the connection

This command also updates the connection data if you need to resave the private key

```go
package main

import (
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

func main() {
	connection := &connect.Connect{...}

	err := kernel.Update(connection, "test")
	if err != nil {
		panic(err)
	}
}
```

### Delete

The command to delete the connection

This command removes the connection from the file and also deletes the private key if it has been saved

```go
package main

import (
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

func main() {
	connection := &connect.Connect{...}

	err := kernel.Delete(connection)
	if err != nil {
		panic(err)
	}
}
```

### List

The command to get a list of connections

This command will list the connections from the previously created connections

```go
package main

import (
	"fmt"

	"github.com/misha-ssh/kernel/pkg/kernel"
)

func main() {
	connections, err := kernel.List()
	if err != nil {
		panic(err)
	}

	fmt.Println(connections)
}
```

### Struct Connection

This structure describes our connection

We pass this structure to the commands:
- Connect
- Create
- Update
- Delete

Description of fields:
* ``Alias`` - unique name (we use it when selecting an exception from the list and create unique connections to identify them)
* ``Login`` - the user's login on the remote device
* ``Password`` - if you have a password connection, then fill in this field, if with a key, then leave an empty line.
* ``Address`` - the address of the remote device
* ``Type`` - there is only one connection type so far - ``connect.TypeSSH``
* ``CreatedAt`` - the creation time is filled in manually
* ``UpdatedAt`` - the update time is filled in manually
* ``SshOptions`` - this is a structure with additional fields for creating software - ``connect.TypeSSH``
* ``Port`` - the port is filled in manually
* ``PrivateKey`` - the path along with the name of the private key

```go
package main

import (
	"fmt"
	"time"

	"github.com/misha-ssh/kernel/pkg/connect"
)

func main() {
	connection := &connect.Connect{
		Alias:     "test",
		Login:     "root",
		Password:  "password",
		Address:   "localhost",
		Type:      connect.TypeSSH,
		CreatedAt: time.Now().Format("2006.01.02 15:04:05"),
		UpdatedAt: time.Now().Format("2006.01.02 15:04:05"),
		SshOptions: &connect.SshOptions{
			Port:       22,
			PrivateKey: "/path/to/private/key",
		},
	}

	fmt.Println(connection)
}
```

### Run ssh server

for local testing, you can raise your ssh servers - there are three types of them.

1) password connection

to run, write the command:

```bash
make up-ssh
```

to install and remove the server:

```bash
make down-ssh
```

Server accesses:
* ``login`` - root
* ``address`` - localhost
* ``password`` - password
* ``port`` - 22

2) connect with a private key

to run, write the command:

to install and remove the server:

3) connecting via a non-standard port

to run, write the command:

to install and remove the server:

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
