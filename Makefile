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
	bash -c "diff -u <(echo -n) <(gofmt -d -s .)"
