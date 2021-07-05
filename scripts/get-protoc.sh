#!/bin/bash -e

OS=$(uname)
if [ "$OS" == "Darwin" ]; then
    OS="osx"
fi

curl -L -o _output/protoc.zip https://github.com/google/protobuf/releases/download/v3.17.3/protoc-3.17.3-${OS}-x86_64.zip
bsdtar -x -f _output/protoc.zip -C _output bin/protoc
