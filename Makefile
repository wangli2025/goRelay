
VERSION ?= "0.0.1"
GITCOMMIT ?= ""

BIN_DIR = ./bin
ARCHIVE_NAME := pipeSource$(VERSION).tar.gz


all: $(BIN_DIR) build

$(BIN_DIR):
	@echo "Creating directory: $(BIN_DIR)"
	mkdir -p  $(BIN_DIR)

build:
	@echo "build project"
	go build -ldflags "-X goRelay/pkg.Version=$(VERSION) -X goRelay/pkg.GitCommit=$(GITCOMMIT)" -o $(BIN_DIR)/pipeServer pipeServer/*.go
	go build -ldflags "-X goRelay/pkg.Version=$(VERSION) -X goRelay/pkg.GitCommit=$(GITCOMMIT)" -o $(BIN_DIR)/pipeClient pipeClient/*.go
	go build -ldflags "-X goRelay/pkg.Version=$(VERSION) -X goRelay/pkg.GitCommit=$(GITCOMMIT)" -o $(BIN_DIR)/relayServer relayServer/*.go
	go build -ldflags "-X goRelay/pkg.Version=$(VERSION) -X goRelay/pkg.GitCommit=$(GITCOMMIT)" -o $(BIN_DIR)/relayClient relayClient/*.go
	tar zcvf $(ARCHIVE_NAME) $(BIN_DIR)

clean:
	@echo "cleaning up..."
	rm -f $(BIN_DIR)/pipeServer
	rm -f $(BIN_DIR)/pipeClient
	rm -f $(BIN_DIR)/relayServer
	rm -f $(BIN_DIR)/relayClient

