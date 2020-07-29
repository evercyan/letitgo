#!/bin/bash

set -e

if [ "$1" != "" ];
then
    f=$(go list ./... | grep $1)
else
    f="./..."
fi

go test -bench=. $f
