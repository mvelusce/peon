#!/usr/bin/env bash

cd cmd/peon

cp main.go main.go.bak
envsubst < main.go > main_with_version.go
mv main_with_version.go main.go

echo "Building for Linux..."
env GOOS=linux GOARCH=386 go build -v -o peon-linux

echo "Building for OSX..."
env GOOS=darwin GOARCH=386 go build -v -o peon-osx

echo "Building for Windows..."
env GOOS=windows GOARCH=386 go build -v -o peon-windows

echo "Zipping artifacts..."
zip peon.zip peon-linux peon-osx peon-windows

mv main.go.bak main.go
