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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"phaseLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dkgLeadLength\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperSetManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"keyBroadcastContractAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DKG_LEAD_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PHASE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentPhase\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIDKGContract.Phase\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cycleLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dkgStart\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasVoted\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyBroadcastContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyperSetManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitAccusation\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"accusedIndices\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitApology\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"accuserIndices\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"polyEvalData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitDealing\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"commitment\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"polyEvals\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitSuccessVote\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"eonPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"succeeded\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voteCount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AccusationSubmitted\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"accusedIndices\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApologySubmitted\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"accuserIndices\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"},{\"name\":\"polyEvalData\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DKGSucceeded\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"eonPublicKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DealingSubmitted\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"commitment\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"polyEvals\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SuccessVoteSubmitted\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"retryCounter\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"eonPublicKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadySucceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyVoted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAccusation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyEonPublicKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrays\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotKeyperAtIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongDKGContract\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLengthParameter\",\"inputs\":[]}]",
	Bin: "0x610100604052348015610010575f5ffd5b506040516120aa3803806120aa833981810160405281019061003291906101ce565b5f8467ffffffffffffffff16148061005357505f8367ffffffffffffffff16145b1561008a576040517f038016d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8367ffffffffffffffff1660808167ffffffffffffffff16815250508267ffffffffffffffff1660a08167ffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff168152505050505050610232565b5f5ffd5b5f67ffffffffffffffff82169050919050565b61015381610137565b811461015d575f5ffd5b50565b5f8151905061016e8161014a565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61019d82610174565b9050919050565b6101ad81610193565b81146101b7575f5ffd5b50565b5f815190506101c8816101a4565b92915050565b5f5f5f5f608085870312156101e6576101e5610133565b5b5f6101f387828801610160565b945050602061020487828801610160565b9350506040610215878288016101ba565b9250506060610226878288016101ba565b91505092959194509250565b60805160a05160c05160e051611e0d61029d5f395f81816109860152610b7901525f818161082501528181610b5201528181610cde01528181610ec9015261103801525f81816103c30152610d9f01525f818161030201528181610b2d0152610ba00152611e0d5ff3fe608060405234801561000f575f5ffd5b50600436106100e8575f3560e01c8063b209e7171161008a578063eac471a011610064578063eac471a014610248578063eb05c28014610266578063f01275b014610296578063f48f3d14146102b2576100e8565b8063b209e717146101ee578063d4b2cedd1461020c578063df6406311461022a576100e8565b80637642d0de116100c65780637642d0de1461015657806378a1b9db1461018657806382b870d3146101b657806394afc12d146101d2576100e8565b80630a3ec4d6146100ec57806316b89c3f1461011c5780633f9800981461013a575b5f5ffd5b610106600480360381019061010191906111f6565b6102e2565b60405161011391906112a7565b60405180910390f35b6101246103c1565b60405161013191906112cf565b60405180910390f35b610154600480360381019061014f919061139e565b6103e5565b005b610170600480360381019061016b91906114af565b610474565b60405161017d9190611519565b60405180910390f35b6101a0600480360381019061019b9190611532565b6104a9565b6040516101ad9190611519565b60405180910390f35b6101d060048036038101906101cb919061155d565b6104c5565b005b6101ec60048036038101906101e79190611636565b610a66565b005b6101f6610b2b565b60405161020391906112cf565b60405180910390f35b610214610b4f565b60405161022191906116c9565b60405180910390f35b610232610b76565b60405161023f91906116c9565b60405180910390f35b610250610b9d565b60405161025d91906112cf565b60405180910390f35b610280600480360381019061027b9190611715565b610bd0565b60405161028d91906112cf565b60405180910390f35b6102b060048036038101906102ab9190611765565b610c0c565b005b6102cc60048036038101906102c791906111f6565b610cda565b6040516102d99190611834565b60405180910390f35b5f5f6102ee8484610cda565b90505f81436102fd919061187a565b90505f7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1690505f821215610340575f93505050506103bb565b8082121561035457600193505050506103bb565b80600261036191906118ba565b82121561037457600293505050506103bb565b80600361038191906118ba565b82121561039457600393505050506103bb565b8060046103a191906118ba565b8212156103b457600493505050506103bb565b5f93505050505b92915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6103ee87610df0565b6103fa87876001610e5b565b61040387610ec6565b61040d8786611035565b8467ffffffffffffffff168667ffffffffffffffff168867ffffffffffffffff167fa5074728b92791250d48ccdc55266aca7b1cca22f89e1ef2ab91ccd7ae0f337d878787876040516104639493929190611ae2565b60405180910390a450505050505050565b6002602052825f5260405f20602052815f5260405f20602052805f5260405f205f92509250509054906101000a900460ff1681565b5f602052805f5260405f205f915054906101000a900460ff1681565b6104d185856004610e5b565b6104da85610ec6565b6104e48584611035565b5f8282905003610520576040517f59acc21d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156105e7576040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160025f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f8282604051610693929190611b49565b604051809103902090505f6001805f8967ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8481526020019081526020015f205f9054906101000a900467ffffffffffffffff166107149190611b61565b90508060015f8967ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f8481526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508467ffffffffffffffff168667ffffffffffffffff168867ffffffffffffffff167facfc0aa3f6ba36ac0ffda53e00d2cdace2fd4a9e41e2fe3214032794332c37c387876040516107e5929190611b9c565b60405180910390a45f5f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610a5d575f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f90f3bed896040518263ffffffff1660e01b815260040161087c91906112cf565b602060405180830381865afa158015610897573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108bb9190611bd2565b90505f8173ffffffffffffffffffffffffffffffffffffffff1663e75235b86040518163ffffffff1660e01b8152600401602060405180830381865afa158015610907573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061092b9190611c11565b90508067ffffffffffffffff168367ffffffffffffffff1610610a5a5760015f5f8b67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663daade8e88a88886040518463ffffffff1660e01b81526004016109e193929190611c3c565b5f604051808303815f87803b1580156109f8575f5ffd5b505af1925050508015610a09575060015b508767ffffffffffffffff168967ffffffffffffffff167fe6c9bd3daa50b9218615605cdf333410d405b7f3b8dbeee5a2373769679aa9f58888604051610a51929190611b9c565b60405180910390a35b50505b50505050505050565b610a6f85610df0565b610a7b85856002610e5b565b610a8485610ec6565b610a8e8584611035565b5f8282905003610aca576040517f2eeffe0b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8267ffffffffffffffff168467ffffffffffffffff168667ffffffffffffffff167fb5cc05e740a503c13e1b7ac3ad2d5abf41b91d2e5fca79455d34e515fb70849a8585604051610b1c929190611d28565b60405180910390a45050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f7f0000000000000000000000000000000000000000000000000000000000000000905090565b5f7f0000000000000000000000000000000000000000000000000000000000000000905090565b5f7f00000000000000000000000000000000000000000000000000000000000000006004610bcb9190611d4a565b905090565b6001602052825f5260405f20602052815f5260405f20602052805f5260405f205f92509250509054906101000a900467ffffffffffffffff1681565b610c1587610df0565b610c2187876003610e5b565b610c2a87610ec6565b610c348786611035565b818190508484905014610c73576040517fa121188700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8467ffffffffffffffff168667ffffffffffffffff168867ffffffffffffffff167f62729f0820cf3ab60fbd63a51cdeae0554d0ab24f843b86bd0d2591594d6913687878787604051610cc99493929190611d86565b60405180910390a450505050505050565b5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663636df979856040518263ffffffff1660e01b8152600401610d3591906112cf565b602060405180830381865afa158015610d50573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d749190611c11565b9050610d7e610b9d565b67ffffffffffffffff168367ffffffffffffffff16610d9d91906118ba565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168267ffffffffffffffff16610ddd919061187a565b610de79190611dbf565b91505092915050565b5f5f8267ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610e58576040517fe0f5ec9d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b806004811115610e6e57610e6d611234565b5b610e7884846102e2565b6004811115610e8a57610e89611234565b5b14610ec1576040517fe2586bcc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f90f3bed836040518263ffffffff1660e01b8152600401610f2091906112cf565b602060405180830381865afa158015610f3b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f5f9190611bd2565b90503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff166266f0a86040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fc0573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fe49190611bd2565b73ffffffffffffffffffffffffffffffffffffffff1614611031576040517ff62f18e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f90f3bed846040518263ffffffff1660e01b815260040161108f91906112cf565b602060405180830381865afa1580156110aa573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110ce9190611bd2565b90503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16632e8e6cad846040518263ffffffff1660e01b815260040161112091906112cf565b602060405180830381865afa15801561113b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061115f9190611bd2565b73ffffffffffffffffffffffffffffffffffffffff16146111ac576040517f7bf9580300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff82169050919050565b6111d5816111b9565b81146111df575f5ffd5b50565b5f813590506111f0816111cc565b92915050565b5f5f6040838503121561120c5761120b6111b1565b5b5f611219858286016111e2565b925050602061122a858286016111e2565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6005811061127257611271611234565b5b50565b5f81905061128282611261565b919050565b5f61129182611275565b9050919050565b6112a181611287565b82525050565b5f6020820190506112ba5f830184611298565b92915050565b6112c9816111b9565b82525050565b5f6020820190506112e25f8301846112c0565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112611309576113086112e8565b5b8235905067ffffffffffffffff811115611326576113256112ec565b5b602083019150836001820283011115611342576113416112f0565b5b9250929050565b5f5f83601f84011261135e5761135d6112e8565b5b8235905067ffffffffffffffff81111561137b5761137a6112ec565b5b602083019150836020820283011115611397576113966112f0565b5b9250929050565b5f5f5f5f5f5f5f60a0888a0312156113b9576113b86111b1565b5b5f6113c68a828b016111e2565b97505060206113d78a828b016111e2565b96505060406113e88a828b016111e2565b955050606088013567ffffffffffffffff811115611409576114086111b5565b5b6114158a828b016112f4565b9450945050608088013567ffffffffffffffff811115611438576114376111b5565b5b6114448a828b01611349565b925092505092959891949750929550565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61147e82611455565b9050919050565b61148e81611474565b8114611498575f5ffd5b50565b5f813590506114a981611485565b92915050565b5f5f5f606084860312156114c6576114c56111b1565b5b5f6114d3868287016111e2565b93505060206114e4868287016111e2565b92505060406114f58682870161149b565b9150509250925092565b5f8115159050919050565b611513816114ff565b82525050565b5f60208201905061152c5f83018461150a565b92915050565b5f60208284031215611547576115466111b1565b5b5f611554848285016111e2565b91505092915050565b5f5f5f5f5f60808688031215611576576115756111b1565b5b5f611583888289016111e2565b9550506020611594888289016111e2565b94505060406115a5888289016111e2565b935050606086013567ffffffffffffffff8111156115c6576115c56111b5565b5b6115d2888289016112f4565b92509250509295509295909350565b5f5f83601f8401126115f6576115f56112e8565b5b8235905067ffffffffffffffff811115611613576116126112ec565b5b60208301915083602082028301111561162f5761162e6112f0565b5b9250929050565b5f5f5f5f5f6080868803121561164f5761164e6111b1565b5b5f61165c888289016111e2565b955050602061166d888289016111e2565b945050604061167e888289016111e2565b935050606086013567ffffffffffffffff81111561169f5761169e6111b5565b5b6116ab888289016115e1565b92509250509295509295909350565b6116c381611474565b82525050565b5f6020820190506116dc5f8301846116ba565b92915050565b5f819050919050565b6116f4816116e2565b81146116fe575f5ffd5b50565b5f8135905061170f816116eb565b92915050565b5f5f5f6060848603121561172c5761172b6111b1565b5b5f611739868287016111e2565b935050602061174a868287016111e2565b925050604061175b86828701611701565b9150509250925092565b5f5f5f5f5f5f5f60a0888a0312156117805761177f6111b1565b5b5f61178d8a828b016111e2565b975050602061179e8a828b016111e2565b96505060406117af8a828b016111e2565b955050606088013567ffffffffffffffff8111156117d0576117cf6111b5565b5b6117dc8a828b016115e1565b9450945050608088013567ffffffffffffffff8111156117ff576117fe6111b5565b5b61180b8a828b01611349565b925092505092959891949750929550565b5f819050919050565b61182e8161181c565b82525050565b5f6020820190506118475f830184611825565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6118848261181c565b915061188f8361181c565b925082820390508181125f8412168282135f8512151617156118b4576118b361184d565b5b92915050565b5f6118c48261181c565b91506118cf8361181c565b92508282026118dd8161181c565b91507f800000000000000000000000000000000000000000000000000000000000000084145f841216156119145761191361184d565b5b82820584148315176119295761192861184d565b5b5092915050565b5f82825260208201905092915050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f6119698385611930565b9350611976838584611940565b61197f8361194e565b840190509392505050565b5f82825260208201905092915050565b5f819050919050565b5f82825260208201905092915050565b5f6119be83856119a3565b93506119cb838584611940565b6119d48361194e565b840190509392505050565b5f6119eb8484846119b3565b90509392505050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83356001602003843603038112611a1c57611a1b6119fc565b5b83810192508235915060208301925067ffffffffffffffff821115611a4457611a436119f4565b5b600182023603831315611a5a57611a596119f8565b5b509250929050565b5f602082019050919050565b5f611a79838561198a565b935083602084028501611a8b8461199a565b805f5b87811015611ad0578484038952611aa58284611a00565b611ab08682846119df565b9550611abb84611a62565b935060208b019a505050600181019050611a8e565b50829750879450505050509392505050565b5f6040820190508181035f830152611afb81868861195e565b90508181036020830152611b10818486611a6e565b905095945050505050565b5f81905092915050565b5f611b308385611b1b565b9350611b3d838584611940565b82840190509392505050565b5f611b55828486611b25565b91508190509392505050565b5f611b6b826111b9565b9150611b76836111b9565b9250828201905067ffffffffffffffff811115611b9657611b9561184d565b5b92915050565b5f6020820190508181035f830152611bb581848661195e565b90509392505050565b5f81519050611bcc81611485565b92915050565b5f60208284031215611be757611be66111b1565b5b5f611bf484828501611bbe565b91505092915050565b5f81519050611c0b816111cc565b92915050565b5f60208284031215611c2657611c256111b1565b5b5f611c3384828501611bfd565b91505092915050565b5f604082019050611c4f5f8301866112c0565b8181036020830152611c6281848661195e565b9050949350505050565b5f82825260208201905092915050565b5f819050919050565b611c8e816111b9565b82525050565b5f611c9f8383611c85565b60208301905092915050565b5f611cb960208401846111e2565b905092915050565b5f602082019050919050565b5f611cd88385611c6c565b9350611ce382611c7c565b805f5b85811015611d1b57611cf88284611cab565b611d028882611c94565b9750611d0d83611cc1565b925050600181019050611ce6565b5085925050509392505050565b5f6020820190508181035f830152611d41818486611ccd565b90509392505050565b5f611d54826111b9565b9150611d5f836111b9565b9250828202611d6d816111b9565b9150808214611d7f57611d7e61184d565b5b5092915050565b5f6040820190508181035f830152611d9f818688611ccd565b90508181036020830152611db4818486611a6e565b905095945050505050565b5f611dc98261181c565b9150611dd48361181c565b92508282019050828112155f8312168382125f841215161715611dfa57611df961184d565b5b9291505056fea164736f6c634300081c000a",
}

// DkgcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use DkgcontractMetaData.ABI instead.
var DkgcontractABI = DkgcontractMetaData.ABI

// DkgcontractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DkgcontractMetaData.Bin instead.
var DkgcontractBin = DkgcontractMetaData.Bin

// DeployDkgcontract deploys a new Ethereum contract, binding an instance of Dkgcontract to it.
func DeployDkgcontract(auth *bind.TransactOpts, backend bind.ContractBackend, phaseLength uint64, dkgLeadLength uint64, keyperSetManagerAddress common.Address, keyBroadcastContractAddress common.Address) (common.Address, *types.Transaction, *Dkgcontract, error) {
	parsed, err := DkgcontractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DkgcontractBin), backend, phaseLength, dkgLeadLength, keyperSetManagerAddress, keyBroadcastContractAddress)
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
