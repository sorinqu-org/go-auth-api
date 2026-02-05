.PHONY: build test run down

build:
	@go build -o bin/app cmd/main.go

test:
	@go test -v ./...

run: build
	@docker compose up -d --build
	@goose up
	@./bin/app

down:
	@docker compose down
	@goose down
