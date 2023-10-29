GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=sumarejidump

all: build
build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v *.go

clean:
	$(GOCLEAN)
	rm -f  bin/$(BINARY_NAME)

mod-tidy:
	$(GOCMD) mod tidy

mod-download:
	$(GOCMD) mod download

mod-verify:
	$(GOCMD) mod verify

# TODO test
