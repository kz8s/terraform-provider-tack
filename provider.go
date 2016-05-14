package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider Function
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"tack_aws_account_id": resourceTackAwsAccountID(),
			"tack_aws_azs":        resourceTackAwsAzs(),
			"tack_coreos":         resourceTackCoreos(),
			"tack_my_ip":          resourceTackMyIP(),
		},
	}
}
