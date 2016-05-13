terraform-provider-tack: $(shell find . -iname "*.go")
	go build -o terraform-provider-tack

build: terraform-provider-tack

clean:
	@rm -rf terraform.tfstate.backup terraform.tfstate terraform-provider-tack ||:

fmt:
	go fmt -x .

get:
	go get -v ./...

gox:
	gox

test:
	TF_ACC=1 go test -v

vet:
	go vet -x ./...

.PHONY: build clean fmt get gox test vet
