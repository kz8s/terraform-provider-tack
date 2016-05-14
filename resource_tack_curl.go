package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTackCurl() *schema.Resource {
	return &schema.Resource{
		Create: resourceTackCurlCreate,
		Read:   resourceTackCurlRead,
		Delete: resourceTackCurlDelete,

		Schema: map[string]*schema.Schema{

			"body": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Body of response",
			},

			"url": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceTackCurlCreate(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling create")

	url := d.Get("url").(string)

	body, err := curl(url)
	if err != nil {
		return
	}
	d.Set("body", body)
	d.SetId(url)
	return
}

func resourceTackCurlRead(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling read")
	return
}

func resourceTackCurlDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] calling delete")
	return nil
}

func curl(url string) (body string, err error) {

	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("failed to GET url: %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("failed to GET url: %s: invalid status code: %d", url, resp.StatusCode)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	body = strings.TrimSpace(string(bodyBytes))

	return
}
