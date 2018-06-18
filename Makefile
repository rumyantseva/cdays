PROJECT?=github.com/rumyantseva/cdays
BUILD_PATH?=cmd/cdays


build: test
	go build  -o ./bin/cdays ${PROJECT}/${BUILD_PATH}

test:
	go test -race ./...
