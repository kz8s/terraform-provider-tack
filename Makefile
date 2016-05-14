terraform-provider-tack: $(shell find . -iname "*.go")
	go build

apply: terraform-provider-tack
	terraform apply

build: terraform-provider-tack

clean:
	@rm -rf terraform.tfstate.backup terraform.tfstate terraform-provider-tack ||:

destroy: terraform-provider-tack
	terraform destroy

fmt:
	go fmt -x .

get:
	go get -v ./...

graph: terraform-provider-tack
	terraform graph

plan: terraform-provider-tack
	terraform plan

test:
	TF_ACC=1 go test -v

vet:
	go vet -x ./...

.PHONY: apply build clean destroy fmt get graph plan test vet
