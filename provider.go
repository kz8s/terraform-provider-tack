package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider Function
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"tack_coreos":         resourceTackCoreos(),
			"tack_aws_azs":        resourceTackAwsAzs(),
			"tack_aws_account_id": resourceTackAwsAccountID(),
		},
	}
}
