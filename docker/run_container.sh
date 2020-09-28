#!/bin/sh
docker stop peon || true
docker rm peon || true
docker run -it --name peon peon "$@"