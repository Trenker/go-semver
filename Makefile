default: test
	go build

test:
	go test

fmt:
	go fmt ./...
