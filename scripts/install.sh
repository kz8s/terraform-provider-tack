#!/bin/bash -eux

# copies (installs) `terraform-provider-tack` from current directory to the
# directory that `terraform` is in. follows symlink if necessary, since
# terraform requires the plugin to reside in the same directory.

TF=$(which terraform)
TF_DIR=$(dirname $TF)

if [[ -L "$TF" ]]; then
  TF_LN=$(dirname "`readlink $TF`")
  TF_DIR=$(cd $TF_DIR && cd $TF_LN && pwd)
fi

echo $TF_DIR

cp terraform-provider-tack $TF_DIR
