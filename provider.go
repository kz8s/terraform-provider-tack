package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider Function
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"tack_coreos": resourceTackCoreos(),
		},
	}
}