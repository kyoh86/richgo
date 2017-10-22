default:
	echo use "test" or "cover"

test:
	go test ./...

cover:
	_bin/cover.sh

.PHONY: default test cover
