
VERSION ?= "0.0.1"
GITCOMMIT ?= ""
BuildAt ?= ""

BIN_DIR = ./bin
ARCHIVE_NAME := pipeSource$(VERSION).tar.gz


all: $(BIN_DIR) build

$(BIN_DIR):
	@echo "Creating directory: $(BIN_DIR)"
	mkdir -p  $(BIN_DIR)

build:
	@echo "build project"
	go build -ldflags "-X goRelay/pkg.Version=$(VERSION) -X goRelay/pkg.BuildAt=$(BuildAt) -X goRelay/pkg.GitCommit=$(GITCOMMIT)" -o $(BIN_DIR)/goRelay 
	tar zcvf $(ARCHIVE_NAME) $(BIN_DIR)

clean:
	@echo "cleaning up..."
	rm -f $(BIN_DIR)/goRelay

