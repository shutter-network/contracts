#!/bin/bash

CONTRACTS=(
    "Sequencer"
    "ValidatorRegistry"
    "KeyperSetManager"
    "KeyperSet"
    "EonKeyPublish"
    "KeyBroadcastContract"
    "Inbox"
)
OUTPUT_DIR="bindings"
PACKAGE_NAME="bindings"

mkdir -p "$OUTPUT_DIR"

forge build

for contract in "${CONTRACTS[@]}"; do
    pkg=$(echo "$contract" | tr '[:upper:]' '[:lower:]')
    d="${OUTPUT_DIR}/${pkg}"
    mkdir -p "${d}"
    abigen --abi <(jq '.["abi"]' "out/${contract}.sol/${contract}.json") --pkg "${pkg}" --out "${d}/${pkg}.go"
done
