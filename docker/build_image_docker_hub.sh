#!/bin/bash

export current_version=$(./get_current_version.sh)

echo "Building the image..."
docker build -t mvelusce/peon:${current_version} -t mvelusce/peon:latest -f ../Dockerfile ../
echo "Image built."

echo "Loading the image on Docker Hub..."
docker push mvelusce/peon:${current_version}
docker push mvelusce/peon:latest
echo "Image uploaded."

mv ../cmd/peon/main.go.bak ../cmd/peon/main.go || true
