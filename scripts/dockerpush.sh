#!/bin/bash

set -e
set -x

export DOCKER_BUILDKIT=1

docker-compose --profile main build
docker-compose --profile main push
