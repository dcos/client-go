#!/usr/bin/env bash

mkdir -p dcos/iam && \
git clean -fdx dcos/iam && \
swagger generate client \
  -t dcos/iam \
  --default-scheme=https \
  --additional-initialism=DCOS \
  --template-dir=swagger/templates \
  -f swagger/iam.yaml

mkdir -p dcos/secrets && \
git clean -fdx dcos/secrets && \
swagger generate client \
  -t dcos/secrets \
  --default-scheme=https \
  --additional-initialism=DCOS \
  --template-dir=swagger/templates \
  -f swagger/secrets.yaml

git clean -fdx dcos/cosmos/client && \
docker run --rm -u $(id -u):$(id -g) -v ${PWD}:/local \
    openapitools/openapi-generator-cli:v4.0.0-beta2 generate \
    -i /local/openapi/cosmos.yaml \
    -g go \
    -o /local/dcos/cosmos/client \
    -t /local/openapi/templates \
    -DpackageName=client \
    -DwithGoCodegenComment \
    -Dmodels,apis,supportingFiles=client.go && \
docker run --rm -u $(id -u):$(id -g) -v ${PWD}:/local \
    openapitools/openapi-generator-cli:v4.0.0-beta2 generate \
    -i /local/openapi/cosmos.yaml \
    -g go \
    -o /local/dcos/cosmos/client \
    -DpackageName=client \
    -DwithGoCodegenComment \
    -DsupportingFiles=configuration.go && \
docker run --rm -u $(id -u):$(id -g) -v ${PWD}:/local \
    openapitools/openapi-generator-cli:v4.0.0-beta2 generate \
    -i /local/openapi/cosmos.yaml \
    -g go \
    -o /local/dcos/cosmos/client \
    -DpackageName=client \
    -DwithGoCodegenComment \
    -DsupportingFiles=response.go && \
gofmt -s -w dcos/cosmos/client

# go get -u -f ./...
