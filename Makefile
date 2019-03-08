.PHONY: test
test: vet
	go test -race -cover ./...

.PHONY: vet
vet: lint
	go vet ./...

.PHONY: lint
lint: fmt
	golint -set_exit_status ./...

.PHONY: fmt
fmt:
	gofmt -w -s .

.PHONY: generate
generate: generate-client fmt

GENERATOR_CMD=docker run --rm -v ${PWD}:/client-go openapitools/openapi-generator-cli:v4.0.0-beta2 generate

.PHONY: generate-client
generate-client:
	${GENERATOR_CMD} -i /client-go/openapi/dcos.yaml -g go -o /client-go/dcos \
		--skip-validate-spec \
		-D packageName=dcos,withGoCodegenComment=true,models,apis \
		-D supportingFiles=client.go \
		-t /client-go/templates
	${GENERATOR_CMD} -i /client-go/openapi/dcos.yaml -g go -o /client-go/dcos \
		--skip-validate-spec \
		-D packageName=dcos,withGoCodegenComment=true,models,apis \
		-D supportingFiles=response.go \
		-t /client-go/templates
	${GENERATOR_CMD} -i /client-go/openapi/dcos.yaml -g go -o /client-go/dcos \
		--skip-validate-spec \
		-D packageName=dcos,withGoCodegenComment=true,models,apis \
		-D supportingFiles=configuration.go \
		-t /client-go/templates
