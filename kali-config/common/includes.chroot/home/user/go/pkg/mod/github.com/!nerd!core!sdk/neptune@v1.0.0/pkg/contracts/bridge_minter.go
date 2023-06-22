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

// BridgeMinterMetaData contains all meta data concerning the BridgeMinter contract.
var BridgeMinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_approver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Bridged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"confirmed\",\"type\":\"bool\"}],\"name\":\"TransferOwnership\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bridgedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"approvedMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"notarizedMessage\",\"type\":\"bytes\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161074238038061074283398101604081905261002f91610167565b6001600160a01b03841661004257600080fd5b6001600160a01b03831661005557600080fd5b6001600160a01b03821661006857600080fd5b600180546001600160a01b03199081166001600160a01b03968716179091556000805482169486169490941790935560028054909316919093161790556003819055604080517fc2f8787176b8ac6bf7215b4adcc1e069bf4ab82d9ab1df05a57a91d425935b6e6020808301919091527fbf362ddb5d91b26633f508c4b8932aac8d6bcea5c25f6f1724c15e83952116d7828401527fae209a0b48f21c054280f2455d32cf309387644879d9acbd8ffc19916381188560608301526080808301949094528251808303909401845260a090910190915281519101206004556101b2565b80516001600160a01b038116811461016257600080fd5b919050565b6000806000806080858703121561017d57600080fd5b6101868561014b565b93506101946020860161014b565b92506101a26040860161014b565b6060959095015193969295505050565b610581806101c16000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80639e07f0db14610030575b600080fd5b61004361003e3660046104a0565b610045565b005b600086815260056020526040902054869060ff161561006357600080fd5b600254600160a01b900460ff161561007a57600080fd5b6002805460ff60a01b1916600160a01b179055604080517f3c1316138cd3c347ee70454f6b80926b84604f3f07629dbed9845a8c06cc9ea360208201529081018890526001600160a01b038a1660608201526080810189905260009060a00160405160208183030381529060405280519060200120905061013f86868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506001548b92508591506001600160a01b0316610292565b61014857600080fd5b61019484848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250548c93508692506001600160a01b03169050610292565b61019d57600080fd5b60008881526005602052604090819020805460ff1916600117905560025490517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b038c81166004830152602482018c9052909116906340c10f1990604401600060405180830381600087803b15801561021e57600080fd5b505af1158015610232573d6000803e3d6000fd5b5050604080516001600160a01b038e168152602081018d90527f48b87fc02925b37a6aefac60c14fa9d8e9988d7dfadf262d4bd845872ca40730935001905060405180910390a150506002805460ff60a01b191690555050505050505050565b6004546040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101919091526042810183905260009081906062016040516020818303038152906040528051906020012090508085146102f957600080fd5b6103038587610323565b6001600160a01b0316836001600160a01b0316149150505b949350505050565b6000806000806103328561034f565b919450925090506103458684848461037e565b9695505050505050565b6000806000835160411461036257600080fd5b5050506020810151604082015160609092015160001a92909190565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156103b05750600061031b565b8360ff16601b141580156103c857508360ff16601c14155b156103d55750600061031b565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015610429573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661044e57600091505061031b565b95945050505050565b60008083601f84011261046957600080fd5b50813567ffffffffffffffff81111561048157600080fd5b60208301915083602082850101111561049957600080fd5b9250929050565b60008060008060008060008060c0898b0312156104bc57600080fd5b88356001600160a01b03811681146104d357600080fd5b9750602089013596506040890135955060608901359450608089013567ffffffffffffffff8082111561050557600080fd5b6105118c838d01610457565b909650945060a08b013591508082111561052a57600080fd5b506105378b828c01610457565b999c989b509699509497939692959450505056fea2646970667358221220cdf67210038ad517c2c46f363f99a1bbe4cef893cd97d87a0370e3770dee4ba464736f6c634300080a0033",
}

// BridgeMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMinterMetaData.ABI instead.
var BridgeMinterABI = BridgeMinterMetaData.ABI

// BridgeMinterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMinterMetaData.Bin instead.
var BridgeMinterBin = BridgeMinterMetaData.Bin

// DeployBridgeMinter deploys a new Ethereum contract, binding an instance of BridgeMinter to it.
func DeployBridgeMinter(auth *bind.TransactOpts, backend bind.ContractBackend, _approver common.Address, _notary common.Address, _tokenAddress common.Address, _chainId *big.Int) (common.Address, *types.Transaction, *BridgeMinter, error) {
	parsed, err := BridgeMinterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeMinterBin), backend, _approver, _notary, _tokenAddress, _chainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeMinter{BridgeMinterCaller: BridgeMinterCaller{contract: contract}, BridgeMinterTransactor: BridgeMinterTransactor{contract: contract}, BridgeMinterFilterer: BridgeMinterFilterer{contract: contract}}, nil
}

// BridgeMinter is an auto generated Go binding around an Ethereum contract.
type BridgeMinter struct {
	BridgeMinterCaller     // Read-only binding to the contract
	BridgeMinterTransactor // Write-only binding to the contract
	BridgeMinterFilterer   // Log filterer for contract events
}

// BridgeMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeMinterSession struct {
	Contract     *BridgeMinter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeMinterCallerSession struct {
	Contract *BridgeMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BridgeMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeMinterTransactorSession struct {
	Contract     *BridgeMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeMinterRaw struct {
	Contract *BridgeMinter // Generic contract binding to access the raw methods on
}

// BridgeMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeMinterCallerRaw struct {
	Contract *BridgeMinterCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeMinterTransactorRaw struct {
	Contract *BridgeMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeMinter creates a new instance of BridgeMinter, bound to a specific deployed contract.
func NewBridgeMinter(address common.Address, backend bind.ContractBackend) (*BridgeMinter, error) {
	contract, err := bindBridgeMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeMinter{BridgeMinterCaller: BridgeMinterCaller{contract: contract}, BridgeMinterTransactor: BridgeMinterTransactor{contract: contract}, BridgeMinterFilterer: BridgeMinterFilterer{contract: contract}}, nil
}

// NewBridgeMinterCaller creates a new read-only instance of BridgeMinter, bound to a specific deployed contract.
func NewBridgeMinterCaller(address common.Address, caller bind.ContractCaller) (*BridgeMinterCaller, error) {
	contract, err := bindBridgeMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeMinterCaller{contract: contract}, nil
}

// NewBridgeMinterTransactor creates a new write-only instance of BridgeMinter, bound to a specific deployed contract.
func NewBridgeMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeMinterTransactor, error) {
	contract, err := bindBridgeMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeMinterTransactor{contract: contract}, nil
}

// NewBridgeMinterFilterer creates a new log filterer instance of BridgeMinter, bound to a specific deployed contract.
func NewBridgeMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeMinterFilterer, error) {
	contract, err := bindBridgeMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeMinterFilterer{contract: contract}, nil
}

// bindBridgeMinter binds a generic wrapper to an already deployed contract.
func bindBridgeMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeMinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeMinter *BridgeMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeMinter.Contract.BridgeMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeMinter *BridgeMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeMinter.Contract.BridgeMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeMinter *BridgeMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeMinter.Contract.BridgeMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeMinter *BridgeMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeMinter *BridgeMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeMinter *BridgeMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeMinter.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a paid mutator transaction binding the contract method 0x9e07f0db.
//
// Solidity: function bridge(address sender, uint256 bridgedAmount, bytes32 nonce, bytes32 messageHash, bytes approvedMessage, bytes notarizedMessage) returns()
func (_BridgeMinter *BridgeMinterTransactor) Bridge(opts *bind.TransactOpts, sender common.Address, bridgedAmount *big.Int, nonce [32]byte, messageHash [32]byte, approvedMessage []byte, notarizedMessage []byte) (*types.Transaction, error) {
	return _BridgeMinter.contract.Transact(opts, "bridge", sender, bridgedAmount, nonce, messageHash, approvedMessage, notarizedMessage)
}

// Bridge is a paid mutator transaction binding the contract method 0x9e07f0db.
//
// Solidity: function bridge(address sender, uint256 bridgedAmount, bytes32 nonce, bytes32 messageHash, bytes approvedMessage, bytes notarizedMessage) returns()
func (_BridgeMinter *BridgeMinterSession) Bridge(sender common.Address, bridgedAmount *big.Int, nonce [32]byte, messageHash [32]byte, approvedMessage []byte, notarizedMessage []byte) (*types.Transaction, error) {
	return _BridgeMinter.Contract.Bridge(&_BridgeMinter.TransactOpts, sender, bridgedAmount, nonce, messageHash, approvedMessage, notarizedMessage)
}

// Bridge is a paid mutator transaction binding the contract method 0x9e07f0db.
//
// Solidity: function bridge(address sender, uint256 bridgedAmount, bytes32 nonce, bytes32 messageHash, bytes approvedMessage, bytes notarizedMessage) returns()
func (_BridgeMinter *BridgeMinterTransactorSession) Bridge(sender common.Address, bridgedAmount *big.Int, nonce [32]byte, messageHash [32]byte, approvedMessage []byte, notarizedMessage []byte) (*types.Transaction, error) {
	return _BridgeMinter.Contract.Bridge(&_BridgeMinter.TransactOpts, sender, bridgedAmount, nonce, messageHash, approvedMessage, notarizedMessage)
}

// BridgeMinterBridgedIterator is returned from FilterBridged and is used to iterate over the raw logs and unpacked data for Bridged events raised by the BridgeMinter contract.
type BridgeMinterBridgedIterator struct {
	Event *BridgeMinterBridged // Event containing the contract specifics and raw log

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
func (it *BridgeMinterBridgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMinterBridged)
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
		it.Event = new(BridgeMinterBridged)
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
func (it *BridgeMinterBridgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMinterBridgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMinterBridged represents a Bridged event raised by the BridgeMinter contract.
type BridgeMinterBridged struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridged is a free log retrieval operation binding the contract event 0x48b87fc02925b37a6aefac60c14fa9d8e9988d7dfadf262d4bd845872ca40730.
//
// Solidity: event Bridged(address receiver, uint256 amount)
func (_BridgeMinter *BridgeMinterFilterer) FilterBridged(opts *bind.FilterOpts) (*BridgeMinterBridgedIterator, error) {

	logs, sub, err := _BridgeMinter.contract.FilterLogs(opts, "Bridged")
	if err != nil {
		return nil, err
	}
	return &BridgeMinterBridgedIterator{contract: _BridgeMinter.contract, event: "Bridged", logs: logs, sub: sub}, nil
}

// WatchBridged is a free log subscription operation binding the contract event 0x48b87fc02925b37a6aefac60c14fa9d8e9988d7dfadf262d4bd845872ca40730.
//
// Solidity: event Bridged(address receiver, uint256 amount)
func (_BridgeMinter *BridgeMinterFilterer) WatchBridged(opts *bind.WatchOpts, sink chan<- *BridgeMinterBridged) (event.Subscription, error) {

	logs, sub, err := _BridgeMinter.contract.WatchLogs(opts, "Bridged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMinterBridged)
				if err := _BridgeMinter.contract.UnpackLog(event, "Bridged", log); err != nil {
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

// ParseBridged is a log parse operation binding the contract event 0x48b87fc02925b37a6aefac60c14fa9d8e9988d7dfadf262d4bd845872ca40730.
//
// Solidity: event Bridged(address receiver, uint256 amount)
func (_BridgeMinter *BridgeMinterFilterer) ParseBridged(log types.Log) (*BridgeMinterBridged, error) {
	event := new(BridgeMinterBridged)
	if err := _BridgeMinter.contract.UnpackLog(event, "Bridged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMinterTransferOwnershipIterator is returned from FilterTransferOwnership and is used to iterate over the raw logs and unpacked data for TransferOwnership events raised by the BridgeMinter contract.
type BridgeMinterTransferOwnershipIterator struct {
	Event *BridgeMinterTransferOwnership // Event containing the contract specifics and raw log

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
func (it *BridgeMinterTransferOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMinterTransferOwnership)
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
		it.Event = new(BridgeMinterTransferOwnership)
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
func (it *BridgeMinterTransferOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMinterTransferOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMinterTransferOwnership represents a TransferOwnership event raised by the BridgeMinter contract.
type BridgeMinterTransferOwnership struct {
	Owner     common.Address
	Confirmed bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferOwnership is a free log retrieval operation binding the contract event 0xcbe2aa772079086290469c944126fdc8e1438a556b92ebdb5044506435217f04.
//
// Solidity: event TransferOwnership(address indexed owner, bool indexed confirmed)
func (_BridgeMinter *BridgeMinterFilterer) FilterTransferOwnership(opts *bind.FilterOpts, owner []common.Address, confirmed []bool) (*BridgeMinterTransferOwnershipIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var confirmedRule []interface{}
	for _, confirmedItem := range confirmed {
		confirmedRule = append(confirmedRule, confirmedItem)
	}

	logs, sub, err := _BridgeMinter.contract.FilterLogs(opts, "TransferOwnership", ownerRule, confirmedRule)
	if err != nil {
		return nil, err
	}
	return &BridgeMinterTransferOwnershipIterator{contract: _BridgeMinter.contract, event: "TransferOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferOwnership is a free log subscription operation binding the contract event 0xcbe2aa772079086290469c944126fdc8e1438a556b92ebdb5044506435217f04.
//
// Solidity: event TransferOwnership(address indexed owner, bool indexed confirmed)
func (_BridgeMinter *BridgeMinterFilterer) WatchTransferOwnership(opts *bind.WatchOpts, sink chan<- *BridgeMinterTransferOwnership, owner []common.Address, confirmed []bool) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var confirmedRule []interface{}
	for _, confirmedItem := range confirmed {
		confirmedRule = append(confirmedRule, confirmedItem)
	}

	logs, sub, err := _BridgeMinter.contract.WatchLogs(opts, "TransferOwnership", ownerRule, confirmedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMinterTransferOwnership)
				if err := _BridgeMinter.contract.UnpackLog(event, "TransferOwnership", log); err != nil {
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

// ParseTransferOwnership is a log parse operation binding the contract event 0xcbe2aa772079086290469c944126fdc8e1438a556b92ebdb5044506435217f04.
//
// Solidity: event TransferOwnership(address indexed owner, bool indexed confirmed)
func (_BridgeMinter *BridgeMinterFilterer) ParseTransferOwnership(log types.Log) (*BridgeMinterTransferOwnership, error) {
	event := new(BridgeMinterTransferOwnership)
	if err := _BridgeMinter.contract.UnpackLog(event, "TransferOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
