.PHONY: test

test:
	mkdir -p test-results
	go install gotest.tools/gotestsum@latest
	gotestsum --format=short-verbose --junitfile=test-results/junit-report.xml --junitfile-testsuite-name=short --junitfile-testcase-classname=short -- -race ./... -timeout 90s
