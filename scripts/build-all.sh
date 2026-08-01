#!/usr/bin/env bash
set -e

# Builds apx for mac/linux/windows (amd64 + arm64) into dist folders
BIN_NAME=apx
mkdir -p dist/Mac dist/Windows-10-11 dist/Linux

echo "Building ${BIN_NAME} darwin amd64..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/Mac/${BIN_NAME}-darwin-amd64 ./cmd

echo "Building darwin arm64..."
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o dist/Mac/${BIN_NAME}-darwin-arm64 ./cmd

echo "Building linux amd64..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/Linux/${BIN_NAME}-linux-amd64 ./cmd

echo "Building linux arm64..."
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o dist/Linux/${BIN_NAME}-linux-arm64 ./cmd

echo "Building windows amd64..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/Windows-10-11/${BIN_NAME}-windows-amd64.exe ./cmd

echo "Building windows arm64..."
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o dist/Windows-10-11/${BIN_NAME}-windows-arm64.exe ./cmd

echo "Done. Dist folders:"
ls -R dist
