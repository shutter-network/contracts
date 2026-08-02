// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dkgcontract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DkgcontractMetaData contains all meta data concerning the Dkgcontract contract.
var DkgcontractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"phaseLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dkgLeadLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRetries\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperSetManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"keyBroadcastContractAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DKG_LEAD_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RETRIES\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PHASE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentPhase\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIDKGContract.Phase\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cycleLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dkgStart\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasVoted\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyBroadcastContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyperSetManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitAccusation\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"accusedIndices\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitApology\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"accuserIndices\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"polyEvalData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitDealing\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"commitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"polyEvals\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitSuccessVote\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"eonPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"succeeded\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"succeededAtRetry\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voteCount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AccusationSubmitted\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"accusedIndices\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApologySubmitted\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"accuserIndices\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"},{\"name\":\"polyEvalData\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DKGSucceeded\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"eonPublicKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DealingSubmitted\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"commitment\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"polyEvals\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SuccessVoteSubmitted\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"eonPublicKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadySucceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyVoted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAccusation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyEonPublicKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxRetriesExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrays\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotKeyperAtIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongDKGContract\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLengthParameter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroMaxRetries\",\"inputs\":[]}]",
	Bin: "0x610120604052348015610010575f5ffd5b506040516123e03803806123e08339818101604052810190610032919061022f565b5f8567ffffffffffffffff16148061005357505f8467ffffffffffffffff16145b1561008a576040517f038016d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8367ffffffffffffffff16036100cd576040517fb8579f9100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8467ffffffffffffffff1660808167ffffffffffffffff16815250508367ffffffffffffffff1660a08167ffffffffffffffff16815250508267ffffffffffffffff1660c08167ffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff166101008173ffffffffffffffffffffffffffffffffffffffff168152505050505050506102a6565b5f5ffd5b5f67ffffffffffffffff82169050919050565b6101b481610198565b81146101be575f5ffd5b50565b5f815190506101cf816101ab565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101fe826101d5565b9050919050565b61020e816101f4565b8114610218575f5ffd5b50565b5f8151905061022981610205565b92915050565b5f5f5f5f5f60a0868803121561024857610247610194565b5b5f610255888289016101c1565b9550506020610266888289016101c1565b9450506040610277888289016101c1565b93505060606102888882890161021b565b92505060806102998882890161021b565b9150509295509295909350565b60805160a05160c05160e051610100516120b76103295f395f8181610b230152610d1f01525f81816109a801528181610cf801528181610e8d015281816110c0015261122f01525f8181610349015281816105c30152610fa101525f818161046a0152610f4e01525f81816103a901528181610cd30152610d4601526120b75ff3fe608060405234801561000f575f5ffd5b50600436106100fe575f3560e01c806394afc12d11610095578063eac471a011610064578063eac471a0146102ac578063eb05c280146102ca578063f01275b0146102fa578063f48f3d1414610316576100fe565b806394afc12d14610236578063b209e71714610252578063d4b2cedd14610270578063df6406311461028e576100fe565b806346ebbdb0116100d157806346ebbdb01461019c5780637642d0de146101ba57806378a1b9db146101ea57806382b870d31461021a576100fe565b80630a3ec4d61461010257806316b89c3f1461013257806316e16061146101505780633f98009814610180575b5f5ffd5b61011c600480360381019061011791906113ed565b610346565b604051610129919061149e565b60405180910390f35b61013a610468565b60405161014791906114c6565b60405180910390f35b61016a600480360381019061016591906114df565b61048c565b60405161017791906114c6565b60405180910390f35b61019a600480360381019061019591906115c0565b610529565b005b6101a46105c1565b6040516101b191906114c6565b60405180910390f35b6101d460048036038101906101cf91906116d1565b6105e5565b6040516101e1919061173b565b60405180910390f35b61020460048036038101906101ff91906114df565b61061a565b604051610211919061173b565b60405180910390f35b610234600480360381019061022f9190611754565b610667565b005b610250600480360381019061024b919061182d565b610c03565b005b61025a610cd1565b60405161026791906114c6565b60405180910390f35b610278610cf5565b60405161028591906118c0565b60405180910390f35b610296610d1c565b6040516102a391906118c0565b60405180910390f35b6102b4610d43565b6040516102c191906114c6565b60405180910390f35b6102e460048036038101906102df919061190c565b610d76565b6040516102f191906114c6565b60405180910390f35b610314600480360381019061030f919061195c565b610db2565b005b610330600480360381019061032b91906113ed565b610e89565b60405161033d9190611a2b565b60405180910390f35b5f7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168267ffffffffffffffff161061038a575f9050610462565b5f6103958484610e89565b90505f81436103a49190611a71565b90505f7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1690505f8212156103e7575f9350505050610462565b808212156103fb5760019350505050610462565b8060026104089190611ab1565b82121561041b5760029350505050610462565b8060036104289190611ab1565b82121561043b5760039350505050610462565b8060046104489190611ab1565b82121561045b5760049350505050610462565b5f93505050505b92915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f5f5f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1690505f8167ffffffffffffffff1603610514576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050b90611b81565b60405180910390fd5b6001816105219190611b9f565b915050919050565b61053286610f9f565b61053b8761100f565b61054787876001611052565b610550876110bd565b61055a878661122c565b8467ffffffffffffffff168667ffffffffffffffff168867ffffffffffffffff167fa5074728b92791250d48ccdc55266aca7b1cca22f89e1ef2ab91ccd7ae0f337d878787876040516105b09493929190611d8c565b60405180910390a450505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6002602052825f5260405f20602052815f5260405f20602052805f5260405f205f92509250509054906101000a900460ff1681565b5f5f5f5f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1614159050919050565b61067084610f9f565b61067c85856004611052565b610685856110bd565b61068f858461122c565b5f82829050036106cb576040517f59acc21d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610792576040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160025f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f828260405161083e929190611df3565b604051809103902090505f6001805f8967ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8481526020019081526020015f205f9054906101000a900467ffffffffffffffff166108bf9190611e0b565b90508060015f8967ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8481526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508467ffffffffffffffff168667ffffffffffffffff168867ffffffffffffffff167facfc0aa3f6ba36ac0ffda53e00d2cdace2fd4a9e41e2fe3214032794332c37c38787604051610990929190611e46565b60405180910390a46109a18761061a565b610bfa575f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f90f3bed896040518263ffffffff1660e01b81526004016109ff91906114c6565b602060405180830381865afa158015610a1a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a3e9190611e7c565b90505f8173ffffffffffffffffffffffffffffffffffffffff1663e75235b86040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a8a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aae9190611ebb565b90508067ffffffffffffffff168367ffffffffffffffff1610610bf757600188610ad89190611e0b565b5f5f8b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663daade8e88a88886040518463ffffffff1660e01b8152600401610b7e93929190611ee6565b5f604051808303815f87803b158015610b95575f5ffd5b505af1925050508015610ba6575060015b508767ffffffffffffffff168967ffffffffffffffff167fe6c9bd3daa50b9218615605cdf333410d405b7f3b8dbeee5a2373769679aa9f58888604051610bee929190611e46565b60405180910390a35b50505b50505050505050565b610c0c84610f9f565b610c158561100f565b610c2185856002611052565b610c2a856110bd565b610c34858461122c565b5f8282905003610c70576040517f2eeffe0b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8267ffffffffffffffff168467ffffffffffffffff168667ffffffffffffffff167fb5cc05e740a503c13e1b7ac3ad2d5abf41b91d2e5fca79455d34e515fb70849a8585604051610cc2929190611fd2565b60405180910390a45050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f7f0000000000000000000000000000000000000000000000000000000000000000905090565b5f7f0000000000000000000000000000000000000000000000000000000000000000905090565b5f7f00000000000000000000000000000000000000000000000000000000000000006004610d719190611ff4565b905090565b6001602052825f5260405f20602052815f5260405f20602052805f5260405f205f92509250509054906101000a900467ffffffffffffffff1681565b610dbb86610f9f565b610dc48761100f565b610dd087876003611052565b610dd9876110bd565b610de3878661122c565b818190508484905014610e22576040517fa121188700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8467ffffffffffffffff168667ffffffffffffffff168867ffffffffffffffff167f62729f0820cf3ab60fbd63a51cdeae0554d0ab24f843b86bd0d2591594d6913687878787604051610e789493929190612030565b60405180910390a450505050505050565b5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663636df979856040518263ffffffff1660e01b8152600401610ee491906114c6565b602060405180830381865afa158015610eff573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f239190611ebb565b9050610f2d610d43565b67ffffffffffffffff168367ffffffffffffffff16610f4c9190611ab1565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168267ffffffffffffffff16610f8c9190611a71565b610f969190612069565b91505092915050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168167ffffffffffffffff161061100c576040517f12f6147700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b6110188161061a565b1561104f576040517fe0f5ec9d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b8060048111156110655761106461142b565b5b61106f8484610346565b60048111156110815761108061142b565b5b146110b8576040517fe2586bcc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f90f3bed836040518263ffffffff1660e01b815260040161111791906114c6565b602060405180830381865afa158015611132573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111569190611e7c565b90503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff166266f0a86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111b7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111db9190611e7c565b73ffffffffffffffffffffffffffffffffffffffff1614611228576040517ff62f18e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f90f3bed846040518263ffffffff1660e01b815260040161128691906114c6565b602060405180830381865afa1580156112a1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112c59190611e7c565b90503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16632e8e6cad846040518263ffffffff1660e01b815260040161131791906114c6565b602060405180830381865afa158015611332573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113569190611e7c565b73ffffffffffffffffffffffffffffffffffffffff16146113a3576040517f7bf9580300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff82169050919050565b6113cc816113b0565b81146113d6575f5ffd5b50565b5f813590506113e7816113c3565b92915050565b5f5f60408385031215611403576114026113a8565b5b5f611410858286016113d9565b9250506020611421858286016113d9565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600581106114695761146861142b565b5b50565b5f81905061147982611458565b919050565b5f6114888261146c565b9050919050565b6114988161147e565b82525050565b5f6020820190506114b15f83018461148f565b92915050565b6114c0816113b0565b82525050565b5f6020820190506114d95f8301846114b7565b92915050565b5f602082840312156114f4576114f36113a8565b5b5f611501848285016113d9565b91505092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261152b5761152a61150a565b5b8235905067ffffffffffffffff8111156115485761154761150e565b5b60208301915083600182028301111561156457611563611512565b5b9250929050565b5f5f83601f8401126115805761157f61150a565b5b8235905067ffffffffffffffff81111561159d5761159c61150e565b5b6020830191508360208202830111156115b9576115b8611512565b5b9250929050565b5f5f5f5f5f5f5f60a0888a0312156115db576115da6113a8565b5b5f6115e88a828b016113d9565b97505060206115f98a828b016113d9565b965050604061160a8a828b016113d9565b955050606088013567ffffffffffffffff81111561162b5761162a6113ac565b5b6116378a828b01611516565b9450945050608088013567ffffffffffffffff81111561165a576116596113ac565b5b6116668a828b0161156b565b925092505092959891949750929550565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6116a082611677565b9050919050565b6116b081611696565b81146116ba575f5ffd5b50565b5f813590506116cb816116a7565b92915050565b5f5f5f606084860312156116e8576116e76113a8565b5b5f6116f5868287016113d9565b9350506020611706868287016113d9565b9250506040611717868287016116bd565b9150509250925092565b5f8115159050919050565b61173581611721565b82525050565b5f60208201905061174e5f83018461172c565b92915050565b5f5f5f5f5f6080868803121561176d5761176c6113a8565b5b5f61177a888289016113d9565b955050602061178b888289016113d9565b945050604061179c888289016113d9565b935050606086013567ffffffffffffffff8111156117bd576117bc6113ac565b5b6117c988828901611516565b92509250509295509295909350565b5f5f83601f8401126117ed576117ec61150a565b5b8235905067ffffffffffffffff81111561180a5761180961150e565b5b60208301915083602082028301111561182657611825611512565b5b9250929050565b5f5f5f5f5f60808688031215611846576118456113a8565b5b5f611853888289016113d9565b9550506020611864888289016113d9565b9450506040611875888289016113d9565b935050606086013567ffffffffffffffff811115611896576118956113ac565b5b6118a2888289016117d8565b92509250509295509295909350565b6118ba81611696565b82525050565b5f6020820190506118d35f8301846118b1565b92915050565b5f819050919050565b6118eb816118d9565b81146118f5575f5ffd5b50565b5f81359050611906816118e2565b92915050565b5f5f5f60608486031215611923576119226113a8565b5b5f611930868287016113d9565b9350506020611941868287016113d9565b9250506040611952868287016118f8565b9150509250925092565b5f5f5f5f5f5f5f60a0888a031215611977576119766113a8565b5b5f6119848a828b016113d9565b97505060206119958a828b016113d9565b96505060406119a68a828b016113d9565b955050606088013567ffffffffffffffff8111156119c7576119c66113ac565b5b6119d38a828b016117d8565b9450945050608088013567ffffffffffffffff8111156119f6576119f56113ac565b5b611a028a828b0161156b565b925092505092959891949750929550565b5f819050919050565b611a2581611a13565b82525050565b5f602082019050611a3e5f830184611a1c565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611a7b82611a13565b9150611a8683611a13565b925082820390508181125f8412168282135f851215161715611aab57611aaa611a44565b5b92915050565b5f611abb82611a13565b9150611ac683611a13565b9250828202611ad481611a13565b91507f800000000000000000000000000000000000000000000000000000000000000084145f84121615611b0b57611b0a611a44565b5b8282058414831517611b2057611b1f611a44565b5b5092915050565b5f82825260208201905092915050565b7f6e6f7420737563636565646564000000000000000000000000000000000000005f82015250565b5f611b6b600d83611b27565b9150611b7682611b37565b602082019050919050565b5f6020820190508181035f830152611b9881611b5f565b9050919050565b5f611ba9826113b0565b9150611bb4836113b0565b9250828203905067ffffffffffffffff811115611bd457611bd3611a44565b5b92915050565b5f82825260208201905092915050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f611c138385611bda565b9350611c20838584611bea565b611c2983611bf8565b840190509392505050565b5f82825260208201905092915050565b5f819050919050565b5f82825260208201905092915050565b5f611c688385611c4d565b9350611c75838584611bea565b611c7e83611bf8565b840190509392505050565b5f611c95848484611c5d565b90509392505050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83356001602003843603038112611cc657611cc5611ca6565b5b83810192508235915060208301925067ffffffffffffffff821115611cee57611ced611c9e565b5b600182023603831315611d0457611d03611ca2565b5b509250929050565b5f602082019050919050565b5f611d238385611c34565b935083602084028501611d3584611c44565b805f5b87811015611d7a578484038952611d4f8284611caa565b611d5a868284611c89565b9550611d6584611d0c565b935060208b019a505050600181019050611d38565b50829750879450505050509392505050565b5f6040820190508181035f830152611da5818688611c08565b90508181036020830152611dba818486611d18565b905095945050505050565b5f81905092915050565b5f611dda8385611dc5565b9350611de7838584611bea565b82840190509392505050565b5f611dff828486611dcf565b91508190509392505050565b5f611e15826113b0565b9150611e20836113b0565b9250828201905067ffffffffffffffff811115611e4057611e3f611a44565b5b92915050565b5f6020820190508181035f830152611e5f818486611c08565b90509392505050565b5f81519050611e76816116a7565b92915050565b5f60208284031215611e9157611e906113a8565b5b5f611e9e84828501611e68565b91505092915050565b5f81519050611eb5816113c3565b92915050565b5f60208284031215611ed057611ecf6113a8565b5b5f611edd84828501611ea7565b91505092915050565b5f604082019050611ef95f8301866114b7565b8181036020830152611f0c818486611c08565b9050949350505050565b5f82825260208201905092915050565b5f819050919050565b611f38816113b0565b82525050565b5f611f498383611f2f565b60208301905092915050565b5f611f6360208401846113d9565b905092915050565b5f602082019050919050565b5f611f828385611f16565b9350611f8d82611f26565b805f5b85811015611fc557611fa28284611f55565b611fac8882611f3e565b9750611fb783611f6b565b925050600181019050611f90565b5085925050509392505050565b5f6020820190508181035f830152611feb818486611f77565b90509392505050565b5f611ffe826113b0565b9150612009836113b0565b9250828202612017816113b0565b915080821461202957612028611a44565b5b5092915050565b5f6040820190508181035f830152612049818688611f77565b9050818103602083015261205e818486611d18565b905095945050505050565b5f61207382611a13565b915061207e83611a13565b92508282019050828112155f8312168382125f8412151617156120a4576120a3611a44565b5b9291505056fea164736f6c634300081c000a",
}

// DkgcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use DkgcontractMetaData.ABI instead.
var DkgcontractABI = DkgcontractMetaData.ABI

// DkgcontractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DkgcontractMetaData.Bin instead.
var DkgcontractBin = DkgcontractMetaData.Bin

// DeployDkgcontract deploys a new Ethereum contract, binding an instance of Dkgcontract to it.
func DeployDkgcontract(auth *bind.TransactOpts, backend bind.ContractBackend, phaseLength uint64, dkgLeadLength uint64, maxRetries uint64, keyperSetManagerAddress common.Address, keyBroadcastContractAddress common.Address) (common.Address, *types.Transaction, *Dkgcontract, error) {
	parsed, err := DkgcontractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DkgcontractBin), backend, phaseLength, dkgLeadLength, maxRetries, keyperSetManagerAddress, keyBroadcastContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dkgcontract{DkgcontractCaller: DkgcontractCaller{contract: contract}, DkgcontractTransactor: DkgcontractTransactor{contract: contract}, DkgcontractFilterer: DkgcontractFilterer{contract: contract}}, nil
}

// Dkgcontract is an auto generated Go binding around an Ethereum contract.
type Dkgcontract struct {
	DkgcontractCaller     // Read-only binding to the contract
	DkgcontractTransactor // Write-only binding to the contract
	DkgcontractFilterer   // Log filterer for contract events
}

// DkgcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DkgcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DkgcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DkgcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DkgcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DkgcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DkgcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DkgcontractSession struct {
	Contract     *Dkgcontract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DkgcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DkgcontractCallerSession struct {
	Contract *DkgcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DkgcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DkgcontractTransactorSession struct {
	Contract     *DkgcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DkgcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DkgcontractRaw struct {
	Contract *Dkgcontract // Generic contract binding to access the raw methods on
}

// DkgcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DkgcontractCallerRaw struct {
	Contract *DkgcontractCaller // Generic read-only contract binding to access the raw methods on
}

// DkgcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DkgcontractTransactorRaw struct {
	Contract *DkgcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDkgcontract creates a new instance of Dkgcontract, bound to a specific deployed contract.
func NewDkgcontract(address common.Address, backend bind.ContractBackend) (*Dkgcontract, error) {
	contract, err := bindDkgcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dkgcontract{DkgcontractCaller: DkgcontractCaller{contract: contract}, DkgcontractTransactor: DkgcontractTransactor{contract: contract}, DkgcontractFilterer: DkgcontractFilterer{contract: contract}}, nil
}

// NewDkgcontractCaller creates a new read-only instance of Dkgcontract, bound to a specific deployed contract.
func NewDkgcontractCaller(address common.Address, caller bind.ContractCaller) (*DkgcontractCaller, error) {
	contract, err := bindDkgcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DkgcontractCaller{contract: contract}, nil
}

// NewDkgcontractTransactor creates a new write-only instance of Dkgcontract, bound to a specific deployed contract.
func NewDkgcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*DkgcontractTransactor, error) {
	contract, err := bindDkgcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DkgcontractTransactor{contract: contract}, nil
}

// NewDkgcontractFilterer creates a new log filterer instance of Dkgcontract, bound to a specific deployed contract.
func NewDkgcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*DkgcontractFilterer, error) {
	contract, err := bindDkgcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DkgcontractFilterer{contract: contract}, nil
}

// bindDkgcontract binds a generic wrapper to an already deployed contract.
func bindDkgcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DkgcontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dkgcontract *DkgcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dkgcontract.Contract.DkgcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dkgcontract *DkgcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dkgcontract.Contract.DkgcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dkgcontract *DkgcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dkgcontract.Contract.DkgcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dkgcontract *DkgcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dkgcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dkgcontract *DkgcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dkgcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dkgcontract *DkgcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dkgcontract.Contract.contract.Transact(opts, method, params...)
}

// DKGLEADLENGTH is a free data retrieval call binding the contract method 0x16b89c3f.
//
// Solidity: function DKG_LEAD_LENGTH() view returns(uint64)
func (_Dkgcontract *DkgcontractCaller) DKGLEADLENGTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "DKG_LEAD_LENGTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DKGLEADLENGTH is a free data retrieval call binding the contract method 0x16b89c3f.
//
// Solidity: function DKG_LEAD_LENGTH() view returns(uint64)
func (_Dkgcontract *DkgcontractSession) DKGLEADLENGTH() (uint64, error) {
	return _Dkgcontract.Contract.DKGLEADLENGTH(&_Dkgcontract.CallOpts)
}

// DKGLEADLENGTH is a free data retrieval call binding the contract method 0x16b89c3f.
//
// Solidity: function DKG_LEAD_LENGTH() view returns(uint64)
func (_Dkgcontract *DkgcontractCallerSession) DKGLEADLENGTH() (uint64, error) {
	return _Dkgcontract.Contract.DKGLEADLENGTH(&_Dkgcontract.CallOpts)
}

// MAXRETRIES is a free data retrieval call binding the contract method 0x46ebbdb0.
//
// Solidity: function MAX_RETRIES() view returns(uint64)
func (_Dkgcontract *DkgcontractCaller) MAXRETRIES(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "MAX_RETRIES")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXRETRIES is a free data retrieval call binding the contract method 0x46ebbdb0.
//
// Solidity: function MAX_RETRIES() view returns(uint64)
func (_Dkgcontract *DkgcontractSession) MAXRETRIES() (uint64, error) {
	return _Dkgcontract.Contract.MAXRETRIES(&_Dkgcontract.CallOpts)
}

// MAXRETRIES is a free data retrieval call binding the contract method 0x46ebbdb0.
//
// Solidity: function MAX_RETRIES() view returns(uint64)
func (_Dkgcontract *DkgcontractCallerSession) MAXRETRIES() (uint64, error) {
	return _Dkgcontract.Contract.MAXRETRIES(&_Dkgcontract.CallOpts)
}

// PHASELENGTH is a free data retrieval call binding the contract method 0xb209e717.
//
// Solidity: function PHASE_LENGTH() view returns(uint64)
func (_Dkgcontract *DkgcontractCaller) PHASELENGTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "PHASE_LENGTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PHASELENGTH is a free data retrieval call binding the contract method 0xb209e717.
//
// Solidity: function PHASE_LENGTH() view returns(uint64)
func (_Dkgcontract *DkgcontractSession) PHASELENGTH() (uint64, error) {
	return _Dkgcontract.Contract.PHASELENGTH(&_Dkgcontract.CallOpts)
}

// PHASELENGTH is a free data retrieval call binding the contract method 0xb209e717.
//
// Solidity: function PHASE_LENGTH() view returns(uint64)
func (_Dkgcontract *DkgcontractCallerSession) PHASELENGTH() (uint64, error) {
	return _Dkgcontract.Contract.PHASELENGTH(&_Dkgcontract.CallOpts)
}

// CurrentPhase is a free data retrieval call binding the contract method 0x0a3ec4d6.
//
// Solidity: function currentPhase(uint64 keyperSetIndex, uint64 retryCounter) view returns(uint8)
func (_Dkgcontract *DkgcontractCaller) CurrentPhase(opts *bind.CallOpts, keyperSetIndex uint64, retryCounter uint64) (uint8, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "currentPhase", keyperSetIndex, retryCounter)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CurrentPhase is a free data retrieval call binding the contract method 0x0a3ec4d6.
//
// Solidity: function currentPhase(uint64 keyperSetIndex, uint64 retryCounter) view returns(uint8)
func (_Dkgcontract *DkgcontractSession) CurrentPhase(keyperSetIndex uint64, retryCounter uint64) (uint8, error) {
	return _Dkgcontract.Contract.CurrentPhase(&_Dkgcontract.CallOpts, keyperSetIndex, retryCounter)
}

// CurrentPhase is a free data retrieval call binding the contract method 0x0a3ec4d6.
//
// Solidity: function currentPhase(uint64 keyperSetIndex, uint64 retryCounter) view returns(uint8)
func (_Dkgcontract *DkgcontractCallerSession) CurrentPhase(keyperSetIndex uint64, retryCounter uint64) (uint8, error) {
	return _Dkgcontract.Contract.CurrentPhase(&_Dkgcontract.CallOpts, keyperSetIndex, retryCounter)
}

// CycleLength is a free data retrieval call binding the contract method 0xeac471a0.
//
// Solidity: function cycleLength() view returns(uint64)
func (_Dkgcontract *DkgcontractCaller) CycleLength(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "cycleLength")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CycleLength is a free data retrieval call binding the contract method 0xeac471a0.
//
// Solidity: function cycleLength() view returns(uint64)
func (_Dkgcontract *DkgcontractSession) CycleLength() (uint64, error) {
	return _Dkgcontract.Contract.CycleLength(&_Dkgcontract.CallOpts)
}

// CycleLength is a free data retrieval call binding the contract method 0xeac471a0.
//
// Solidity: function cycleLength() view returns(uint64)
func (_Dkgcontract *DkgcontractCallerSession) CycleLength() (uint64, error) {
	return _Dkgcontract.Contract.CycleLength(&_Dkgcontract.CallOpts)
}

// DkgStart is a free data retrieval call binding the contract method 0xf48f3d14.
//
// Solidity: function dkgStart(uint64 keyperSetIndex, uint64 retryCounter) view returns(int256)
func (_Dkgcontract *DkgcontractCaller) DkgStart(opts *bind.CallOpts, keyperSetIndex uint64, retryCounter uint64) (*big.Int, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "dkgStart", keyperSetIndex, retryCounter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DkgStart is a free data retrieval call binding the contract method 0xf48f3d14.
//
// Solidity: function dkgStart(uint64 keyperSetIndex, uint64 retryCounter) view returns(int256)
func (_Dkgcontract *DkgcontractSession) DkgStart(keyperSetIndex uint64, retryCounter uint64) (*big.Int, error) {
	return _Dkgcontract.Contract.DkgStart(&_Dkgcontract.CallOpts, keyperSetIndex, retryCounter)
}

// DkgStart is a free data retrieval call binding the contract method 0xf48f3d14.
//
// Solidity: function dkgStart(uint64 keyperSetIndex, uint64 retryCounter) view returns(int256)
func (_Dkgcontract *DkgcontractCallerSession) DkgStart(keyperSetIndex uint64, retryCounter uint64) (*big.Int, error) {
	return _Dkgcontract.Contract.DkgStart(&_Dkgcontract.CallOpts, keyperSetIndex, retryCounter)
}

// HasVoted is a free data retrieval call binding the contract method 0x7642d0de.
//
// Solidity: function hasVoted(uint64 , uint64 , address ) view returns(bool)
func (_Dkgcontract *DkgcontractCaller) HasVoted(opts *bind.CallOpts, arg0 uint64, arg1 uint64, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "hasVoted", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x7642d0de.
//
// Solidity: function hasVoted(uint64 , uint64 , address ) view returns(bool)
func (_Dkgcontract *DkgcontractSession) HasVoted(arg0 uint64, arg1 uint64, arg2 common.Address) (bool, error) {
	return _Dkgcontract.Contract.HasVoted(&_Dkgcontract.CallOpts, arg0, arg1, arg2)
}

// HasVoted is a free data retrieval call binding the contract method 0x7642d0de.
//
// Solidity: function hasVoted(uint64 , uint64 , address ) view returns(bool)
func (_Dkgcontract *DkgcontractCallerSession) HasVoted(arg0 uint64, arg1 uint64, arg2 common.Address) (bool, error) {
	return _Dkgcontract.Contract.HasVoted(&_Dkgcontract.CallOpts, arg0, arg1, arg2)
}

// KeyBroadcastContract is a free data retrieval call binding the contract method 0xdf640631.
//
// Solidity: function keyBroadcastContract() view returns(address)
func (_Dkgcontract *DkgcontractCaller) KeyBroadcastContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "keyBroadcastContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyBroadcastContract is a free data retrieval call binding the contract method 0xdf640631.
//
// Solidity: function keyBroadcastContract() view returns(address)
func (_Dkgcontract *DkgcontractSession) KeyBroadcastContract() (common.Address, error) {
	return _Dkgcontract.Contract.KeyBroadcastContract(&_Dkgcontract.CallOpts)
}

// KeyBroadcastContract is a free data retrieval call binding the contract method 0xdf640631.
//
// Solidity: function keyBroadcastContract() view returns(address)
func (_Dkgcontract *DkgcontractCallerSession) KeyBroadcastContract() (common.Address, error) {
	return _Dkgcontract.Contract.KeyBroadcastContract(&_Dkgcontract.CallOpts)
}

// KeyperSetManager is a free data retrieval call binding the contract method 0xd4b2cedd.
//
// Solidity: function keyperSetManager() view returns(address)
func (_Dkgcontract *DkgcontractCaller) KeyperSetManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "keyperSetManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyperSetManager is a free data retrieval call binding the contract method 0xd4b2cedd.
//
// Solidity: function keyperSetManager() view returns(address)
func (_Dkgcontract *DkgcontractSession) KeyperSetManager() (common.Address, error) {
	return _Dkgcontract.Contract.KeyperSetManager(&_Dkgcontract.CallOpts)
}

// KeyperSetManager is a free data retrieval call binding the contract method 0xd4b2cedd.
//
// Solidity: function keyperSetManager() view returns(address)
func (_Dkgcontract *DkgcontractCallerSession) KeyperSetManager() (common.Address, error) {
	return _Dkgcontract.Contract.KeyperSetManager(&_Dkgcontract.CallOpts)
}

// Succeeded is a free data retrieval call binding the contract method 0x78a1b9db.
//
// Solidity: function succeeded(uint64 keyperSetIndex) view returns(bool)
func (_Dkgcontract *DkgcontractCaller) Succeeded(opts *bind.CallOpts, keyperSetIndex uint64) (bool, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "succeeded", keyperSetIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Succeeded is a free data retrieval call binding the contract method 0x78a1b9db.
//
// Solidity: function succeeded(uint64 keyperSetIndex) view returns(bool)
func (_Dkgcontract *DkgcontractSession) Succeeded(keyperSetIndex uint64) (bool, error) {
	return _Dkgcontract.Contract.Succeeded(&_Dkgcontract.CallOpts, keyperSetIndex)
}

// Succeeded is a free data retrieval call binding the contract method 0x78a1b9db.
//
// Solidity: function succeeded(uint64 keyperSetIndex) view returns(bool)
func (_Dkgcontract *DkgcontractCallerSession) Succeeded(keyperSetIndex uint64) (bool, error) {
	return _Dkgcontract.Contract.Succeeded(&_Dkgcontract.CallOpts, keyperSetIndex)
}

// SucceededAtRetry is a free data retrieval call binding the contract method 0x16e16061.
//
// Solidity: function succeededAtRetry(uint64 keyperSetIndex) view returns(uint64)
func (_Dkgcontract *DkgcontractCaller) SucceededAtRetry(opts *bind.CallOpts, keyperSetIndex uint64) (uint64, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "succeededAtRetry", keyperSetIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SucceededAtRetry is a free data retrieval call binding the contract method 0x16e16061.
//
// Solidity: function succeededAtRetry(uint64 keyperSetIndex) view returns(uint64)
func (_Dkgcontract *DkgcontractSession) SucceededAtRetry(keyperSetIndex uint64) (uint64, error) {
	return _Dkgcontract.Contract.SucceededAtRetry(&_Dkgcontract.CallOpts, keyperSetIndex)
}

// SucceededAtRetry is a free data retrieval call binding the contract method 0x16e16061.
//
// Solidity: function succeededAtRetry(uint64 keyperSetIndex) view returns(uint64)
func (_Dkgcontract *DkgcontractCallerSession) SucceededAtRetry(keyperSetIndex uint64) (uint64, error) {
	return _Dkgcontract.Contract.SucceededAtRetry(&_Dkgcontract.CallOpts, keyperSetIndex)
}

// VoteCount is a free data retrieval call binding the contract method 0xeb05c280.
//
// Solidity: function voteCount(uint64 , uint64 , bytes32 ) view returns(uint64)
func (_Dkgcontract *DkgcontractCaller) VoteCount(opts *bind.CallOpts, arg0 uint64, arg1 uint64, arg2 [32]byte) (uint64, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "voteCount", arg0, arg1, arg2)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VoteCount is a free data retrieval call binding the contract method 0xeb05c280.
//
// Solidity: function voteCount(uint64 , uint64 , bytes32 ) view returns(uint64)
func (_Dkgcontract *DkgcontractSession) VoteCount(arg0 uint64, arg1 uint64, arg2 [32]byte) (uint64, error) {
	return _Dkgcontract.Contract.VoteCount(&_Dkgcontract.CallOpts, arg0, arg1, arg2)
}

// VoteCount is a free data retrieval call binding the contract method 0xeb05c280.
//
// Solidity: function voteCount(uint64 , uint64 , bytes32 ) view returns(uint64)
func (_Dkgcontract *DkgcontractCallerSession) VoteCount(arg0 uint64, arg1 uint64, arg2 [32]byte) (uint64, error) {
	return _Dkgcontract.Contract.VoteCount(&_Dkgcontract.CallOpts, arg0, arg1, arg2)
}

// SubmitAccusation is a paid mutator transaction binding the contract method 0x94afc12d.
//
// Solidity: function submitAccusation(uint64 keyperSetIndex, uint64 retryCounter, uint64 keyperIndex, uint64[] accusedIndices) returns()
func (_Dkgcontract *DkgcontractTransactor) SubmitAccusation(opts *bind.TransactOpts, keyperSetIndex uint64, retryCounter uint64, keyperIndex uint64, accusedIndices []uint64) (*types.Transaction, error) {
	return _Dkgcontract.contract.Transact(opts, "submitAccusation", keyperSetIndex, retryCounter, keyperIndex, accusedIndices)
}

// SubmitAccusation is a paid mutator transaction binding the contract method 0x94afc12d.
//
// Solidity: function submitAccusation(uint64 keyperSetIndex, uint64 retryCounter, uint64 keyperIndex, uint64[] accusedIndices) returns()
func (_Dkgcontract *DkgcontractSession) SubmitAccusation(keyperSetIndex uint64, retryCounter uint64, keyperIndex uint64, accusedIndices []uint64) (*types.Transaction, error) {
	return _Dkgcontract.Contract.SubmitAccusation(&_Dkgcontract.TransactOpts, keyperSetIndex, retryCounter, keyperIndex, accusedIndices)
}

// SubmitAccusation is a paid mutator transaction binding the contract method 0x94afc12d.
//
// Solidity: function submitAccusation(uint64 keyperSetIndex, uint64 retryCounter, uint64 keyperIndex, uint64[] accusedIndices) returns()
func (_Dkgcontract *DkgcontractTransactorSession) SubmitAccusation(keyperSetIndex uint64, retryCounter uint64, keyperIndex uint64, accusedIndices []uint64) (*types.Transaction, error) {
	return _Dkgcontract.Contract.SubmitAccusation(&_Dkgcontract.TransactOpts, keyperSetIndex, retryCounter, keyperIndex, accusedIndices)
}

// SubmitApology is a paid mutator transaction binding the contract method 0xf01275b0.
//
// Solidity: function submitApology(uint64 keyperSetIndex, uint64 retryCounter, uint64 keyperIndex, uint64[] accuserIndices, bytes[] polyEvalData) returns()
func (_Dkgcontract *DkgcontractTransactor) SubmitApology(opts *bind.TransactOpts, keyperSetIndex uint64, retryCounter uint64, keyperIndex uint64, accuserIndices []uint64, polyEvalData [][]byte) (*types.Transaction, error) {
	return _Dkgcontract.contract.Transact(opts, "submitApology", keyperSetIndex, retryCounter, keyperIndex, accuserIndices, polyEvalData)
}

// SubmitApology is a paid mutator transaction binding the contract method 0xf01275b0.
//
// Solidity: function submitApology(uint64 keyperSetIndex, uint64 retryCounter, uint64 keyperIndex, uint64[] accuserIndices, bytes[] polyEvalData) returns()
func (_Dkgcontract *DkgcontractSession) SubmitApology(keyperSetIndex uint64, retryCounter uint64, keyperIndex uint64, accuserIndices []uint64, polyEvalData [][]byte) (*types.Transaction, error) {
	return _Dkgcontract.Contract.SubmitApology(&_Dkgcontract.TransactOpts, keyperSetIndex, retryCounter, keyperIndex, accuserIndices, polyEvalData)
}

// SubmitApology is a paid mutator transaction binding the contract method 0xf01275b0.
//
// Solidity: function submitApology(uint64 keyperSetIndex, uint64 retryCounter, uint64 keyperIndex, uint64[] accuserIndices, bytes[] polyEvalData) returns()
func (_Dkgcontract *DkgcontractTransactorSession) SubmitApology(keyperSetIndex uint64, retryCounter uint64, keyperIndex uint64, accuserIndices []uint64, polyEvalData [][]byte) (*types.Transaction, error) {
	return _Dkgcontract.Contract.SubmitApology(&_Dkgcontract.TransactOpts, keyperSetIndex, retryCounter, keyperIndex, accuserIndices, polyEvalData)
}

// SubmitDealing is a paid mutator transaction binding the contract method 0x3f980098.
//
// Solidity: function submitDealing(uint64 keyperSetIndex, uint64 retryCounter, uint64 keyperIndex, bytes commitment, bytes[] polyEvals) returns()
func (_Dkgcontract *DkgcontractTransactor) SubmitDealing(opts *bind.TransactOpts, keyperSetIndex uint64, retryCounter uint64, keyperIndex uint64, commitment []byte, polyEvals [][]byte) (*types.Transaction, error) {
	return _Dkgcontract.contract.Transact(opts, "submitDealing", keyperSetIndex, retryCounter, keyperIndex, commitment, polyEvals)
}

// SubmitDealing is a paid mutator transaction binding the contract method 0x3f980098.
//
// Solidity: function submitDealing(uint64 keyperSetIndex, uint64 retryCounter, uint64 keyperIndex, bytes commitment, bytes[] polyEvals) returns()
func (_Dkgcontract *DkgcontractSession) SubmitDealing(keyperSetIndex uint64, retryCounter uint64, keyperIndex uint64, commitment []byte, polyEvals [][]byte) (*types.Transaction, error) {
	return _Dkgcontract.Contract.SubmitDealing(&_Dkgcontract.TransactOpts, keyperSetIndex, retryCounter, keyperIndex, commitment, polyEvals)
}

// SubmitDealing is a paid mutator transaction binding the contract method 0x3f980098.
//
// Solidity: function submitDealing(uint64 keyperSetIndex, uint64 retryCounter, uint64 keyperIndex, bytes commitment, bytes[] polyEvals) returns()
func (_Dkgcontract *DkgcontractTransactorSession) SubmitDealing(keyperSetIndex uint64, retryCounter uint64, keyperIndex uint64, commitment []byte, polyEvals [][]byte) (*types.Transaction, error) {
	return _Dkgcontract.Contract.SubmitDealing(&_Dkgcontract.TransactOpts, keyperSetIndex, retryCounter, keyperIndex, commitment, polyEvals)
}

// SubmitSuccessVote is a paid mutator transaction binding the contract method 0x82b870d3.
//
// Solidity: function submitSuccessVote(uint64 keyperSetIndex, uint64 retryCounter, uint64 keyperIndex, bytes eonPublicKey) returns()
func (_Dkgcontract *DkgcontractTransactor) SubmitSuccessVote(opts *bind.TransactOpts, keyperSetIndex uint64, retryCounter uint64, keyperIndex uint64, eonPublicKey []byte) (*types.Transaction, error) {
	return _Dkgcontract.contract.Transact(opts, "submitSuccessVote", keyperSetIndex, retryCounter, keyperIndex, eonPublicKey)
}

// SubmitSuccessVote is a paid mutator transaction binding the contract method 0x82b870d3.
//
// Solidity: function submitSuccessVote(uint64 keyperSetIndex, uint64 retryCounter, uint64 keyperIndex, bytes eonPublicKey) returns()
func (_Dkgcontract *DkgcontractSession) SubmitSuccessVote(keyperSetIndex uint64, retryCounter uint64, keyperIndex uint64, eonPublicKey []byte) (*types.Transaction, error) {
	return _Dkgcontract.Contract.SubmitSuccessVote(&_Dkgcontract.TransactOpts, keyperSetIndex, retryCounter, keyperIndex, eonPublicKey)
}

// SubmitSuccessVote is a paid mutator transaction binding the contract method 0x82b870d3.
//
// Solidity: function submitSuccessVote(uint64 keyperSetIndex, uint64 retryCounter, uint64 keyperIndex, bytes eonPublicKey) returns()
func (_Dkgcontract *DkgcontractTransactorSession) SubmitSuccessVote(keyperSetIndex uint64, retryCounter uint64, keyperIndex uint64, eonPublicKey []byte) (*types.Transaction, error) {
	return _Dkgcontract.Contract.SubmitSuccessVote(&_Dkgcontract.TransactOpts, keyperSetIndex, retryCounter, keyperIndex, eonPublicKey)
}

// DkgcontractAccusationSubmittedIterator is returned from FilterAccusationSubmitted and is used to iterate over the raw logs and unpacked data for AccusationSubmitted events raised by the Dkgcontract contract.
type DkgcontractAccusationSubmittedIterator struct {
	Event *DkgcontractAccusationSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DkgcontractAccusationSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DkgcontractAccusationSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DkgcontractAccusationSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DkgcontractAccusationSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DkgcontractAccusationSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DkgcontractAccusationSubmitted represents a AccusationSubmitted event raised by the Dkgcontract contract.
type DkgcontractAccusationSubmitted struct {
	KeyperSetIndex uint64
	RetryCounter   uint64
	KeyperIndex    uint64
	AccusedIndices []uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccusationSubmitted is a free log retrieval operation binding the contract event 0xb5cc05e740a503c13e1b7ac3ad2d5abf41b91d2e5fca79455d34e515fb70849a.
//
// Solidity: event AccusationSubmitted(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, uint64 indexed keyperIndex, uint64[] accusedIndices)
func (_Dkgcontract *DkgcontractFilterer) FilterAccusationSubmitted(opts *bind.FilterOpts, keyperSetIndex []uint64, retryCounter []uint64, keyperIndex []uint64) (*DkgcontractAccusationSubmittedIterator, error) {

	var keyperSetIndexRule []interface{}
	for _, keyperSetIndexItem := range keyperSetIndex {
		keyperSetIndexRule = append(keyperSetIndexRule, keyperSetIndexItem)
	}
	var retryCounterRule []interface{}
	for _, retryCounterItem := range retryCounter {
		retryCounterRule = append(retryCounterRule, retryCounterItem)
	}
	var keyperIndexRule []interface{}
	for _, keyperIndexItem := range keyperIndex {
		keyperIndexRule = append(keyperIndexRule, keyperIndexItem)
	}

	logs, sub, err := _Dkgcontract.contract.FilterLogs(opts, "AccusationSubmitted", keyperSetIndexRule, retryCounterRule, keyperIndexRule)
	if err != nil {
		return nil, err
	}
	return &DkgcontractAccusationSubmittedIterator{contract: _Dkgcontract.contract, event: "AccusationSubmitted", logs: logs, sub: sub}, nil
}

// WatchAccusationSubmitted is a free log subscription operation binding the contract event 0xb5cc05e740a503c13e1b7ac3ad2d5abf41b91d2e5fca79455d34e515fb70849a.
//
// Solidity: event AccusationSubmitted(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, uint64 indexed keyperIndex, uint64[] accusedIndices)
func (_Dkgcontract *DkgcontractFilterer) WatchAccusationSubmitted(opts *bind.WatchOpts, sink chan<- *DkgcontractAccusationSubmitted, keyperSetIndex []uint64, retryCounter []uint64, keyperIndex []uint64) (event.Subscription, error) {

	var keyperSetIndexRule []interface{}
	for _, keyperSetIndexItem := range keyperSetIndex {
		keyperSetIndexRule = append(keyperSetIndexRule, keyperSetIndexItem)
	}
	var retryCounterRule []interface{}
	for _, retryCounterItem := range retryCounter {
		retryCounterRule = append(retryCounterRule, retryCounterItem)
	}
	var keyperIndexRule []interface{}
	for _, keyperIndexItem := range keyperIndex {
		keyperIndexRule = append(keyperIndexRule, keyperIndexItem)
	}

	logs, sub, err := _Dkgcontract.contract.WatchLogs(opts, "AccusationSubmitted", keyperSetIndexRule, retryCounterRule, keyperIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DkgcontractAccusationSubmitted)
				if err := _Dkgcontract.contract.UnpackLog(event, "AccusationSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccusationSubmitted is a log parse operation binding the contract event 0xb5cc05e740a503c13e1b7ac3ad2d5abf41b91d2e5fca79455d34e515fb70849a.
//
// Solidity: event AccusationSubmitted(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, uint64 indexed keyperIndex, uint64[] accusedIndices)
func (_Dkgcontract *DkgcontractFilterer) ParseAccusationSubmitted(log types.Log) (*DkgcontractAccusationSubmitted, error) {
	event := new(DkgcontractAccusationSubmitted)
	if err := _Dkgcontract.contract.UnpackLog(event, "AccusationSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DkgcontractApologySubmittedIterator is returned from FilterApologySubmitted and is used to iterate over the raw logs and unpacked data for ApologySubmitted events raised by the Dkgcontract contract.
type DkgcontractApologySubmittedIterator struct {
	Event *DkgcontractApologySubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DkgcontractApologySubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DkgcontractApologySubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DkgcontractApologySubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DkgcontractApologySubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DkgcontractApologySubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DkgcontractApologySubmitted represents a ApologySubmitted event raised by the Dkgcontract contract.
type DkgcontractApologySubmitted struct {
	KeyperSetIndex uint64
	RetryCounter   uint64
	KeyperIndex    uint64
	AccuserIndices []uint64
	PolyEvalData   [][]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterApologySubmitted is a free log retrieval operation binding the contract event 0x62729f0820cf3ab60fbd63a51cdeae0554d0ab24f843b86bd0d2591594d69136.
//
// Solidity: event ApologySubmitted(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, uint64 indexed keyperIndex, uint64[] accuserIndices, bytes[] polyEvalData)
func (_Dkgcontract *DkgcontractFilterer) FilterApologySubmitted(opts *bind.FilterOpts, keyperSetIndex []uint64, retryCounter []uint64, keyperIndex []uint64) (*DkgcontractApologySubmittedIterator, error) {

	var keyperSetIndexRule []interface{}
	for _, keyperSetIndexItem := range keyperSetIndex {
		keyperSetIndexRule = append(keyperSetIndexRule, keyperSetIndexItem)
	}
	var retryCounterRule []interface{}
	for _, retryCounterItem := range retryCounter {
		retryCounterRule = append(retryCounterRule, retryCounterItem)
	}
	var keyperIndexRule []interface{}
	for _, keyperIndexItem := range keyperIndex {
		keyperIndexRule = append(keyperIndexRule, keyperIndexItem)
	}

	logs, sub, err := _Dkgcontract.contract.FilterLogs(opts, "ApologySubmitted", keyperSetIndexRule, retryCounterRule, keyperIndexRule)
	if err != nil {
		return nil, err
	}
	return &DkgcontractApologySubmittedIterator{contract: _Dkgcontract.contract, event: "ApologySubmitted", logs: logs, sub: sub}, nil
}

// WatchApologySubmitted is a free log subscription operation binding the contract event 0x62729f0820cf3ab60fbd63a51cdeae0554d0ab24f843b86bd0d2591594d69136.
//
// Solidity: event ApologySubmitted(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, uint64 indexed keyperIndex, uint64[] accuserIndices, bytes[] polyEvalData)
func (_Dkgcontract *DkgcontractFilterer) WatchApologySubmitted(opts *bind.WatchOpts, sink chan<- *DkgcontractApologySubmitted, keyperSetIndex []uint64, retryCounter []uint64, keyperIndex []uint64) (event.Subscription, error) {

	var keyperSetIndexRule []interface{}
	for _, keyperSetIndexItem := range keyperSetIndex {
		keyperSetIndexRule = append(keyperSetIndexRule, keyperSetIndexItem)
	}
	var retryCounterRule []interface{}
	for _, retryCounterItem := range retryCounter {
		retryCounterRule = append(retryCounterRule, retryCounterItem)
	}
	var keyperIndexRule []interface{}
	for _, keyperIndexItem := range keyperIndex {
		keyperIndexRule = append(keyperIndexRule, keyperIndexItem)
	}

	logs, sub, err := _Dkgcontract.contract.WatchLogs(opts, "ApologySubmitted", keyperSetIndexRule, retryCounterRule, keyperIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DkgcontractApologySubmitted)
				if err := _Dkgcontract.contract.UnpackLog(event, "ApologySubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApologySubmitted is a log parse operation binding the contract event 0x62729f0820cf3ab60fbd63a51cdeae0554d0ab24f843b86bd0d2591594d69136.
//
// Solidity: event ApologySubmitted(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, uint64 indexed keyperIndex, uint64[] accuserIndices, bytes[] polyEvalData)
func (_Dkgcontract *DkgcontractFilterer) ParseApologySubmitted(log types.Log) (*DkgcontractApologySubmitted, error) {
	event := new(DkgcontractApologySubmitted)
	if err := _Dkgcontract.contract.UnpackLog(event, "ApologySubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DkgcontractDKGSucceededIterator is returned from FilterDKGSucceeded and is used to iterate over the raw logs and unpacked data for DKGSucceeded events raised by the Dkgcontract contract.
type DkgcontractDKGSucceededIterator struct {
	Event *DkgcontractDKGSucceeded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DkgcontractDKGSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DkgcontractDKGSucceeded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DkgcontractDKGSucceeded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DkgcontractDKGSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DkgcontractDKGSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DkgcontractDKGSucceeded represents a DKGSucceeded event raised by the Dkgcontract contract.
type DkgcontractDKGSucceeded struct {
	KeyperSetIndex uint64
	RetryCounter   uint64
	EonPublicKey   []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDKGSucceeded is a free log retrieval operation binding the contract event 0xe6c9bd3daa50b9218615605cdf333410d405b7f3b8dbeee5a2373769679aa9f5.
//
// Solidity: event DKGSucceeded(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, bytes eonPublicKey)
func (_Dkgcontract *DkgcontractFilterer) FilterDKGSucceeded(opts *bind.FilterOpts, keyperSetIndex []uint64, retryCounter []uint64) (*DkgcontractDKGSucceededIterator, error) {

	var keyperSetIndexRule []interface{}
	for _, keyperSetIndexItem := range keyperSetIndex {
		keyperSetIndexRule = append(keyperSetIndexRule, keyperSetIndexItem)
	}
	var retryCounterRule []interface{}
	for _, retryCounterItem := range retryCounter {
		retryCounterRule = append(retryCounterRule, retryCounterItem)
	}

	logs, sub, err := _Dkgcontract.contract.FilterLogs(opts, "DKGSucceeded", keyperSetIndexRule, retryCounterRule)
	if err != nil {
		return nil, err
	}
	return &DkgcontractDKGSucceededIterator{contract: _Dkgcontract.contract, event: "DKGSucceeded", logs: logs, sub: sub}, nil
}

// WatchDKGSucceeded is a free log subscription operation binding the contract event 0xe6c9bd3daa50b9218615605cdf333410d405b7f3b8dbeee5a2373769679aa9f5.
//
// Solidity: event DKGSucceeded(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, bytes eonPublicKey)
func (_Dkgcontract *DkgcontractFilterer) WatchDKGSucceeded(opts *bind.WatchOpts, sink chan<- *DkgcontractDKGSucceeded, keyperSetIndex []uint64, retryCounter []uint64) (event.Subscription, error) {

	var keyperSetIndexRule []interface{}
	for _, keyperSetIndexItem := range keyperSetIndex {
		keyperSetIndexRule = append(keyperSetIndexRule, keyperSetIndexItem)
	}
	var retryCounterRule []interface{}
	for _, retryCounterItem := range retryCounter {
		retryCounterRule = append(retryCounterRule, retryCounterItem)
	}

	logs, sub, err := _Dkgcontract.contract.WatchLogs(opts, "DKGSucceeded", keyperSetIndexRule, retryCounterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DkgcontractDKGSucceeded)
				if err := _Dkgcontract.contract.UnpackLog(event, "DKGSucceeded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDKGSucceeded is a log parse operation binding the contract event 0xe6c9bd3daa50b9218615605cdf333410d405b7f3b8dbeee5a2373769679aa9f5.
//
// Solidity: event DKGSucceeded(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, bytes eonPublicKey)
func (_Dkgcontract *DkgcontractFilterer) ParseDKGSucceeded(log types.Log) (*DkgcontractDKGSucceeded, error) {
	event := new(DkgcontractDKGSucceeded)
	if err := _Dkgcontract.contract.UnpackLog(event, "DKGSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DkgcontractDealingSubmittedIterator is returned from FilterDealingSubmitted and is used to iterate over the raw logs and unpacked data for DealingSubmitted events raised by the Dkgcontract contract.
type DkgcontractDealingSubmittedIterator struct {
	Event *DkgcontractDealingSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DkgcontractDealingSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DkgcontractDealingSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DkgcontractDealingSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DkgcontractDealingSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DkgcontractDealingSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DkgcontractDealingSubmitted represents a DealingSubmitted event raised by the Dkgcontract contract.
type DkgcontractDealingSubmitted struct {
	KeyperSetIndex uint64
	RetryCounter   uint64
	KeyperIndex    uint64
	Commitment     []byte
	PolyEvals      [][]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDealingSubmitted is a free log retrieval operation binding the contract event 0xa5074728b92791250d48ccdc55266aca7b1cca22f89e1ef2ab91ccd7ae0f337d.
//
// Solidity: event DealingSubmitted(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, uint64 indexed keyperIndex, bytes commitment, bytes[] polyEvals)
func (_Dkgcontract *DkgcontractFilterer) FilterDealingSubmitted(opts *bind.FilterOpts, keyperSetIndex []uint64, retryCounter []uint64, keyperIndex []uint64) (*DkgcontractDealingSubmittedIterator, error) {

	var keyperSetIndexRule []interface{}
	for _, keyperSetIndexItem := range keyperSetIndex {
		keyperSetIndexRule = append(keyperSetIndexRule, keyperSetIndexItem)
	}
	var retryCounterRule []interface{}
	for _, retryCounterItem := range retryCounter {
		retryCounterRule = append(retryCounterRule, retryCounterItem)
	}
	var keyperIndexRule []interface{}
	for _, keyperIndexItem := range keyperIndex {
		keyperIndexRule = append(keyperIndexRule, keyperIndexItem)
	}

	logs, sub, err := _Dkgcontract.contract.FilterLogs(opts, "DealingSubmitted", keyperSetIndexRule, retryCounterRule, keyperIndexRule)
	if err != nil {
		return nil, err
	}
	return &DkgcontractDealingSubmittedIterator{contract: _Dkgcontract.contract, event: "DealingSubmitted", logs: logs, sub: sub}, nil
}

// WatchDealingSubmitted is a free log subscription operation binding the contract event 0xa5074728b92791250d48ccdc55266aca7b1cca22f89e1ef2ab91ccd7ae0f337d.
//
// Solidity: event DealingSubmitted(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, uint64 indexed keyperIndex, bytes commitment, bytes[] polyEvals)
func (_Dkgcontract *DkgcontractFilterer) WatchDealingSubmitted(opts *bind.WatchOpts, sink chan<- *DkgcontractDealingSubmitted, keyperSetIndex []uint64, retryCounter []uint64, keyperIndex []uint64) (event.Subscription, error) {

	var keyperSetIndexRule []interface{}
	for _, keyperSetIndexItem := range keyperSetIndex {
		keyperSetIndexRule = append(keyperSetIndexRule, keyperSetIndexItem)
	}
	var retryCounterRule []interface{}
	for _, retryCounterItem := range retryCounter {
		retryCounterRule = append(retryCounterRule, retryCounterItem)
	}
	var keyperIndexRule []interface{}
	for _, keyperIndexItem := range keyperIndex {
		keyperIndexRule = append(keyperIndexRule, keyperIndexItem)
	}

	logs, sub, err := _Dkgcontract.contract.WatchLogs(opts, "DealingSubmitted", keyperSetIndexRule, retryCounterRule, keyperIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DkgcontractDealingSubmitted)
				if err := _Dkgcontract.contract.UnpackLog(event, "DealingSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDealingSubmitted is a log parse operation binding the contract event 0xa5074728b92791250d48ccdc55266aca7b1cca22f89e1ef2ab91ccd7ae0f337d.
//
// Solidity: event DealingSubmitted(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, uint64 indexed keyperIndex, bytes commitment, bytes[] polyEvals)
func (_Dkgcontract *DkgcontractFilterer) ParseDealingSubmitted(log types.Log) (*DkgcontractDealingSubmitted, error) {
	event := new(DkgcontractDealingSubmitted)
	if err := _Dkgcontract.contract.UnpackLog(event, "DealingSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DkgcontractSuccessVoteSubmittedIterator is returned from FilterSuccessVoteSubmitted and is used to iterate over the raw logs and unpacked data for SuccessVoteSubmitted events raised by the Dkgcontract contract.
type DkgcontractSuccessVoteSubmittedIterator struct {
	Event *DkgcontractSuccessVoteSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DkgcontractSuccessVoteSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DkgcontractSuccessVoteSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DkgcontractSuccessVoteSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DkgcontractSuccessVoteSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DkgcontractSuccessVoteSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DkgcontractSuccessVoteSubmitted represents a SuccessVoteSubmitted event raised by the Dkgcontract contract.
type DkgcontractSuccessVoteSubmitted struct {
	KeyperSetIndex uint64
	RetryCounter   uint64
	KeyperIndex    uint64
	EonPublicKey   []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSuccessVoteSubmitted is a free log retrieval operation binding the contract event 0xacfc0aa3f6ba36ac0ffda53e00d2cdace2fd4a9e41e2fe3214032794332c37c3.
//
// Solidity: event SuccessVoteSubmitted(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, uint64 indexed keyperIndex, bytes eonPublicKey)
func (_Dkgcontract *DkgcontractFilterer) FilterSuccessVoteSubmitted(opts *bind.FilterOpts, keyperSetIndex []uint64, retryCounter []uint64, keyperIndex []uint64) (*DkgcontractSuccessVoteSubmittedIterator, error) {

	var keyperSetIndexRule []interface{}
	for _, keyperSetIndexItem := range keyperSetIndex {
		keyperSetIndexRule = append(keyperSetIndexRule, keyperSetIndexItem)
	}
	var retryCounterRule []interface{}
	for _, retryCounterItem := range retryCounter {
		retryCounterRule = append(retryCounterRule, retryCounterItem)
	}
	var keyperIndexRule []interface{}
	for _, keyperIndexItem := range keyperIndex {
		keyperIndexRule = append(keyperIndexRule, keyperIndexItem)
	}

	logs, sub, err := _Dkgcontract.contract.FilterLogs(opts, "SuccessVoteSubmitted", keyperSetIndexRule, retryCounterRule, keyperIndexRule)
	if err != nil {
		return nil, err
	}
	return &DkgcontractSuccessVoteSubmittedIterator{contract: _Dkgcontract.contract, event: "SuccessVoteSubmitted", logs: logs, sub: sub}, nil
}

// WatchSuccessVoteSubmitted is a free log subscription operation binding the contract event 0xacfc0aa3f6ba36ac0ffda53e00d2cdace2fd4a9e41e2fe3214032794332c37c3.
//
// Solidity: event SuccessVoteSubmitted(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, uint64 indexed keyperIndex, bytes eonPublicKey)
func (_Dkgcontract *DkgcontractFilterer) WatchSuccessVoteSubmitted(opts *bind.WatchOpts, sink chan<- *DkgcontractSuccessVoteSubmitted, keyperSetIndex []uint64, retryCounter []uint64, keyperIndex []uint64) (event.Subscription, error) {

	var keyperSetIndexRule []interface{}
	for _, keyperSetIndexItem := range keyperSetIndex {
		keyperSetIndexRule = append(keyperSetIndexRule, keyperSetIndexItem)
	}
	var retryCounterRule []interface{}
	for _, retryCounterItem := range retryCounter {
		retryCounterRule = append(retryCounterRule, retryCounterItem)
	}
	var keyperIndexRule []interface{}
	for _, keyperIndexItem := range keyperIndex {
		keyperIndexRule = append(keyperIndexRule, keyperIndexItem)
	}

	logs, sub, err := _Dkgcontract.contract.WatchLogs(opts, "SuccessVoteSubmitted", keyperSetIndexRule, retryCounterRule, keyperIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DkgcontractSuccessVoteSubmitted)
				if err := _Dkgcontract.contract.UnpackLog(event, "SuccessVoteSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSuccessVoteSubmitted is a log parse operation binding the contract event 0xacfc0aa3f6ba36ac0ffda53e00d2cdace2fd4a9e41e2fe3214032794332c37c3.
//
// Solidity: event SuccessVoteSubmitted(uint64 indexed keyperSetIndex, uint64 indexed retryCounter, uint64 indexed keyperIndex, bytes eonPublicKey)
func (_Dkgcontract *DkgcontractFilterer) ParseSuccessVoteSubmitted(log types.Log) (*DkgcontractSuccessVoteSubmitted, error) {
	event := new(DkgcontractSuccessVoteSubmitted)
	if err := _Dkgcontract.contract.UnpackLog(event, "SuccessVoteSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
