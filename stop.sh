#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GOPATH=$DIR
export GOBIN=$DIR/bin
kill `cat $DIR/go.pid`
rm $DIR/go.pid
rm -rf $GOBIN