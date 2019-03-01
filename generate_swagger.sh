#!/usr/bin/env bash

mkdir -p dcos/iam
swagger generate client -t dcos/iam --default-scheme=https --additional-initialism=DCOS -f swagger/iam.yaml
go get -u -f ./dcos/iam/...
