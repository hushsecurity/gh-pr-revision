.PHONY: all
all: lint build

.PHONY: lint
lint: version
	@golangci-lint run

.PHONY: version
version:
	@git describe --always --dirty --tags | tr -d "\n" > version.txt

.PHONY: show-version
show-version: version
	@cat version.txt

.PHONY: build
build: version
	@go build

.PHONY: clean
clean:
	@rm -f gh-pr-revision version.txt
