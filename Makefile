.PHONY: all
all: lint build

.PHONY: lint
lint: shellcheck gofmt version
	@golangci-lint run

.PHONY: shellcheck
shellcheck:
	@find . -name "*.sh" -exec shellcheck {} +

.PHONY: gofmt
gofmt:
	@./ci/gofmt.sh .

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
