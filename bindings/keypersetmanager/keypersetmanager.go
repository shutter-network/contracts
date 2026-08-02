// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package keypersetmanager

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

// KeypersetmanagerMetaData contains all meta data concerning the Keypersetmanager contract.
var KeypersetmanagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initializer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addKeyperSet\",\"inputs\":[{\"name\":\"activationBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperSetContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getKeyperSetActivationBlock\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeyperSetAddress\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeyperSetIndexByBlock\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumKeyperSets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"KeyperSetAdded\",\"inputs\":[{\"name\":\"activationBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"keyperSetContract\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"members\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeyperSetAdded\",\"inputs\":[{\"name\":\"activationSlot\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"keyperSetContract\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadyHaveKeyperSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DKGContractManagerMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DKGContractNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyperSetNotFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoActiveKeyperSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedInitializer\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051611ab8380380611ab8833981810160405281019061003191906100d7565b80806001806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050610102565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a68261007d565b9050919050565b6100b68161009c565b81146100c0575f5ffd5b50565b5f815190506100d1816100ad565b92915050565b5f602082840312156100ec576100eb610079565b5b5f6100f9848285016100c3565b91505092915050565b6119a98061010f5f395ff3fe608060405234801561000f575f5ffd5b5060043610610114575f3560e01c80638456cb59116100a0578063d3877c431161006f578063d3877c43146102ca578063d547741f146102e6578063e63ab1e914610302578063f2e6100a14610320578063f90f3bed1461033e57610114565b80638456cb591461025457806391d148541461025e5780639ce110d71461028e578063a217fddf146102ac57610114565b806336568abe116100e757806336568abe146101c45780633f4ba83a146101e0578063485cc955146101ea5780635c975abb14610206578063636df9791461022457610114565b806301ffc9a714610118578063035cef1514610148578063248a9ca3146101785780632f2ff15d146101a8575b5f5ffd5b610132600480360381019061012d919061125c565b61036e565b60405161013f91906112a1565b60405180910390f35b610162600480360381019061015d91906112f7565b6103e7565b60405161016f9190611331565b60405180910390f35b610192600480360381019061018d919061137d565b6104b2565b60405161019f91906113b7565b60405180910390f35b6101c260048036038101906101bd919061142a565b6104ce565b005b6101de60048036038101906101d9919061142a565b6104f0565b005b6101e861056b565b005b61020460048036038101906101ff9190611468565b610582565b005b61020e610708565b60405161021b91906112a1565b60405180910390f35b61023e600480360381019061023991906112f7565b61071d565b60405161024b9190611331565b60405180910390f35b61025c610761565b005b6102786004803603810190610273919061142a565b610796565b60405161028591906112a1565b60405180910390f35b6102966107f9565b6040516102a391906114b5565b60405180910390f35b6102b461081d565b6040516102c191906113b7565b60405180910390f35b6102e460048036038101906102df91906114ce565b610823565b005b61030060048036038101906102fb919061142a565b610d2f565b005b61030a610d51565b60405161031791906113b7565b60405180910390f35b610328610d75565b6040516103359190611331565b60405180910390f35b610358600480360381019061035391906112f7565b610d81565b60405161036591906114b5565b60405180910390f35b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103e057506103df82610dd2565b5b9050919050565b5f5f60028054905090505b5f81111561047a578267ffffffffffffffff1660026001836104149190611542565b8154811061042557610424611575565b5b905f5260205f20015f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16116104675760018161045f9190611542565b9150506104ad565b8080610472906115a2565b9150506103f2565b506040517fcf93fa3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b5f5f5f8381526020019081526020015f20600101549050919050565b6104d7826104b2565b6104e081610e3b565b6104ea8383610e4f565b50505050565b6104f8610f38565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461055c576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105668282610f3f565b505050565b5f5f1b61057781610e3b565b61057f611028565b50565b5f73ffffffffffffffffffffffffffffffffffffffff1660018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610607576040517f0dc149f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461068c576040517f0d622feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106985f5f1b83610e4f565b506106c37f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610e4f565b505f6001806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b5f60015f9054906101000a900460ff16905090565b5f60028267ffffffffffffffff168154811061073c5761073b611575565b5b905f5260205f20015f015f9054906101000a900467ffffffffffffffff169050919050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61078b81610e3b565b610793611089565b50565b5f5f5f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f1b81565b5f5f1b61082f81610e3b565b5f6002805490501180156108ad57506108a0600260016002805490506108559190611542565b8154811061086657610865611575565b5b905f5260205f20015f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1660014361089b91906115c9565b6110ea565b8367ffffffffffffffff16105b156108e4576040517fc396e95600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff16638d4e40836040518163ffffffff1660e01b8152600401602060405180830381865afa15801561092d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109519190611626565b610987576040517feac631aa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff166266f0a86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109d0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109f49190611665565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a5b576040517f73dd647200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1663d4b2cedd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610abb573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610adf9190611665565b73ffffffffffffffffffffffffffffffffffffffff1614610b2c576040517f7cfd996a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260405180604001604052808667ffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003905f5260205f20015f909190919091505f820151815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151815f0160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505f8390507fa940387dac06ebd336730f1d14b21629a9d137069a9137e871f95313e101016585858373ffffffffffffffffffffffffffffffffffffffff16639eab52536040518163ffffffff1660e01b81526004015f60405180830381865afa158015610c68573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610c9091906117e0565b8473ffffffffffffffffffffffffffffffffffffffff1663e75235b86040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cd9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cfd919061183b565b6001600280549050610d0f9190611542565b604051610d2095949392919061191d565b60405180910390a15050505050565b610d38826104b2565b610d4181610e3b565b610d4b8383610f3f565b50505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b5f600280549050905090565b5f60028267ffffffffffffffff1681548110610da057610d9f611575565b5b905f5260205f20015f0160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b610e4c81610e47610f38565b611100565b50565b5f610e5a8383610796565b610f2e5760015f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610ecb610f38565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019050610f32565b5f90505b92915050565b5f33905090565b5f610f4a8383610796565b1561101e575f5f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610fbb610f38565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a460019050611022565b5f90505b92915050565b611030611151565b5f60015f6101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa611072610f38565b60405161107f91906114b5565b60405180910390a1565b611091611191565b6001805f6101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586110d3610f38565b6040516110e091906114b5565b60405180910390a1565b5f6110f882841184846111d2565b905092915050565b61110a8282610796565b61114d5780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401611144929190611975565b60405180910390fd5b5050565b611159610708565b61118f576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b611199610708565b156111d0576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f6111dc846111eb565b82841802821890509392505050565b5f8115159050919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61123b81611207565b8114611245575f5ffd5b50565b5f8135905061125681611232565b92915050565b5f60208284031215611271576112706111ff565b5b5f61127e84828501611248565b91505092915050565b5f8115159050919050565b61129b81611287565b82525050565b5f6020820190506112b45f830184611292565b92915050565b5f67ffffffffffffffff82169050919050565b6112d6816112ba565b81146112e0575f5ffd5b50565b5f813590506112f1816112cd565b92915050565b5f6020828403121561130c5761130b6111ff565b5b5f611319848285016112e3565b91505092915050565b61132b816112ba565b82525050565b5f6020820190506113445f830184611322565b92915050565b5f819050919050565b61135c8161134a565b8114611366575f5ffd5b50565b5f8135905061137781611353565b92915050565b5f60208284031215611392576113916111ff565b5b5f61139f84828501611369565b91505092915050565b6113b18161134a565b82525050565b5f6020820190506113ca5f8301846113a8565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6113f9826113d0565b9050919050565b611409816113ef565b8114611413575f5ffd5b50565b5f8135905061142481611400565b92915050565b5f5f604083850312156114405761143f6111ff565b5b5f61144d85828601611369565b925050602061145e85828601611416565b9150509250929050565b5f5f6040838503121561147e5761147d6111ff565b5b5f61148b85828601611416565b925050602061149c85828601611416565b9150509250929050565b6114af816113ef565b82525050565b5f6020820190506114c85f8301846114a6565b92915050565b5f5f604083850312156114e4576114e36111ff565b5b5f6114f1858286016112e3565b925050602061150285828601611416565b9150509250929050565b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61154c8261150c565b91506115578361150c565b925082820390508181111561156f5761156e611515565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6115ac8261150c565b91505f82036115be576115bd611515565b5b600182039050919050565b5f6115d38261150c565b91506115de8361150c565b92508282019050808211156115f6576115f5611515565b5b92915050565b61160581611287565b811461160f575f5ffd5b50565b5f81519050611620816115fc565b92915050565b5f6020828403121561163b5761163a6111ff565b5b5f61164884828501611612565b91505092915050565b5f8151905061165f81611400565b92915050565b5f6020828403121561167a576116796111ff565b5b5f61168784828501611651565b91505092915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6116da82611694565b810181811067ffffffffffffffff821117156116f9576116f86116a4565b5b80604052505050565b5f61170b6111f6565b905061171782826116d1565b919050565b5f67ffffffffffffffff821115611736576117356116a4565b5b602082029050602081019050919050565b5f5ffd5b5f61175d6117588461171c565b611702565b905080838252602082019050602084028301858111156117805761177f611747565b5b835b818110156117a957806117958882611651565b845260208401935050602081019050611782565b5050509392505050565b5f82601f8301126117c7576117c6611690565b5b81516117d784826020860161174b565b91505092915050565b5f602082840312156117f5576117f46111ff565b5b5f82015167ffffffffffffffff81111561181257611811611203565b5b61181e848285016117b3565b91505092915050565b5f81519050611835816112cd565b92915050565b5f602082840312156118505761184f6111ff565b5b5f61185d84828501611827565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611898816113ef565b82525050565b5f6118a9838361188f565b60208301905092915050565b5f602082019050919050565b5f6118cb82611866565b6118d58185611870565b93506118e083611880565b805f5b838110156119105781516118f7888261189e565b9750611902836118b5565b9250506001810190506118e3565b5085935050505092915050565b5f60a0820190506119305f830188611322565b61193d60208301876114a6565b818103604083015261194f81866118c1565b905061195e6060830185611322565b61196b6080830184611322565b9695505050505050565b5f6040820190506119885f8301856114a6565b61199560208301846113a8565b939250505056fea164736f6c634300081c000a",
}

// KeypersetmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use KeypersetmanagerMetaData.ABI instead.
var KeypersetmanagerABI = KeypersetmanagerMetaData.ABI

// KeypersetmanagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KeypersetmanagerMetaData.Bin instead.
var KeypersetmanagerBin = KeypersetmanagerMetaData.Bin

// DeployKeypersetmanager deploys a new Ethereum contract, binding an instance of Keypersetmanager to it.
func DeployKeypersetmanager(auth *bind.TransactOpts, backend bind.ContractBackend, initializer common.Address) (common.Address, *types.Transaction, *Keypersetmanager, error) {
	parsed, err := KeypersetmanagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KeypersetmanagerBin), backend, initializer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Keypersetmanager{KeypersetmanagerCaller: KeypersetmanagerCaller{contract: contract}, KeypersetmanagerTransactor: KeypersetmanagerTransactor{contract: contract}, KeypersetmanagerFilterer: KeypersetmanagerFilterer{contract: contract}}, nil
}

// Keypersetmanager is an auto generated Go binding around an Ethereum contract.
type Keypersetmanager struct {
	KeypersetmanagerCaller     // Read-only binding to the contract
	KeypersetmanagerTransactor // Write-only binding to the contract
	KeypersetmanagerFilterer   // Log filterer for contract events
}

// KeypersetmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeypersetmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeypersetmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeypersetmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeypersetmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeypersetmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeypersetmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeypersetmanagerSession struct {
	Contract     *Keypersetmanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeypersetmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeypersetmanagerCallerSession struct {
	Contract *KeypersetmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// KeypersetmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeypersetmanagerTransactorSession struct {
	Contract     *KeypersetmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// KeypersetmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeypersetmanagerRaw struct {
	Contract *Keypersetmanager // Generic contract binding to access the raw methods on
}

// KeypersetmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeypersetmanagerCallerRaw struct {
	Contract *KeypersetmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// KeypersetmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeypersetmanagerTransactorRaw struct {
	Contract *KeypersetmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeypersetmanager creates a new instance of Keypersetmanager, bound to a specific deployed contract.
func NewKeypersetmanager(address common.Address, backend bind.ContractBackend) (*Keypersetmanager, error) {
	contract, err := bindKeypersetmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Keypersetmanager{KeypersetmanagerCaller: KeypersetmanagerCaller{contract: contract}, KeypersetmanagerTransactor: KeypersetmanagerTransactor{contract: contract}, KeypersetmanagerFilterer: KeypersetmanagerFilterer{contract: contract}}, nil
}

// NewKeypersetmanagerCaller creates a new read-only instance of Keypersetmanager, bound to a specific deployed contract.
func NewKeypersetmanagerCaller(address common.Address, caller bind.ContractCaller) (*KeypersetmanagerCaller, error) {
	contract, err := bindKeypersetmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerCaller{contract: contract}, nil
}

// NewKeypersetmanagerTransactor creates a new write-only instance of Keypersetmanager, bound to a specific deployed contract.
func NewKeypersetmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*KeypersetmanagerTransactor, error) {
	contract, err := bindKeypersetmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerTransactor{contract: contract}, nil
}

// NewKeypersetmanagerFilterer creates a new log filterer instance of Keypersetmanager, bound to a specific deployed contract.
func NewKeypersetmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*KeypersetmanagerFilterer, error) {
	contract, err := bindKeypersetmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerFilterer{contract: contract}, nil
}

// bindKeypersetmanager binds a generic wrapper to an already deployed contract.
func bindKeypersetmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeypersetmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Keypersetmanager *KeypersetmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Keypersetmanager.Contract.KeypersetmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Keypersetmanager *KeypersetmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.KeypersetmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Keypersetmanager *KeypersetmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.KeypersetmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Keypersetmanager *KeypersetmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Keypersetmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Keypersetmanager *KeypersetmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Keypersetmanager *KeypersetmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Keypersetmanager.Contract.DEFAULTADMINROLE(&_Keypersetmanager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Keypersetmanager.Contract.DEFAULTADMINROLE(&_Keypersetmanager.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerSession) PAUSERROLE() ([32]byte, error) {
	return _Keypersetmanager.Contract.PAUSERROLE(&_Keypersetmanager.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Keypersetmanager.Contract.PAUSERROLE(&_Keypersetmanager.CallOpts)
}

// GetKeyperSetActivationBlock is a free data retrieval call binding the contract method 0x636df979.
//
// Solidity: function getKeyperSetActivationBlock(uint64 index) view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerCaller) GetKeyperSetActivationBlock(opts *bind.CallOpts, index uint64) (uint64, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "getKeyperSetActivationBlock", index)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetKeyperSetActivationBlock is a free data retrieval call binding the contract method 0x636df979.
//
// Solidity: function getKeyperSetActivationBlock(uint64 index) view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerSession) GetKeyperSetActivationBlock(index uint64) (uint64, error) {
	return _Keypersetmanager.Contract.GetKeyperSetActivationBlock(&_Keypersetmanager.CallOpts, index)
}

// GetKeyperSetActivationBlock is a free data retrieval call binding the contract method 0x636df979.
//
// Solidity: function getKeyperSetActivationBlock(uint64 index) view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerCallerSession) GetKeyperSetActivationBlock(index uint64) (uint64, error) {
	return _Keypersetmanager.Contract.GetKeyperSetActivationBlock(&_Keypersetmanager.CallOpts, index)
}

// GetKeyperSetAddress is a free data retrieval call binding the contract method 0xf90f3bed.
//
// Solidity: function getKeyperSetAddress(uint64 index) view returns(address)
func (_Keypersetmanager *KeypersetmanagerCaller) GetKeyperSetAddress(opts *bind.CallOpts, index uint64) (common.Address, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "getKeyperSetAddress", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetKeyperSetAddress is a free data retrieval call binding the contract method 0xf90f3bed.
//
// Solidity: function getKeyperSetAddress(uint64 index) view returns(address)
func (_Keypersetmanager *KeypersetmanagerSession) GetKeyperSetAddress(index uint64) (common.Address, error) {
	return _Keypersetmanager.Contract.GetKeyperSetAddress(&_Keypersetmanager.CallOpts, index)
}

// GetKeyperSetAddress is a free data retrieval call binding the contract method 0xf90f3bed.
//
// Solidity: function getKeyperSetAddress(uint64 index) view returns(address)
func (_Keypersetmanager *KeypersetmanagerCallerSession) GetKeyperSetAddress(index uint64) (common.Address, error) {
	return _Keypersetmanager.Contract.GetKeyperSetAddress(&_Keypersetmanager.CallOpts, index)
}

// GetKeyperSetIndexByBlock is a free data retrieval call binding the contract method 0x035cef15.
//
// Solidity: function getKeyperSetIndexByBlock(uint64 blockNumber) view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerCaller) GetKeyperSetIndexByBlock(opts *bind.CallOpts, blockNumber uint64) (uint64, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "getKeyperSetIndexByBlock", blockNumber)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetKeyperSetIndexByBlock is a free data retrieval call binding the contract method 0x035cef15.
//
// Solidity: function getKeyperSetIndexByBlock(uint64 blockNumber) view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerSession) GetKeyperSetIndexByBlock(blockNumber uint64) (uint64, error) {
	return _Keypersetmanager.Contract.GetKeyperSetIndexByBlock(&_Keypersetmanager.CallOpts, blockNumber)
}

// GetKeyperSetIndexByBlock is a free data retrieval call binding the contract method 0x035cef15.
//
// Solidity: function getKeyperSetIndexByBlock(uint64 blockNumber) view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerCallerSession) GetKeyperSetIndexByBlock(blockNumber uint64) (uint64, error) {
	return _Keypersetmanager.Contract.GetKeyperSetIndexByBlock(&_Keypersetmanager.CallOpts, blockNumber)
}

// GetNumKeyperSets is a free data retrieval call binding the contract method 0xf2e6100a.
//
// Solidity: function getNumKeyperSets() view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerCaller) GetNumKeyperSets(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "getNumKeyperSets")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNumKeyperSets is a free data retrieval call binding the contract method 0xf2e6100a.
//
// Solidity: function getNumKeyperSets() view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerSession) GetNumKeyperSets() (uint64, error) {
	return _Keypersetmanager.Contract.GetNumKeyperSets(&_Keypersetmanager.CallOpts)
}

// GetNumKeyperSets is a free data retrieval call binding the contract method 0xf2e6100a.
//
// Solidity: function getNumKeyperSets() view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerCallerSession) GetNumKeyperSets() (uint64, error) {
	return _Keypersetmanager.Contract.GetNumKeyperSets(&_Keypersetmanager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Keypersetmanager.Contract.GetRoleAdmin(&_Keypersetmanager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Keypersetmanager.Contract.GetRoleAdmin(&_Keypersetmanager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Keypersetmanager *KeypersetmanagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Keypersetmanager *KeypersetmanagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Keypersetmanager.Contract.HasRole(&_Keypersetmanager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Keypersetmanager *KeypersetmanagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Keypersetmanager.Contract.HasRole(&_Keypersetmanager.CallOpts, role, account)
}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Keypersetmanager *KeypersetmanagerCaller) Initializer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "initializer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Keypersetmanager *KeypersetmanagerSession) Initializer() (common.Address, error) {
	return _Keypersetmanager.Contract.Initializer(&_Keypersetmanager.CallOpts)
}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Keypersetmanager *KeypersetmanagerCallerSession) Initializer() (common.Address, error) {
	return _Keypersetmanager.Contract.Initializer(&_Keypersetmanager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Keypersetmanager *KeypersetmanagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Keypersetmanager *KeypersetmanagerSession) Paused() (bool, error) {
	return _Keypersetmanager.Contract.Paused(&_Keypersetmanager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Keypersetmanager *KeypersetmanagerCallerSession) Paused() (bool, error) {
	return _Keypersetmanager.Contract.Paused(&_Keypersetmanager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Keypersetmanager *KeypersetmanagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Keypersetmanager *KeypersetmanagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Keypersetmanager.Contract.SupportsInterface(&_Keypersetmanager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Keypersetmanager *KeypersetmanagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Keypersetmanager.Contract.SupportsInterface(&_Keypersetmanager.CallOpts, interfaceId)
}

// AddKeyperSet is a paid mutator transaction binding the contract method 0xd3877c43.
//
// Solidity: function addKeyperSet(uint64 activationBlock, address keyperSetContract) returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) AddKeyperSet(opts *bind.TransactOpts, activationBlock uint64, keyperSetContract common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "addKeyperSet", activationBlock, keyperSetContract)
}

// AddKeyperSet is a paid mutator transaction binding the contract method 0xd3877c43.
//
// Solidity: function addKeyperSet(uint64 activationBlock, address keyperSetContract) returns()
func (_Keypersetmanager *KeypersetmanagerSession) AddKeyperSet(activationBlock uint64, keyperSetContract common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.AddKeyperSet(&_Keypersetmanager.TransactOpts, activationBlock, keyperSetContract)
}

// AddKeyperSet is a paid mutator transaction binding the contract method 0xd3877c43.
//
// Solidity: function addKeyperSet(uint64 activationBlock, address keyperSetContract) returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) AddKeyperSet(activationBlock uint64, keyperSetContract common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.AddKeyperSet(&_Keypersetmanager.TransactOpts, activationBlock, keyperSetContract)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Keypersetmanager *KeypersetmanagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.GrantRole(&_Keypersetmanager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.GrantRole(&_Keypersetmanager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address pauser) returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, pauser common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "initialize", admin, pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address pauser) returns()
func (_Keypersetmanager *KeypersetmanagerSession) Initialize(admin common.Address, pauser common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.Initialize(&_Keypersetmanager.TransactOpts, admin, pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address pauser) returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) Initialize(admin common.Address, pauser common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.Initialize(&_Keypersetmanager.TransactOpts, admin, pauser)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Keypersetmanager *KeypersetmanagerSession) Pause() (*types.Transaction, error) {
	return _Keypersetmanager.Contract.Pause(&_Keypersetmanager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) Pause() (*types.Transaction, error) {
	return _Keypersetmanager.Contract.Pause(&_Keypersetmanager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Keypersetmanager *KeypersetmanagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.RenounceRole(&_Keypersetmanager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.RenounceRole(&_Keypersetmanager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Keypersetmanager *KeypersetmanagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.RevokeRole(&_Keypersetmanager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.RevokeRole(&_Keypersetmanager.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Keypersetmanager *KeypersetmanagerSession) Unpause() (*types.Transaction, error) {
	return _Keypersetmanager.Contract.Unpause(&_Keypersetmanager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _Keypersetmanager.Contract.Unpause(&_Keypersetmanager.TransactOpts)
}

// KeypersetmanagerKeyperSetAddedIterator is returned from FilterKeyperSetAdded and is used to iterate over the raw logs and unpacked data for KeyperSetAdded events raised by the Keypersetmanager contract.
type KeypersetmanagerKeyperSetAddedIterator struct {
	Event *KeypersetmanagerKeyperSetAdded // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerKeyperSetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerKeyperSetAdded)
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
		it.Event = new(KeypersetmanagerKeyperSetAdded)
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
func (it *KeypersetmanagerKeyperSetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerKeyperSetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerKeyperSetAdded represents a KeyperSetAdded event raised by the Keypersetmanager contract.
type KeypersetmanagerKeyperSetAdded struct {
	ActivationBlock   uint64
	KeyperSetContract common.Address
	Members           []common.Address
	Threshold         uint64
	Eon               uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterKeyperSetAdded is a free log retrieval operation binding the contract event 0xa940387dac06ebd336730f1d14b21629a9d137069a9137e871f95313e1010165.
//
// Solidity: event KeyperSetAdded(uint64 activationBlock, address keyperSetContract, address[] members, uint64 threshold, uint64 eon)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterKeyperSetAdded(opts *bind.FilterOpts) (*KeypersetmanagerKeyperSetAddedIterator, error) {

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "KeyperSetAdded")
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerKeyperSetAddedIterator{contract: _Keypersetmanager.contract, event: "KeyperSetAdded", logs: logs, sub: sub}, nil
}

// WatchKeyperSetAdded is a free log subscription operation binding the contract event 0xa940387dac06ebd336730f1d14b21629a9d137069a9137e871f95313e1010165.
//
// Solidity: event KeyperSetAdded(uint64 activationBlock, address keyperSetContract, address[] members, uint64 threshold, uint64 eon)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchKeyperSetAdded(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerKeyperSetAdded) (event.Subscription, error) {

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "KeyperSetAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerKeyperSetAdded)
				if err := _Keypersetmanager.contract.UnpackLog(event, "KeyperSetAdded", log); err != nil {
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

// ParseKeyperSetAdded is a log parse operation binding the contract event 0xa940387dac06ebd336730f1d14b21629a9d137069a9137e871f95313e1010165.
//
// Solidity: event KeyperSetAdded(uint64 activationBlock, address keyperSetContract, address[] members, uint64 threshold, uint64 eon)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParseKeyperSetAdded(log types.Log) (*KeypersetmanagerKeyperSetAdded, error) {
	event := new(KeypersetmanagerKeyperSetAdded)
	if err := _Keypersetmanager.contract.UnpackLog(event, "KeyperSetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeypersetmanagerKeyperSetAdded0Iterator is returned from FilterKeyperSetAdded0 and is used to iterate over the raw logs and unpacked data for KeyperSetAdded0 events raised by the Keypersetmanager contract.
type KeypersetmanagerKeyperSetAdded0Iterator struct {
	Event *KeypersetmanagerKeyperSetAdded0 // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerKeyperSetAdded0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerKeyperSetAdded0)
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
		it.Event = new(KeypersetmanagerKeyperSetAdded0)
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
func (it *KeypersetmanagerKeyperSetAdded0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerKeyperSetAdded0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerKeyperSetAdded0 represents a KeyperSetAdded0 event raised by the Keypersetmanager contract.
type KeypersetmanagerKeyperSetAdded0 struct {
	ActivationSlot    uint64
	KeyperSetContract common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterKeyperSetAdded0 is a free log retrieval operation binding the contract event 0x6605cb866297050f9f49ae7e0b38e0e4c8178d4b176e24332bc01672818d707b.
//
// Solidity: event KeyperSetAdded(uint64 activationSlot, address keyperSetContract)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterKeyperSetAdded0(opts *bind.FilterOpts) (*KeypersetmanagerKeyperSetAdded0Iterator, error) {

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "KeyperSetAdded0")
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerKeyperSetAdded0Iterator{contract: _Keypersetmanager.contract, event: "KeyperSetAdded0", logs: logs, sub: sub}, nil
}

// WatchKeyperSetAdded0 is a free log subscription operation binding the contract event 0x6605cb866297050f9f49ae7e0b38e0e4c8178d4b176e24332bc01672818d707b.
//
// Solidity: event KeyperSetAdded(uint64 activationSlot, address keyperSetContract)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchKeyperSetAdded0(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerKeyperSetAdded0) (event.Subscription, error) {

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "KeyperSetAdded0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerKeyperSetAdded0)
				if err := _Keypersetmanager.contract.UnpackLog(event, "KeyperSetAdded0", log); err != nil {
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

// ParseKeyperSetAdded0 is a log parse operation binding the contract event 0x6605cb866297050f9f49ae7e0b38e0e4c8178d4b176e24332bc01672818d707b.
//
// Solidity: event KeyperSetAdded(uint64 activationSlot, address keyperSetContract)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParseKeyperSetAdded0(log types.Log) (*KeypersetmanagerKeyperSetAdded0, error) {
	event := new(KeypersetmanagerKeyperSetAdded0)
	if err := _Keypersetmanager.contract.UnpackLog(event, "KeyperSetAdded0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeypersetmanagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Keypersetmanager contract.
type KeypersetmanagerPausedIterator struct {
	Event *KeypersetmanagerPaused // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerPaused)
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
		it.Event = new(KeypersetmanagerPaused)
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
func (it *KeypersetmanagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerPaused represents a Paused event raised by the Keypersetmanager contract.
type KeypersetmanagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterPaused(opts *bind.FilterOpts) (*KeypersetmanagerPausedIterator, error) {

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerPausedIterator{contract: _Keypersetmanager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerPaused) (event.Subscription, error) {

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerPaused)
				if err := _Keypersetmanager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParsePaused(log types.Log) (*KeypersetmanagerPaused, error) {
	event := new(KeypersetmanagerPaused)
	if err := _Keypersetmanager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeypersetmanagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Keypersetmanager contract.
type KeypersetmanagerRoleAdminChangedIterator struct {
	Event *KeypersetmanagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerRoleAdminChanged)
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
		it.Event = new(KeypersetmanagerRoleAdminChanged)
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
func (it *KeypersetmanagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerRoleAdminChanged represents a RoleAdminChanged event raised by the Keypersetmanager contract.
type KeypersetmanagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*KeypersetmanagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerRoleAdminChangedIterator{contract: _Keypersetmanager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerRoleAdminChanged)
				if err := _Keypersetmanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParseRoleAdminChanged(log types.Log) (*KeypersetmanagerRoleAdminChanged, error) {
	event := new(KeypersetmanagerRoleAdminChanged)
	if err := _Keypersetmanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeypersetmanagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Keypersetmanager contract.
type KeypersetmanagerRoleGrantedIterator struct {
	Event *KeypersetmanagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerRoleGranted)
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
		it.Event = new(KeypersetmanagerRoleGranted)
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
func (it *KeypersetmanagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerRoleGranted represents a RoleGranted event raised by the Keypersetmanager contract.
type KeypersetmanagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KeypersetmanagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerRoleGrantedIterator{contract: _Keypersetmanager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerRoleGranted)
				if err := _Keypersetmanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParseRoleGranted(log types.Log) (*KeypersetmanagerRoleGranted, error) {
	event := new(KeypersetmanagerRoleGranted)
	if err := _Keypersetmanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeypersetmanagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Keypersetmanager contract.
type KeypersetmanagerRoleRevokedIterator struct {
	Event *KeypersetmanagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerRoleRevoked)
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
		it.Event = new(KeypersetmanagerRoleRevoked)
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
func (it *KeypersetmanagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerRoleRevoked represents a RoleRevoked event raised by the Keypersetmanager contract.
type KeypersetmanagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KeypersetmanagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerRoleRevokedIterator{contract: _Keypersetmanager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerRoleRevoked)
				if err := _Keypersetmanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParseRoleRevoked(log types.Log) (*KeypersetmanagerRoleRevoked, error) {
	event := new(KeypersetmanagerRoleRevoked)
	if err := _Keypersetmanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeypersetmanagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Keypersetmanager contract.
type KeypersetmanagerUnpausedIterator struct {
	Event *KeypersetmanagerUnpaused // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerUnpaused)
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
		it.Event = new(KeypersetmanagerUnpaused)
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
func (it *KeypersetmanagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerUnpaused represents a Unpaused event raised by the Keypersetmanager contract.
type KeypersetmanagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*KeypersetmanagerUnpausedIterator, error) {

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerUnpausedIterator{contract: _Keypersetmanager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerUnpaused)
				if err := _Keypersetmanager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParseUnpaused(log types.Log) (*KeypersetmanagerUnpaused, error) {
	event := new(KeypersetmanagerUnpaused)
	if err := _Keypersetmanager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
