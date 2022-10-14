# formapp
This repository provides a base project to learn a simple HTML form application.

The project can be built with either Docker or Go 1.17.

It depends on [Gin Web Framework](https://pkg.go.dev/github.com/gin-gonic/gin) for ease of implementation.
Gin will be donloaded automatically.

You can run the server with the following commands;
```sh
$ docker-compose up -d
$ docker-compose exec app go run main.go
```
Or, you can directly execute `go run main.go` if you have Go development tools.

The execution command downloads Gin framework from the Internet, thus your PC is required to be online and the server may require a couple of minutes until start working.
