#!/bin/bash

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

GODIR=$DIR/../go/bluetoothpb/
mkdir -p $GODIR
DARTDIR=$DIR/../flutter/dichess/lib/bluetoothpb/
mkdir -p $DARTDIR
protoc -I=$DIR --go_out=$GODIR --dart_out=$DARTDIR $DIR/bluetoothpb.proto
