#!/usr/bin/env bash

GO111MODULE=on

  go version
  go clean -testcache
  go build ./...
  go test -v ./...
  
