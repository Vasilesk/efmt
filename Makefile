tidy:
	go mod tidy

lint:
	golangci-lint run -c .golangci.yml ./...

lint-fix:
	golangci-lint run --fix -c .golangci.yml ./...

test:
	go test -count=1 ./... -covermode=atomic -v -race

.PHONY: lint test tidy
