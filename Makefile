BIN_DIR = bin
build: $(BIN_DIR) $(BIN_DIR)/gnmap

$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

GOSOURCES = $(wildcard *.go cmd/*.go) Makefile
$(BIN_DIR)/gnmap: $(GOSOURCES)
	@go build -o $(BIN_DIR)/gnmap

test:
	go test ./...

GITIGNORE ?= go
gitignore:
	curl -Ls "http://www.gitignore.io/api/$(GITIGNORE)" | tee .gitignore
	@if [ -f .gitignore.custom ]; then \
		cat .gitignore.custom >> .gitignore; \
	fi

clean:
	@rm -rf $(BIN_DIR)/

.PHONY: clean
