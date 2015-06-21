GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test

all: build

build:
	$(GOBUILD) -v -o bin/wasteband

install:
	$(GOINSTALL)

clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/wasteband
	rm -rf bin/wasteband

test:
	$(GOTEST) -v -cover

.PHONY: all clean