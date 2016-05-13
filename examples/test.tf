resource "tack_coreos" "foo" {
  channel = "stable"
  region  = "us-west-1"
  vmtype  = "hvm"
}

resource "tack_aws_azs" "foo" { region  = "us-west-2" }

output "ami" { value = "${ tack_coreos.foo.ami }" }
output "azs" { value = "${ join(",", tack_aws_azs.foo.*.azs) }" }
output "azs_string" { value = "${ tack_aws_azs.foo.azs_string }" }
