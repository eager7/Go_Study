// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CryptoKitties

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CryptoKittiesABI is the input ABI used to generate the binding from.
const CryptoKittiesABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"startedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nftAddress\",\"type\":\"address\"},{\"name\":\"_cut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// CryptoKittiesBin is the compiled bytecode used for deploying new contracts.
const CryptoKittiesBin = `606060405260008060146101000a81548160ff021916908315150217905550341561002957600080fd5b60405160408061176d833981016040528080519060200190919080519060200190919050506000336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061271082111515156100a157600080fd5b816002819055508290508073ffffffffffffffffffffffffffffffffffffffff166301ffc9a7639a20483d7c0100000000000000000000000000000000000000000000000000000000026000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001915050602060405180830381600087803b151561018357600080fd5b6102c65a03f1151561019457600080fd5b5050506040518051905015156101a957600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050611571806101fc6000396000f3006060604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806327ebe40a146100d55780633f4ba83a14610132578063454a2ab31461015f5780635c975abb146101775780635fd8c710146101a457806378bd7935146101b957806383b5ff8b146102385780638456cb5914610261578063878eb3681461028e5780638da5cb5b146102b157806396b5a75514610306578063c55d0f5614610329578063dd1b7a0f14610360578063f2fde38b146103b5575b600080fd5b34156100e057600080fd5b610130600480803590602001909190803590602001909190803590602001909190803590602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506103ee565b005b341561013d57600080fd5b610145610513565b604051808215151515815260200191505060405180910390f35b61017560048080359060200190919050506105d8565b005b341561018257600080fd5b61018a61060c565b604051808215151515815260200191505060405180910390f35b34156101af57600080fd5b6101b761061f565b005b34156101c457600080fd5b6101da600480803590602001909190505061072a565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020018281526020019550505050505060405180910390f35b341561024357600080fd5b61024b61084d565b6040518082815260200191505060405180910390f35b341561026c57600080fd5b610274610853565b604051808215151515815260200191505060405180910390f35b341561029957600080fd5b6102af600480803590602001909190505061091a565b005b34156102bc57600080fd5b6102c46109ec565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561031157600080fd5b6103276004808035906020019091905050610a11565b005b341561033457600080fd5b61034a6004808035906020019091905050610aad565b6040518082815260200191505060405180910390f35b341561036b57600080fd5b610373610aea565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103c057600080fd5b6103ec600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610b10565b005b6103f66114c7565b600060149054906101000a900460ff1615151561041257600080fd5b846fffffffffffffffffffffffffffffffff168514151561043257600080fd5b836fffffffffffffffffffffffffffffffff168414151561045257600080fd5b8267ffffffffffffffff168314151561046a57600080fd5b6104743387610be5565b151561047f57600080fd5b6104893387610cd1565b60a0604051908101604052808373ffffffffffffffffffffffffffffffffffffffff168152602001866fffffffffffffffffffffffffffffffff168152602001856fffffffffffffffffffffffffffffffff1681526020018467ffffffffffffffff1681526020014267ffffffffffffffff16815250905061050b8682610de1565b505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561057057600080fd5b600060149054906101000a900460ff16151561058b57600080fd5b60008060146101000a81548160ff0219169083151502179055507f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a16001905090565b600060149054906101000a900460ff161515156105f457600080fd5b6105fe8134610fc6565b506106093382611165565b50565b600060149054906101000a900460ff1681565b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806106cd57508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b15156106d857600080fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f1935050505090505050565b60008060008060008060036000888152602001908152602001600020905061075181611241565b151561075c57600080fd5b8060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160010160009054906101000a90046fffffffffffffffffffffffffffffffff168260010160109054906101000a90046fffffffffffffffffffffffffffffffff168360020160009054906101000a900467ffffffffffffffff168460020160089054906101000a900467ffffffffffffffff16836fffffffffffffffffffffffffffffffff169350826fffffffffffffffffffffffffffffffff1692508167ffffffffffffffff1691508067ffffffffffffffff169050955095509550955095505091939590929450565b60025481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108b057600080fd5b600060149054906101000a900460ff161515156108cc57600080fd5b6001600060146101000a81548160ff0219169083151502179055507f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a16001905090565b60008060149054906101000a900460ff16151561093657600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561099157600080fd5b6003600083815260200190815260200160002090506109af81611241565b15156109ba57600080fd5b6109e8828260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661126f565b5050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600360008481526020019081526020016000209150610a3282611241565b1515610a3d57600080fd5b8160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a9e57600080fd5b610aa8838261126f565b505050565b600080600360008481526020019081526020016000209050610ace81611241565b1515610ad957600080fd5b610ae2816112bd565b915050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b6b57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515610be257806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b60008273ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e846000604051602001526040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1515610c9757600080fd5b6102c65a03f11515610ca857600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614905092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd8330846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b1515610dc957600080fd5b6102c65a03f11515610dda57600080fd5b5050505050565b603c816060015167ffffffffffffffff1610151515610dff57600080fd5b806003600084815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060408201518160010160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060608201518160020160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160020160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050507fa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba78282602001516fffffffffffffffffffffffffffffffff1683604001516fffffffffffffffffffffffffffffffff16846060015167ffffffffffffffff166040518085815260200184815260200183815260200182815260200194505050505060405180910390a15050565b6000806000806000806000600360008a81526020019081526020016000209550610fef86611241565b1515610ffa57600080fd5b611003866112bd565b945084881015151561101457600080fd5b8560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169350611044896113b0565b600085111561109e576110568561146c565b925082850391508373ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050151561109d57600080fd5b5b84880390503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015156110e357600080fd5b7f4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2898633604051808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390a184965050505050505092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b151561122957600080fd5b6102c65a03f1151561123a57600080fd5b5050505050565b6000808260020160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16119050919050565b611278826113b0565b6112828183611165565b7f2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df826040518082815260200191505060405180910390a15050565b600080600090508260020160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16421115611316578260020160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16420390505b6113a88360010160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff168460010160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff168560020160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1684611487565b915050919050565b60036000828152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a8154906fffffffffffffffffffffffffffffffff02191690556001820160106101000a8154906fffffffffffffffffffffffffffffffff02191690556002820160006101000a81549067ffffffffffffffff02191690556002820160086101000a81549067ffffffffffffffff0219169055505050565b6000612710600254830281151561147f57fe5b049050919050565b600080600080858510151561149e578693506114bc565b8787039250858584028115156114b057fe5b05915081880190508093505b505050949350505050565b60a060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815250905600a165627a7a72305820030b5a941725951cc77daa1158f74567b93b59abad87a0bed4c2dde8265cf79e0029`

// DeployCryptoKitties deploys a new Ethereum contract, binding an instance of CryptoKitties to it.
func DeployCryptoKitties(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddress common.Address, _cut *big.Int) (common.Address, *types.Transaction, *CryptoKitties, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoKittiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CryptoKittiesBin), backend, _nftAddress, _cut)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CryptoKitties{CryptoKittiesCaller: CryptoKittiesCaller{contract: contract}, CryptoKittiesTransactor: CryptoKittiesTransactor{contract: contract}, CryptoKittiesFilterer: CryptoKittiesFilterer{contract: contract}}, nil
}

// CryptoKitties is an auto generated Go binding around an Ethereum contract.
type CryptoKitties struct {
	CryptoKittiesCaller     // Read-only binding to the contract
	CryptoKittiesTransactor // Write-only binding to the contract
	CryptoKittiesFilterer   // Log filterer for contract events
}

// CryptoKittiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type CryptoKittiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoKittiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CryptoKittiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoKittiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CryptoKittiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoKittiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CryptoKittiesSession struct {
	Contract     *CryptoKitties    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CryptoKittiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CryptoKittiesCallerSession struct {
	Contract *CryptoKittiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CryptoKittiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CryptoKittiesTransactorSession struct {
	Contract     *CryptoKittiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CryptoKittiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type CryptoKittiesRaw struct {
	Contract *CryptoKitties // Generic contract binding to access the raw methods on
}

// CryptoKittiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CryptoKittiesCallerRaw struct {
	Contract *CryptoKittiesCaller // Generic read-only contract binding to access the raw methods on
}

// CryptoKittiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CryptoKittiesTransactorRaw struct {
	Contract *CryptoKittiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCryptoKitties creates a new instance of CryptoKitties, bound to a specific deployed contract.
func NewCryptoKitties(address common.Address, backend bind.ContractBackend) (*CryptoKitties, error) {
	contract, err := bindCryptoKitties(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CryptoKitties{CryptoKittiesCaller: CryptoKittiesCaller{contract: contract}, CryptoKittiesTransactor: CryptoKittiesTransactor{contract: contract}, CryptoKittiesFilterer: CryptoKittiesFilterer{contract: contract}}, nil
}

// NewCryptoKittiesCaller creates a new read-only instance of CryptoKitties, bound to a specific deployed contract.
func NewCryptoKittiesCaller(address common.Address, caller bind.ContractCaller) (*CryptoKittiesCaller, error) {
	contract, err := bindCryptoKitties(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoKittiesCaller{contract: contract}, nil
}

// NewCryptoKittiesTransactor creates a new write-only instance of CryptoKitties, bound to a specific deployed contract.
func NewCryptoKittiesTransactor(address common.Address, transactor bind.ContractTransactor) (*CryptoKittiesTransactor, error) {
	contract, err := bindCryptoKitties(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoKittiesTransactor{contract: contract}, nil
}

// NewCryptoKittiesFilterer creates a new log filterer instance of CryptoKitties, bound to a specific deployed contract.
func NewCryptoKittiesFilterer(address common.Address, filterer bind.ContractFilterer) (*CryptoKittiesFilterer, error) {
	contract, err := bindCryptoKitties(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CryptoKittiesFilterer{contract: contract}, nil
}

// bindCryptoKitties binds a generic wrapper to an already deployed contract.
func bindCryptoKitties(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoKittiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoKitties *CryptoKittiesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CryptoKitties.Contract.CryptoKittiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoKitties *CryptoKittiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoKitties.Contract.CryptoKittiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoKitties *CryptoKittiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoKitties.Contract.CryptoKittiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoKitties *CryptoKittiesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CryptoKitties.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoKitties *CryptoKittiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoKitties.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoKitties *CryptoKittiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoKitties.Contract.contract.Transact(opts, method, params...)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) constant returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_CryptoKitties *CryptoKittiesCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	ret := new(struct {
		Seller        common.Address
		StartingPrice *big.Int
		EndingPrice   *big.Int
		Duration      *big.Int
		StartedAt     *big.Int
	})
	out := ret
	err := _CryptoKitties.contract.Call(opts, out, "getAuction", _tokenId)
	return *ret, err
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) constant returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_CryptoKitties *CryptoKittiesSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _CryptoKitties.Contract.GetAuction(&_CryptoKitties.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) constant returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_CryptoKitties *CryptoKittiesCallerSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _CryptoKitties.Contract.GetAuction(&_CryptoKitties.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) constant returns(uint256)
func (_CryptoKitties *CryptoKittiesCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CryptoKitties.contract.Call(opts, out, "getCurrentPrice", _tokenId)
	return *ret0, err
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) constant returns(uint256)
func (_CryptoKitties *CryptoKittiesSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _CryptoKitties.Contract.GetCurrentPrice(&_CryptoKitties.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) constant returns(uint256)
func (_CryptoKitties *CryptoKittiesCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _CryptoKitties.Contract.GetCurrentPrice(&_CryptoKitties.CallOpts, _tokenId)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_CryptoKitties *CryptoKittiesCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoKitties.contract.Call(opts, out, "nonFungibleContract")
	return *ret0, err
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_CryptoKitties *CryptoKittiesSession) NonFungibleContract() (common.Address, error) {
	return _CryptoKitties.Contract.NonFungibleContract(&_CryptoKitties.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() constant returns(address)
func (_CryptoKitties *CryptoKittiesCallerSession) NonFungibleContract() (common.Address, error) {
	return _CryptoKitties.Contract.NonFungibleContract(&_CryptoKitties.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CryptoKitties *CryptoKittiesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoKitties.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CryptoKitties *CryptoKittiesSession) Owner() (common.Address, error) {
	return _CryptoKitties.Contract.Owner(&_CryptoKitties.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CryptoKitties *CryptoKittiesCallerSession) Owner() (common.Address, error) {
	return _CryptoKitties.Contract.Owner(&_CryptoKitties.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_CryptoKitties *CryptoKittiesCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CryptoKitties.contract.Call(opts, out, "ownerCut")
	return *ret0, err
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_CryptoKitties *CryptoKittiesSession) OwnerCut() (*big.Int, error) {
	return _CryptoKitties.Contract.OwnerCut(&_CryptoKitties.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() constant returns(uint256)
func (_CryptoKitties *CryptoKittiesCallerSession) OwnerCut() (*big.Int, error) {
	return _CryptoKitties.Contract.OwnerCut(&_CryptoKitties.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CryptoKitties *CryptoKittiesCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CryptoKitties.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CryptoKitties *CryptoKittiesSession) Paused() (bool, error) {
	return _CryptoKitties.Contract.Paused(&_CryptoKitties.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CryptoKitties *CryptoKittiesCallerSession) Paused() (bool, error) {
	return _CryptoKitties.Contract.Paused(&_CryptoKitties.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _tokenId) returns()
func (_CryptoKitties *CryptoKittiesTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoKitties.contract.Transact(opts, "bid", _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _tokenId) returns()
func (_CryptoKitties *CryptoKittiesSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoKitties.Contract.Bid(&_CryptoKitties.TransactOpts, _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _tokenId) returns()
func (_CryptoKitties *CryptoKittiesTransactorSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoKitties.Contract.Bid(&_CryptoKitties.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_CryptoKitties *CryptoKittiesTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoKitties.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_CryptoKitties *CryptoKittiesSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoKitties.Contract.CancelAuction(&_CryptoKitties.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_CryptoKitties *CryptoKittiesTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoKitties.Contract.CancelAuction(&_CryptoKitties.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_CryptoKitties *CryptoKittiesTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoKitties.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_CryptoKitties *CryptoKittiesSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoKitties.Contract.CancelAuctionWhenPaused(&_CryptoKitties.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_CryptoKitties *CryptoKittiesTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoKitties.Contract.CancelAuctionWhenPaused(&_CryptoKitties.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration, address _seller) returns()
func (_CryptoKitties *CryptoKittiesTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _CryptoKitties.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration, address _seller) returns()
func (_CryptoKitties *CryptoKittiesSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _CryptoKitties.Contract.CreateAuction(&_CryptoKitties.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration, address _seller) returns()
func (_CryptoKitties *CryptoKittiesTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _CryptoKitties.Contract.CreateAuction(&_CryptoKitties.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_CryptoKitties *CryptoKittiesTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoKitties.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_CryptoKitties *CryptoKittiesSession) Pause() (*types.Transaction, error) {
	return _CryptoKitties.Contract.Pause(&_CryptoKitties.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_CryptoKitties *CryptoKittiesTransactorSession) Pause() (*types.Transaction, error) {
	return _CryptoKitties.Contract.Pause(&_CryptoKitties.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CryptoKitties *CryptoKittiesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CryptoKitties.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CryptoKitties *CryptoKittiesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CryptoKitties.Contract.TransferOwnership(&_CryptoKitties.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CryptoKitties *CryptoKittiesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CryptoKitties.Contract.TransferOwnership(&_CryptoKitties.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_CryptoKitties *CryptoKittiesTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoKitties.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_CryptoKitties *CryptoKittiesSession) Unpause() (*types.Transaction, error) {
	return _CryptoKitties.Contract.Unpause(&_CryptoKitties.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_CryptoKitties *CryptoKittiesTransactorSession) Unpause() (*types.Transaction, error) {
	return _CryptoKitties.Contract.Unpause(&_CryptoKitties.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_CryptoKitties *CryptoKittiesTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoKitties.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_CryptoKitties *CryptoKittiesSession) WithdrawBalance() (*types.Transaction, error) {
	return _CryptoKitties.Contract.WithdrawBalance(&_CryptoKitties.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_CryptoKitties *CryptoKittiesTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _CryptoKitties.Contract.WithdrawBalance(&_CryptoKitties.TransactOpts)
}

// CryptoKittiesAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the CryptoKitties contract.
type CryptoKittiesAuctionCancelledIterator struct {
	Event *CryptoKittiesAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *CryptoKittiesAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoKittiesAuctionCancelled)
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
		it.Event = new(CryptoKittiesAuctionCancelled)
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
func (it *CryptoKittiesAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoKittiesAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoKittiesAuctionCancelled represents a AuctionCancelled event raised by the CryptoKitties contract.
type CryptoKittiesAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_CryptoKitties *CryptoKittiesFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*CryptoKittiesAuctionCancelledIterator, error) {

	logs, sub, err := _CryptoKitties.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &CryptoKittiesAuctionCancelledIterator{contract: _CryptoKitties.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_CryptoKitties *CryptoKittiesFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *CryptoKittiesAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _CryptoKitties.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoKittiesAuctionCancelled)
				if err := _CryptoKitties.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// CryptoKittiesAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the CryptoKitties contract.
type CryptoKittiesAuctionCreatedIterator struct {
	Event *CryptoKittiesAuctionCreated // Event containing the contract specifics and raw log

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
func (it *CryptoKittiesAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoKittiesAuctionCreated)
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
		it.Event = new(CryptoKittiesAuctionCreated)
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
func (it *CryptoKittiesAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoKittiesAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoKittiesAuctionCreated represents a AuctionCreated event raised by the CryptoKitties contract.
type CryptoKittiesAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_CryptoKitties *CryptoKittiesFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*CryptoKittiesAuctionCreatedIterator, error) {

	logs, sub, err := _CryptoKitties.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &CryptoKittiesAuctionCreatedIterator{contract: _CryptoKitties.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_CryptoKitties *CryptoKittiesFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *CryptoKittiesAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _CryptoKitties.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoKittiesAuctionCreated)
				if err := _CryptoKitties.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// CryptoKittiesAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the CryptoKitties contract.
type CryptoKittiesAuctionSuccessfulIterator struct {
	Event *CryptoKittiesAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *CryptoKittiesAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoKittiesAuctionSuccessful)
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
		it.Event = new(CryptoKittiesAuctionSuccessful)
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
func (it *CryptoKittiesAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoKittiesAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoKittiesAuctionSuccessful represents a AuctionSuccessful event raised by the CryptoKitties contract.
type CryptoKittiesAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_CryptoKitties *CryptoKittiesFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*CryptoKittiesAuctionSuccessfulIterator, error) {

	logs, sub, err := _CryptoKitties.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &CryptoKittiesAuctionSuccessfulIterator{contract: _CryptoKitties.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_CryptoKitties *CryptoKittiesFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *CryptoKittiesAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _CryptoKitties.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoKittiesAuctionSuccessful)
				if err := _CryptoKitties.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// CryptoKittiesPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the CryptoKitties contract.
type CryptoKittiesPauseIterator struct {
	Event *CryptoKittiesPause // Event containing the contract specifics and raw log

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
func (it *CryptoKittiesPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoKittiesPause)
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
		it.Event = new(CryptoKittiesPause)
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
func (it *CryptoKittiesPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoKittiesPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoKittiesPause represents a Pause event raised by the CryptoKitties contract.
type CryptoKittiesPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CryptoKitties *CryptoKittiesFilterer) FilterPause(opts *bind.FilterOpts) (*CryptoKittiesPauseIterator, error) {

	logs, sub, err := _CryptoKitties.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &CryptoKittiesPauseIterator{contract: _CryptoKitties.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CryptoKitties *CryptoKittiesFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *CryptoKittiesPause) (event.Subscription, error) {

	logs, sub, err := _CryptoKitties.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoKittiesPause)
				if err := _CryptoKitties.contract.UnpackLog(event, "Pause", log); err != nil {
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

// CryptoKittiesUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the CryptoKitties contract.
type CryptoKittiesUnpauseIterator struct {
	Event *CryptoKittiesUnpause // Event containing the contract specifics and raw log

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
func (it *CryptoKittiesUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoKittiesUnpause)
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
		it.Event = new(CryptoKittiesUnpause)
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
func (it *CryptoKittiesUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoKittiesUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoKittiesUnpause represents a Unpause event raised by the CryptoKitties contract.
type CryptoKittiesUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CryptoKitties *CryptoKittiesFilterer) FilterUnpause(opts *bind.FilterOpts) (*CryptoKittiesUnpauseIterator, error) {

	logs, sub, err := _CryptoKitties.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &CryptoKittiesUnpauseIterator{contract: _CryptoKitties.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CryptoKitties *CryptoKittiesFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *CryptoKittiesUnpause) (event.Subscription, error) {

	logs, sub, err := _CryptoKitties.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoKittiesUnpause)
				if err := _CryptoKitties.contract.UnpackLog(event, "Unpause", log); err != nil {
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
