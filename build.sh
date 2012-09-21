#!/bin/sh

APP=diary-server
pushd go > /dev/null
go build
popd > /dev/null

mv go/go $APP
