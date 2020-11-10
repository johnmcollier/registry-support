#!/bin/sh
IMAGE_TAG=$1
docker tag registry-proxy:latest $IMAGE_TAG
docker push $1