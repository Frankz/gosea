APP_NAME = gosea
PATH := $(GOPATH)/bin:$(PATH)

TAG = $(shell git describe --tags --always --dirty)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
REVISION = $(shell git rev-parse HEAD)
USER = $(shell whoami)
GOVERSION = $(shell go version | awk '{print $$3}')
NOW	= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GOSEA_VERSION = "\
	-X github.com/komand/gosea/version.appName=${APP_NAME} \
	-X github.com/komand/gosea/version.tag=${TAG} \
	-X github.com/komand/gosea/version.branch=${BRANCH} \
	-X github.com/komand/gosea/version.revision=${REVISION} \
	-X github.com/komand/gosea/version.goVersion=${GOVERSION} \
	-X github.com/komand/gosea/version.buildUser=${USER} \
	-X github.com/komand/gosea/version.buildDate=${NOW}"

all: build

clean:
	rm -rf build

build: clean
	go build -i -o build/${APP_NAME} -ldflags ${GOSEA_VERSION} ./cmd/gosea
	
key:
	openssl genrsa -out server.key 2048
	openssl ecparam -genkey -name secp384r1 -out server.key

cert:
	openssl req -new -x509 -sha256 -key server.key -out server.pem -days 365

install: key cert
	go install

run:
	@build/gosea
