#!/usr/bin/env bash

cd cmd/peon

echo "Building for Linux..."
env GOOS=linux GOARCH=arm go build -v -o peon-linux

echo "Building for OSX..."
env GOOS=darwin GOARCH=386 go build -v -o peon-osx

echo "Building for Windows..."
env GOOS=windows go build -v -o peon-windows

echo "Zipping artifacts..."
zip peon.zip peon-linux peon-osx peon-windows