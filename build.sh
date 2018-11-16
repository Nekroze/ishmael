#!/bin/sh
set -euf
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o "$GOPATH/bin/ishmael" github.com/Nekroze/ishmael
echo "Generated static binary at $GOPATH/bin/ishmael"
