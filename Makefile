
install:
	go install gotest.tools/gotestsum@latest && \
	go mod tidy && \
	go mod vendor

test:
	gotestsum --format testname
