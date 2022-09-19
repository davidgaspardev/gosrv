PACKAGE_NAME = github.com/davidgaspardev/gosrv

DIR_EXAMPLES = examples

run_examples:
	go run $(DIR_EXAMPLES)/main.go

test:
	python tests/multi_requests.py