#!/opt/homebrew/bin/bash
set -e

# Mainnet Gnosis Set Transition Nov 10
private_key=<KEYPERSETMANAGER_DEPLOYER_PRIVATE_KEY>
rpc_url=https://rpc.gnosischain.com
keypers=0x4ba8E8e0061EA7CE9353e7B641A116F1d48d9d46,0xcB43AEAA8c029Da18A499dC684FdD250ab3FCd13,0x46c24A9B4C92Aeb416d732801208ADD353635EC4,0x772213aaD4c6B37d2cc0Fa3B74F9ace5a028B8C8,0x2711d53B3dF060D08AC30C6517D4B8EED50A54B3,0xe03472CCb8e011b7Dfb3343837D75Bf6C9c3324C,0x4B5E2356b666898e101627BdDc518956bcd90a03,0x48A0e1789C82084aE28c179bd5742454f8CD4ed6
threshold=5
activation_block_number=<>
keyper_set_manager_address=0x7C2337f9bFce19d8970661DA50dE8DD7d3D34abb
key_broadcast_address=0x626dB87f9a9aC47070016A50e802dd5974341301

export PRIVATE_KEY=$private_key
export KEYPER_ADDRESSES=$keypers
export THRESHOLD=$threshold
export ACTIVATION_BLOCK_NUMBER=$activation_block_number
export KEYPERSETMANAGER_ADDRESS=$keyper_set_manager_address
export KEYBROADCAST_ADDRESS=$key_broadcast_address

echo "Running AddKeyperSet Forge script..."

contracts_root_dir=../
cd $contracts_root_dir

forge script script/AddKeyperSetWithActivationBlock.s.sol \
  --rpc-url "$rpc_url" \
  --broadcast \
  -vvvv