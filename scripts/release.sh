#!/bin/bash -eux

FN=terraform-provider-tack
VER=$(git describe --dirty)

buildOsArch() {
  GOOS=$1
  GOARCH=$2
  mkdir -p release/$GOOS/$GOARCH
  go build -o release/$GOOS/$GOARCH/$FN

  tar pczf release/${FN}_${VER}_${GOOS}-${GOARCH}.tar.gz -C release/$GOOS/$GOARCH $FN
  rm -rf release/$GOOS
}

buildOsArch darwin amd64
buildOsArch linux amd64

cd release && shasum -a 256 *.tar.gz > terraform-provider-tack_${VER}.SHA256SUMS
