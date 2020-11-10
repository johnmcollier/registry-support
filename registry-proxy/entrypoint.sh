#!/usr/bin/env sh
set -eu

# Set the `REGISTRY_SERVICE` env var in the nginx config template
envsubst '${REGISTRY_SERVICE}' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

exec "$@"