// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gen

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

// ASMetaData contains all meta data concerning the AS contract.
var ASMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"A11GiveETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ercaddress_\",\"type\":\"address\"}],\"name\":\"A11Smile_setErc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"Startuse\",\"type\":\"address\"}],\"name\":\"A11SmileGiveETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Addtokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"account\",\"type\":\"uint256\"}],\"name\":\"AsgetETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"EthgetAs\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"AddEth\",\"type\":\"address\"}],\"name\":\"EthGetAs\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"A11Smile1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ASABI is the input ABI used to generate the binding from.
// Deprecated: Use ASMetaData.ABI instead.
var ASABI = ASMetaData.ABI

// AS is an auto generated Go binding around an Ethereum contract.
type AS struct {
	ASCaller     // Read-only binding to the contract
	ASTransactor // Write-only binding to the contract
	ASFilterer   // Log filterer for contract events
}

// ASCaller is an auto generated read-only Go binding around an Ethereum contract.
type ASCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ASTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ASTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ASFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ASFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ASSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ASSession struct {
	Contract     *AS               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ASCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ASCallerSession struct {
	Contract *ASCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ASTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ASTransactorSession struct {
	Contract     *ASTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ASRaw is an auto generated low-level Go binding around an Ethereum contract.
type ASRaw struct {
	Contract *AS // Generic contract binding to access the raw methods on
}

// ASCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ASCallerRaw struct {
	Contract *ASCaller // Generic read-only contract binding to access the raw methods on
}

// ASTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ASTransactorRaw struct {
	Contract *ASTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAS creates a new instance of AS, bound to a specific deployed contract.
func NewAS(address common.Address, backend bind.ContractBackend) (*AS, error) {
	contract, err := bindAS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AS{ASCaller: ASCaller{contract: contract}, ASTransactor: ASTransactor{contract: contract}, ASFilterer: ASFilterer{contract: contract}}, nil
}

// NewASCaller creates a new read-only instance of AS, bound to a specific deployed contract.
func NewASCaller(address common.Address, caller bind.ContractCaller) (*ASCaller, error) {
	contract, err := bindAS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ASCaller{contract: contract}, nil
}

// NewASTransactor creates a new write-only instance of AS, bound to a specific deployed contract.
func NewASTransactor(address common.Address, transactor bind.ContractTransactor) (*ASTransactor, error) {
	contract, err := bindAS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ASTransactor{contract: contract}, nil
}

// NewASFilterer creates a new log filterer instance of AS, bound to a specific deployed contract.
func NewASFilterer(address common.Address, filterer bind.ContractFilterer) (*ASFilterer, error) {
	contract, err := bindAS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ASFilterer{contract: contract}, nil
}

// bindAS binds a generic wrapper to an already deployed contract.
func bindAS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ASABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AS *ASRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AS.Contract.ASCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AS *ASRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AS.Contract.ASTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AS *ASRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AS.Contract.ASTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AS *ASCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AS *ASTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AS *ASTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AS.Contract.contract.Transact(opts, method, params...)
}

// A11Smile1 is a free data retrieval call binding the contract method 0x8f403c7d.
//
// Solidity: function A11Smile1() view returns(address)
func (_AS *ASCaller) A11Smile1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AS.contract.Call(opts, &out, "A11Smile1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// A11Smile1 is a free data retrieval call binding the contract method 0x8f403c7d.
//
// Solidity: function A11Smile1() view returns(address)
func (_AS *ASSession) A11Smile1() (common.Address, error) {
	return _AS.Contract.A11Smile1(&_AS.CallOpts)
}

// A11Smile1 is a free data retrieval call binding the contract method 0x8f403c7d.
//
// Solidity: function A11Smile1() view returns(address)
func (_AS *ASCallerSession) A11Smile1() (common.Address, error) {
	return _AS.Contract.A11Smile1(&_AS.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AS *ASCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AS.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AS *ASSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AS.Contract.Allowance(&_AS.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AS *ASCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AS.Contract.Allowance(&_AS.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AS *ASCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AS.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AS *ASSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AS.Contract.BalanceOf(&_AS.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AS *ASCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AS.Contract.BalanceOf(&_AS.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AS *ASCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AS.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AS *ASSession) Decimals() (uint8, error) {
	return _AS.Contract.Decimals(&_AS.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AS *ASCallerSession) Decimals() (uint8, error) {
	return _AS.Contract.Decimals(&_AS.CallOpts)
}

// GetUserAS is a free data retrieval call binding the contract method 0xda0b61f7.
//
// Solidity: function getUserAS() view returns(uint256)
func (_AS *ASCaller) GetUserAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AS.contract.Call(opts, &out, "getUserAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAS is a free data retrieval call binding the contract method 0xda0b61f7.
//
// Solidity: function getUserAS() view returns(uint256)
func (_AS *ASSession) GetUserAS() (*big.Int, error) {
	return _AS.Contract.GetUserAS(&_AS.CallOpts)
}

// GetUserAS is a free data retrieval call binding the contract method 0xda0b61f7.
//
// Solidity: function getUserAS() view returns(uint256)
func (_AS *ASCallerSession) GetUserAS() (*big.Int, error) {
	return _AS.Contract.GetUserAS(&_AS.CallOpts)
}

// GetUserETH is a free data retrieval call binding the contract method 0xbc7edaf2.
//
// Solidity: function getUserETH() view returns(uint256)
func (_AS *ASCaller) GetUserETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AS.contract.Call(opts, &out, "getUserETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserETH is a free data retrieval call binding the contract method 0xbc7edaf2.
//
// Solidity: function getUserETH() view returns(uint256)
func (_AS *ASSession) GetUserETH() (*big.Int, error) {
	return _AS.Contract.GetUserETH(&_AS.CallOpts)
}

// GetUserETH is a free data retrieval call binding the contract method 0xbc7edaf2.
//
// Solidity: function getUserETH() view returns(uint256)
func (_AS *ASCallerSession) GetUserETH() (*big.Int, error) {
	return _AS.Contract.GetUserETH(&_AS.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AS *ASCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AS.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AS *ASSession) Name() (string, error) {
	return _AS.Contract.Name(&_AS.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AS *ASCallerSession) Name() (string, error) {
	return _AS.Contract.Name(&_AS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AS *ASCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AS.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AS *ASSession) Symbol() (string, error) {
	return _AS.Contract.Symbol(&_AS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AS *ASCallerSession) Symbol() (string, error) {
	return _AS.Contract.Symbol(&_AS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AS *ASCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AS.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AS *ASSession) TotalSupply() (*big.Int, error) {
	return _AS.Contract.TotalSupply(&_AS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AS *ASCallerSession) TotalSupply() (*big.Int, error) {
	return _AS.Contract.TotalSupply(&_AS.CallOpts)
}

// A11SmileGiveETH is a paid mutator transaction binding the contract method 0xae622b60.
//
// Solidity: function A11SmileGiveETH(address Startuse) payable returns()
func (_AS *ASTransactor) A11SmileGiveETH(opts *bind.TransactOpts, Startuse common.Address) (*types.Transaction, error) {
	return _AS.contract.Transact(opts, "A11SmileGiveETH", Startuse)
}

// A11SmileGiveETH is a paid mutator transaction binding the contract method 0xae622b60.
//
// Solidity: function A11SmileGiveETH(address Startuse) payable returns()
func (_AS *ASSession) A11SmileGiveETH(Startuse common.Address) (*types.Transaction, error) {
	return _AS.Contract.A11SmileGiveETH(&_AS.TransactOpts, Startuse)
}

// A11SmileGiveETH is a paid mutator transaction binding the contract method 0xae622b60.
//
// Solidity: function A11SmileGiveETH(address Startuse) payable returns()
func (_AS *ASTransactorSession) A11SmileGiveETH(Startuse common.Address) (*types.Transaction, error) {
	return _AS.Contract.A11SmileGiveETH(&_AS.TransactOpts, Startuse)
}

// A11SmileSetErc is a paid mutator transaction binding the contract method 0xe7809062.
//
// Solidity: function A11Smile_setErc(address ercaddress_) returns()
func (_AS *ASTransactor) A11SmileSetErc(opts *bind.TransactOpts, ercaddress_ common.Address) (*types.Transaction, error) {
	return _AS.contract.Transact(opts, "A11Smile_setErc", ercaddress_)
}

// A11SmileSetErc is a paid mutator transaction binding the contract method 0xe7809062.
//
// Solidity: function A11Smile_setErc(address ercaddress_) returns()
func (_AS *ASSession) A11SmileSetErc(ercaddress_ common.Address) (*types.Transaction, error) {
	return _AS.Contract.A11SmileSetErc(&_AS.TransactOpts, ercaddress_)
}

// A11SmileSetErc is a paid mutator transaction binding the contract method 0xe7809062.
//
// Solidity: function A11Smile_setErc(address ercaddress_) returns()
func (_AS *ASTransactorSession) A11SmileSetErc(ercaddress_ common.Address) (*types.Transaction, error) {
	return _AS.Contract.A11SmileSetErc(&_AS.TransactOpts, ercaddress_)
}

// AsgetETH is a paid mutator transaction binding the contract method 0xeb35746f.
//
// Solidity: function AsgetETH(address to, uint256 account) returns()
func (_AS *ASTransactor) AsgetETH(opts *bind.TransactOpts, to common.Address, account *big.Int) (*types.Transaction, error) {
	return _AS.contract.Transact(opts, "AsgetETH", to, account)
}

// AsgetETH is a paid mutator transaction binding the contract method 0xeb35746f.
//
// Solidity: function AsgetETH(address to, uint256 account) returns()
func (_AS *ASSession) AsgetETH(to common.Address, account *big.Int) (*types.Transaction, error) {
	return _AS.Contract.AsgetETH(&_AS.TransactOpts, to, account)
}

// AsgetETH is a paid mutator transaction binding the contract method 0xeb35746f.
//
// Solidity: function AsgetETH(address to, uint256 account) returns()
func (_AS *ASTransactorSession) AsgetETH(to common.Address, account *big.Int) (*types.Transaction, error) {
	return _AS.Contract.AsgetETH(&_AS.TransactOpts, to, account)
}

// EthGetAs is a paid mutator transaction binding the contract method 0xc97c1b44.
//
// Solidity: function EthGetAs(address AddEth) payable returns()
func (_AS *ASTransactor) EthGetAs(opts *bind.TransactOpts, AddEth common.Address) (*types.Transaction, error) {
	return _AS.contract.Transact(opts, "EthGetAs", AddEth)
}

// EthGetAs is a paid mutator transaction binding the contract method 0xc97c1b44.
//
// Solidity: function EthGetAs(address AddEth) payable returns()
func (_AS *ASSession) EthGetAs(AddEth common.Address) (*types.Transaction, error) {
	return _AS.Contract.EthGetAs(&_AS.TransactOpts, AddEth)
}

// EthGetAs is a paid mutator transaction binding the contract method 0xc97c1b44.
//
// Solidity: function EthGetAs(address AddEth) payable returns()
func (_AS *ASTransactorSession) EthGetAs(AddEth common.Address) (*types.Transaction, error) {
	return _AS.Contract.EthGetAs(&_AS.TransactOpts, AddEth)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AS *ASTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AS.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AS *ASSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AS.Contract.Approve(&_AS.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AS *ASTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AS.Contract.Approve(&_AS.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AS *ASTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AS.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AS *ASSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AS.Contract.DecreaseAllowance(&_AS.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AS *ASTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AS.Contract.DecreaseAllowance(&_AS.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AS *ASTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AS.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AS *ASSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AS.Contract.IncreaseAllowance(&_AS.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AS *ASTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AS.Contract.IncreaseAllowance(&_AS.TransactOpts, spender, addedValue)
}

// Mint1 is a paid mutator transaction binding the contract method 0x25cccb94.
//
// Solidity: function mint1(uint256 amount) returns()
func (_AS *ASTransactor) Mint1(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _AS.contract.Transact(opts, "mint1", amount)
}

// Mint1 is a paid mutator transaction binding the contract method 0x25cccb94.
//
// Solidity: function mint1(uint256 amount) returns()
func (_AS *ASSession) Mint1(amount *big.Int) (*types.Transaction, error) {
	return _AS.Contract.Mint1(&_AS.TransactOpts, amount)
}

// Mint1 is a paid mutator transaction binding the contract method 0x25cccb94.
//
// Solidity: function mint1(uint256 amount) returns()
func (_AS *ASTransactorSession) Mint1(amount *big.Int) (*types.Transaction, error) {
	return _AS.Contract.Mint1(&_AS.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AS *ASTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AS.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AS *ASSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AS.Contract.Transfer(&_AS.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AS *ASTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AS.Contract.Transfer(&_AS.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AS *ASTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AS.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AS *ASSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AS.Contract.TransferFrom(&_AS.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AS *ASTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AS.Contract.TransferFrom(&_AS.TransactOpts, sender, recipient, amount)
}

// ASA11GiveETHIterator is returned from FilterA11GiveETH and is used to iterate over the raw logs and unpacked data for A11GiveETH events raised by the AS contract.
type ASA11GiveETHIterator struct {
	Event *ASA11GiveETH // Event containing the contract specifics and raw log

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
func (it *ASA11GiveETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ASA11GiveETH)
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
		it.Event = new(ASA11GiveETH)
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
func (it *ASA11GiveETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ASA11GiveETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ASA11GiveETH represents a A11GiveETH event raised by the AS contract.
type ASA11GiveETH struct {
	User1 common.Address
	User2 common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterA11GiveETH is a free log retrieval operation binding the contract event 0x41003739cb7adda22b5e7029b076372e289d38e49fcb784b70e3f2883dfb0265.
//
// Solidity: event A11GiveETH(address indexed user1, address indexed user2, uint256 indexed price)
func (_AS *ASFilterer) FilterA11GiveETH(opts *bind.FilterOpts, user1 []common.Address, user2 []common.Address, price []*big.Int) (*ASA11GiveETHIterator, error) {

	var user1Rule []interface{}
	for _, user1Item := range user1 {
		user1Rule = append(user1Rule, user1Item)
	}
	var user2Rule []interface{}
	for _, user2Item := range user2 {
		user2Rule = append(user2Rule, user2Item)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _AS.contract.FilterLogs(opts, "A11GiveETH", user1Rule, user2Rule, priceRule)
	if err != nil {
		return nil, err
	}
	return &ASA11GiveETHIterator{contract: _AS.contract, event: "A11GiveETH", logs: logs, sub: sub}, nil
}

// WatchA11GiveETH is a free log subscription operation binding the contract event 0x41003739cb7adda22b5e7029b076372e289d38e49fcb784b70e3f2883dfb0265.
//
// Solidity: event A11GiveETH(address indexed user1, address indexed user2, uint256 indexed price)
func (_AS *ASFilterer) WatchA11GiveETH(opts *bind.WatchOpts, sink chan<- *ASA11GiveETH, user1 []common.Address, user2 []common.Address, price []*big.Int) (event.Subscription, error) {

	var user1Rule []interface{}
	for _, user1Item := range user1 {
		user1Rule = append(user1Rule, user1Item)
	}
	var user2Rule []interface{}
	for _, user2Item := range user2 {
		user2Rule = append(user2Rule, user2Item)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _AS.contract.WatchLogs(opts, "A11GiveETH", user1Rule, user2Rule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ASA11GiveETH)
				if err := _AS.contract.UnpackLog(event, "A11GiveETH", log); err != nil {
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

// ParseA11GiveETH is a log parse operation binding the contract event 0x41003739cb7adda22b5e7029b076372e289d38e49fcb784b70e3f2883dfb0265.
//
// Solidity: event A11GiveETH(address indexed user1, address indexed user2, uint256 indexed price)
func (_AS *ASFilterer) ParseA11GiveETH(log types.Log) (*ASA11GiveETH, error) {
	event := new(ASA11GiveETH)
	if err := _AS.contract.UnpackLog(event, "A11GiveETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ASAddtokensIterator is returned from FilterAddtokens and is used to iterate over the raw logs and unpacked data for Addtokens events raised by the AS contract.
type ASAddtokensIterator struct {
	Event *ASAddtokens // Event containing the contract specifics and raw log

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
func (it *ASAddtokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ASAddtokens)
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
		it.Event = new(ASAddtokens)
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
func (it *ASAddtokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ASAddtokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ASAddtokens represents a Addtokens event raised by the AS contract.
type ASAddtokens struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddtokens is a free log retrieval operation binding the contract event 0x140232c0c4b8ad28248113c4185b3b3af0ba6f0589381f850b8731ad8af29db8.
//
// Solidity: event Addtokens(address indexed owner, uint256 indexed amount)
func (_AS *ASFilterer) FilterAddtokens(opts *bind.FilterOpts, owner []common.Address, amount []*big.Int) (*ASAddtokensIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AS.contract.FilterLogs(opts, "Addtokens", ownerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &ASAddtokensIterator{contract: _AS.contract, event: "Addtokens", logs: logs, sub: sub}, nil
}

// WatchAddtokens is a free log subscription operation binding the contract event 0x140232c0c4b8ad28248113c4185b3b3af0ba6f0589381f850b8731ad8af29db8.
//
// Solidity: event Addtokens(address indexed owner, uint256 indexed amount)
func (_AS *ASFilterer) WatchAddtokens(opts *bind.WatchOpts, sink chan<- *ASAddtokens, owner []common.Address, amount []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _AS.contract.WatchLogs(opts, "Addtokens", ownerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ASAddtokens)
				if err := _AS.contract.UnpackLog(event, "Addtokens", log); err != nil {
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

// ParseAddtokens is a log parse operation binding the contract event 0x140232c0c4b8ad28248113c4185b3b3af0ba6f0589381f850b8731ad8af29db8.
//
// Solidity: event Addtokens(address indexed owner, uint256 indexed amount)
func (_AS *ASFilterer) ParseAddtokens(log types.Log) (*ASAddtokens, error) {
	event := new(ASAddtokens)
	if err := _AS.contract.UnpackLog(event, "Addtokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ASApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AS contract.
type ASApprovalIterator struct {
	Event *ASApproval // Event containing the contract specifics and raw log

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
func (it *ASApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ASApproval)
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
		it.Event = new(ASApproval)
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
func (it *ASApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ASApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ASApproval represents a Approval event raised by the AS contract.
type ASApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AS *ASFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ASApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AS.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ASApprovalIterator{contract: _AS.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AS *ASFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ASApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AS.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ASApproval)
				if err := _AS.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AS *ASFilterer) ParseApproval(log types.Log) (*ASApproval, error) {
	event := new(ASApproval)
	if err := _AS.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ASEthgetAsIterator is returned from FilterEthgetAs and is used to iterate over the raw logs and unpacked data for EthgetAs events raised by the AS contract.
type ASEthgetAsIterator struct {
	Event *ASEthgetAs // Event containing the contract specifics and raw log

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
func (it *ASEthgetAsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ASEthgetAs)
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
		it.Event = new(ASEthgetAs)
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
func (it *ASEthgetAsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ASEthgetAsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ASEthgetAs represents a EthgetAs event raised by the AS contract.
type ASEthgetAs struct {
	User1 common.Address
	User2 common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEthgetAs is a free log retrieval operation binding the contract event 0x4e28aa9ff438c656a77043806d0d6a21a098eda7102c9b52698d2a6f321fc3e6.
//
// Solidity: event EthgetAs(address indexed user1, address indexed user2, uint256 indexed price)
func (_AS *ASFilterer) FilterEthgetAs(opts *bind.FilterOpts, user1 []common.Address, user2 []common.Address, price []*big.Int) (*ASEthgetAsIterator, error) {

	var user1Rule []interface{}
	for _, user1Item := range user1 {
		user1Rule = append(user1Rule, user1Item)
	}
	var user2Rule []interface{}
	for _, user2Item := range user2 {
		user2Rule = append(user2Rule, user2Item)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _AS.contract.FilterLogs(opts, "EthgetAs", user1Rule, user2Rule, priceRule)
	if err != nil {
		return nil, err
	}
	return &ASEthgetAsIterator{contract: _AS.contract, event: "EthgetAs", logs: logs, sub: sub}, nil
}

// WatchEthgetAs is a free log subscription operation binding the contract event 0x4e28aa9ff438c656a77043806d0d6a21a098eda7102c9b52698d2a6f321fc3e6.
//
// Solidity: event EthgetAs(address indexed user1, address indexed user2, uint256 indexed price)
func (_AS *ASFilterer) WatchEthgetAs(opts *bind.WatchOpts, sink chan<- *ASEthgetAs, user1 []common.Address, user2 []common.Address, price []*big.Int) (event.Subscription, error) {

	var user1Rule []interface{}
	for _, user1Item := range user1 {
		user1Rule = append(user1Rule, user1Item)
	}
	var user2Rule []interface{}
	for _, user2Item := range user2 {
		user2Rule = append(user2Rule, user2Item)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _AS.contract.WatchLogs(opts, "EthgetAs", user1Rule, user2Rule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ASEthgetAs)
				if err := _AS.contract.UnpackLog(event, "EthgetAs", log); err != nil {
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

// ParseEthgetAs is a log parse operation binding the contract event 0x4e28aa9ff438c656a77043806d0d6a21a098eda7102c9b52698d2a6f321fc3e6.
//
// Solidity: event EthgetAs(address indexed user1, address indexed user2, uint256 indexed price)
func (_AS *ASFilterer) ParseEthgetAs(log types.Log) (*ASEthgetAs, error) {
	event := new(ASEthgetAs)
	if err := _AS.contract.UnpackLog(event, "EthgetAs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ASTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AS contract.
type ASTransferIterator struct {
	Event *ASTransfer // Event containing the contract specifics and raw log

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
func (it *ASTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ASTransfer)
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
		it.Event = new(ASTransfer)
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
func (it *ASTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ASTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ASTransfer represents a Transfer event raised by the AS contract.
type ASTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AS *ASFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ASTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AS.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ASTransferIterator{contract: _AS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AS *ASFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ASTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AS.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ASTransfer)
				if err := _AS.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AS *ASFilterer) ParseTransfer(log types.Log) (*ASTransfer, error) {
	event := new(ASTransfer)
	if err := _AS.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
