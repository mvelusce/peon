#!/usr/bin/env bash

cd cmd/peon

cp main.go main.go.bak
envsubst < main.go > main_with_version.go
mv main_with_version.go main.go

echo "Building for Linux..."
env GOOS=linux GOARCH=386 go build -i -v -o ../../bin/peon-linux

echo "Building for OSX..."
env GOOS=darwin GOARCH=386 go build -i -v -o ../../bin/peon-osx

echo "Building for Windows..."
env GOOS=windows GOARCH=386 go build -i -v -o ../../bin/peon-windows

mv main.go.bak main.go

echo "Zipping artifacts..."
cd ../../bin
zip peon.zip peon-linux peon-osx peon-windows
