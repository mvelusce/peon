#!/usr/bin/env bash

current_version=$(git -c 'versionsort.suffix=-' \
    ls-remote --exit-code --refs --sort='version:refname' --tags origin '*.*.*' \
    | egrep -o '[0-9]+\.[0-9]+\.[0-9]+' \
    | tail -1)
echo "Current version: $current_version"

sed -i .bak 's/$PROG_VERSION/'"$current_version"'/g' ../cmd/peon/main.go

export CURRENT_VERSION=$current_version
