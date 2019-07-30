PKG = github.com/anthonynsimon/bild
BENCHMARK = BenchmarkRAWInput
VERSION ?= dev
LDFLAGS = -ldflags "-X $(PKG)/cmd.Version=$(VERSION) -extldflags \"-static\""
MAC_LDFLAGS = -ldflags "-X $(PKG)/cmd.Version=$(VERSION)"

deps:
	go get ./...

install:
	go install $(MAC_LDFLAGS)

test: deps
	go test ./... -timeout 60s $(LDFLAGS) -v

cover: deps
	go test ./... -race -v -timeout 15s -coverprofile=coverage.out
	go tool cover -html=coverage.out

fmt:
	go fmt ./...

bench: deps
	go test ./... $(LDFLAGS) -v -run NOT_EXISTING -bench $(BENCHMARK) -benchtime 5s

race: deps
	go test ./... -v -race -timeout 15s

release: release-x64 release-mac

ensure-dist: deps
	mkdir -p dist

release-bin: ensure-dist
	GOOS=linux GOARCH=amd64 go build -o dist/bild $(LDFLAGS)

release-x64: ensure-dist
	GOOS=linux GOARCH=amd64 go build -o dist/bild $(LDFLAGS) && cd dist && tar -czf bild_$(VERSION)_x64.tar.gz bild && rm bild

release-x86: ensure-dist
	GOOS=linux GOARCH=386 go build -o dist/bild $(LDFLAGS) && cd dist && tar -czf bild_$(VERSION)_x86.tar.gz bild && rm bild

release-mac: ensure-dist
	go build $(MAC_LDFLAGS) -o dist/bild && cd dist && tar -czf bild_$(VERSION)_mac.tar.gz bild && rm bild
