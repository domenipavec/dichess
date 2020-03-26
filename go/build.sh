#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o dichess
tar czf linux_amd64.tar.gz dichess
GOOS=linux GOARCH=arm go build -o dichess
tar czf linux_arm.tar.gz dichess
rm dichess
