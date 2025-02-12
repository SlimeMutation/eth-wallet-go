eth-wallet-go:
	env GO111MODULE=on go build

clean:
	rm eth-wallet-go

test:
	go test -v ./...

lint:
	golangci-lint run ./...

.PHONY: \
	eth-wallet-go \
	clean \
	test \
	lint