#!/bin/bash -eux

FN=terraform-provider-tack

buildOsArch() {
  GOOS=$1
  GOARCH=$2
  mkdir -p release/$GOOS/$GOARCH
  go build -o release/$GOOS/$GOARCH/$FN
}

buildOsArch darwin amd64
buildOsArch linux amd64
