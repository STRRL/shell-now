#!/usr/bin/env bash

set -euxo pipefail

HASH=$(git rev-parse --short HEAD)
if [[ $(git status --porcelain) ]]; then
  HASH=${HASH}-dirty
fi 

echo $HASH
