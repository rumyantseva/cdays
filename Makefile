PROJECT?=github.com/rumyantseva/cdays
BUILD_PATH?=cmd/cdays
APP?=cdays

PORT?=8887
INTERNAL_PORT?=8888

GOOS?=linux
GOARCH?=amd64

REGISTRY?=docker.io/webdeva
NAMESPACE?=rumyantseva
CONTAINER_NAME?=${NAMESPACE}-${APP}
CONTAINER_IMAGE?=${REGISTRY}/${CONTAINER_NAME}

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

clean:
	rm -f ./bin/${GOOS}-${GOARCH}/${APP}

build: test clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/internal/version.Release=${RELEASE} \
		-X ${PROJECT}/internal/version.Commit=${COMMIT} \
		-X ${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
	-o ./bin/${GOOS}-${GOARCH}/${APP} ${PROJECT}/${BUILD_PATH}

image: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

run: image
	docker run --name ${APP} -p ${PORT}:${PORT} -p ${INTERNAL_PORT}:${INTERNAL_PORT} --rm \
		-e "PORT=${PORT}" -e "INTERNAL_PORT=${INTERNAL_PORT}" \
		$(CONTAINER_IMAGE):$(RELEASE)

test: clean
	go test -race ./...
