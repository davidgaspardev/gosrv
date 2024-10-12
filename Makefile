PACKAGE_NAME = github.com/davidgaspardev/gosrv
DIR_EXAMPLES = examples
PYTHON = $(shell if command -v python3 > /dev/null 2>&1; then echo python3; else echo python; fi)

.PHONY: help examples test

help:
	@echo "Usage: make <target>"
	@echo
	@egrep "^(.+)\:\ .*##\ (.+)" ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'
	@echo
	@echo "Options (can be passed to 'make test'):"
	@echo "  max=NUMBER    Set the max value (default is 512)"
	@echo "  method=METHOD Set the method (default is empty)"
	@echo "  path=PATH     Set the path (default is empty)"

examples: ## Run examples
	@go run $(DIR_EXAMPLES)/main.go

test: ## Run tests
	@$(PYTHON) tests/multi_requests.py --max $(or $(max), 512) $(if $(method), --method $(method)) $(if $(path), --path $(path))
