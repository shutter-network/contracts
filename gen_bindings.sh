#!/bin/bash

set -Eeuo pipefail

CONTRACTS=(
    "src/gnosh/Sequencer.sol:Sequencer"
    "src/gnosh/ValidatorRegistry.sol:ValidatorRegistry"
    "src/common/KeyperSetManager.sol:KeyperSetManager"
    "src/common/KeyperSet.sol:KeyperSet"
    "src/common/EonKeyPublish.sol:EonKeyPublish"
    "src/common/KeyBroadcastContract.sol:KeyBroadcastContract"
    "src/common/DKGContract.sol:DKGContract"
    "src/common/ECIESKeyRegistry.sol:ECIESKeyRegistry"
    "src/shop/Inbox.sol:Inbox"
    "src/shutter-service/ShutterRegistry.sol:ShutterRegistry"
    "src/shutter-service/ShutterEventTriggerRegistry.sol:ShutterEventTriggerRegistryV1"
    "src/shutter-service/EventTriggerTestHelper.sol:EventTriggerTestHelper"
)
OUTPUT_DIR="bindings"

mkdir -p "$OUTPUT_DIR"

forge build

for contract_entry in "${CONTRACTS[@]}"; do
    contract_path="${contract_entry%%:*}"
    contract_name="${contract_entry##*:}"

    pkg=$(echo "$contract_name" | tr '[:upper:]' '[:lower:]')
    d="${OUTPUT_DIR}/${pkg}"
    mkdir -p "${d}"

    contract_file=$(basename "$contract_path")
    out_path="out/${contract_file}/${contract_name}.json"

    abigen  \
        --abi <(jq -c '.abi' "$out_path")  \
        --bin <(jq -r '.bytecode.object' "$out_path")  \
        --pkg "${pkg}"  \
        --out "${d}/${pkg}.go"
done
