########################################################################################################################
### Makefile for building and testing                                                                                ###
########################################################################################################################
GOOS := linux
GOARCH := amd64
GO111MODULE ?= on
PLATFORM := $(GOOS)_$(GOARCH)
SOURCES := $(shell find . -name "*.go")
all: kepler demo

build/:
	@mkdir -p $@
kepler: build/kepler-cli
build/kepler-cli: build/ $(SOURCES)
	@echo "Building kepler-cli ..."
	@CGO_ENABLED=0 go build -tags release -o build/kepler-cli github.com/extvos/kepler

demo: build/kepler-demo
build/kepler-demo: build/ $(SOURCES)
	@echo "Building kepler-demo ..."
	@CGO_ENABLED=0 go build -tags release -o build/kepler-demo github.com/extvos/kepler-demo

run: demo
	@echo "Starting kepler-demo ..."
	@./build/kepler-demo serve --config demo/demo.yml

clean:
	@echo "Clean up built ..."
	@rm -rf build/