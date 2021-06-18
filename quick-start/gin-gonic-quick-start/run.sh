#!/bin/bash

set -xeuo  pipefail

go mod tidy
go build ./
./gin-example


