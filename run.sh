#!/usr/bin/env bash
set -e

go test ./... -coverprofile=c.out && go tool cover -html=c.out
go run cmd/main.go