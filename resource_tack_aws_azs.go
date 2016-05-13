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

			"azs_string": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "azs",
			},

			"azs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

	d.Set("azs_string", strings.Join(azs, ","))
	d.Set("azs", azs)

	d.SetId(region + "!")
	return
}

func resourceTackAwsAzsRead(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling read")
	return
}

func resourceTackAwsAzsDelete(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling delete")
	return
}

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
