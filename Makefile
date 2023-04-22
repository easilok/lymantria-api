BINARY_NAME=lymantria_api.out

build:
	go build -o ${BINARY_NAME} main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

test:
	go test ./...

dev:
	go run .

setup:
	go mod tidy

clean:
	go clean
	rm ${BINARY_NAME}
