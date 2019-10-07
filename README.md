# first-api-go

[![Actions Status](https://github.com/richardbertozzo/first-api-go/workflows/build/badge.svg)](https://github.com/richardbertozzo/first-api-go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/richardbertozzo/first-api-go)](https://goreportcard.com/report/github.com/richardbertozzo/first-api-go)

Example of API in Golang

## Getting started

### Prerequisites

You're going to need:

- [**Go**](https://golang.org/doc/install) or [**Docker**](https://docs.docker.com/install/). (*Its important configure the $GOPATH, in the link has an explain about*)

### Setting up Dev
- Create `config.toml` like `config-example.toml`. Put all envs values like **PORT**.

### Building / Run

**Go**
```shell
# build and run
go run main.go

# Build and generate exec
go build
# exec the app builded
./first-api-go
```

**Docker**
```shell
make build && make run
```