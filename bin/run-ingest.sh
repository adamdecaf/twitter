#!/bin/bash

platform=`uname -s | tr '[:upper:]' '[:lower:]'`
if [[ "$platform" == "darwin" ]]
then
    platform="osx"
fi

export CONSUMER_KEY=''
export CONSUMER_SECRET=''
export ACCESS_TOKEN=''
export ACCESS_SECRET=''

./bin/ingest-$platform
