#!/usr/bin/env bash

cd cmd/peon
rm peon
go build

./peon -r ../../test/data/project -p python3 test
./peon clean
rm .peon-config.json
rm peon
