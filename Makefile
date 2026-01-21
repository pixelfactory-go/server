.PHONY: fmt test lint fuzz

SHELL := /bin/bash

fmt:
	@diff -u <(echo -n) <(gofmt -d -s .)

test:
	@go test -v -race -coverprofile coverage.txt -covermode atomic ./...

lint:
	@golangci-lint run ./...

fuzz:
ifdef FUZZ_TEST
	@echo "Running fuzz test: $(FUZZ_TEST)..."
	@go test -run=^$$ -fuzz=$(FUZZ_TEST) -fuzztime=30s .
	@echo "Done!"
else
	@echo "Running all fuzz tests..."
	@for test in $$(grep -h '^func Fuzz' *_test.go 2>/dev/null | sed 's/func \(Fuzz[^(]*\).*/\1/'); do \
		echo "Running $$test..."; \
		go test -run=^$$ -fuzz=$$test -fuzztime=30s . || exit 1; \
	done
	@echo "Done!"
endif
