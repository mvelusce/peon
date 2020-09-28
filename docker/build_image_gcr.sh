#!/bin/bash
export PROJECT_ID=skyita-da-daita-dev

export current_version=$(./get_current_version.sh)

echo "Building the image..."
docker build -t eu.gcr.io/${PROJECT_ID}/peon:${current_version} -f ../Dockerfile ../
echo "Image built."

echo "Loading the image on GCR..."
docker push eu.gcr.io/${PROJECT_ID}/peon:${current_version}
echo "Image uploaded."

gcloud container images list-tags eu.gcr.io/${PROJECT_ID}/peon

mv ../cmd/peon/main.go.bak ../cmd/peon/main.go
