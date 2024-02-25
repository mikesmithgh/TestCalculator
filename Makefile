.PHONY: help
help:
	@echo "==> describe make commands"
	@echo ""
	@echo "test      ==> run unit tests and integration tests"
	@echo "test_unit ==> run unit tests"
	@echo "cover     ==> run test coverage"

.PHONY: test
test:
	@go clean --testcache && go test ./...

.PHONY: test_unit
test_unit:
	@go clean --testcache && go test ./calculator/...

.PHONY: cover
cover:
	@go clean --testcache
	@go test ./... --coverprofile=cov.out
	@go tool cover -html=cov.out -o cov.html 
	@echo ""
	@echo "coverage report generated"
	@echo "  cov.out"
	@echo "  cov.html"

