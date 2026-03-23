#!/opt/homebrew/bin/bash
set -e

# Shutter API Set deployment 1002 | Monday 23, 2026
private_key=<KEYPERSETMANAGER_DEPLOYER_PRIVATE_KEY>
rpc_url=https://rpc.gnosischain.com
keypers=0x7FfC32AE6270DA85Da831c631Bcc437B99973aCc,0xf211186332B072A1b145F46F91E92C8EEc1Cb49c,0xEbe0BE11161e8aea85733D4ff09De6470E6558Da,0x4B5E2356b666898e101627BdDc518956bcd90a03,0x23d33956940083e0E92Dd608D6E576AfbEcc83a9,0x48A0e1789C82084aE28c179bd5742454f8CD4ed6,0x00D82BAc88c5E60fDAfac7e534A13D0E7F3e145a,0xcc7cd01106951B4809e640873C15363609d2C58e,0x0c8f3E3912F35a59ffddc9Ff1ABB8FafC89b29de,0xcB43AEAA8c029Da18A499dC684FdD250ab3FCd13,0xfc7d75e4bb6D18591cDc1E766CE7cF231bc08fBc
threshold=6
activation_block_number=45298404
keyper_set_manager_address=0xc09B516d87f53ca3C090F85f27A558b221692eb2
key_broadcast_address=0x70bD6A075e8664c9c7Cd15d29Ac6C35dA7D0e41b

contracts_root_dir=$(git rev-parse --show-toplevel)
cd $contracts_root_dir

export PRIVATE_KEY=$private_key
export KEYPER_ADDRESSES=$keypers
export THRESHOLD=$threshold
export ACTIVATION_BLOCK_NUMBER=$activation_block_number
export KEYPERSETMANAGER_ADDRESS=$keyper_set_manager_address
export KEYBROADCAST_ADDRESS=$key_broadcast_address

echo "Running AddKeyperSet Forge script..."

forge script script/AddKeyperSetWithActivationBlock.s.sol \
  --rpc-url "$rpc_url" \
  --broadcast \
  -vvvv