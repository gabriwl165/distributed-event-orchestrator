install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.1.6
	pip install pre-commit

format:
	golangci-lint fmt

lint:
	golangci-lint run

test:
	go test ./...

docker-build:
	docker build -t event-orchestrator .

server:
	go run cmd/server/main.go
