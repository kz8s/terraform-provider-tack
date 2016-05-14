#!/bin/bash -eux

FN=terraform-provider-tack
VER=$(git describe --dirty)

buildOsArch() {
  GOOS=$1
  GOARCH=$2
  mkdir -p release/$GOOS/$GOARCH
  go build -o release/$GOOS/$GOARCH/$FN

  tar pczf release/${FN}_${VER}_${GOOS}-${GOARCH}.tgz -C release/$GOOS/$GOARCH $FN
}

buildOsArch darwin amd64
buildOsArch linux amd64
