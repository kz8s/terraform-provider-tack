package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestTackAwsAzsRecord_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAwsAzsRecord_basic,
				Check: resource.ComposeTestCheckFunc(
					testCoreOSAmiRecordExists("tack_coreos.foo"),
				),
			},
		},
	})
}

func testAwsAzsRecordExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource not found: %s", n)
		}

		azs_string := s.RootModule().Outputs["azs_string"]
		match, err := regexp.MatchString("us-west-1.*", azs_string)
		if err != nil {
			return err
		}
		if !match {
			return fmt.Errorf("azs_string doesn't contain us-west-1: %s", azs_string)
		}

		fmt.Println(azs_string)

		return nil
	}
}

const testAwsAzsRecord_basic = `
resource "tack_aws_azs" "foo" {
  region  = "us-west-1"
}

output "azs_string" {
  value = "${ tack_aws_azs.foo.azs_string }"
}
`
