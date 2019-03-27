
GOOS=$(shell go env | grep GOOS | awk -F "=" '{print $$NF}' | awk -F "\"" '{print $$2}')
GOARCH=$(shell go env | grep GOARCH | awk -F "=" '{print $$NF}' | awk -F "\"" '{print $$2}')
SERVICE=$(shell pwd | awk -F "/" '{print $$NF}')
BINARY=$(SERVICE)-$(GOOS)-$(GOARCH)
TAG?=latest
EXENAME=gomm

GOROOT?=/usr/local/

default: build
.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH)  $(GOROOT)/bin/go build -o ./bin/$(EXENAME) ./main.go
linux:
	GOOS=linux GOARCH=amd64   $(GOROOT)/bin/go build -o ./bin/$(EXENAME)  ./main.go
darwin:
	GOOS=darwin GOARCH=amd64   $(GOROOT)/bin/go build -o ./bin/$(EXENAME)  ./main.go
run: build
	./bin/$(EXENAME)

