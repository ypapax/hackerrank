#!/usr/bin/env bash
set -ex

go install
cat 19.txt
cat 19.txt | hour_glasses