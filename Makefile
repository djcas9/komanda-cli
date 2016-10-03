WEBSITE="http://komanda.io"
DESCRIPTION="Komanda IRC Client"
NAME="komanda"

BUILDVERSION=$(shell cat VERSION)
GO_VERSION=$(shell go version)

# Get the git commit
SHA=$(shell git rev-parse --short HEAD)
BUILD_COUNT=$(shell git rev-list --count HEAD)

BUILD_TAG="${BUILD_COUNT}.${SHA}"

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
	@echo "[*] Done building $(NAME)..."

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

setup:
	@mkdir -p package/root/usr/bin/
	@mkdir -p dist/
	@cp -R ./bin/$(NAME) package/root/usr/bin/$(NAME)
	@./bin/$(NAME) --version > VERSION

cc:
	@echo "[-] Cross Compiling $(NAME)"
	@mkdir -p dist/
	@echo "[-] $(NAME) freebsd"
	@GOOS=freebsd go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-freebsd
	@echo "[-] $(NAME) linux"
	@GOOS=linux go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-linux-amd64
	@GOOS=linux GOARCH=386 go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-linux-386
	@echo "[-] $(NAME) darwin"
	@GOOS=darwin go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-darwin-amd64
	@echo "[-] $(NAME) arm"
	@GOOS=linux GOARCH=arm GOARM=5 go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-linux-arm-5
	@GOOS=linux GOARCH=arm GOARM=6 go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-linux-arm-6
	@GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-X main.build=${BUILD_TAG}" -o dist/$(NAME)-linux-arm-7

dist: clean banner lint build setup cc tar rpm64 deb64
	@echo "[*] Dist Build Done..."

tar:
	@for bin in dist/*; do \
		tar -cvf $$bin.tar.xz -C dist $$(basename $$bin); \
		rm -rf $$bin; \
	done

test:
	go list ./... | xargs -n1 go test

rpm64:
	fpm -s dir -t rpm -n $(NAME) -v $(BUILDVERSION) -p dist/$(NAME)-amd64.rpm \
		--rpm-compression xz --rpm-os linux \
		--force \
		--url $(WEBSITE) \
		--description $(DESCRIPTION) \
		-m "$(NAME) <dev@$(NAME).io>" \
		--vendor "$(NAME)" -a amd64 \
		--exclude */**.gitkeep \
		package/root/=/

deb64:
	fpm -s dir -t deb -n $(NAME) -v $(BUILDVERSION) -p dist/$(NAME)-amd64.deb \
		--force \
		--deb-compression xz \
		--url $(WEBSITE) \
		--description $(DESCRIPTION) \
		-m "$(NAME) <dev@$(NAME).io>" \
		--vendor "$(NAME)" -a amd64 \
		--exclude */**.gitkeep \
		package/root/=/

clean:
	@rm -rf bin/
	@rm -rf .$(NAME)-version
	@rm -rf doc/
	@rm -rf package/
	@rm -rf dist/
	@rm -rf bin/
	@rm -rf tmp/


.PHONY: build
