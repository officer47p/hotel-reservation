build:
	@ go build -o bin/api

run: build
	@./bin/api --listenAddr :3001

test:
	@go test -v ./...