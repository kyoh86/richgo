#!/bin/sh

bin="$(cd -- "$(dirname -- "${BASH_SOURCE:-${(%):-%N}}")"; pwd)"
cd ${bin}/..

find ./sample -type f -name '*.go' -exec gsed -e 's|.*// *COMMENT:|//&|' -e 's|//\(.*// *UNCOMMENT:\)|\1|' -i {} \;

echo "=====================================   go test   ====================================="
\go test ./sample/... -v

echo ""
echo "===================================== richgo test ====================================="
export RICHGO_LOCAL=1
go run richgo.go test ./sample/... -v

find ./sample -type f -name '*.go' -exec gsed -e 's|.*// *UNCOMMENT:|//&|' -e 's|//\(.*// *COMMENT:\)|\1|' -i {} \;
