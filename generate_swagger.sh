#!/usr/bin/env bash

mkdir -p dcos/iam && \
swagger generate client \
  -t dcos/iam \
  --default-scheme=https \
  --additional-initialism=DCOS \
  --template-dir=swagger/templates \
  -f swagger/iam.yaml

mkdir -p dcos/secrets && \
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/openapi/secrets.yaml \
  -g go \
  -o /local/dcos/secrets \
  -c /local/openapi/secrets_config.json

mkdir -p dcos/cosmos && \
swagger generate client \
  -t dcos/cosmos \
  --default-scheme=https \
  --additional-initialism=DCOS \
  --template-dir=swagger/templates \
  -f swagger/cosmos.yaml

# go get -u -f ./...
