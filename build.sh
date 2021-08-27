#!/bin/sh
mkdir -p target
echo "Building elvish..."
GOBIN=$(pwd)/target CGO_ENABLED=1 go install -trimpath -ldflags="-s -w" -buildmode=pie src.elv.sh/cmd/elvish@$(go list -f '{{.Version}}' -m src.elv.sh)
echo "Building crypto plugin..."
go build -buildmode=plugin -trimpath -ldflags="-s -w" -o target/crypto.so plugins/crypto/crypto.go
echo "Building fun plugin..."
go build -buildmode=plugin -trimpath -ldflags="-s -w" -o target/fun.so plugins/fun/fun.go
echo "Building io plugin..."
go build -buildmode=plugin -trimpath -ldflags="-s -w" -o target/io.so plugins/io/io.go
echo "Building memo plugin..."
go build -buildmode=plugin -trimpath -ldflags="-s -w" -o target/memo.so plugins/memo/memo.go
echo "Building net plugin..."
go build -buildmode=plugin -trimpath -ldflags="-s -w" -o target/net.so plugins/net/net.go
echo "Building proc plugin..."
go build -buildmode=plugin -trimpath -ldflags="-s -w" -o target/proc.so plugins/proc/proc.go
