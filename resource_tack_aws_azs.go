package main

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTackAwsAzs() *schema.Resource {
	return &schema.Resource{
		Create: resourceTackAwsAzsCreate,
		Read:   resourceTackAwsAzsRead,
		// Update: resourceTackAwsAzsUpdate,
		Delete: resourceTackAwsAzsDelete,
		// Exists: resourceTackAwsAzsExists,

		Schema: map[string]*schema.Schema{
			"azs": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "azs",
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceTackAwsAzsCreate(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling create")

	region := d.Get("region").(string)

	azs, err := getAvailabilityZones(region)
	if err != nil {
		return
	}

	d.Set("azs", strings.Join(azs, ","))

	d.SetId(region + "!")
	return
}

func resourceTackAwsAzsRead(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling read")
	return
}

// func resourceTackAwsAzsUpdate(d *schema.ResourceData, m interface{}) error {
// 	return nil
// }

func resourceTackAwsAzsDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] calling delete")
	return nil
}

// func resourceTackAwsAzsExists(d *schema.ResourceData, meta interface{}) (bool, error) {
// 	log.Println("[INFO] calling exists")
// 	return d.Get("region").(string) == d.Id(), nil
// }

func getAvailabilityZones(region string) (azs []string, err error) {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String(region)})

	var params *ec2.DescribeAvailabilityZonesInput
	resp, err := svc.DescribeAvailabilityZones(params)

	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, az := range resp.AvailabilityZones {
		azs = append(azs, *az.ZoneName)
	}
	return
}
