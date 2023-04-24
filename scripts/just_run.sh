#!/bin/sh
#
# Runs main
#
set -e

go mod tidy
go run main.go "$@"