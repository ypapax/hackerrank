#!/usr/bin/env bash
set -ex

generate(){
	create-react-app fe
}

deps(){
	npm install jquery --save
}

run(){
	npm start
}

$@