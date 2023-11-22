#!/usr/bin/env bash

# It's BASH script (Not SHELL), because GitLab is using this script within golang docker container.
set -euxo pipefail

go install github.com/t-yuki/gocover-cobertura@latest
go install github.com/jstemmer/go-junit-report@v0.9.1 
go test -v -race -coverpkg=./internal/... -coverprofile=test/coverage.txt ./internal/... 2>&1 | tee /dev/stderr | go-junit-report -set-exit-code > test/junit-report.xml
gocover-cobertura < test/coverage.txt > test/coverage.xml
go tool cover -func=test/coverage.txt