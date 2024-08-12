// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package predeploy

import "github.com/ethereum/go-ethereum/common"

var (
	inbox                    = "0x4200000000000000000000000000000000000066"
	InboxAddr                = common.HexToAddress(inbox)
	keyperSetManager         = "0x4200000000000000000000000000000000000067"
	KeyperSetManagerAddr     = common.HexToAddress(keyperSetManager)
	keyBroadcastContract     = "0x4200000000000000000000000000000000000068"
	KeyBroadcastContractAddr = common.HexToAddress(keyBroadcastContract)

	Predeploys        = make(map[string]*common.Address)
	PredeploysAddrSet = make(map[common.Address]string)
)

func init() {

	Predeploys["Inbox"] = &InboxAddr
	Predeploys["KeyperSetManager"] = &KeyperSetManagerAddr
	Predeploys["KeyBroadcastContract"] = &KeyBroadcastContractAddr

	for name, addr := range Predeploys {
		PredeploysAddrSet[*addr] = name
	}
}
