# GoatCLI
[![Go Report Card](https://goreportcard.com/badge/github.com/goatcms/goatcli)](https://goreportcard.com/report/github.com/goatcms/goatcli)
[![GoDoc](https://godoc.org/github.com/goatcms/goatcli?status.svg)](https://godoc.org/github.com/goatcms/goatcli)

## About
GoatCLI is set of console tools.

## Install
```
go install github.com/goatcms/goatcli
```

## Commands
* clone - clone project (and its modules) from remote repository.
* data:add  - create new build data by commmand line.
* build - build current project. It should be run in root directory (contains ".goat direcotry"). You can use "cwd" argument to change current directory.
* help - show command line help

## Arguments
* cwd - set Current Working Directory 
