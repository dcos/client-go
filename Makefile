export GO111MODULE := on

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

define run_generator
docker run -u $(shell id -u):$(shell id -g) \
	-v $(CURDIR):/local -w /local \
	openapitools/openapi-generator-cli:v4.0.0-beta2 \
	generate -i openapi/dcos.yaml -g go -o dcos \
	-t templates \
	--skip-validate-spec \
	-DpackageName=dcos -DwithGoCodegenComment=true -Dmodels -Dapis \
	$(1)
endef

.PHONY: generate-client
generate-client:
	$(call run_generator,-DsupportingFiles=client.go)
	$(call run_generator,-DsupportingFiles=response.go)
	$(call run_generator,-DsupportingFiles=configuration.go)
	$(call run_generator,-DsupportingFiles=README.md)

