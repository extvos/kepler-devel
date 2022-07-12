########################################################################################################################
### Makefile for building and testing                                                                                ###
########################################################################################################################
GOOS := linux
GOARCH := amd64
GO111MODULE ?= on
PLATFORM := $(GOOS)_$(GOARCH)

all: kepler demo

build/:
	@mkdir -p $@
kepler: build/kepler-cli
build/kepler-cli: build/
	@echo "Building kepler-cli ..."
	@CGO_ENABLED=0 go build -tags release -o build/kepler-cli github.com/extvos/kepler

demo: build/kepler-demo
build/kepler-demo: build/
	@echo "Building kepler-demo ..."
	@CGO_ENABLED=0 go build -tags release -o build/kepler-demo github.com/extvos/kepler-demo

run: demo
	@echo "Starting kepler-demo ..."
	@./build/kepler-demo serve

clean:
	@echo "Clean up built ..."
	@rm -rf build/