SHELL := /usr/bin/env bash

lower = $(shell echo '$1' | tr '[:upper:]' '[:lower:]')

CONTRACTS := $(basename $(notdir $(wildcard ./src/*.sol)))

all: forge gen
.PHONY: all forge gen

forge:
	forge build

gen:
	go run gen/gen.go -config genconfig.toml -base-path ./

$(CONTRACTS:%=bindings/%.go): forge
	abigen --abi <(jq '.["abi"]' "out/$(basename $(notdir $@)).sol/$(basename $(notdir $@)).json") --pkg bindings --type $(basename $(notdir $@)) --out $@

bindings: $(CONTRACTS:%=bindings/%.go)
