.PHONY: help install run update ensure

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install:	## installs the project to $GOPATH/bin
	go fmt ./...
	go build -i -o "$$GOPATH/bin/golang-rest-test" cmd/golang-rest-test/main.go

run:	## runs the project
	go fmt ./...
	go run cmd/golang-rest-test/main.go

update: ## Update the locked versions of all dependencies
	dep ensure -update

ensure:	## Install the project's dependencies
	dep ensure