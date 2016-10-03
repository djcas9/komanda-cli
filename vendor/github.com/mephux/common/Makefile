DEPS=$(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)

all: deps test

deps:
	@go get -d -v ./...
	@echo $(DEPS) | xargs -n1 go get -d

test: deps
	@go list ./... | xargs -n1 go test

cov:
	@gocov test ./... | gocov-html > /tmp/coverage.html
	@open /tmp/coverage.html

lint:
	@golint .
	@go tool vet .

clean:

.PHONY: deps test
