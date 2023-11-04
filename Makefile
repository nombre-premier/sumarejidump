GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=sumarejidump

all: build

build: vendor
	$(GOBUILD) -mod=vendor -o bin/$(BINARY_NAME) -v *.go

clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)

vendor:
	$(GOCMD) mod vendor

# TODO test
