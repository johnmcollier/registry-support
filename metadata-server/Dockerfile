# Build the registry metadata server
FROM golang:alpine3.11 AS build
WORKDIR /tools
COPY . .
RUN go build -o metadata-server main.go

FROM alpine:3.12.0

# Install and configure dependencies
RUN apk add --no-cache git curl

COPY entrypoint.sh /scripts/entrypoint.sh

# Git clone the devfiles
# ToDo: Switch to github.com/devfile/catalog/
RUN git clone https://github.com/devfile/registry /registry
COPY --from=build /tools/metadata-server /registry/metadata-server
RUN chgrp -R 0 /registry && \
    chmod -R g=u /registry

# Load the devfiles into the registry and serve the index.json
ENTRYPOINT ["/scripts/entrypoint.sh"]
CMD ["/registry/metadata-server"]