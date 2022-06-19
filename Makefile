OS ?= darwin
MAIN ?= cmd/server/server.go
BUILD_DIR ?= build

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=${OS} go build -a -installsuffix cgo -o ${BUILD_DIR}/main ${MAIN}

.PHONY: run
run: 
	./${BUILD_DIR}/main

.PHONY: all
all: 
	build run