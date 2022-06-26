GO = go
PACKAGE = github.com/HimanshuM/go_mt5

run:
	$(GO) run main.go

test:
	$(GO) test -run ^TestMain$$ $(PACKAGE)/mt5tests