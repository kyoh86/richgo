test:
	go test $(shell go list ./... | grep -vFe'/vendor/')

example:
	go run cmd/xdg-example/main.go

gen:
	go generate $(shell go list ./... | grep -vFe'/vendor/')

.PHONY: test example gen
