#!/bin/sh

cd "$(dirname "$0")"

go build -o ./build/tam4-server -ldflags="-s -w" main.go
