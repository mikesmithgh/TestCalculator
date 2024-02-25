.PHONY: help
help:
	@echo "==> describe make commands"
	@echo ""
	@echo "test             ==> run unit tests and integration tests"
	@echo "testjsonfmt      ==> run unit tests and integration tests with json output"
	@echo "test_unit        ==> run unit tests"
	@echo "test_integration ==> run integration tests"
	@echo "cover            ==> run test coverage"

.PHONY: test
test:
	@go clean --testcache && go test -v ./...

.PHONY: testjsonfmt
testjsonfmt:
	@go clean --testcache && go test -json -v ./...

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

