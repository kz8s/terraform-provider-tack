#!/bin/bash -ux

SHA=sha256sum

hash $SHA &> /dev/null
if [ $? -eq 1 ]; then
    SHA="shasum -a 256"
fi

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

cd release && $SHA *.tar.gz > terraform-provider-tack_${VER}.SHA256SUMS
