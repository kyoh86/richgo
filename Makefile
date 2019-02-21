.PHONY: gen test lint install sample

install: gen test lint
	go install ./...

lint: test
	gometalinter ./...

test: gen
	go test -v -race ./...

gen:
	goblet -g -p test -o editor/test/output_test.go --ignore-dotfiles ./sample/out_*.txt
	gofmt -w editor/test/output_test.go

sample:
	sample/run.sh
