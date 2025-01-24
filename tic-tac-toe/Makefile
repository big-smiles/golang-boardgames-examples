.DEFAULT_GOAL= build
.PHONY= fmt vet build
BINARY_NAME=tictactoe.exe

fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	 go build -o bin/${BINARY_NAME} cmd/main.go
clean: rm -f bin/myapp