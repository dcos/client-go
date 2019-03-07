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

.PHONY: generate-client
generate-client:
	openapi-generator generate -i openapi/dcos.yaml -g go -o dcos --skip-validate-spec \
		-D packageName=dcos,withGoCodegenComment=true,models,apis \
		-D supportingFiles=client.go \
		-t templates
	openapi-generator generate -i openapi/dcos.yaml -g go -o dcos --skip-validate-spec \
		-D packageName=dcos,withGoCodegenComment=true,models,apis \
		-D supportingFiles=response.go \
		-t templates
	openapi-generator generate -i openapi/dcos.yaml -g go -o dcos --skip-validate-spec \
		-D packageName=dcos,withGoCodegenComment=true,models,apis \
		-D supportingFiles=configuration.go \
		-t templates