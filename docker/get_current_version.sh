#!/usr/bin/env bash

current_version=$(git -c 'versionsort.suffix=-' \
    ls-remote --exit-code --refs --sort='version:refname' --tags origin '*.*.*' \
    | egrep -o '[0-9]+\.[0-9]+\.[0-9]+' \
    | tail -1)

sed -i .bak 's/$PROG_VERSION/'"$current_version"'/g' ../cmd/peon/main.go

echo "$current_version"
