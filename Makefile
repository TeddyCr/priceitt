.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo "Usage: [make <target>]. Needs to be run from the root directory.\n"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":"}; {printf "\033[35m%-35s\033[0m %s\n", $$2, $$3}'

.PHONY: format
format: ## format -- Format go code using golangci-lint
	cd service && golangci-lint fmt
	cd service && golangci-lint run --fix --issues-exit-code=0 --timeout=10m0s

.PHONY: format_check
format_check: ## format_check -- Check go code using golangci-lint
	cd service && golangci-lint fmt --diff
	cd service && golangci-lint run --issues-exit-code=0 --timeout=10m0s
	
.PHONY: run_unit_tests
run_unit_tests: ## run_unit_tests -- Run unit tests
	cd service && go test -tags=unit -v -cover -count=1 ./...

.PHONY: run_integration_tests
run_integration_tests: ## run_integration_tests -- Run integration tests
	cd service && go test -tags=integration -v -cover -count=1 ./...

.PHONY: run_migration_tests
run_migration_tests: ## run_migration_tests -- Run migration tests
	cd service && go test -tags=migration -v -cover -count=1 ./...

.PHONY: run_all_tests
run_all_tests: ## run_all_tests -- Run all tests
	cd service && go test -v -cover -tags=unit,integration -count=1 -coverprofile=./coverage.out ./...