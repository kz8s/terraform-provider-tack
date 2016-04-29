#!/bin/bash -eux

FN=terraform-provider-tack

buildOsArch() {
  GOOS=$1
  GOARCH=$2
  mkdir -p bin/$GOOS/$GOARCH
  go build -o bin/$GOOS/$GOARCH/$FN
}

buildOsArch darwin amd64
buildOsArch linux amd64
