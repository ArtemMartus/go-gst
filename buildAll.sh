#!/usr/bin/env bash

# this script builds all go packages in the current directory

packages=$(go list ./...)
for package in $packages; do
    go build -v -o /dev/null "$package" || exit 1
done
