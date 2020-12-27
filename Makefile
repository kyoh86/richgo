.PHONY: gen lint test install man sample

VERSION := `git vertag get`
COMMIT  := `git rev-parse HEAD`

gen:
	go run github.com/rakyll/statik -src=./sample -dest editor/test -include='*.txt' -f
	gofmt -w editor/test/output_test.go

lint: gen
	golangci-lint run

test: lint
	go test -v --race ./...

install: test
	go install -a -ldflags "-X=main.version=$(VERSION) -X=main.commit=$(COMMIT)" ./...

sample:
	sample/run.sh
