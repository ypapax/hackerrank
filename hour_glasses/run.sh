#!/usr/bin/env bash
set -ex

go install
cat 19.txt >&2
cat 19.txt | hour_glasses