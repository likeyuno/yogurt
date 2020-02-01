#!/usr/bin/env bash

echo "build gateway_linux_amd64"
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o build/gateway_linux_amd64 service/gateway/cmd/main.go

echo "build reverse_linux_amd64"
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o build/reverse_linux_amd64 service/reverse/cmd/main.go

echo "build static_linux_amd64"
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o build/static_linux_amd64 service/static/cmd/main.go

echo "build download_linux_amd64"
dart2native service/download/bin/download.dart -o build/download_linux_amd64
