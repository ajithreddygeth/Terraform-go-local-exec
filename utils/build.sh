#!/usr/bin/env bash
#docker run -v  $(pwd):/app/ -e CGO_ENABLED=0 -w /app/src/  golang:1.14 go build -a -o bin/utils .

go build -a -o bin/utils src/main.go
chmod +x $(pwd)/bin/utils