terraform-provider-tack: $(shell find . -iname "*.go")
	go build -o $@

build: terraform-provider-tack

clean:
	@rm terraform-provider-tack ||:

fmt:
	go fmt -x .

get:
	go get -v ./...

test:
	TF_ACC=1 go test -v

vet:
	go vet -x ./...

.PHONY: build clean fmt get test vet
