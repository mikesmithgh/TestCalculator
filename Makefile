.PHONY: help
help:
	@echo "==> describe make commands"
	@echo ""
	@echo "test      ==> run unit tests and integration tests"
	@echo "test_unit ==> run unit tests"
	@echo "cover     ==> run test coverage"

.PHONY: test
test:
	@go clean --testcache && go test -v ./...

.PHONY: test_unit
test_unit:
	@go clean --testcache && go test -v ./calculator/...

.PHONY: test_integration
test_integration:
	@go clean --testcache && go test -v ./integration/...

.PHONY: cover
cover:
	@go clean --testcache
	@go test ./... --coverprofile=cov.out
	@go tool cover -html=cov.out -o cov.html 
	@echo ""
	@echo "coverage report generated"
	@echo "  cov.out"
	@echo "  cov.html"

