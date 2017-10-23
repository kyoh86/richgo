#!/bin/sh

bin="$(cd -- "$(dirname -- "${BASH_SOURCE:-${(%):-%N}}")"; pwd)"
cd ${bin}/..

OPTIONS="test ./sample/... -v -cover"

find ./sample -type f -name '*.go' -exec gsed -e 's|.*// *COMMENT:|//&|' -e 's|//\(.*// *UNCOMMENT:\)|\1|' -i {} \;

echo "=====================================   go test   ====================================="
\go ${OPTIONS}

echo ""
echo "===================================== richgo test ====================================="
export RICHGO_LOCAL=1
go run richgo.go ${OPTIONS}

find ./sample -type f -name '*.go' -exec gsed -e 's|.*// *UNCOMMENT:|//&|' -e 's|//\(.*// *COMMENT:\)|\1|' -i {} \;

git checkout sample
