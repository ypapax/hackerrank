#!/usr/bin/env bash
set -ex

go install
cat in.txt >&2
cat in.txt | balls_containers $@