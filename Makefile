default: build

build:
	go build

test:
	go test -v --race `go list ./... | grep -v vendor`
