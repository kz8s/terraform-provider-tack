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
	urlGetIP = "http://myip.vg"
)

func resourceTackMyIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceTackMyIPCreate,
		Read:   resourceTackMyIPRead,
		Delete: resourceTackMyIPDelete,

		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current IP address",
			},
		},
	}
}

func resourceTackMyIPCreate(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling create")
	ip, err := getMyIP()
	if err != nil {
		return
	}
	d.Set("ip", ip)
	d.SetId("myip!")
	return
}

func resourceTackMyIPRead(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling read")
	return
}

func resourceTackMyIPDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] calling delete")
	return nil
}

func getMyIP() (ip string, err error) {

	resp, err := http.Get(urlGetIP)
	if err != nil {
		err = fmt.Errorf("failed to get IP: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("failed to get IP: invalid status code: %d", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	ip = strings.TrimSpace(string(body))

	return
}
