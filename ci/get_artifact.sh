#!/bin/bash

set -e -u

progname="$(jq -r .progname < oas.proj)"
lang="$(jq -r .lang < oas.proj)"

case "$lang"
in
  go)
    export GOOS="${GOOS:-linux}"
    export GOARCH="${GOARCH:-amd64}"
    go_build_path="$(jq -r .go.build_path < oas.proj)"
    bin="${progname}-${GOOS}-${GOARCH}"
    go build -v -o "${bin}" "${go_build_path:-.}" &> /dev/stderr
    echo -n "${bin}"
    ;;
  *)
    echo "lang ${lang} not supported" > /dev/stderr
    exit 1
    ;;
esac
