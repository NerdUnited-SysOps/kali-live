// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_approver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bridgedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bridgedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"approvedMessage\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"notarizeMessage\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"Notarized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bridgedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePaid\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"submission\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"verifyHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610cb1380380610cb183398101604081905261002f91610156565b6001600160a01b03851661004257600080fd5b6001600160a01b03841661005557600080fd5b600080546001600160a01b03199081166001600160a01b039788161790915560018054821695871695909517909455600280549094169290941691909117909155600655604080517fc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e6020808301919091527fbf362ddb5d91b26633f508c4b8932aac8d6bcea5c25f6f1724c15e83952116d7828401527fae209a0b48f21c054280f2455d32cf309387644879d9acbd8ffc19916381188560608301526080808301949094528251808303909401845260a090910190915281519101206005556101ac565b80516001600160a01b038116811461015157600080fd5b919050565b600080600080600060a0868803121561016e57600080fd5b6101778661013a565b94506101856020870161013a565b93506101936040870161013a565b6060870151608090970151959894975095949392505050565b610af6806101bb6000396000f3fe60806040526004361061002d5760003560e01c80638f2a8b3314610044578063ddca3f43146100645761003c565b3661003c5761003a61008c565b005b61003a61008c565b34801561005057600080fd5b5061003a61005f3660046108fb565b6101d3565b34801561007057600080fd5b5061007a60065481565b60405190815260200160405180910390f35b32331461009857600080fd5b60065434116100a657600080fd5b6000600654346100b691906109bf565b9050600081116100c557600080fd5b6040805143602082015242918101919091523360608201526080810182905260009060a00160408051601f198184030181529181528151602092830120600081815260039093529120600101549091506001600160a01b03161561012857600080fd5b60408051606081018252838152336020808301828152600084860181815287825260038452908690209451855590516001909401805491511515600160a01b0274ffffffffffffffffffffffffffffffffffffffffff199092166001600160a01b0390951694909417179092556006548351868152928301529183917fc175201a7969fcffd556969f39dd05f68a7b2f985c83e95a9a01aafb7b8bf188910160405180910390a35050565b6000838152600360205260409020600101546001600160a01b03166101f757600080fd5b6000546001600160a01b031633141561030f57600083815260036020526040902060010154600160a01b900460ff161561023057600080fd5b61023b828285610584565b61024457600080fd5b600083815260036020908152604080832060010180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16600160a01b1790556004825290912082516102999284019061080f565b5060008381526003602052604080822060018101549054925491516001600160a01b039182169387937f1d34f6e322b862a5cd4c93e569a792834508633caede2653b219b9e12ac6a7f893610302939116909182526001600160a01b0316602082015260400190565b60405180910390a3505050565b6001546001600160a01b031633141561057f5760008381526003602090815260408083208151606081018352815481526001909101546001600160a01b03811682850152600160a01b900460ff1615158183015286845260049092528220805491929161037b906109e4565b80601f01602080910402602001604051908101604052809291908181526020018280546103a7906109e4565b80156103f45780601f106103c9576101008083540402835291602001916103f4565b820191906000526020600020905b8154815290600101906020018083116103d757829003601f168201915b50505050509050816040015161040957600080fd5b610414848487610584565b61041d57600080fd5b6000858152600360209081526040808320838155600101805474ffffffffffffffffffffffffffffffffffffffffff191690556004909152812061046091610893565b60025460065460405160009283926001600160a01b03909116918381818185875af1925050503d80600081146104b2576040519150601f19603f3d011682016040523d82523d6000602084013e6104b7565b606091505b5091509150816104c657600080fd5b8351604051600091908281818185825af1925050503d8060008114610507576040519150601f19603f3d011682016040523d82523d6000602084013e61050c565b606091505b5090925090508161051c57600080fd5b602084015184516001546040516001600160a01b03938416937fb53a4163819a38354d70815c4f3a89de1b9c2cc32f076d0374f0905d6100365f9361056e9390928d928d928b928e9290911690610a6c565b60405180910390a250505050505050565b600080fd5b600081815260036020908152604080832081516060808201845282548083526001909301546001600160a01b038116838701819052600160a01b90910460ff1615158386015284517f3c1316138cd3c347ee70454f6b80926b84604f3f07629dbed9845a8c06cc9ea381880152808601899052918201526080808201939093528351808203909301835260a08101909352815191909301206005547f190100000000000000000000000000000000000000000000000000000000000060c084015260c283015260e282018190529083906101020160405160208183030381529060405280519060200120905080871461067c57600080fd5b60408301516106a75761068f87876106d9565b6000546001600160a01b0390811691161493506106cf565b8260400151156106cf576106bb87876106d9565b6001546001600160a01b0390811691161493505b5050509392505050565b6000806000806106e885610705565b919450925090506106fb86848484610734565b9695505050505050565b6000806000835160411461071857600080fd5b5050506020810151604082015160609092015160001a92909190565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561076657506000610807565b8360ff16601b1415801561077e57508360ff16601c14155b1561078b57506000610807565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa1580156107df573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610804576000915050610807565b90505b949350505050565b82805461081b906109e4565b90600052602060002090601f01602090048101928261083d5760008555610883565b82601f1061085657805160ff1916838001178555610883565b82800160010185558215610883579182015b82811115610883578251825591602001919060010190610868565b5061088f9291506108d0565b5090565b50805461089f906109e4565b6000825580601f106108af575050565b601f0160209004906000526020600020908101906108cd91906108d0565b50565b5b8082111561088f57600081556001016108d1565b634e487b7160e01b600052604160045260246000fd5b60008060006060848603121561091057600080fd5b8335925060208401359150604084013567ffffffffffffffff8082111561093657600080fd5b818601915086601f83011261094a57600080fd5b81358181111561095c5761095c6108e5565b604051601f8201601f19908116603f01168101908382118183101715610984576109846108e5565b8160405282815289602084870101111561099d57600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b6000828210156109df57634e487b7160e01b600052601160045260246000fd5b500390565b600181811c908216806109f857607f821691505b60208210811415610a1957634e487b7160e01b600052602260045260246000fd5b50919050565b6000815180845260005b81811015610a4557602081850181015186830182015201610a29565b81811115610a57576000602083870101525b50601f01601f19169290920160200192915050565b86815285602082015284604082015260c060608201526000610a9160c0830186610a1f565b8281036080840152610aa38186610a1f565b9150506001600160a01b03831660a083015297965050505050505056fea26469706673582212209414163ce51955bbe32b21f3b54a4a6614304141c6d07a9490a0b8d1553e3b4964736f6c634300080a0033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _approver common.Address, _notary common.Address, _feeReceiver common.Address, _fee *big.Int, _chainId *big.Int) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend, _approver, _notary, _feeReceiver, _fee, _chainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Bridge *BridgeCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Bridge *BridgeSession) Fee() (*big.Int, error) {
	return _Bridge.Contract.Fee(&_Bridge.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Bridge *BridgeCallerSession) Fee() (*big.Int, error) {
	return _Bridge.Contract.Fee(&_Bridge.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x8f2a8b33.
//
// Solidity: function approve(bytes32 submission, bytes32 verifyHash, bytes sig) returns()
func (_Bridge *BridgeTransactor) Approve(opts *bind.TransactOpts, submission [32]byte, verifyHash [32]byte, sig []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "approve", submission, verifyHash, sig)
}

// Approve is a paid mutator transaction binding the contract method 0x8f2a8b33.
//
// Solidity: function approve(bytes32 submission, bytes32 verifyHash, bytes sig) returns()
func (_Bridge *BridgeSession) Approve(submission [32]byte, verifyHash [32]byte, sig []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Approve(&_Bridge.TransactOpts, submission, verifyHash, sig)
}

// Approve is a paid mutator transaction binding the contract method 0x8f2a8b33.
//
// Solidity: function approve(bytes32 submission, bytes32 verifyHash, bytes sig) returns()
func (_Bridge *BridgeTransactorSession) Approve(submission [32]byte, verifyHash [32]byte, sig []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Approve(&_Bridge.TransactOpts, submission, verifyHash, sig)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Bridge *BridgeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Bridge *BridgeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Fallback(&_Bridge.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Bridge *BridgeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Fallback(&_Bridge.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// BridgeApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the Bridge contract.
type BridgeApprovedIterator struct {
	Event *BridgeApproved // Event containing the contract specifics and raw log

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
func (it *BridgeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeApproved)
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
		it.Event = new(BridgeApproved)
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
func (it *BridgeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeApproved represents a Approved event raised by the Bridge contract.
type BridgeApproved struct {
	Key           [32]byte
	Sender        common.Address
	BridgedAmount *big.Int
	Approver      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0x1d34f6e322b862a5cd4c93e569a792834508633caede2653b219b9e12ac6a7f8.
//
// Solidity: event Approved(bytes32 indexed key, address indexed sender, uint256 bridgedAmount, address approver)
func (_Bridge *BridgeFilterer) FilterApproved(opts *bind.FilterOpts, key [][32]byte, sender []common.Address) (*BridgeApprovedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Approved", keyRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeApprovedIterator{contract: _Bridge.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0x1d34f6e322b862a5cd4c93e569a792834508633caede2653b219b9e12ac6a7f8.
//
// Solidity: event Approved(bytes32 indexed key, address indexed sender, uint256 bridgedAmount, address approver)
func (_Bridge *BridgeFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *BridgeApproved, key [][32]byte, sender []common.Address) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Approved", keyRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeApproved)
				if err := _Bridge.contract.UnpackLog(event, "Approved", log); err != nil {
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

// ParseApproved is a log parse operation binding the contract event 0x1d34f6e322b862a5cd4c93e569a792834508633caede2653b219b9e12ac6a7f8.
//
// Solidity: event Approved(bytes32 indexed key, address indexed sender, uint256 bridgedAmount, address approver)
func (_Bridge *BridgeFilterer) ParseApproved(log types.Log) (*BridgeApproved, error) {
	event := new(BridgeApproved)
	if err := _Bridge.contract.UnpackLog(event, "Approved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeNotarizedIterator is returned from FilterNotarized and is used to iterate over the raw logs and unpacked data for Notarized events raised by the Bridge contract.
type BridgeNotarizedIterator struct {
	Event *BridgeNotarized // Event containing the contract specifics and raw log

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
func (it *BridgeNotarizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeNotarized)
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
		it.Event = new(BridgeNotarized)
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
func (it *BridgeNotarizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeNotarizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeNotarized represents a Notarized event raised by the Bridge contract.
type BridgeNotarized struct {
	Sender          common.Address
	BridgedAmount   *big.Int
	Nonce           [32]byte
	MessageHash     [32]byte
	ApprovedMessage []byte
	NotarizeMessage []byte
	Notary          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNotarized is a free log retrieval operation binding the contract event 0xb53a4163819a38354d70815c4f3a89de1b9c2cc32f076d0374f0905d6100365f.
//
// Solidity: event Notarized(address indexed sender, uint256 bridgedAmount, bytes32 nonce, bytes32 messageHash, bytes approvedMessage, bytes notarizeMessage, address notary)
func (_Bridge *BridgeFilterer) FilterNotarized(opts *bind.FilterOpts, sender []common.Address) (*BridgeNotarizedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Notarized", senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeNotarizedIterator{contract: _Bridge.contract, event: "Notarized", logs: logs, sub: sub}, nil
}

// WatchNotarized is a free log subscription operation binding the contract event 0xb53a4163819a38354d70815c4f3a89de1b9c2cc32f076d0374f0905d6100365f.
//
// Solidity: event Notarized(address indexed sender, uint256 bridgedAmount, bytes32 nonce, bytes32 messageHash, bytes approvedMessage, bytes notarizeMessage, address notary)
func (_Bridge *BridgeFilterer) WatchNotarized(opts *bind.WatchOpts, sink chan<- *BridgeNotarized, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Notarized", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeNotarized)
				if err := _Bridge.contract.UnpackLog(event, "Notarized", log); err != nil {
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

// ParseNotarized is a log parse operation binding the contract event 0xb53a4163819a38354d70815c4f3a89de1b9c2cc32f076d0374f0905d6100365f.
//
// Solidity: event Notarized(address indexed sender, uint256 bridgedAmount, bytes32 nonce, bytes32 messageHash, bytes approvedMessage, bytes notarizeMessage, address notary)
func (_Bridge *BridgeFilterer) ParseNotarized(log types.Log) (*BridgeNotarized, error) {
	event := new(BridgeNotarized)
	if err := _Bridge.contract.UnpackLog(event, "Notarized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Bridge contract.
type BridgeReceivedIterator struct {
	Event *BridgeReceived // Event containing the contract specifics and raw log

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
func (it *BridgeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeReceived)
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
		it.Event = new(BridgeReceived)
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
func (it *BridgeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeReceived represents a Received event raised by the Bridge contract.
type BridgeReceived struct {
	Key           [32]byte
	Sender        common.Address
	BridgedAmount *big.Int
	FeePaid       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0xc175201a7969fcffd556969f39dd05f68a7b2f985c83e95a9a01aafb7b8bf188.
//
// Solidity: event Received(bytes32 indexed key, address indexed sender, uint256 bridgedAmount, uint256 feePaid)
func (_Bridge *BridgeFilterer) FilterReceived(opts *bind.FilterOpts, key [][32]byte, sender []common.Address) (*BridgeReceivedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Received", keyRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeReceivedIterator{contract: _Bridge.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0xc175201a7969fcffd556969f39dd05f68a7b2f985c83e95a9a01aafb7b8bf188.
//
// Solidity: event Received(bytes32 indexed key, address indexed sender, uint256 bridgedAmount, uint256 feePaid)
func (_Bridge *BridgeFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *BridgeReceived, key [][32]byte, sender []common.Address) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Received", keyRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeReceived)
				if err := _Bridge.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0xc175201a7969fcffd556969f39dd05f68a7b2f985c83e95a9a01aafb7b8bf188.
//
// Solidity: event Received(bytes32 indexed key, address indexed sender, uint256 bridgedAmount, uint256 feePaid)
func (_Bridge *BridgeFilterer) ParseReceived(log types.Log) (*BridgeReceived, error) {
	event := new(BridgeReceived)
	if err := _Bridge.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
