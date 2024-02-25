.PHONY: help
help:
	@echo "==> describe make commands"
	@echo ""
	@echo "test ==> run unit tests"

.PHONY: test
test:
	@go clean --testcache && go test ./...
