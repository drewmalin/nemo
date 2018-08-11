APP=nemo
VERSION=1.0.0

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt

DEPCMD=dep

PROJECT_ROOT=github.com/$(APP)

build: clean fmt build_osx test

build_osx:
		env GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/osx/$(APP) $(PROJECT_ROOT)/cmd

clean:
		$(GOCLEAN)
		rm -rf bin/

deps:
		$(DEPCMD) ensure

fmt:
		$(GOFMT) ./...

test:
		$(GOTEST) -v ./...
