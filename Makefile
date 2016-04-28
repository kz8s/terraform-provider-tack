terraform-provider-tack: $(shell find . -iname "*.go")
	go build -o $@

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

.PHONY: clean fmt get test vet
