#!/usr/bin/env bash

set -exo pipefail

DIR=$(dirname $0)

COMMIT_HASH=$(bash "${DIR}"/commit-hash.sh)

cd ${DIR}/../ && \
    DOCKER_BUILDKIT=1 DOCKER_DEFAULT_PLATFORM=linux/amd64 docker build -t ghcr.io/strrl/shell-now:"${COMMIT_HASH}" \
    -f ./Dockerfile ./

docker tag ghcr.io/strrl/shell-now:"${COMMIT_HASH}" ghcr.io/strrl/shell-now:latest

if [ ! -z ${IMAGE_TAG} ]; then
    docker tag ghcr.io/strrl/shell-now:"${COMMIT_HASH}" ghcr.io/strrl/shell-now:${IMAGE_TAG}
fi
