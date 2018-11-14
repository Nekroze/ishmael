#!/bin/sh
export COMPOSE_FILE='testing/docker-compose.yml'
export COMPOSE_PROJECT_NAME='ishmael'

trap 'docker-compose down' EXIT
docker-compose build
docker-compose run --rm tests
