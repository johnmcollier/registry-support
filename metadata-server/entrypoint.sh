#!/bin/sh

## Simple proof of concept bootstrap script to load devfiles into an oci registry
DEVFILES=/registry/stacks

# Generate the index.json from the devfiles
cd /registry
./index-generator $DEVFILES /usr/local/apache2/htdocs/devfiles/index.json

# Push the devfiles to the registry
cd $DEVFILES

# Wait for the registry to start
until $(curl --output /dev/null --silent --head --fail http://localhost:5000); do
    printf 'Waiting for the registry at localhost:5000 to start\n'
    sleep 0.5
done

# Launch the server hosting the index.json
exec "${@}"