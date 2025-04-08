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
	
	