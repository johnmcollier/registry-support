#!/bin/sh

# Build the registry proxy container for the registry
buildfolder="$(basename "$(dirname "$0")")"
docker build -t registry-proxy $buildfolder