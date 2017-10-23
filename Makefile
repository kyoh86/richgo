default:
	echo use "test", "cover" or "sample"

test:
	go test ./...

cover:
	_bin/cover.sh

sample:
	sample/run.sh

.PHONY: default test cover sample
