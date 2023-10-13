setup-dev:
	go get -u gotest.tools/gotestsum
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	go get -u github.com/securego/gosec/cmd/gosec
	go get -u github.com/jstemmer/go-junit-report

lint:
	golangci-lint run

test: lint sec
	mkdir -p output
	gotestsum --format=testname -- -coverprofile=output/coverage.out ./...

sec:
	gosec ./...

test-watch:
	gotestsum --format=testname --watch ./...
