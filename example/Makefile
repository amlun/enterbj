BINARY=enterbj
GOARCH = amd64

VERSION=0.0.1
BUILD=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.PHONY: help build linux darwin windows clean

help:
	@echo "usage: make <option>"
	@echo "options and effects:"
	@echo "    build  : build the binary"
	@echo "    linux  : build the linux binary"
	@echo "    darwin : build the darwin binary"
	@echo "    windows: build the windows binary"
	@echo "    clean  : clean binary"

all: linux darwin windows

build:
	@go build ${LDFLAGS} -o ${BINARY}

linux:
	@GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-linux-${GOARCH}

darwin:
	@GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-darwin-${GOARCH}

windows:
	@GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-windows-${GOARCH}

clean:
	@rm -f ${BINARY}
	@rm -f ${BINARY}-*