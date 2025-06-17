#show readme
#go mod tidy
#go mod verify
#go test
#go run

all: tidy test-cover test run


tidy:
	go mod tidy
	go mod verify
test-cover:
	go test -cover ./...
test:
	go test ./... -v
run:
	go run cmd/main.go
