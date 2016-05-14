#!/bin/bash -eux

terraform plan
terraform graph
terraform apply
terraform output
terraform destroy
terraform plan
