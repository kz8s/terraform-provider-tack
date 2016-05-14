#!/bin/bash -ex

# usage
[ -z "$1" ] || cd $1

set -o nounset

pwd

terraform plan
terraform graph
terraform apply
terraform output
terraform destroy -force
terraform plan
