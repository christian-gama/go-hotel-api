#!/bin/bash
CACHE_DIR="$PWD/.cache/linter"
DEFAULT_CONFIG="$HOME/.golangci.yml"
mkdir -p "$CACHE_DIR"

docker run \
  --rm -t \
  --user "$(id -u):$(id -g)" \
  -v "$CACHE_DIR:/.cache" \
  -v "$(go env GOPATH)/pkg:/go/pkg" \
  -v "$PWD:/app" \
  --workdir /app \
  golangci/golangci-lint:v1.47.3 \
  golangci-lint run --config .golangci.yml "$@"