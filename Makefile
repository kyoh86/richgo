.PHONY: gen lint test install man sample

VERSION := $(shell git describe --tags --abbrev=0)
COMMIT  := $(shell git rev-parse HEAD)

gen:
	go run github.com/rakyll/statik -src=./sample -dest editor/test -include='*.txt' -f

lint: gen
	golangci-lint run

test: lint
	go test -v --race ./...

install: test
	go install -a -ldflags "-X=main.version=$(VERSION) -X=main.commit=$(COMMIT)" ./...

sample:
	sample/run.sh
