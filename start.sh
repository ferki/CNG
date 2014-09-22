#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GOPATH=$DIR
export GOBIN=$DIR/bin

mkdir $GOBIN

cd $DIR
go install src/player-service.go

$DIR/bin/player-service &
echo $! > $DIR/go.pid