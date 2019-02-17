.PHONY: default gen test vendor install sample

default:
	echo use gen, test, vendor or install

gen:
	goblet -g -p test -o editor/test/output_test.go --ignore-dotfiles ./sample/out_*.txt
	gofmt -w editor/test/output_test.go

test:
	go test -v -race ./...

lint:
	gometalinter ./...

sample:
	sample/run.sh

install:
	go install ./...
