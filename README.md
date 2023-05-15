# practicalgo
A template for Go projects. Includes datastore, Go routines, web server, CLI commands, unit tests, and GitHub Actions with goreleaser.

### Overview

This is a template for Go projects.
I use this to bootstrap new projects, and to demonstrate Go concepts.

It is intended to have all the basic components required for a Go project.
In a layout that is easy to understand, and easy to extend.

Always a work in progress...

#### Getting Started

To run as a webserver locally (requires Redis):

```
$ go run -tags dev ./...                                                      
2023/05/15 10:01:25 main.go:23: Version dev, built at unknown, commit none 
2023/05/15 10:01:25 main.go:33: Datastore connecting, host: localhost:6379, username: 
2023/05/15 10:01:25 main.go:40: Datastore connected.
2023/05/15 10:01:25 main.go:52: webserver.Start(): listening on port 8080
```

The web server is now running on port 8080:

```
$ curl http://localhost:8080/listusers/
No users found.
```

To run as a CLI tool:

```
 $ go run -tags dev ./... -action listusers -debug true
2023/05/15 11:28:30 main.go:25: Version dev, built at unknown, commit none 
2023/05/15 11:28:30 main.go:35: Datastore connecting, host: localhost:6379, username: 
2023/05/15 11:28:30 main.go:42: Datastore connected.
2023/05/15 11:28:30 listusers.go:38: CLI_GetUsers(): Users: []
$
```

#### Dynamic Config

The *config* module enables configuration to be set via environment variables.
This makes it easy to deploy to different environments, without changing code.

Code usage:
```
log.Println("webserver.Start(): listening on port", cfg.ValueAsStr("API_PORT"))
```

#### Tasks

The *tasks* module enables actions to run via cli or Go routines in a loop.

Lightweight entry points using the same underlying methods.
The benefit is that a task can be developed and tested as a cli command, and then called from a Go routine.

Tasks are implemented using types, with interfaces that isolate datastore methods.
This makes mocking simpler, and cleaner.