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

dep:
	dep ensure

dep-update:
	dep update

# TODO test
