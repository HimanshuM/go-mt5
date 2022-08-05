GO = go
PACKAGE = github.com/HimanshuM/go-mt5

run:
	$(GO) run main.go

test:
	$(GO) test -run ^TestMain$$ $(PACKAGE)/mt5tests -v

lint:
	go fmt ./...
	golangci-lint run