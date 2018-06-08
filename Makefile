GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=amoeba

all: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
run:
	$(GOBUILD) -o $(BINARY_NAME) - v
	./$(BINARY_NAME)