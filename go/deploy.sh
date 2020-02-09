#!/bin/bash

set -e

mkdir -p ./bin/

export GOARCH=arm
go build -o ./bin/dichess
scp ./bin/dichess pi@172.29.163.160:~
