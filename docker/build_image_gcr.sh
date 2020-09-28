#!/bin/bash
export PROJECT_ID=skyita-da-daita-dev

./get_current_version.sh

echo "Building the image..."
docker build -t eu.gcr.io/${PROJECT_ID}/peon:${CURRENT_VERSION} -f ../Dockerfile ../
echo "Image built."

echo "Loading the image on GCR..."
docker push eu.gcr.io/${PROJECT_ID}/peon:${CURRENT_VERSION}
echo "Image uploaded."

gcloud container images list-tags eu.gcr.io/${PROJECT_ID}/peon
