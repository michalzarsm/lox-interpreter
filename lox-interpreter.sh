#!/bin/sh

set -e

(
  cd "$(dirname "$0")"
  go build -o /tmp/interpreter-target ./cmd/lox-interpreter
)

exec /tmp/interpreter-target "$@"
