package main

import (
	"testing"

	"github.com/hashicorp/terraform/terraform"
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

var testProviders = map[string]terraform.ResourceProvider{
	"tack": Provider(),
}
