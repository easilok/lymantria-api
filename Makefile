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

deploy:
       docker-compose build lymantria_api
       docker-compose up -d

clean-bd: 
       docker-compose down
	   docker volume rm lymantria-api_lymantria-db-data

clean:
	go clean
	rm ${BINARY_NAME}
