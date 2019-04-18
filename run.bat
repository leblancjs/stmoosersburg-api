@echo off

rem Set environment variables
rem (see the "Configuration" section in the README.md for more information)
call .env

rem Enable Go Modules (in case the repo was cloned in the $GOPATH/src directory)
set GO111MODULE=on

rem Run!
go run main.go