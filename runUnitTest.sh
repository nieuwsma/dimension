#!/usr/bin/env bash
#
# Copyright
#

set -x

# Configure docker compose
export COMPOSE_PROJECT_NAME=$RANDOM
export COMPOSE_FILE=docker-compose.test.unit.yaml

echo "COMPOSE_PROJECT_NAME: ${COMPOSE_PROJECT_NAME}"
echo "COMPOSE_FILE: $COMPOSE_FILE"


function cleanup() {
  docker-compose down
  if ! [[ $? -eq 0 ]]; then
    echo "Failed to decompose environment!"
    exit 1
  fi
  exit $1
}

echo "Starting containers..."
docker-compose build
docker-compose up -d dummy #we use dummy to make sure all our dependencies are up
docker-compose ps # To improve debuggability display the current state of the containers.
docker-compose up --exit-code-from unit-tests unit-tests

test_result=$?

# Clean up
echo "Cleaning up containers..."
if [[ $test_result -ne 0 ]]; then
  echo "Unit tests FAILED!"
  cleanup 1
fi

echo "Unit tests PASSED!"
cleanup 0
