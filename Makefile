DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
WEBSITE="http://komanda.io"
DESCRIPTION="Komanda IRC Client"
NAME="komanda"

BUILDVERSION=$(shell cat VERSION)

# Get the git commit
SHA=$(git rev-parse --short HEAD)

build: lint generate
	@echo "Building..."
	@mkdir -p bin/
	@godep go build \
    -ldflags "-X main.Build=${SHA}" \
    -o bin/${NAME} cmd/main.go

generate:
	@echo "Running go generate..."
	@godep go generate ./...

lint:
	@godep go vet ./...
	# @golint ./...

updatedeps:
	@godep update ...

test: deps
	go list ./... | xargs -n1 go test

deps:
	@go get -u github.com/tools/godep

clean:
	@rm -rf bin/

.PHONY: build
