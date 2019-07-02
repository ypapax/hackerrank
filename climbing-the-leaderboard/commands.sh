#!/usr/bin/env bash

test(){
	GO111MODULE=off go test -v
}

$@