BIN_DIR := bin

.PHONY: build build-cdd start start-analyser test cover

build:
	go build -o $(BIN_DIR)/api ./cmd/api/...

build-cdd:
	go build -o $(BIN_DIR)/metrics ./cmd/metrics/...

start: build
	./$(BIN_DIR)/api

start-analyser: build-cdd
	./$(BIN_DIR)/metrics

test:
	go test ./internal/author/... -coverprofile=coverage.out
	go tool cover -func=coverage.out

cover: test
	go tool cover -html=coverage.out
