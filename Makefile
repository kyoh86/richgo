.PHONY: gen lint test install man sample

VERSION := `git vertag get`
COMMIT  := `git rev-parse HEAD`

gen:
	goblet -g -p test -o editor/test/output_test.go --ignore-dotfiles ./sample/out_*.txt
	gofmt -w editor/test/output_test.go

lint: gen
	golangci-lint run

test: lint
	go test -v --race ./...

install: test
	go install -a -ldflags "-X=main.version=$(VERSION) -X=main.commit=$(COMMIT)" ./...

man: test
	go run main.go --help-man > richgo.1

sample:
	sample/run.sh
