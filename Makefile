PROJECT?=github.com/rumyantseva/cdays
BUILD_PATH?=cmd/cdays

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

build: test
	go build \
		-ldflags "-s -w -X ${PROJECT}/internal/version.Release=${RELEASE} \
		-X ${PROJECT}/internal/version.Commit=${COMMIT} \
		-X ${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
	-o ./bin/cdays ${PROJECT}/${BUILD_PATH}

test:
	go test -race ./...
