#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o linux_amd64
GOOS=linux GOARCH=arm go build -o linux_arm
