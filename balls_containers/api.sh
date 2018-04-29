#!/usr/bin/env bash
set -ex
go install
balls_containers -bind 0.0.0.0:8083
