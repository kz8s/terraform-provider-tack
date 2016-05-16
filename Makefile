terraform-provider-tack: $(shell find . -iname "*.go")
	go build

apply: terraform-provider-tack
	terraform apply examples/

build: terraform-provider-tack

clean:
	@rm -rf terraform.tfstate.backup terraform.tfstate terraform-provider-tack release ||:

destroy: terraform-provider-tack
	terraform destroy examples/

fmt:
	go fmt -x .

get:
  go get -v -d -t ./...
	git submodule update --init --recursive # vendor/ terraform v0.6.16

graph: terraform-provider-tack
	terraform graph examples/

install: terraform-provider-tack
	scripts/install.sh

plan: terraform-provider-tack
	terraform plan examples/

release:
	scripts/release.sh

test:
	TF_ACC=1 go test -v

vet:
	go vet -x .

.PHONY: apply build clean destroy fmt get graph plan release test vet
