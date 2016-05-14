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
				Config: testAwsAzsRecordBasic,
				Check: resource.ComposeTestCheckFunc(
					testAwsAzsRecordExists("tack_aws_azs.foo"),
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

		azsString := s.RootModule().Outputs["azs_string"]
		match, err := regexp.MatchString("us-west-1.*", azsString)
		if err != nil {
			return err
		}
		if !match {
			return fmt.Errorf("azs_string doesn't contain us-west-1: %s", azsString)
		}

		fmt.Println(azsString)

		return nil
	}
}

const testAwsAzsRecordBasic = `
resource "tack_aws_azs" "foo" {
  region  = "us-west-1"
}

output "azs_string" {
  value = "${ tack_aws_azs.foo.azs_string }"
}
`
