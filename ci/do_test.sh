#!/bin/bash

set -e -u

lang="$(jq -r .lang < oas.proj)"

case "$lang"
in
  go)
    go test -v ./...
    ;;
  *)
    echo "lang ${lang} not supported" > /dev/stderr
    exit 1
    ;;
esac
