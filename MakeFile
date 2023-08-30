lint:
	golangci-lint run -c .golangci.yml
test:
	go test ./... -coverprofile=coverage.out