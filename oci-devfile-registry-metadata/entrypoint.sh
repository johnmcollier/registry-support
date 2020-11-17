#!/bin/sh

## Simple proof of concept bootstrap script to load devfiles into an oci registry
DEVFILES=/stacks

# Wait for the registry to start
until $(curl --output /dev/null --silent --head --fail http://localhost:5000); do
    printf 'Waiting for the registry at localhost:5000 to start\n'
    sleep 0.5
done

# Push the devfiles to the registry
cd $DEVFILES
for devfileDir in "$DEVFILES"/*
do
  devfile="$devfileDir/devfile.yaml"
  stackName=`basename $devfileDir`
  # Push the devfile to the registry
  
  # ToDo:
  # 1) Discover the service name for the registry
  # 2) Getting the stack name (need to be reading the meta.yaml)
  # 3) Getting the proper stack version
  # 4) Not pushing over plain http
  # 5) Do in Golang
  echo "Pushing $stackName to $REGISTRY_HOST"
  cd $stackName
  oras push localhost:5000/devfile-catalog/$stackName:latest --manifest-config /dev/null:application/vnd.devfileio.devfile.config.v2+json ./devfile.yaml:application/vnd.devfileio.devfile.layer.v1 --plain-http
  cd $DEVFILES
done

# Copy the index.json over to /www/data
cp /index.json /www/data/

# Start the nginx server
exec "$@"