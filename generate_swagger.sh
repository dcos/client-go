#!/usr/bin/env bash

mkdir -p dcos/iam && \
swagger generate client \
  -t dcos/iam \
  --default-scheme=https \
  --additional-initialism=DCOS \
  --template-dir=swagger/templates \
  -f swagger/iam.yaml

mkdir -p dcos/secrets && \
swagger generate client \
  -t dcos/secrets \
  --default-scheme=https \
  --additional-initialism=DCOS \
  --template-dir=swagger/templates \
  -f swagger/secrets.yaml

mkdir -p dcos/cosmos && \
swagger generate client \
  -t dcos/cosmos \
  --default-scheme=https \
  --additional-initialism=DCOS \
  --template-dir=swagger/templates \
  -f swagger/cosmos.yaml

# go get -u -f ./...
