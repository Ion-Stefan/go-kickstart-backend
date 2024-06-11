build:
	@go build -o bin/go-kickstart-backend cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-kickstart-backend
