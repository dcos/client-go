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
generate:
	openapi-generator generate -i openapi/dcos.yaml -g go -o dcos --skip-validate-spec -D packageName=dcos,withGoCodegenComment=true -t templates
