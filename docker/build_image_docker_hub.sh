#!/bin/bash

export current_version=$(./get_current_version.sh)

echo "Building the image..."
docker build -t skyveluscekm/peon:${current_version} -t skyveluscekm/peon:latest -f ../Dockerfile ../
echo "Image built."

echo "Loading the image on Docker Hub..."
#docker push skyveluscekm/peon:${current_version}
#docker push skyveluscekm/peon:latest
echo "Image uploaded."

mv ../cmd/peon/main.go.bak ../cmd/peon/main.go || true
