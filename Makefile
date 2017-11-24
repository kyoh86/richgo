default:
	echo use "test", "cover" or "sample"

gen:
	go-bindata -o editor/test/output_test.go -pkg test -prefix sample/out_ ./sample/out_*.txt

test:
	go test ./...

cover:
	goveralls -race -package ./... -repotoken $(COVERALLS_TOKEN)

sample:
	sample/run.sh

.PHONY: default gen test cover sample
