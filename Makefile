PACKAGE_NAME = gosrv

DIR_EXAMPLES = $(PACKAGE_NAME)/examples

run_examples:
	go run $(DIR_EXAMPLES)

test:
	python tests/multi_requests.py