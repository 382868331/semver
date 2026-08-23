#!/bin/bash
set -euo pipefail
IMAGE_NAME=${1:-goletalab-task}
docker build --platform linux/amd64 -f goletalab.Dockerfile -t "$IMAGE_NAME" .
docker run --rm "$IMAGE_NAME" go test ./...
