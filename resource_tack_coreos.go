package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

const (
	urlTemplate = "http://%s.release.core-os.net/amd64-usr/current/coreos_production_ami_%s_%s.txt"
)

func resourceTackCoreos() *schema.Resource {
	return &schema.Resource{
		Create: resourceTackCoreosCreate,
		Read:   resourceTackCoreosRead,
		Update: resourceTackCoreosUpdate,
		Delete: resourceTackCoreosDelete,
		Exists: resourceTackCoreosExists,

		Schema: map[string]*schema.Schema{
			"ami": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ami",
			},
			"channel": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vmtype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceTackCoreosCreate(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling create")
	ami, err := getAmi(d)
	if err != nil {
		return
	}
	d.Set("ami", ami)

	id := getId(d)
	d.SetId(id)
	return
}

func resourceTackCoreosRead(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling read")
	return
}

func resourceTackCoreosUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceTackCoreosDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] calling delete")
	return nil
}

func resourceTackCoreosExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	log.Println("[INFO] calling exists")
	return getId(d) == d.Id(), nil
}

func getAmi(d *schema.ResourceData) (ami string, err error) {
	c, r, v := getStrings(d)
	amiURL := fmt.Sprintf(urlTemplate, c, v, r)

	log.Println(amiURL)

	resp, err := http.Get(amiURL)
	if err != nil {
		err = fmt.Errorf("failed to get AMI data: %s: %v", c, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("failed to get AMI data: %s: invalid status code: %d", c, resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	ami = strings.TrimSpace(string(body))

	return
}

func getId(d *schema.ResourceData) string {
	c, r, v := getStrings(d)
	return fmt.Sprintf("%s:%s:%s", c, r, v)
}

func getStrings(d *schema.ResourceData) (string, string, string) {
	c := d.Get("channel").(string)
	r := d.Get("region").(string)
	v := d.Get("vmtype").(string)
	return c, r, v
}
