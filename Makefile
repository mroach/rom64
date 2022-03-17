BIN_NAME := rom64
BUILD_DIR := build
PREFIX ?= /opt/local
BUILD_TIME := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
CURRENT_TAG = $(shell git tag -l --sort=-creatordate 'v*' | head -n1)
CURRENT_SHA = $(shell git rev-parse --short HEAD)

ifeq ($(shell git rev-parse --short $(CURRENT_TAG)),$(CURRENT_SHA))
	VERSION = $(CURRENT_TAG)
else
	VERSION = $(CURRENT_TAG)+$(CURRENT_SHA)
endif

PKGNAME = github.com/mroach/rom64
SETVARS = '$(PKGNAME)/version.BuildTime=$(BUILD_TIME)' \
		  '$(PKGNAME)/version.Version=$(VERSION)' \
		  '$(PKGNAME)/cmd.binName=$(BIN_NAME)'

LDFLAGS = $(addprefix -X , $(SETVARS))

LOCAL_GOOS_ARCH := $(shell go version | cut -d' ' -f4 | tr '/' ' ')
LOCAL_GOOS = $(firstword $(LOCAL_GOOS_ARCH))
LOCAL_GOARCH = $(lastword $(LOCAL_GOOS_ARCH))


dynamic_target = $(subst -, , $@)
derived_os = $(word 2, $(dynamic_target))
derived_arch = $(word 3, $(dynamic_target))

.PHONY: clean fresh install lint release

all: $(BUILD_DIR)/$(BIN_NAME)-linux-amd64 \
	 $(BUILD_DIR)/$(BIN_NAME)-linux-arm64 \
	 $(BUILD_DIR)/$(BIN_NAME)-darwin-amd64 \
	 $(BUILD_DIR)/$(BIN_NAME)-darwin-arm64 \
	 $(BUILD_DIR)/$(BIN_NAME)-windows-amd64

$(BUILD_DIR)/$(BIN_NAME)-%:
	@test -d $(BUILD_DIR) || mkdir $(BUILD_DIR)
	env GOOS=$(derived_os) GOARCH=$(derived_arch) go build -v -o $@ -ldflags "$(LDFLAGS)"

clean:
	rm -f $(BUILD_DIR)/$(BIN_NAME)-*

fresh: clean
	$(MAKE) all

install:
	cp build/$(BIN_NAME)-$(LOCAL_GOOS)-$(LOCAL_GOARCH) $(PREFIX)/bin/$(BIN_NAME)

lint:
	gofmt -s -w .
	golangci-lint run

# Before making a release, create a tag with `git tag vX.Y.Z`
release: clean all
	git push --tags
	gh release create $(CURRENT_TAG) --title $(CURRENT_TAG) --generate-notes
