########################################################################################################################
### Makefile for building and testing                                                                                ###
########################################################################################################################
GOOS := linux
GOARCH := amd64
GO111MODULE ?= on
PLATFORM := $(GOOS)_$(GOARCH)
BUILD_OPTS :=
SUFFIX :=
SOURCES := $(shell find . -name "*.go")

ifeq ($(OS), Windows_NT)
	SUFFIX := .exe
	BUILD_OPTS := -buildvcs=false
endif


all: kepler demo

build/:
	@mkdir -p $@
kepler: build/kepler-cli$(SUFFIX)
build/kepler-cli$(SUFFIX): build/ $(SOURCES)
	@echo "Building kepler-cli ..."
	@CGO_ENABLED=0 go build $(BUILD_OPTS) -tags release -o $@ github.com/extvos/kepler

demo: build/kepler-demo$(SUFFIX)
build/kepler-demo$(SUFFIX): build/ $(SOURCES)
	@echo "Building kepler-demo ..."
	@CGO_ENABLED=0 go build  $(BUILD_OPTS) -tags release -o $@ github.com/extvos/kepler-demo

run: demo
	@echo "Starting kepler-demo ..."
	@./build/kepler-demo$(SUFFIX) serve --config demo/demo.yml

clean:
	@echo "Clean up built ..."
	@rm -rf build/