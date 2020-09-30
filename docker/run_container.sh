#!/bin/sh
docker stop peon || true
docker rm peon || true
docker run -it --name peon \
    --volume=$(cd .. && pwd)/test/data:/opt/peon/app/test/data \
    peon peon "$@"