#!/bin/bash

if [ -z "$1" ]
  then
    echo "You must provide module name to installation."
fi

MODULE_NAME=$1

PROJECT_SRC=$GOPATH/src/aws-vpc/$MODULE_NAME
mkdir -p $PROJECT_SRC
cp $MODULE_NAME/$MODULE_NAME.go $PROJECT_SRC/$MODULE_NAME.go
go install $PROJECT_SRC/$MODULE_NAME.go
nohup $GOBIN/$MODULE_NAME > $MODULE_NAME-log.txt 2>&1 &