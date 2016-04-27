resource "tack_coreos" "foo" {
  channel = "stable"
  region  = "us-west-1"
  vmtype  = "hvm"
}

output "ami" {
  value = "${ tack_coreos.foo.ami }"
}
