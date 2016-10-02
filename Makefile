DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
WEBSITE="http://komanda.io"
DESCRIPTION="Komanda IRC Client"
NAME="komanda"

BUILDVERSION=$(shell cat VERSION)

# Get the git commit
SHA=$(shell git rev-parse --short HEAD)

build: lint generate
	@echo "Building..."
	@mkdir -p bin/
	@go build \
    -ldflags "-X main.Build=${SHA}" \
    -o bin/${NAME} cmd/main.go

generate:
	@echo "Running go generate..."
	@go generate ./...

lint:
	@go vet ./...
	# @golint ./...

test:
	go list ./... | xargs -n1 go test

clean:
	@rm -rf bin/

.PHONY: build
