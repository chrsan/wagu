#!/bin/bash

for go in *.go; do
    if [ "${go}" = "test_test.go" ]; then
        continue
    fi
    mv -- "${go}" "${go%.go}_test.go"
done
