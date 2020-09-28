#!/usr/bin/env bash

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        ./peon-linux "$@"
elif [[ "$OSTYPE" == "darwin"* ]]; then
        ./peon-osx "$@"
elif [[ "$OSTYPE" == "cygwin" ]]; then
        ./peon-windows "$@"
elif [[ "$OSTYPE" == "msys" ]]; then
        ./peon-windows "$@"
elif [[ "$OSTYPE" == "win32" ]]; then
        ./peon-windows "$@"
elif [[ "$OSTYPE" == "freebsd"* ]]; then
        ./peon-linux "$@"
else
        ./peon "$@"
fi
