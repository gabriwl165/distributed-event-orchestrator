test:
	go test ./...

docker-build:
	docker build -t event-orchestrator .

server:
	go run cmd/server/main.go