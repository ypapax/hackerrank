#!/usr/bin/env bash
set -ex
go test -v -run TestArrangeMatrix/4_big.txt -timeout 60m
