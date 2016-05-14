resource "tack_aws_account_id" "foo" {}

resource "tack_aws_azs" "foo" { region  = "us-west-2" }

resource "tack_coreos" "foo" {
  channel = "stable"
  region  = "us-west-1"
  vmtype  = "hvm"
}

resource "tack_curl" "foo" {
  url = "http://myip.vg"
}

resource "tack_my_ip" "foo" {}

output "account_id" { value = "${ tack_aws_account_id.foo.id }" }
output "ami" { value = "${ tack_coreos.foo.ami }" }
output "body" { value = "${ tack_curl.foo.body }" }
output "azs" { value = "${ join(",", tack_aws_azs.foo.*.azs) }" }
output "azs_string" { value = "${ tack_aws_azs.foo.azs_string }" }
output "my_ip" { value = "${ tack_my_ip.foo.ip }" }
