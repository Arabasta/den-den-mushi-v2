GO ?= go
CMD_DIR := cmd
BIN_DIR := bin

# space separated list
BINARIES := main

.PHONY: all build

all: build

build:
	go mod tidy
	mkdir -p $(BIN_DIR)
	@for b in $(BINARIES); do \
		echo "Building $$b..."; \
		cd $(CMD_DIR) && $(GO) build -buildvcs=false -o ../$(BIN_DIR)/$$b; \
		cd - > /dev/null; \
	done
