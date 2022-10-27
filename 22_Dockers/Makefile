dep:
	go mod tidy
	go mod download

run: dep
	go run main.go

docker-build:

docker-run: docker-build

test:
	go test ./...