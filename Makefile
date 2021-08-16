BIN_NAME := rom64
BUILD_DIR := build
BUILD_DATE := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
VERSION ?= $(shell git tag -l --sort=-creatordate 'v*' | head -n1)+$(shell git rev-parse --short HEAD)

LOCAL_GOOS_ARCH := $(shell go version | cut -d' ' -f4 | tr '/' ' ')
LOCAL_GOOS = $(firstword $(LOCAL_GOOS_ARCH))
LOCAL_GOARCH = $(lastword $(LOCAL_GOOS_ARCH))

dynamic_target = $(subst -, , $@)
derived_os = $(word 2, $(dynamic_target))
derived_arch = $(word 3, $(dynamic_target))

.PHONY: clean fresh install lint

all: $(BUILD_DIR)/$(BIN_NAME)-linux-amd64 \
	 $(BUILD_DIR)/$(BIN_NAME)-linux-arm64 \
	 $(BUILD_DIR)/$(BIN_NAME)-darwin-amd64 \
	 $(BUILD_DIR)/$(BIN_NAME)-darwin-arm64 \
	 $(BUILD_DIR)/$(BIN_NAME)-windows-amd64

$(BUILD_DIR)/$(BIN_NAME)-%:
	env GOOS=$(derived_os) GOARCH=$(derived_arch) go build -o $@ \
		-ldflags "-X 'github.com/mroach/n64-go/version.BuildTime=$(BUILD_DATE)'" \
		-ldflags "-X 'github.com/mroach/n64-go/version.Version=$(VERSION)'"

clean:
	rm $(BUILD_DIR)/$(BIN_NAME)-*

fresh: clean all

install:
	cp build/$(BIN_NAME)-$(LOCAL_GOOS)-$(LOCAL_GOARCH) $(HOME)/bin/$(BIN_NAME)

lint:
	gofmt -s -w .
	golangci-lint run
