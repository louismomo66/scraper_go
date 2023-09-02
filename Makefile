hello:
	@echo "Hello, world!"
test:
	@go test ./... -coverprofile=coverage.out
lint:
	@golangci-lint run -c .golangci.yml