#!/bin/sh
#
# Runs main
#
set -e

SCRIPTS="$(dirname "$0")"

go generate ./...

$SCRIPTS/just_run.sh "$@"