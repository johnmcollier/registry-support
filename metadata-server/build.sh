#!/bin/sh

# Build the metadata container for the registry
buildfolder="$(basename "$(dirname "$0")")"
docker build -t devfile-registry-metadata:latest $buildfolder