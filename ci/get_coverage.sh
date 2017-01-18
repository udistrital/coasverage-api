#!/bin/bash

set -e -u

cleanup() {
  rm -rf /tmp/$$.out
}

trap cleanup EXIT

lang="$(jq -r .lang < oas.proj)"

case "$lang"
in
  go)
    go test -v -cover ./... &> /tmp/$$.out
    cat /tmp/$$.out > /dev/stderr
    coverage="$(awk '$NF=="statements" && $(NF-1)=="of"{print $(NF-2);exit}' /tmp/$$.out)"
    ;;
  *)
    echo "lang ${lang} not supported"
    exit 1
    ;;
esac

if [ -n "$coverage" ]
then
  echo -n "${coverage%\%}"
fi
