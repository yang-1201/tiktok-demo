#!/bin/bash

set -e
cd "$(dirname "$0")"

go env -w GOPROXY=https://goproxy.cn,direct

go test -v ./...
#go test -v ./test_package/likes_docker_test.go