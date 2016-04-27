terraform-provider-tack: $(shell find . -iname "*.go")
	go build -o $@

clean:
	@rm terraform-provider-tack ||:

test:
	TF_ACC=1 go test -v

.PHONY: clean test
