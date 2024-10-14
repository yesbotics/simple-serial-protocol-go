SHELL:=/bin/bash

.ONESHELL:

all: setup

setup: FORCE setup-lib setup-example

setup-lib: FORCE
	go mod vendor
	go mod tidy

setup-example: FORCE
	cd examples/echo_example
	go mod vendor
	go mod tidy

update: FORCE
	go get -u ./...
	go mod vendor
	go mod tidy

test: FORCE

run-example: FORCE setup
	cd examples/echo_example
	DEBUG=1 go run .

dev: FORCE run-example

test: FORCE
	go test ./...

.PHONY: FORCE
FORCE:
