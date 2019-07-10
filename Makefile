.PHONY: install vet test

install:
	go get ./...

vet:
	go vet ./...

test:
	go test ./...

check: vet test
