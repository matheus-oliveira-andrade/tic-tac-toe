build:
	@go build -o bin/tic-tac-toe cmd/main.go

run:
	@go run cmd/main.go

test:
	@go test ./... -v