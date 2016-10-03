WEBSITE="http://komanda.io"
DESCRIPTION="Komanda IRC Client"
NAME="komanda"

BUILDVERSION=$(shell cat package/VERSION)
GO_VERSION=$(shell go version)

UCHITD_VERSION=$(shell cat .uchitd-version)
UCHIT_AGENT_VERSION=$(shell cat .uchit-agent-version)

# Get the git commit
SHA=$(shell git rev-parse --short HEAD)
BUILD_COUNT=$(shell git rev-list --count HEAD)

BUILD_TAG="${BUILD_COUNT}.${SHA}"

PROTOC=$(shell which protoc)

CCOS=windows darwin linux
CCARCH=amd64
CCOUTPUT="package/output/{{.OS}}-{{.Arch}}/"

BIN_FILES=$(NAME)d $(NAME)-agent

build: banner lint
	@echo "Building..."
	@mkdir -p bin/
	@go build \
    -ldflags "-X main.Build=${SHA}" \
    -o bin/${NAME} .

banner: deps
	@echo "Project:    $(NAME)"
	@echo "Go Version: ${GO_VERSION}"
	@echo "Go Path:    ${GOPATH}"
	@echo

generate:
	@echo "Running go generate..."
	@go generate ./...

lint:
	# @go vet  $$(go list ./... | grep -v /vendor/)
	# @for pkg in $$(go list ./... |grep -v /vendor/ |grep -v /services/) ; do \
		# golint -min_confidence=1 $$pkg ; \
		# done

deps:
	@go get -u github.com/golang/lint/golint

cc:
	@echo "[-] Cross Compiling $(NAME)"
	@mkdir -p dist/
	@echo "[-] $(NAME) (freebsd,linux,darwin,arm7)"
	@GOOS=freebsd go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-freebsd
	@GOOS=linux go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-linux-amd64
	@GOOS=linux GOARCH=386 go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-linux-386
	@GOOS=darwin go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-darwin-amd64
	@GOOS=linux GOARCH=arm GOARM=5 go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-linux-arm-5
	@GOOS=linux GOARCH=arm GOARM=6 go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-linux-arm-6
	@GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-linux-arm-7

dist: clean banner lint cc
	@for bin in dist/*; do \
		tar -cvf $$bin.tar.xz -C dist $$(basename $$bin); \
		rm -rf $$bin; \
	done

test:
	go list ./... | xargs -n1 go test

clean:
	@rm -rf bin/
	@rm -rf .$(NAME)-version
	@rm -rf doc/
	@rm -rf package/
	@rm -rf dist/
	@rm -rf bin/
	@rm -rf tmp/


.PHONY: build
