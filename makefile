version=0.2.4

.PHONY: all

all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  build         - build the dist binary"
	@echo "  clean         - clean the dist build"
	@echo "  coverage      - generates test coverage report"
	@echo "  deps          - install deps"
	@echo "  test          - standard go test"
	@echo "  tools         - go gets a bunch of tools for dev"

build: clean deps
	@go vet ./...
	@golint ./...
	@go install

clean:
	@rm -rf ./bin

coverage:
	@ginkgo -cover -r -v

deps:
	@go get github.com/garyburd/go-oauth/oauth

test:
	@ginkgo -r -cover -race

tools:
	go get golang.org/x/lint/golint
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega
