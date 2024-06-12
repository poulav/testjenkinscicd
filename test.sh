#!/usr/bin/env sh

go_ver = 'go version'

if go_ver; then
    echo "Go Exists!"
else
    apt update
    apt install golang-go
    go version
fi