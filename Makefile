.PHONY: setup build test run down

setup:
	@chmod +x env.sh

build:
	@go build -o bin/app cmd/main.go

test:
	@go test -v ./...

run: build
	@docker compose up -d --build
	@goose up
	@./env.sh

down:
	@docker compose down
	@goose down-to 0
