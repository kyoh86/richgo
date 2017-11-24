default:
	echo use "gen", "test" or "sample"

gen:
	go-bindata -o editor/test/output_test.go -pkg test -prefix sample/out_ ./sample/out_*.txt

test:
	go test ./...

sample:
	sample/run.sh

.PHONY: default gen test sample
