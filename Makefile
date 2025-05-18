tidy:
	go mod tidy

lint:
	golangci-lint run -c .golangci.yml ./...

lint-fix:
	golangci-lint run --fix -c .golangci.yml ./...

test:
	go test -race -v ./...

coverage:
	go test -v -coverprofile=cover.out -covermode=atomic ./...
	go tool cover -html=cover.out -o cover.html

build:
	go build -v ./...


.PHONY: tidy lint lint-fix test coverage build
