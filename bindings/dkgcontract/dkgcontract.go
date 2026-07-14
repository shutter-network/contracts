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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"phaseLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dkgLeadLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxRetries\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperSetManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"keyBroadcastContractAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DKG_LEAD_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RETRIES\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PHASE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentPhase\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIDKGContract.Phase\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cycleLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dkgStart\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasVoted\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyBroadcastContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyperSetManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitAccusation\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"accusedIndices\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitApology\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"accuserIndices\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"polyEvalData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitDealing\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"commitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"polyEvals\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitSuccessVote\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"eonPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"succeeded\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voteCount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AccusationSubmitted\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"accusedIndices\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApologySubmitted\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"accuserIndices\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"},{\"name\":\"polyEvalData\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DKGSucceeded\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"eonPublicKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DealingSubmitted\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"commitment\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"polyEvals\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SuccessVoteSubmitted\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"eonPublicKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadySucceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyVoted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAccusation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyEonPublicKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxRetriesExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrays\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotKeyperAtIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongDKGContract\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLengthParameter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroMaxRetries\",\"inputs\":[]}]",
	Bin: "0x610120604052348015610010575f5ffd5b5060405161225a38038061225a8339818101604052810190610032919061022f565b5f8567ffffffffffffffff16148061005357505f8467ffffffffffffffff16145b1561008a576040517f038016d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8367ffffffffffffffff16036100cd576040517fb8579f9100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8467ffffffffffffffff1660808167ffffffffffffffff16815250508367ffffffffffffffff1660a08167ffffffffffffffff16815250508267ffffffffffffffff1660c08167ffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff166101008173ffffffffffffffffffffffffffffffffffffffff168152505050505050506102a6565b5f5ffd5b5f67ffffffffffffffff82169050919050565b6101b481610198565b81146101be575f5ffd5b50565b5f815190506101cf816101ab565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101fe826101d5565b9050919050565b61020e816101f4565b8114610218575f5ffd5b50565b5f8151905061022981610205565b92915050565b5f5f5f5f5f60a0868803121561024857610247610194565b5b5f610255888289016101c1565b9550506020610266888289016101c1565b9450506040610277888289016101c1565b93505060606102888882890161021b565b92505060806102998882890161021b565b9150509295509295909350565b60805160a05160c05160e05161010051611f316103295f395f8181610a280152610c2401525f81816108c701528181610bfd01528181610d9201528181610fed015261115c01525f818161030e015281816104eb0152610ea601525f818161042f0152610e5301525f818161036e01528181610bd80152610c4b0152611f315ff3fe608060405234801561000f575f5ffd5b50600436106100f3575f3560e01c806394afc12d11610095578063eac471a011610064578063eac471a014610271578063eb05c2801461028f578063f01275b0146102bf578063f48f3d14146102db576100f3565b806394afc12d146101fb578063b209e71714610217578063d4b2cedd14610235578063df64063114610253576100f3565b806346ebbdb0116100d157806346ebbdb0146101615780637642d0de1461017f57806378a1b9db146101af57806382b870d3146101df576100f3565b80630a3ec4d6146100f757806316b89c3f146101275780633f98009814610145575b5f5ffd5b610111600480360381019061010c919061131a565b61030b565b60405161011e91906113cb565b60405180910390f35b61012f61042d565b60405161013c91906113f3565b60405180910390f35b61015f600480360381019061015a91906114c2565b610451565b005b6101696104e9565b60405161017691906113f3565b60405180910390f35b610199600480360381019061019491906115d3565b61050d565b6040516101a6919061163d565b60405180910390f35b6101c960048036038101906101c49190611656565b610542565b6040516101d6919061163d565b60405180910390f35b6101f960048036038101906101f49190611681565b61055e565b005b6102156004803603810190610210919061175a565b610b08565b005b61021f610bd6565b60405161022c91906113f3565b60405180910390f35b61023d610bfa565b60405161024a91906117ed565b60405180910390f35b61025b610c21565b60405161026891906117ed565b60405180910390f35b610279610c48565b60405161028691906113f3565b60405180910390f35b6102a960048036038101906102a49190611839565b610c7b565b6040516102b691906113f3565b60405180910390f35b6102d960048036038101906102d49190611889565b610cb7565b005b6102f560048036038101906102f0919061131a565b610d8e565b6040516103029190611958565b60405180910390f35b5f7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168267ffffffffffffffff161061034f575f9050610427565b5f61035a8484610d8e565b90505f8143610369919061199e565b90505f7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1690505f8212156103ac575f9350505050610427565b808212156103c05760019350505050610427565b8060026103cd91906119de565b8212156103e05760029350505050610427565b8060036103ed91906119de565b8212156104005760039350505050610427565b80600461040d91906119de565b8212156104205760049350505050610427565b5f93505050505b92915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b61045a86610ea4565b61046387610f14565b61046f87876001610f7f565b61047887610fea565b6104828786611159565b8467ffffffffffffffff168667ffffffffffffffff168867ffffffffffffffff167fa5074728b92791250d48ccdc55266aca7b1cca22f89e1ef2ab91ccd7ae0f337d878787876040516104d89493929190611c06565b60405180910390a450505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6002602052825f5260405f20602052815f5260405f20602052805f5260405f205f92509250509054906101000a900460ff1681565b5f602052805f5260405f205f915054906101000a900460ff1681565b61056784610ea4565b61057385856004610f7f565b61057c85610fea565b6105868584611159565b5f82829050036105c2576040517f59acc21d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610689576040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160025f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f8282604051610735929190611c6d565b604051809103902090505f6001805f8967ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8481526020019081526020015f205f9054906101000a900467ffffffffffffffff166107b69190611c85565b90508060015f8967ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8481526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508467ffffffffffffffff168667ffffffffffffffff168867ffffffffffffffff167facfc0aa3f6ba36ac0ffda53e00d2cdace2fd4a9e41e2fe3214032794332c37c38787604051610887929190611cc0565b60405180910390a45f5f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610aff575f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f90f3bed896040518263ffffffff1660e01b815260040161091e91906113f3565b602060405180830381865afa158015610939573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061095d9190611cf6565b90505f8173ffffffffffffffffffffffffffffffffffffffff1663e75235b86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109a9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109cd9190611d35565b90508067ffffffffffffffff168367ffffffffffffffff1610610afc5760015f5f8b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663daade8e88a88886040518463ffffffff1660e01b8152600401610a8393929190611d60565b5f604051808303815f87803b158015610a9a575f5ffd5b505af1925050508015610aab575060015b508767ffffffffffffffff168967ffffffffffffffff167fe6c9bd3daa50b9218615605cdf333410d405b7f3b8dbeee5a2373769679aa9f58888604051610af3929190611cc0565b60405180910390a35b50505b50505050505050565b610b1184610ea4565b610b1a85610f14565b610b2685856002610f7f565b610b2f85610fea565b610b398584611159565b5f8282905003610b75576040517f2eeffe0b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8267ffffffffffffffff168467ffffffffffffffff168667ffffffffffffffff167fb5cc05e740a503c13e1b7ac3ad2d5abf41b91d2e5fca79455d34e515fb70849a8585604051610bc7929190611e4c565b60405180910390a45050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f7f0000000000000000000000000000000000000000000000000000000000000000905090565b5f7f0000000000000000000000000000000000000000000000000000000000000000905090565b5f7f00000000000000000000000000000000000000000000000000000000000000006004610c769190611e6e565b905090565b6001602052825f5260405f20602052815f5260405f20602052805f5260405f205f92509250509054906101000a900467ffffffffffffffff1681565b610cc086610ea4565b610cc987610f14565b610cd587876003610f7f565b610cde87610fea565b610ce88786611159565b818190508484905014610d27576040517fa121188700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8467ffffffffffffffff168667ffffffffffffffff168867ffffffffffffffff167f62729f0820cf3ab60fbd63a51cdeae0554d0ab24f843b86bd0d2591594d6913687878787604051610d7d9493929190611eaa565b60405180910390a450505050505050565b5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663636df979856040518263ffffffff1660e01b8152600401610de991906113f3565b602060405180830381865afa158015610e04573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e289190611d35565b9050610e32610c48565b67ffffffffffffffff168367ffffffffffffffff16610e5191906119de565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168267ffffffffffffffff16610e91919061199e565b610e9b9190611ee3565b91505092915050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168167ffffffffffffffff1610610f11576040517f12f6147700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b5f5f8267ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610f7c576040517fe0f5ec9d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b806004811115610f9257610f91611358565b5b610f9c848461030b565b6004811115610fae57610fad611358565b5b14610fe5576040517fe2586bcc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f90f3bed836040518263ffffffff1660e01b815260040161104491906113f3565b602060405180830381865afa15801561105f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110839190611cf6565b90503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff166266f0a86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110e4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111089190611cf6565b73ffffffffffffffffffffffffffffffffffffffff1614611155576040517ff62f18e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f90f3bed846040518263ffffffff1660e01b81526004016111b391906113f3565b602060405180830381865afa1580156111ce573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111f29190611cf6565b90503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16632e8e6cad846040518263ffffffff1660e01b815260040161124491906113f3565b602060405180830381865afa15801561125f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112839190611cf6565b73ffffffffffffffffffffffffffffffffffffffff16146112d0576040517f7bf9580300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff82169050919050565b6112f9816112dd565b8114611303575f5ffd5b50565b5f81359050611314816112f0565b92915050565b5f5f604083850312156113305761132f6112d5565b5b5f61133d85828601611306565b925050602061134e85828601611306565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6005811061139657611395611358565b5b50565b5f8190506113a682611385565b919050565b5f6113b582611399565b9050919050565b6113c5816113ab565b82525050565b5f6020820190506113de5f8301846113bc565b92915050565b6113ed816112dd565b82525050565b5f6020820190506114065f8301846113e4565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261142d5761142c61140c565b5b8235905067ffffffffffffffff81111561144a57611449611410565b5b60208301915083600182028301111561146657611465611414565b5b9250929050565b5f5f83601f8401126114825761148161140c565b5b8235905067ffffffffffffffff81111561149f5761149e611410565b5b6020830191508360208202830111156114bb576114ba611414565b5b9250929050565b5f5f5f5f5f5f5f60a0888a0312156114dd576114dc6112d5565b5b5f6114ea8a828b01611306565b97505060206114fb8a828b01611306565b965050604061150c8a828b01611306565b955050606088013567ffffffffffffffff81111561152d5761152c6112d9565b5b6115398a828b01611418565b9450945050608088013567ffffffffffffffff81111561155c5761155b6112d9565b5b6115688a828b0161146d565b925092505092959891949750929550565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6115a282611579565b9050919050565b6115b281611598565b81146115bc575f5ffd5b50565b5f813590506115cd816115a9565b92915050565b5f5f5f606084860312156115ea576115e96112d5565b5b5f6115f786828701611306565b935050602061160886828701611306565b9250506040611619868287016115bf565b9150509250925092565b5f8115159050919050565b61163781611623565b82525050565b5f6020820190506116505f83018461162e565b92915050565b5f6020828403121561166b5761166a6112d5565b5b5f61167884828501611306565b91505092915050565b5f5f5f5f5f6080868803121561169a576116996112d5565b5b5f6116a788828901611306565b95505060206116b888828901611306565b94505060406116c988828901611306565b935050606086013567ffffffffffffffff8111156116ea576116e96112d9565b5b6116f688828901611418565b92509250509295509295909350565b5f5f83601f84011261171a5761171961140c565b5b8235905067ffffffffffffffff81111561173757611736611410565b5b60208301915083602082028301111561175357611752611414565b5b9250929050565b5f5f5f5f5f60808688031215611773576117726112d5565b5b5f61178088828901611306565b955050602061179188828901611306565b94505060406117a288828901611306565b935050606086013567ffffffffffffffff8111156117c3576117c26112d9565b5b6117cf88828901611705565b92509250509295509295909350565b6117e781611598565b82525050565b5f6020820190506118005f8301846117de565b92915050565b5f819050919050565b61181881611806565b8114611822575f5ffd5b50565b5f813590506118338161180f565b92915050565b5f5f5f606084860312156118505761184f6112d5565b5b5f61185d86828701611306565b935050602061186e86828701611306565b925050604061187f86828701611825565b9150509250925092565b5f5f5f5f5f5f5f60a0888a0312156118a4576118a36112d5565b5b5f6118b18a828b01611306565b97505060206118c28a828b01611306565b96505060406118d38a828b01611306565b955050606088013567ffffffffffffffff8111156118f4576118f36112d9565b5b6119008a828b01611705565b9450945050608088013567ffffffffffffffff811115611923576119226112d9565b5b61192f8a828b0161146d565b925092505092959891949750929550565b5f819050919050565b61195281611940565b82525050565b5f60208201905061196b5f830184611949565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6119a882611940565b91506119b383611940565b925082820390508181125f8412168282135f8512151617156119d8576119d7611971565b5b92915050565b5f6119e882611940565b91506119f383611940565b9250828202611a0181611940565b91507f800000000000000000000000000000000000000000000000000000000000000084145f84121615611a3857611a37611971565b5b8282058414831517611a4d57611a4c611971565b5b5092915050565b5f82825260208201905092915050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f611a8d8385611a54565b9350611a9a838584611a64565b611aa383611a72565b840190509392505050565b5f82825260208201905092915050565b5f819050919050565b5f82825260208201905092915050565b5f611ae28385611ac7565b9350611aef838584611a64565b611af883611a72565b840190509392505050565b5f611b0f848484611ad7565b90509392505050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83356001602003843603038112611b4057611b3f611b20565b5b83810192508235915060208301925067ffffffffffffffff821115611b6857611b67611b18565b5b600182023603831315611b7e57611b7d611b1c565b5b509250929050565b5f602082019050919050565b5f611b9d8385611aae565b935083602084028501611baf84611abe565b805f5b87811015611bf4578484038952611bc98284611b24565b611bd4868284611b03565b9550611bdf84611b86565b935060208b019a505050600181019050611bb2565b50829750879450505050509392505050565b5f6040820190508181035f830152611c1f818688611a82565b90508181036020830152611c34818486611b92565b905095945050505050565b5f81905092915050565b5f611c548385611c3f565b9350611c61838584611a64565b82840190509392505050565b5f611c79828486611c49565b91508190509392505050565b5f611c8f826112dd565b9150611c9a836112dd565b9250828201905067ffffffffffffffff811115611cba57611cb9611971565b5b92915050565b5f6020820190508181035f830152611cd9818486611a82565b90509392505050565b5f81519050611cf0816115a9565b92915050565b5f60208284031215611d0b57611d0a6112d5565b5b5f611d1884828501611ce2565b91505092915050565b5f81519050611d2f816112f0565b92915050565b5f60208284031215611d4a57611d496112d5565b5b5f611d5784828501611d21565b91505092915050565b5f604082019050611d735f8301866113e4565b8181036020830152611d86818486611a82565b9050949350505050565b5f82825260208201905092915050565b5f819050919050565b611db2816112dd565b82525050565b5f611dc38383611da9565b60208301905092915050565b5f611ddd6020840184611306565b905092915050565b5f602082019050919050565b5f611dfc8385611d90565b9350611e0782611da0565b805f5b85811015611e3f57611e1c8284611dcf565b611e268882611db8565b9750611e3183611de5565b925050600181019050611e0a565b5085925050509392505050565b5f6020820190508181035f830152611e65818486611df1565b90509392505050565b5f611e78826112dd565b9150611e83836112dd565b9250828202611e91816112dd565b9150808214611ea357611ea2611971565b5b5092915050565b5f6040820190508181035f830152611ec3818688611df1565b90508181036020830152611ed8818486611b92565b905095945050505050565b5f611eed82611940565b9150611ef883611940565b92508282019050828112155f8312168382125f841215161715611f1e57611f1d611971565b5b9291505056fea164736f6c634300081c000a",
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
// Solidity: function succeeded(uint64 ) view returns(bool)
func (_Dkgcontract *DkgcontractCaller) Succeeded(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _Dkgcontract.contract.Call(opts, &out, "succeeded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Succeeded is a free data retrieval call binding the contract method 0x78a1b9db.
//
// Solidity: function succeeded(uint64 ) view returns(bool)
func (_Dkgcontract *DkgcontractSession) Succeeded(arg0 uint64) (bool, error) {
	return _Dkgcontract.Contract.Succeeded(&_Dkgcontract.CallOpts, arg0)
}

// Succeeded is a free data retrieval call binding the contract method 0x78a1b9db.
//
// Solidity: function succeeded(uint64 ) view returns(bool)
func (_Dkgcontract *DkgcontractCallerSession) Succeeded(arg0 uint64) (bool, error) {
	return _Dkgcontract.Contract.Succeeded(&_Dkgcontract.CallOpts, arg0)
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
