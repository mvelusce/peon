#!/usr/bin/env bash

cd cmd/peon
rm peon
go build

./peon test
rm peon
