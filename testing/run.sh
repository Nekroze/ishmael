#!/bin/sh
set -euf
export COMPOSE_FILE='testing/docker-compose.yml'
export COMPOSE_PROJECT_NAME='ishmael'

cleanup() {
    docker-compose down --volumes --remove-orphans
    mess="$(docker ps --filter='name=ishmael_tests_' -q)"
    [ -z "$mess" ] || echo "$mess" | xargs docker rm -f
}
trap cleanup EXIT

docker-compose build
docker-compose up --exit-code-from tests tests
