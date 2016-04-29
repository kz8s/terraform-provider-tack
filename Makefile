terraform-provider-tack: $(shell find . -iname "*.go")
	scripts/build.sh

build: terraform-provider-tack

clean:
	@rm -rf bin/ ||:

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
