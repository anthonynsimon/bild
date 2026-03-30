PKG = github.com/anthonynsimon/bild
VERSION ?= dev
LDFLAGS = -ldflags "-X $(PKG)/cmd.Version=$(VERSION) -extldflags \"-static\""
MAC_LDFLAGS = -ldflags "-X $(PKG)/cmd.Version=$(VERSION)"

deps:
	go get ./...

build: ensure-dist
	go build $(MAC_LDFLAGS) -o dist/ ./...

install:
	go install $(MAC_LDFLAGS)

test:
	go test ./... -timeout 60s -v

cover:
	go test ./... -race -v -timeout 15s -coverprofile=coverage.out
	go tool cover -html=coverage.out

fmt:
	go fmt ./...

bench:
	go test -benchmem -bench=. -benchtime=5s ./...

race:
	go test ./... -v -race -timeout 15s

release: release-linux release-darwin

ensure-dist: deps
	mkdir -p dist

release-linux: ensure-dist
	GOOS=linux GOARCH=amd64 go build -o dist/bild $(LDFLAGS) && cd dist && tar -czf bild_linux_amd64.tar.gz bild && rm bild
	GOOS=linux GOARCH=arm64 go build -o dist/bild $(LDFLAGS) && cd dist && tar -czf bild_linux_arm64.tar.gz bild && rm bild

release-darwin: ensure-dist
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/bild && cd dist && tar -czf bild_darwin_amd64.tar.gz bild && rm bild
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/bild && cd dist && tar -czf bild_darwin_arm64.tar.gz bild && rm bild