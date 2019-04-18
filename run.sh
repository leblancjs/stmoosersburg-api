#!/bin/bash

# Set environment variables
# (see the "Configuration" section in the README.md for more information)
set -o allexport
source .env
set -o allexport

# Enable Go Modules (in case the repo was cloned in the $GOPATH/src directory)
GO111MODULE=on

# Run!
go run main.go