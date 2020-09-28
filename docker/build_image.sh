#!/bin/bash

current_version=$(./get_current_version.sh)

docker build -t peon -f ../Dockerfile ../

mv ../cmd/peon/main.go.bak ../cmd/peon/main.go
