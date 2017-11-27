.PHONY: default gen test vendor install sample

default:
	echo use gen, test, vendor or install

gen:
	go-bindata -o editor/test/output_test.go -pkg test -prefix sample/out_ ./sample/out_*.txt

test:
	go test ./...

sample:
	sample/run.sh

vendor:
	dep ensure

install:
	go install ./...
