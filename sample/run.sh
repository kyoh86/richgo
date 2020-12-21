#!/bin/bash

bin="$(cd -- "$(dirname -- "${BASH_SOURCE:-${(%):-%N}}")"; pwd)"
cd "${bin}/.."

export RICHGO_LOCAL=1

OPTIONS=(test -tags=sample ./sample/... -cover)

echo "=====================================   go test   ====================================="
\go "${OPTIONS[@]}"

echo ""
echo "===================================== richgo test ====================================="
go run . "${OPTIONS[@]}"

echo "====================================   go test -v  ===================================="
\go "${OPTIONS[@]}" -v

echo ""
echo "==================================== richgo test -v ==================================="
go run . "${OPTIONS[@]}" -v
