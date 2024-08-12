package predeploy

import (
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrContractAddressExists = errors.New("contract address already exists")
	ErrContractNameExists    = errors.New("contract name already exists")
)

func MergeShutterPredeploys(predeploy map[string]*common.Address) error {
	if err := ValidatePredeploys(predeploy); err != nil {
		return err
	}
	for name, addr := range Predeploys {
		predeploy[name] = addr
	}
	return nil
}

func ValidatePredeploys(predeploy map[string]*common.Address) error {
	existingAddrs := make(map[common.Address]string)
	for name, addr := range predeploy {
		existingAddrs[*addr] = name
	}

	// TODO: check that the shutter addresses fall in the range of the
	// Optimism predeploy namespace! (not 100% sure, but it needs to have
	// the 0x420000000... prefix [:18] and the remainder has to be unique etc.)

	for addr, name := range PredeploysAddrSet {
		_, ok := predeploy[name]
		if ok {
			return errors.Wrapf(ErrContractNameExists, "contract `%s` at address `%s` already predeployed", name, addr.Hex())
		}

		name, ok = existingAddrs[addr]
		if ok {
			return errors.Wrapf(ErrContractAddressExists, "contract `%s` at address `%s` already predeployed", name, addr.Hex())
		}
	}
	return nil
}
