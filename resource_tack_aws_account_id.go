package main

import (
	"errors"
	"log"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	region = "us-west-1"
)

//golint:ignore
func resourceTackAwsAccountID() *schema.Resource {
	return &schema.Resource{
		Create: resourceTackAwsAccountIDCreate,
		Read:   resourceTackAwsAccountIDRead,
		Delete: resourceTackAwsAccountIDDelete,
	}
}

func resourceTackAwsAccountIDCreate(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling create")

	accountid, err := getAccountID()
	if err != nil {
		return
	}

	d.SetId(accountid)
	return
}

func resourceTackAwsAccountIDRead(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling read")
	return
}

func resourceTackAwsAccountIDDelete(d *schema.ResourceData, m interface{}) (err error) {
	log.Println("[INFO] calling delete")
	return
}

func getAccountID() (id string, err error) {
	svc := iam.New(session.New(), &aws.Config{Region: aws.String(region)})

	var params *iam.GetUserInput
	resp, err := svc.GetUser(params)

	if err != nil {
		log.Println(err.Error())
		return
	}

	re1, _ := regexp.Compile(`arn:aws:iam::(\d*):`)
	result := re1.FindStringSubmatch(*resp.User.Arn)
	if len(result) != 2 {
		log.Println("arn: *resp.User.Arn")
		log.Println(result)
		err = errors.New("couldn't parse account id")
		return
	}

	id = result[1]
	return
}
