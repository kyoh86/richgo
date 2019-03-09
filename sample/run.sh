#!/bin/sh

bin="$(cd -- "$(dirname -- "${BASH_SOURCE:-${(%):-%N}}")"; pwd)"
cd ${bin}/..

OPTIONS="test -tags=sample ./sample/... -v -cover"

echo "=====================================   go test   ====================================="
\go ${OPTIONS}

echo ""
echo "===================================== richgo test ====================================="
export RICHGO_LOCAL=1
go run main.go ${OPTIONS}
