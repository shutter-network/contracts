// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package inbox

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

// InboxTransaction is an auto generated low-level Go binding around an user-defined struct.
type InboxTransaction struct {
	EncryptedTransaction []byte
	Sender               common.Address
	GasLimit             uint64
	CumulativeGasLimit   uint64
}

// InboxMetaData contains all meta data concerning the Inbox contract.
var InboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_blockGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"initializer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BLOCK_GAS_LIMIT_SETTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SEQUENCER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAW_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clear\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBlockGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTransactions\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structInbox.Transaction[]\",\"components\":[{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cumulativeGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"dao\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sequencer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlockGasLimit\",\"inputs\":[{\"name\":\"newBlockGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitEncryptedTransaction\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"excessFeeRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EncryptedTransactionSubmitted\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"block\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"cumulativeGasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeesWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockAlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockGasLimitExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFunds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferEtherFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedInitializer\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161239b38038061239b8339818101604052810190610031919061013d565b80806001806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508160035f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505061017b565b5f5ffd5b5f67ffffffffffffffff82169050919050565b6100c2816100a6565b81146100cc575f5ffd5b50565b5f815190506100dd816100b9565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61010c826100e3565b9050919050565b61011c81610102565b8114610126575f5ffd5b50565b5f8151905061013781610113565b92915050565b5f5f60408385031215610153576101526100a2565b5b5f610160858286016100cf565b925050602061017185828601610129565b9150509250929050565b612213806101885f395ff3fe608060405260043610610134575f3560e01c80635c975abb116100aa578063a217fddf1161006e578063a217fddf146103dc578063a44e7cfe14610406578063ae0dad6014610422578063d547741f1461044a578063e02023a114610472578063e63ab1e91461049c57610134565b80635c975abb1461030c5780638456cb591461033657806391d148541461034c57806391de737e146103885780639ce110d7146103b257610134565b80633f4ba83a116100fc5780633f4ba83a1461022a5780634842855c14610240578063485cc9551461026a57806351cff8d91461029257806352efea6e146102ba578063552fd4aa146102d057610134565b806301ffc9a714610138578063248a9ca3146101745780632cc8377d146101b05780632f2ff15d146101da57806336568abe14610202575b5f5ffd5b348015610143575f5ffd5b5061015e600480360381019061015991906116c4565b6104c6565b60405161016b9190611709565b60405180910390f35b34801561017f575f5ffd5b5061019a60048036038101906101959190611755565b61053f565b6040516101a7919061178f565b60405180910390f35b3480156101bb575f5ffd5b506101c461055b565b6040516101d191906117ca565b60405180910390f35b3480156101e5575f5ffd5b5061020060048036038101906101fb919061183d565b610577565b005b34801561020d575f5ffd5b506102286004803603810190610223919061183d565b610599565b005b348015610235575f5ffd5b5061023e610614565b005b34801561024b575f5ffd5b5061025461062b565b604051610261919061178f565b60405180910390f35b348015610275575f5ffd5b50610290600480360381019061028b919061187b565b61064f565b005b34801561029d575f5ffd5b506102b860048036038101906102b391906118b9565b610792565b005b3480156102c5575f5ffd5b506102ce610801565b005b3480156102db575f5ffd5b506102f660048036038101906102f1919061190e565b61085e565b6040516103039190611ae2565b60405180910390f35b348015610317575f5ffd5b50610320610a26565b60405161032d9190611709565b60405180910390f35b348015610341575f5ffd5b5061034a610a3b565b005b348015610357575f5ffd5b50610372600480360381019061036d919061183d565b610a70565b60405161037f9190611709565b60405180910390f35b348015610393575f5ffd5b5061039c610ad3565b6040516103a9919061178f565b60405180910390f35b3480156103bd575f5ffd5b506103c6610af7565b6040516103d39190611b11565b60405180910390f35b3480156103e7575f5ffd5b506103f0610b1b565b6040516103fd919061178f565b60405180910390f35b610420600480360381019061041b9190611c56565b610b21565b005b34801561042d575f5ffd5b506104486004803603810190610443919061190e565b610e91565b005b348015610455575f5ffd5b50610470600480360381019061046b919061183d565b610ee7565b005b34801561047d575f5ffd5b50610486610f09565b604051610493919061178f565b60405180910390f35b3480156104a7575f5ffd5b506104b0610f2d565b6040516104bd919061178f565b60405180910390f35b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610538575061053782610f51565b5b9050919050565b5f5f5f8381526020019081526020015f20600101549050919050565b5f60035f9054906101000a900467ffffffffffffffff16905090565b6105808261053f565b61058981610fba565b6105938383610fce565b50505050565b6105a16110b7565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610605576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61060f82826110be565b505050565b5f5f1b61062081610fba565b6106286111a7565b50565b7fac4f1890dc96c9a02330d1fa696648a38f3b282d2449c2d8e6f10507488c84c881565b5f73ffffffffffffffffffffffffffffffffffffffff1660018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036106d4576040517f0dc149f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610759576040517f0d622feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107837fac4f1890dc96c9a02330d1fa696648a38f3b282d2449c2d8e6f10507488c84c882610fce565b5061078e8282611208565b5050565b7f5d8e12c39142ff96d79d04d15d1ba1269e4fe57bb9d26f43523628b34ba108ec6107bc81610fba565b6107c6824761138e565b7f792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c7826040516107f59190611b11565b60405180910390a15050565b7fac4f1890dc96c9a02330d1fa696648a38f3b282d2449c2d8e6f10507488c84c861082b81610fba565b60025f4367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f61085b9190611565565b50565b606060025f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610a1b578382905f5260205f2090600302016040518060800160405290815f820180546108d490611d03565b80601f016020809104026020016040519081016040528092919081815260200182805461090090611d03565b801561094b5780601f106109225761010080835404028352916020019161094b565b820191905f5260205f20905b81548152906001019060200180831161092e57829003601f168201915b50505050508152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600282015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050815260200190600101906108a4565b505050509050919050565b5f60015f9054906101000a900460ff16905090565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a6581610fba565b610a6d611432565b50565b5f5f5f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b7ff265e6f4485879ff06917e0b78d882e64c54d090917c00d06bf8720a31ba125681565b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f1b81565b610b29611493565b4367ffffffffffffffff168467ffffffffffffffff1611610b76576040517fa393699e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f488367ffffffffffffffff16610b8d9190611d69565b905080341015610bc9576040517f356680b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60025f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2090505f818054905090505f5f8267ffffffffffffffff1611610c11575f610c5e565b82600183610c1f9190611daa565b67ffffffffffffffff1681548110610c3a57610c39611de5565b5b905f5260205f2090600302016002015f9054906101000a900467ffffffffffffffff165b905060035f9054906101000a900467ffffffffffffffff1667ffffffffffffffff168682610c8c9190611e12565b67ffffffffffffffff161115610cce576040517fa026cc8600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8260405180608001604052808981526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018867ffffffffffffffff1681526020018884610d189190611e12565b67ffffffffffffffff16815250908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f019081610d5e9190611fed565b506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151816002015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050508767ffffffffffffffff168267ffffffffffffffff167f0ac44a68e7048007d82fa113a6c5d84bec8a110ed9e953d4f234fab3c9ecac5389338a8b87610e4a9190611e12565b8a604051610e5c959493929190612113565b60405180910390a35f8434610e71919061216b565b90505f811115610e8657610e85868261138e565b5b505050505050505050565b7ff265e6f4485879ff06917e0b78d882e64c54d090917c00d06bf8720a31ba1256610ebb81610fba565b8160035f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050565b610ef08261053f565b610ef981610fba565b610f0383836110be565b50505050565b7f5d8e12c39142ff96d79d04d15d1ba1269e4fe57bb9d26f43523628b34ba108ec81565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b610fcb81610fc66110b7565b6114d4565b50565b5f610fd98383610a70565b6110ad5760015f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555061104a6110b7565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600190506110b1565b5f90505b92915050565b5f33905090565b5f6110c98383610a70565b1561119d575f5f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555061113a6110b7565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a4600190506111a1565b5f90505b92915050565b6111af611525565b5f60015f6101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6111f16110b7565b6040516111fe9190611b11565b60405180910390a1565b5f73ffffffffffffffffffffffffffffffffffffffff1660018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361128d576040517f0dc149f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611312576040517f0d622feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61131e5f5f1b83610fce565b506113497f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610fce565b505f6001806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b5f8273ffffffffffffffffffffffffffffffffffffffff16826040516113b3906121cb565b5f6040518083038185875af1925050503d805f81146113ed576040519150601f19603f3d011682016040523d82523d5f602084013e6113f2565b606091505b505090508061142d576040517f560fbc9400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b61143a611493565b6001805f6101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861147c6110b7565b6040516114899190611b11565b60405180910390a1565b61149b610a26565b156114d2576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6114de8282610a70565b6115215780826040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526004016115189291906121df565b60405180910390fd5b5050565b61152d610a26565b611563576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5080545f8255600302905f5260205f20908101906115839190611586565b50565b5b80821115611602575f5f82015f61159e9190611606565b600182015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160146101000a81549067ffffffffffffffff0219169055600282015f6101000a81549067ffffffffffffffff021916905550600301611587565b5090565b50805461161290611d03565b5f825580601f106116235750611640565b601f0160209004905f5260205f209081019061163f9190611643565b5b50565b5b8082111561165a575f815f905550600101611644565b5090565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6116a38161166f565b81146116ad575f5ffd5b50565b5f813590506116be8161169a565b92915050565b5f602082840312156116d9576116d8611667565b5b5f6116e6848285016116b0565b91505092915050565b5f8115159050919050565b611703816116ef565b82525050565b5f60208201905061171c5f8301846116fa565b92915050565b5f819050919050565b61173481611722565b811461173e575f5ffd5b50565b5f8135905061174f8161172b565b92915050565b5f6020828403121561176a57611769611667565b5b5f61177784828501611741565b91505092915050565b61178981611722565b82525050565b5f6020820190506117a25f830184611780565b92915050565b5f67ffffffffffffffff82169050919050565b6117c4816117a8565b82525050565b5f6020820190506117dd5f8301846117bb565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61180c826117e3565b9050919050565b61181c81611802565b8114611826575f5ffd5b50565b5f8135905061183781611813565b92915050565b5f5f6040838503121561185357611852611667565b5b5f61186085828601611741565b925050602061187185828601611829565b9150509250929050565b5f5f6040838503121561189157611890611667565b5b5f61189e85828601611829565b92505060206118af85828601611829565b9150509250929050565b5f602082840312156118ce576118cd611667565b5b5f6118db84828501611829565b91505092915050565b6118ed816117a8565b81146118f7575f5ffd5b50565b5f81359050611908816118e4565b92915050565b5f6020828403121561192357611922611667565b5b5f611930848285016118fa565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6119a482611962565b6119ae818561196c565b93506119be81856020860161197c565b6119c78161198a565b840191505092915050565b6119db81611802565b82525050565b6119ea816117a8565b82525050565b5f608083015f8301518482035f860152611a0a828261199a565b9150506020830151611a1f60208601826119d2565b506040830151611a3260408601826119e1565b506060830151611a4560608601826119e1565b508091505092915050565b5f611a5b83836119f0565b905092915050565b5f602082019050919050565b5f611a7982611939565b611a838185611943565b935083602082028501611a9585611953565b805f5b85811015611ad05784840389528151611ab18582611a50565b9450611abc83611a63565b925060208a01995050600181019050611a98565b50829750879550505050505092915050565b5f6020820190508181035f830152611afa8184611a6f565b905092915050565b611b0b81611802565b82525050565b5f602082019050611b245f830184611b02565b92915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611b688261198a565b810181811067ffffffffffffffff82111715611b8757611b86611b32565b5b80604052505050565b5f611b9961165e565b9050611ba58282611b5f565b919050565b5f67ffffffffffffffff821115611bc457611bc3611b32565b5b611bcd8261198a565b9050602081019050919050565b828183375f83830152505050565b5f611bfa611bf584611baa565b611b90565b905082815260208101848484011115611c1657611c15611b2e565b5b611c21848285611bda565b509392505050565b5f82601f830112611c3d57611c3c611b2a565b5b8135611c4d848260208601611be8565b91505092915050565b5f5f5f5f60808587031215611c6e57611c6d611667565b5b5f611c7b878288016118fa565b945050602085013567ffffffffffffffff811115611c9c57611c9b61166b565b5b611ca887828801611c29565b9350506040611cb9878288016118fa565b9250506060611cca87828801611829565b91505092959194509250565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611d1a57607f821691505b602082108103611d2d57611d2c611cd6565b5b50919050565b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611d7382611d33565b9150611d7e83611d33565b9250828202611d8c81611d33565b91508282048414831517611da357611da2611d3c565b5b5092915050565b5f611db4826117a8565b9150611dbf836117a8565b9250828203905067ffffffffffffffff811115611ddf57611dde611d3c565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f611e1c826117a8565b9150611e27836117a8565b9250828201905067ffffffffffffffff811115611e4757611e46611d3c565b5b92915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611ea97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611e6e565b611eb38683611e6e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611eee611ee9611ee484611d33565b611ecb565b611d33565b9050919050565b5f819050919050565b611f0783611ed4565b611f1b611f1382611ef5565b848454611e7a565b825550505050565b5f5f905090565b611f32611f23565b611f3d818484611efe565b505050565b5b81811015611f6057611f555f82611f2a565b600181019050611f43565b5050565b601f821115611fa557611f7681611e4d565b611f7f84611e5f565b81016020851015611f8e578190505b611fa2611f9a85611e5f565b830182611f42565b50505b505050565b5f82821c905092915050565b5f611fc55f1984600802611faa565b1980831691505092915050565b5f611fdd8383611fb6565b9150826002028217905092915050565b611ff682611962565b67ffffffffffffffff81111561200f5761200e611b32565b5b6120198254611d03565b612024828285611f64565b5f60209050601f831160018114612055575f8415612043578287015190505b61204d8582611fd2565b8655506120b4565b601f19841661206386611e4d565b5f5b8281101561208a57848901518255600182019150602085019450602081019050612065565b868310156120a757848901516120a3601f891682611fb6565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b5f6120d682611962565b6120e081856120bc565b93506120f081856020860161197c565b6120f98161198a565b840191505092915050565b61210d81611d33565b82525050565b5f60a0820190508181035f83015261212b81886120cc565b905061213a6020830187611b02565b61214760408301866117bb565b61215460608301856117bb565b6121616080830184612104565b9695505050505050565b5f61217582611d33565b915061218083611d33565b925082820390508181111561219857612197611d3c565b5b92915050565b5f81905092915050565b50565b5f6121b65f8361219e565b91506121c1826121a8565b5f82019050919050565b5f6121d5826121ab565b9150819050919050565b5f6040820190506121f25f830185611b02565b6121ff6020830184611780565b939250505056fea164736f6c634300081c000a",
}

// InboxABI is the input ABI used to generate the binding from.
// Deprecated: Use InboxMetaData.ABI instead.
var InboxABI = InboxMetaData.ABI

// InboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InboxMetaData.Bin instead.
var InboxBin = InboxMetaData.Bin

// DeployInbox deploys a new Ethereum contract, binding an instance of Inbox to it.
func DeployInbox(auth *bind.TransactOpts, backend bind.ContractBackend, _blockGasLimit uint64, initializer common.Address) (common.Address, *types.Transaction, *Inbox, error) {
	parsed, err := InboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InboxBin), backend, _blockGasLimit, initializer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// Inbox is an auto generated Go binding around an Ethereum contract.
type Inbox struct {
	InboxCaller     // Read-only binding to the contract
	InboxTransactor // Write-only binding to the contract
	InboxFilterer   // Log filterer for contract events
}

// InboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxSession struct {
	Contract     *Inbox            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxCallerSession struct {
	Contract *InboxCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// InboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxTransactorSession struct {
	Contract     *InboxTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxRaw struct {
	Contract *Inbox // Generic contract binding to access the raw methods on
}

// InboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxCallerRaw struct {
	Contract *InboxCaller // Generic read-only contract binding to access the raw methods on
}

// InboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxTransactorRaw struct {
	Contract *InboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInbox creates a new instance of Inbox, bound to a specific deployed contract.
func NewInbox(address common.Address, backend bind.ContractBackend) (*Inbox, error) {
	contract, err := bindInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// NewInboxCaller creates a new read-only instance of Inbox, bound to a specific deployed contract.
func NewInboxCaller(address common.Address, caller bind.ContractCaller) (*InboxCaller, error) {
	contract, err := bindInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxCaller{contract: contract}, nil
}

// NewInboxTransactor creates a new write-only instance of Inbox, bound to a specific deployed contract.
func NewInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxTransactor, error) {
	contract, err := bindInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTransactor{contract: contract}, nil
}

// NewInboxFilterer creates a new log filterer instance of Inbox, bound to a specific deployed contract.
func NewInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxFilterer, error) {
	contract, err := bindInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxFilterer{contract: contract}, nil
}

// bindInbox binds a generic wrapper to an already deployed contract.
func bindInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.InboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transact(opts, method, params...)
}

// BLOCKGASLIMITSETTERROLE is a free data retrieval call binding the contract method 0x91de737e.
//
// Solidity: function BLOCK_GAS_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) BLOCKGASLIMITSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "BLOCK_GAS_LIMIT_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKGASLIMITSETTERROLE is a free data retrieval call binding the contract method 0x91de737e.
//
// Solidity: function BLOCK_GAS_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) BLOCKGASLIMITSETTERROLE() ([32]byte, error) {
	return _Inbox.Contract.BLOCKGASLIMITSETTERROLE(&_Inbox.CallOpts)
}

// BLOCKGASLIMITSETTERROLE is a free data retrieval call binding the contract method 0x91de737e.
//
// Solidity: function BLOCK_GAS_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) BLOCKGASLIMITSETTERROLE() ([32]byte, error) {
	return _Inbox.Contract.BLOCKGASLIMITSETTERROLE(&_Inbox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Inbox.Contract.DEFAULTADMINROLE(&_Inbox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Inbox.Contract.DEFAULTADMINROLE(&_Inbox.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) PAUSERROLE() ([32]byte, error) {
	return _Inbox.Contract.PAUSERROLE(&_Inbox.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Inbox.Contract.PAUSERROLE(&_Inbox.CallOpts)
}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) SEQUENCERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "SEQUENCER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) SEQUENCERROLE() ([32]byte, error) {
	return _Inbox.Contract.SEQUENCERROLE(&_Inbox.CallOpts)
}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) SEQUENCERROLE() ([32]byte, error) {
	return _Inbox.Contract.SEQUENCERROLE(&_Inbox.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) WITHDRAWROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "WITHDRAW_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) WITHDRAWROLE() ([32]byte, error) {
	return _Inbox.Contract.WITHDRAWROLE(&_Inbox.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) WITHDRAWROLE() ([32]byte, error) {
	return _Inbox.Contract.WITHDRAWROLE(&_Inbox.CallOpts)
}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint64)
func (_Inbox *InboxCaller) GetBlockGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "getBlockGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint64)
func (_Inbox *InboxSession) GetBlockGasLimit() (uint64, error) {
	return _Inbox.Contract.GetBlockGasLimit(&_Inbox.CallOpts)
}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint64)
func (_Inbox *InboxCallerSession) GetBlockGasLimit() (uint64, error) {
	return _Inbox.Contract.GetBlockGasLimit(&_Inbox.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Inbox *InboxCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Inbox *InboxSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Inbox.Contract.GetRoleAdmin(&_Inbox.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Inbox *InboxCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Inbox.Contract.GetRoleAdmin(&_Inbox.CallOpts, role)
}

// GetTransactions is a free data retrieval call binding the contract method 0x552fd4aa.
//
// Solidity: function getTransactions(uint64 blockNumber) view returns((bytes,address,uint64,uint64)[])
func (_Inbox *InboxCaller) GetTransactions(opts *bind.CallOpts, blockNumber uint64) ([]InboxTransaction, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "getTransactions", blockNumber)

	if err != nil {
		return *new([]InboxTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]InboxTransaction)).(*[]InboxTransaction)

	return out0, err

}

// GetTransactions is a free data retrieval call binding the contract method 0x552fd4aa.
//
// Solidity: function getTransactions(uint64 blockNumber) view returns((bytes,address,uint64,uint64)[])
func (_Inbox *InboxSession) GetTransactions(blockNumber uint64) ([]InboxTransaction, error) {
	return _Inbox.Contract.GetTransactions(&_Inbox.CallOpts, blockNumber)
}

// GetTransactions is a free data retrieval call binding the contract method 0x552fd4aa.
//
// Solidity: function getTransactions(uint64 blockNumber) view returns((bytes,address,uint64,uint64)[])
func (_Inbox *InboxCallerSession) GetTransactions(blockNumber uint64) ([]InboxTransaction, error) {
	return _Inbox.Contract.GetTransactions(&_Inbox.CallOpts, blockNumber)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Inbox *InboxCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Inbox *InboxSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Inbox.Contract.HasRole(&_Inbox.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Inbox *InboxCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Inbox.Contract.HasRole(&_Inbox.CallOpts, role, account)
}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Inbox *InboxCaller) Initializer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "initializer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Inbox *InboxSession) Initializer() (common.Address, error) {
	return _Inbox.Contract.Initializer(&_Inbox.CallOpts)
}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Inbox *InboxCallerSession) Initializer() (common.Address, error) {
	return _Inbox.Contract.Initializer(&_Inbox.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxSession) Paused() (bool, error) {
	return _Inbox.Contract.Paused(&_Inbox.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxCallerSession) Paused() (bool, error) {
	return _Inbox.Contract.Paused(&_Inbox.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Inbox *InboxCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Inbox *InboxSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Inbox.Contract.SupportsInterface(&_Inbox.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Inbox *InboxCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Inbox.Contract.SupportsInterface(&_Inbox.CallOpts, interfaceId)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns()
func (_Inbox *InboxTransactor) Clear(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "clear")
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns()
func (_Inbox *InboxSession) Clear() (*types.Transaction, error) {
	return _Inbox.Contract.Clear(&_Inbox.TransactOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns()
func (_Inbox *InboxTransactorSession) Clear() (*types.Transaction, error) {
	return _Inbox.Contract.Clear(&_Inbox.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Inbox *InboxTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Inbox *InboxSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.GrantRole(&_Inbox.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Inbox *InboxTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.GrantRole(&_Inbox.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address dao, address sequencer) returns()
func (_Inbox *InboxTransactor) Initialize(opts *bind.TransactOpts, dao common.Address, sequencer common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "initialize", dao, sequencer)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address dao, address sequencer) returns()
func (_Inbox *InboxSession) Initialize(dao common.Address, sequencer common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Initialize(&_Inbox.TransactOpts, dao, sequencer)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address dao, address sequencer) returns()
func (_Inbox *InboxTransactorSession) Initialize(dao common.Address, sequencer common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Initialize(&_Inbox.TransactOpts, dao, sequencer)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxSession) Pause() (*types.Transaction, error) {
	return _Inbox.Contract.Pause(&_Inbox.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxTransactorSession) Pause() (*types.Transaction, error) {
	return _Inbox.Contract.Pause(&_Inbox.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Inbox *InboxTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Inbox *InboxSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.RenounceRole(&_Inbox.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Inbox *InboxTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.RenounceRole(&_Inbox.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Inbox *InboxTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Inbox *InboxSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.RevokeRole(&_Inbox.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Inbox *InboxTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.RevokeRole(&_Inbox.TransactOpts, role, account)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xae0dad60.
//
// Solidity: function setBlockGasLimit(uint64 newBlockGasLimit) returns()
func (_Inbox *InboxTransactor) SetBlockGasLimit(opts *bind.TransactOpts, newBlockGasLimit uint64) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "setBlockGasLimit", newBlockGasLimit)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xae0dad60.
//
// Solidity: function setBlockGasLimit(uint64 newBlockGasLimit) returns()
func (_Inbox *InboxSession) SetBlockGasLimit(newBlockGasLimit uint64) (*types.Transaction, error) {
	return _Inbox.Contract.SetBlockGasLimit(&_Inbox.TransactOpts, newBlockGasLimit)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xae0dad60.
//
// Solidity: function setBlockGasLimit(uint64 newBlockGasLimit) returns()
func (_Inbox *InboxTransactorSession) SetBlockGasLimit(newBlockGasLimit uint64) (*types.Transaction, error) {
	return _Inbox.Contract.SetBlockGasLimit(&_Inbox.TransactOpts, newBlockGasLimit)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0xa44e7cfe.
//
// Solidity: function submitEncryptedTransaction(uint64 blockNumber, bytes encryptedTransaction, uint64 gasLimit, address excessFeeRecipient) payable returns()
func (_Inbox *InboxTransactor) SubmitEncryptedTransaction(opts *bind.TransactOpts, blockNumber uint64, encryptedTransaction []byte, gasLimit uint64, excessFeeRecipient common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "submitEncryptedTransaction", blockNumber, encryptedTransaction, gasLimit, excessFeeRecipient)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0xa44e7cfe.
//
// Solidity: function submitEncryptedTransaction(uint64 blockNumber, bytes encryptedTransaction, uint64 gasLimit, address excessFeeRecipient) payable returns()
func (_Inbox *InboxSession) SubmitEncryptedTransaction(blockNumber uint64, encryptedTransaction []byte, gasLimit uint64, excessFeeRecipient common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.SubmitEncryptedTransaction(&_Inbox.TransactOpts, blockNumber, encryptedTransaction, gasLimit, excessFeeRecipient)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0xa44e7cfe.
//
// Solidity: function submitEncryptedTransaction(uint64 blockNumber, bytes encryptedTransaction, uint64 gasLimit, address excessFeeRecipient) payable returns()
func (_Inbox *InboxTransactorSession) SubmitEncryptedTransaction(blockNumber uint64, encryptedTransaction []byte, gasLimit uint64, excessFeeRecipient common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.SubmitEncryptedTransaction(&_Inbox.TransactOpts, blockNumber, encryptedTransaction, gasLimit, excessFeeRecipient)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxSession) Unpause() (*types.Transaction, error) {
	return _Inbox.Contract.Unpause(&_Inbox.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxTransactorSession) Unpause() (*types.Transaction, error) {
	return _Inbox.Contract.Unpause(&_Inbox.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Inbox *InboxTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "withdraw", recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Inbox *InboxSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Withdraw(&_Inbox.TransactOpts, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Inbox *InboxTransactorSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Withdraw(&_Inbox.TransactOpts, recipient)
}

// InboxEncryptedTransactionSubmittedIterator is returned from FilterEncryptedTransactionSubmitted and is used to iterate over the raw logs and unpacked data for EncryptedTransactionSubmitted events raised by the Inbox contract.
type InboxEncryptedTransactionSubmittedIterator struct {
	Event *InboxEncryptedTransactionSubmitted // Event containing the contract specifics and raw log

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
func (it *InboxEncryptedTransactionSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxEncryptedTransactionSubmitted)
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
		it.Event = new(InboxEncryptedTransactionSubmitted)
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
func (it *InboxEncryptedTransactionSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxEncryptedTransactionSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxEncryptedTransactionSubmitted represents a EncryptedTransactionSubmitted event raised by the Inbox contract.
type InboxEncryptedTransactionSubmitted struct {
	Index                uint64
	Block                uint64
	EncryptedTransaction []byte
	Sender               common.Address
	GasLimit             uint64
	CumulativeGasLimit   uint64
	Fee                  *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterEncryptedTransactionSubmitted is a free log retrieval operation binding the contract event 0x0ac44a68e7048007d82fa113a6c5d84bec8a110ed9e953d4f234fab3c9ecac53.
//
// Solidity: event EncryptedTransactionSubmitted(uint64 indexed index, uint64 indexed block, bytes encryptedTransaction, address sender, uint64 gasLimit, uint64 cumulativeGasLimit, uint256 fee)
func (_Inbox *InboxFilterer) FilterEncryptedTransactionSubmitted(opts *bind.FilterOpts, index []uint64, block []uint64) (*InboxEncryptedTransactionSubmittedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var blockRule []interface{}
	for _, blockItem := range block {
		blockRule = append(blockRule, blockItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "EncryptedTransactionSubmitted", indexRule, blockRule)
	if err != nil {
		return nil, err
	}
	return &InboxEncryptedTransactionSubmittedIterator{contract: _Inbox.contract, event: "EncryptedTransactionSubmitted", logs: logs, sub: sub}, nil
}

// WatchEncryptedTransactionSubmitted is a free log subscription operation binding the contract event 0x0ac44a68e7048007d82fa113a6c5d84bec8a110ed9e953d4f234fab3c9ecac53.
//
// Solidity: event EncryptedTransactionSubmitted(uint64 indexed index, uint64 indexed block, bytes encryptedTransaction, address sender, uint64 gasLimit, uint64 cumulativeGasLimit, uint256 fee)
func (_Inbox *InboxFilterer) WatchEncryptedTransactionSubmitted(opts *bind.WatchOpts, sink chan<- *InboxEncryptedTransactionSubmitted, index []uint64, block []uint64) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var blockRule []interface{}
	for _, blockItem := range block {
		blockRule = append(blockRule, blockItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "EncryptedTransactionSubmitted", indexRule, blockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxEncryptedTransactionSubmitted)
				if err := _Inbox.contract.UnpackLog(event, "EncryptedTransactionSubmitted", log); err != nil {
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

// ParseEncryptedTransactionSubmitted is a log parse operation binding the contract event 0x0ac44a68e7048007d82fa113a6c5d84bec8a110ed9e953d4f234fab3c9ecac53.
//
// Solidity: event EncryptedTransactionSubmitted(uint64 indexed index, uint64 indexed block, bytes encryptedTransaction, address sender, uint64 gasLimit, uint64 cumulativeGasLimit, uint256 fee)
func (_Inbox *InboxFilterer) ParseEncryptedTransactionSubmitted(log types.Log) (*InboxEncryptedTransactionSubmitted, error) {
	event := new(InboxEncryptedTransactionSubmitted)
	if err := _Inbox.contract.UnpackLog(event, "EncryptedTransactionSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxFeesWithdrawnIterator is returned from FilterFeesWithdrawn and is used to iterate over the raw logs and unpacked data for FeesWithdrawn events raised by the Inbox contract.
type InboxFeesWithdrawnIterator struct {
	Event *InboxFeesWithdrawn // Event containing the contract specifics and raw log

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
func (it *InboxFeesWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxFeesWithdrawn)
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
		it.Event = new(InboxFeesWithdrawn)
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
func (it *InboxFeesWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxFeesWithdrawn represents a FeesWithdrawn event raised by the Inbox contract.
type InboxFeesWithdrawn struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesWithdrawn is a free log retrieval operation binding the contract event 0x792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c7.
//
// Solidity: event FeesWithdrawn(address recipient)
func (_Inbox *InboxFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*InboxFeesWithdrawnIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &InboxFeesWithdrawnIterator{contract: _Inbox.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFeesWithdrawn is a free log subscription operation binding the contract event 0x792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c7.
//
// Solidity: event FeesWithdrawn(address recipient)
func (_Inbox *InboxFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *InboxFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxFeesWithdrawn)
				if err := _Inbox.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

// ParseFeesWithdrawn is a log parse operation binding the contract event 0x792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c7.
//
// Solidity: event FeesWithdrawn(address recipient)
func (_Inbox *InboxFilterer) ParseFeesWithdrawn(log types.Log) (*InboxFeesWithdrawn, error) {
	event := new(InboxFeesWithdrawn)
	if err := _Inbox.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Inbox contract.
type InboxPausedIterator struct {
	Event *InboxPaused // Event containing the contract specifics and raw log

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
func (it *InboxPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxPaused)
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
		it.Event = new(InboxPaused)
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
func (it *InboxPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxPaused represents a Paused event raised by the Inbox contract.
type InboxPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Inbox *InboxFilterer) FilterPaused(opts *bind.FilterOpts) (*InboxPausedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &InboxPausedIterator{contract: _Inbox.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Inbox *InboxFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *InboxPaused) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxPaused)
				if err := _Inbox.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Inbox *InboxFilterer) ParsePaused(log types.Log) (*InboxPaused, error) {
	event := new(InboxPaused)
	if err := _Inbox.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Inbox contract.
type InboxRoleAdminChangedIterator struct {
	Event *InboxRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *InboxRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxRoleAdminChanged)
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
		it.Event = new(InboxRoleAdminChanged)
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
func (it *InboxRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxRoleAdminChanged represents a RoleAdminChanged event raised by the Inbox contract.
type InboxRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Inbox *InboxFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*InboxRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &InboxRoleAdminChangedIterator{contract: _Inbox.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Inbox *InboxFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *InboxRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxRoleAdminChanged)
				if err := _Inbox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseRoleAdminChanged(log types.Log) (*InboxRoleAdminChanged, error) {
	event := new(InboxRoleAdminChanged)
	if err := _Inbox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Inbox contract.
type InboxRoleGrantedIterator struct {
	Event *InboxRoleGranted // Event containing the contract specifics and raw log

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
func (it *InboxRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxRoleGranted)
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
		it.Event = new(InboxRoleGranted)
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
func (it *InboxRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxRoleGranted represents a RoleGranted event raised by the Inbox contract.
type InboxRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Inbox *InboxFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*InboxRoleGrantedIterator, error) {

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

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &InboxRoleGrantedIterator{contract: _Inbox.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Inbox *InboxFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *InboxRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxRoleGranted)
				if err := _Inbox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseRoleGranted(log types.Log) (*InboxRoleGranted, error) {
	event := new(InboxRoleGranted)
	if err := _Inbox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Inbox contract.
type InboxRoleRevokedIterator struct {
	Event *InboxRoleRevoked // Event containing the contract specifics and raw log

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
func (it *InboxRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxRoleRevoked)
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
		it.Event = new(InboxRoleRevoked)
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
func (it *InboxRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxRoleRevoked represents a RoleRevoked event raised by the Inbox contract.
type InboxRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Inbox *InboxFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*InboxRoleRevokedIterator, error) {

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

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &InboxRoleRevokedIterator{contract: _Inbox.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Inbox *InboxFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *InboxRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxRoleRevoked)
				if err := _Inbox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseRoleRevoked(log types.Log) (*InboxRoleRevoked, error) {
	event := new(InboxRoleRevoked)
	if err := _Inbox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Inbox contract.
type InboxUnpausedIterator struct {
	Event *InboxUnpaused // Event containing the contract specifics and raw log

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
func (it *InboxUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxUnpaused)
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
		it.Event = new(InboxUnpaused)
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
func (it *InboxUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxUnpaused represents a Unpaused event raised by the Inbox contract.
type InboxUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Inbox *InboxFilterer) FilterUnpaused(opts *bind.FilterOpts) (*InboxUnpausedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &InboxUnpausedIterator{contract: _Inbox.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Inbox *InboxFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *InboxUnpaused) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxUnpaused)
				if err := _Inbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseUnpaused(log types.Log) (*InboxUnpaused, error) {
	event := new(InboxUnpaused)
	if err := _Inbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
