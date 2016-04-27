package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var testProviders = map[string]terraform.ResourceProvider{
	"tack": Provider(),
}

func TestTackCoreOSAmiRecord_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testCoreOSAmiRecord_basic,
				Check: resource.ComposeTestCheckFunc(
					testCoreOSAmiRecordExists("tack_coreos.foo"),
				),
			},
		},
	})
}

func testCoreOSAmiRecordExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource not found: %s", n)
		}

		ami := s.RootModule().Outputs["ami"]
		match, err := regexp.MatchString("ami-.*", ami)
		if err != nil {
			return err
		}
		if !match {
			return fmt.Errorf("ami not an ami: %s", ami)
		}

		fmt.Println(ami)

		return nil
	}
}

const testCoreOSAmiRecord_basic = `
resource "tack_coreos" "foo" {
  channel = "stable"
  region  = "us-west-1"
  vmtype  = "hvm"
}

output "ami" {
  value = "${ tack_coreos.foo.ami }"
}
`
