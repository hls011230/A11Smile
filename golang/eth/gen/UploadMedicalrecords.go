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

// UploadMedicalrecordsMetaData contains all meta data concerning the UploadMedicalrecords contract.
var UploadMedicalrecordsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"A11GiveETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Addtokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"EthgetAs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"route\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"}],\"name\":\"Uploadmedicaldata\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"route\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"}],\"name\":\"gainerUploadmedicaldata\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"route\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"reward\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ercaddress_\",\"type\":\"address\"}],\"name\":\"A11Smile_setErc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"Startuse\",\"type\":\"address\"}],\"name\":\"A11SmileGiveETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_Medicalrecordrequirements\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_Requirementdescription\",\"type\":\"string\"}],\"name\":\"gainer_AddMedicalInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"}],\"name\":\"gainer_setDoctor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity1\",\"type\":\"uint256\"}],\"name\":\"gainer_transfer_AS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_soliciter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"PictureRoute\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"whether\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"ercnum_\",\"type\":\"uint256\"}],\"name\":\"gainer_whether\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Proute\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_soliciter\",\"type\":\"address\"}],\"name\":\"user_AddMedicalInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"user_Adduser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"user_UPMedicalExaminationReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"user_UPMedicalinformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"A11Smile\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"examineTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"}],\"name\":\"gainer_doctorsee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"}],\"name\":\"gainer_isDoctor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"}],\"name\":\"gainer_SeeOwnMedicalInformation\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"PictureRoute\",\"type\":\"string\"}],\"name\":\"seeMedicaldata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"user_IsUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"user_MedicalExaminationReportstrucrName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"user_MedicalinformationrName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"user_ViewMedicalExaminationReport\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"user_ViewMedicalInformation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"user_ViewMedicalRecords\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewMedicalExaminationReportCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewMedicalinformationCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UploadMedicalrecordsABI is the input ABI used to generate the binding from.
// Deprecated: Use UploadMedicalrecordsMetaData.ABI instead.
var UploadMedicalrecordsABI = UploadMedicalrecordsMetaData.ABI

// UploadMedicalrecords is an auto generated Go binding around an Ethereum contract.
type UploadMedicalrecords struct {
	UploadMedicalrecordsCaller     // Read-only binding to the contract
	UploadMedicalrecordsTransactor // Write-only binding to the contract
	UploadMedicalrecordsFilterer   // Log filterer for contract events
}

// UploadMedicalrecordsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UploadMedicalrecordsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UploadMedicalrecordsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UploadMedicalrecordsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UploadMedicalrecordsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UploadMedicalrecordsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UploadMedicalrecordsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UploadMedicalrecordsSession struct {
	Contract     *UploadMedicalrecords // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UploadMedicalrecordsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UploadMedicalrecordsCallerSession struct {
	Contract *UploadMedicalrecordsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// UploadMedicalrecordsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UploadMedicalrecordsTransactorSession struct {
	Contract     *UploadMedicalrecordsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// UploadMedicalrecordsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UploadMedicalrecordsRaw struct {
	Contract *UploadMedicalrecords // Generic contract binding to access the raw methods on
}

// UploadMedicalrecordsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UploadMedicalrecordsCallerRaw struct {
	Contract *UploadMedicalrecordsCaller // Generic read-only contract binding to access the raw methods on
}

// UploadMedicalrecordsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UploadMedicalrecordsTransactorRaw struct {
	Contract *UploadMedicalrecordsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUploadMedicalrecords creates a new instance of UploadMedicalrecords, bound to a specific deployed contract.
func NewUploadMedicalrecords(address common.Address, backend bind.ContractBackend) (*UploadMedicalrecords, error) {
	contract, err := bindUploadMedicalrecords(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UploadMedicalrecords{UploadMedicalrecordsCaller: UploadMedicalrecordsCaller{contract: contract}, UploadMedicalrecordsTransactor: UploadMedicalrecordsTransactor{contract: contract}, UploadMedicalrecordsFilterer: UploadMedicalrecordsFilterer{contract: contract}}, nil
}

// NewUploadMedicalrecordsCaller creates a new read-only instance of UploadMedicalrecords, bound to a specific deployed contract.
func NewUploadMedicalrecordsCaller(address common.Address, caller bind.ContractCaller) (*UploadMedicalrecordsCaller, error) {
	contract, err := bindUploadMedicalrecords(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UploadMedicalrecordsCaller{contract: contract}, nil
}

// NewUploadMedicalrecordsTransactor creates a new write-only instance of UploadMedicalrecords, bound to a specific deployed contract.
func NewUploadMedicalrecordsTransactor(address common.Address, transactor bind.ContractTransactor) (*UploadMedicalrecordsTransactor, error) {
	contract, err := bindUploadMedicalrecords(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UploadMedicalrecordsTransactor{contract: contract}, nil
}

// NewUploadMedicalrecordsFilterer creates a new log filterer instance of UploadMedicalrecords, bound to a specific deployed contract.
func NewUploadMedicalrecordsFilterer(address common.Address, filterer bind.ContractFilterer) (*UploadMedicalrecordsFilterer, error) {
	contract, err := bindUploadMedicalrecords(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UploadMedicalrecordsFilterer{contract: contract}, nil
}

// bindUploadMedicalrecords binds a generic wrapper to an already deployed contract.
func bindUploadMedicalrecords(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UploadMedicalrecordsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UploadMedicalrecords *UploadMedicalrecordsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UploadMedicalrecords.Contract.UploadMedicalrecordsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UploadMedicalrecords *UploadMedicalrecordsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.UploadMedicalrecordsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UploadMedicalrecords *UploadMedicalrecordsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.UploadMedicalrecordsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UploadMedicalrecords *UploadMedicalrecordsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UploadMedicalrecords.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.contract.Transact(opts, method, params...)
}

// A11Smile is a free data retrieval call binding the contract method 0xee580fd8.
//
// Solidity: function A11Smile() view returns(address)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) A11Smile(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "A11Smile")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// A11Smile is a free data retrieval call binding the contract method 0xee580fd8.
//
// Solidity: function A11Smile() view returns(address)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) A11Smile() (common.Address, error) {
	return _UploadMedicalrecords.Contract.A11Smile(&_UploadMedicalrecords.CallOpts)
}

// A11Smile is a free data retrieval call binding the contract method 0xee580fd8.
//
// Solidity: function A11Smile() view returns(address)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) A11Smile() (common.Address, error) {
	return _UploadMedicalrecords.Contract.A11Smile(&_UploadMedicalrecords.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UploadMedicalrecords.Contract.Allowance(&_UploadMedicalrecords.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UploadMedicalrecords.Contract.Allowance(&_UploadMedicalrecords.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UploadMedicalrecords.Contract.BalanceOf(&_UploadMedicalrecords.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UploadMedicalrecords.Contract.BalanceOf(&_UploadMedicalrecords.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) Decimals() (uint8, error) {
	return _UploadMedicalrecords.Contract.Decimals(&_UploadMedicalrecords.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) Decimals() (uint8, error) {
	return _UploadMedicalrecords.Contract.Decimals(&_UploadMedicalrecords.CallOpts)
}

// ExamineTime is a free data retrieval call binding the contract method 0x495c184e.
//
// Solidity: function examineTime() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) ExamineTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "examineTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExamineTime is a free data retrieval call binding the contract method 0x495c184e.
//
// Solidity: function examineTime() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) ExamineTime() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.ExamineTime(&_UploadMedicalrecords.CallOpts)
}

// ExamineTime is a free data retrieval call binding the contract method 0x495c184e.
//
// Solidity: function examineTime() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) ExamineTime() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.ExamineTime(&_UploadMedicalrecords.CallOpts)
}

// GainerSeeOwnMedicalInformation is a free data retrieval call binding the contract method 0xc4de85e0.
//
// Solidity: function gainer_SeeOwnMedicalInformation(string MedicalName) view returns(string[], uint256, uint256, string, string, address)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) GainerSeeOwnMedicalInformation(opts *bind.CallOpts, MedicalName string) ([]string, *big.Int, *big.Int, string, string, common.Address, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "gainer_SeeOwnMedicalInformation", MedicalName)

	if err != nil {
		return *new([]string), *new(*big.Int), *new(*big.Int), *new(string), *new(string), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, out4, out5, err

}

// GainerSeeOwnMedicalInformation is a free data retrieval call binding the contract method 0xc4de85e0.
//
// Solidity: function gainer_SeeOwnMedicalInformation(string MedicalName) view returns(string[], uint256, uint256, string, string, address)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) GainerSeeOwnMedicalInformation(MedicalName string) ([]string, *big.Int, *big.Int, string, string, common.Address, error) {
	return _UploadMedicalrecords.Contract.GainerSeeOwnMedicalInformation(&_UploadMedicalrecords.CallOpts, MedicalName)
}

// GainerSeeOwnMedicalInformation is a free data retrieval call binding the contract method 0xc4de85e0.
//
// Solidity: function gainer_SeeOwnMedicalInformation(string MedicalName) view returns(string[], uint256, uint256, string, string, address)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) GainerSeeOwnMedicalInformation(MedicalName string) ([]string, *big.Int, *big.Int, string, string, common.Address, error) {
	return _UploadMedicalrecords.Contract.GainerSeeOwnMedicalInformation(&_UploadMedicalrecords.CallOpts, MedicalName)
}

// GainerDoctorsee is a free data retrieval call binding the contract method 0xfece0fe8.
//
// Solidity: function gainer_doctorsee(address soliciter) view returns(address, string[], address)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) GainerDoctorsee(opts *bind.CallOpts, soliciter common.Address) (common.Address, []string, common.Address, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "gainer_doctorsee", soliciter)

	if err != nil {
		return *new(common.Address), *new([]string), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]string)).(*[]string)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// GainerDoctorsee is a free data retrieval call binding the contract method 0xfece0fe8.
//
// Solidity: function gainer_doctorsee(address soliciter) view returns(address, string[], address)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) GainerDoctorsee(soliciter common.Address) (common.Address, []string, common.Address, error) {
	return _UploadMedicalrecords.Contract.GainerDoctorsee(&_UploadMedicalrecords.CallOpts, soliciter)
}

// GainerDoctorsee is a free data retrieval call binding the contract method 0xfece0fe8.
//
// Solidity: function gainer_doctorsee(address soliciter) view returns(address, string[], address)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) GainerDoctorsee(soliciter common.Address) (common.Address, []string, common.Address, error) {
	return _UploadMedicalrecords.Contract.GainerDoctorsee(&_UploadMedicalrecords.CallOpts, soliciter)
}

// GainerIsDoctor is a free data retrieval call binding the contract method 0x93172bf7.
//
// Solidity: function gainer_isDoctor(address soliciter) view returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) GainerIsDoctor(opts *bind.CallOpts, soliciter common.Address) (bool, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "gainer_isDoctor", soliciter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GainerIsDoctor is a free data retrieval call binding the contract method 0x93172bf7.
//
// Solidity: function gainer_isDoctor(address soliciter) view returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) GainerIsDoctor(soliciter common.Address) (bool, error) {
	return _UploadMedicalrecords.Contract.GainerIsDoctor(&_UploadMedicalrecords.CallOpts, soliciter)
}

// GainerIsDoctor is a free data retrieval call binding the contract method 0x93172bf7.
//
// Solidity: function gainer_isDoctor(address soliciter) view returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) GainerIsDoctor(soliciter common.Address) (bool, error) {
	return _UploadMedicalrecords.Contract.GainerIsDoctor(&_UploadMedicalrecords.CallOpts, soliciter)
}

// GetUserAS is a free data retrieval call binding the contract method 0xda0b61f7.
//
// Solidity: function getUserAS() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) GetUserAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "getUserAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAS is a free data retrieval call binding the contract method 0xda0b61f7.
//
// Solidity: function getUserAS() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) GetUserAS() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.GetUserAS(&_UploadMedicalrecords.CallOpts)
}

// GetUserAS is a free data retrieval call binding the contract method 0xda0b61f7.
//
// Solidity: function getUserAS() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) GetUserAS() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.GetUserAS(&_UploadMedicalrecords.CallOpts)
}

// GetUserETH is a free data retrieval call binding the contract method 0xbc7edaf2.
//
// Solidity: function getUserETH() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) GetUserETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "getUserETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserETH is a free data retrieval call binding the contract method 0xbc7edaf2.
//
// Solidity: function getUserETH() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) GetUserETH() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.GetUserETH(&_UploadMedicalrecords.CallOpts)
}

// GetUserETH is a free data retrieval call binding the contract method 0xbc7edaf2.
//
// Solidity: function getUserETH() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) GetUserETH() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.GetUserETH(&_UploadMedicalrecords.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) LastUpdateTime() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.LastUpdateTime(&_UploadMedicalrecords.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) LastUpdateTime() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.LastUpdateTime(&_UploadMedicalrecords.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) Name() (string, error) {
	return _UploadMedicalrecords.Contract.Name(&_UploadMedicalrecords.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) Name() (string, error) {
	return _UploadMedicalrecords.Contract.Name(&_UploadMedicalrecords.CallOpts)
}

// SeeMedicaldata is a free data retrieval call binding the contract method 0x600d9056.
//
// Solidity: function seeMedicaldata(string PictureRoute) view returns(string, bool)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) SeeMedicaldata(opts *bind.CallOpts, PictureRoute string) (string, bool, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "seeMedicaldata", PictureRoute)

	if err != nil {
		return *new(string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// SeeMedicaldata is a free data retrieval call binding the contract method 0x600d9056.
//
// Solidity: function seeMedicaldata(string PictureRoute) view returns(string, bool)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) SeeMedicaldata(PictureRoute string) (string, bool, error) {
	return _UploadMedicalrecords.Contract.SeeMedicaldata(&_UploadMedicalrecords.CallOpts, PictureRoute)
}

// SeeMedicaldata is a free data retrieval call binding the contract method 0x600d9056.
//
// Solidity: function seeMedicaldata(string PictureRoute) view returns(string, bool)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) SeeMedicaldata(PictureRoute string) (string, bool, error) {
	return _UploadMedicalrecords.Contract.SeeMedicaldata(&_UploadMedicalrecords.CallOpts, PictureRoute)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) Symbol() (string, error) {
	return _UploadMedicalrecords.Contract.Symbol(&_UploadMedicalrecords.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) Symbol() (string, error) {
	return _UploadMedicalrecords.Contract.Symbol(&_UploadMedicalrecords.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) TotalSupply() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.TotalSupply(&_UploadMedicalrecords.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) TotalSupply() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.TotalSupply(&_UploadMedicalrecords.CallOpts)
}

// UserIsUser is a free data retrieval call binding the contract method 0xd5f03d3c.
//
// Solidity: function user_IsUser(address _user) view returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) UserIsUser(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "user_IsUser", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserIsUser is a free data retrieval call binding the contract method 0xd5f03d3c.
//
// Solidity: function user_IsUser(address _user) view returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) UserIsUser(_user common.Address) (bool, error) {
	return _UploadMedicalrecords.Contract.UserIsUser(&_UploadMedicalrecords.CallOpts, _user)
}

// UserIsUser is a free data retrieval call binding the contract method 0xd5f03d3c.
//
// Solidity: function user_IsUser(address _user) view returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) UserIsUser(_user common.Address) (bool, error) {
	return _UploadMedicalrecords.Contract.UserIsUser(&_UploadMedicalrecords.CallOpts, _user)
}

// UserMedicalExaminationReportstrucrName is a free data retrieval call binding the contract method 0xeae38791.
//
// Solidity: function user_MedicalExaminationReportstrucrName(address , uint256 ) view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) UserMedicalExaminationReportstrucrName(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "user_MedicalExaminationReportstrucrName", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UserMedicalExaminationReportstrucrName is a free data retrieval call binding the contract method 0xeae38791.
//
// Solidity: function user_MedicalExaminationReportstrucrName(address , uint256 ) view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) UserMedicalExaminationReportstrucrName(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _UploadMedicalrecords.Contract.UserMedicalExaminationReportstrucrName(&_UploadMedicalrecords.CallOpts, arg0, arg1)
}

// UserMedicalExaminationReportstrucrName is a free data retrieval call binding the contract method 0xeae38791.
//
// Solidity: function user_MedicalExaminationReportstrucrName(address , uint256 ) view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) UserMedicalExaminationReportstrucrName(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _UploadMedicalrecords.Contract.UserMedicalExaminationReportstrucrName(&_UploadMedicalrecords.CallOpts, arg0, arg1)
}

// UserMedicalinformationrName is a free data retrieval call binding the contract method 0x0d121e4e.
//
// Solidity: function user_MedicalinformationrName(address , uint256 ) view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) UserMedicalinformationrName(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "user_MedicalinformationrName", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UserMedicalinformationrName is a free data retrieval call binding the contract method 0x0d121e4e.
//
// Solidity: function user_MedicalinformationrName(address , uint256 ) view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) UserMedicalinformationrName(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _UploadMedicalrecords.Contract.UserMedicalinformationrName(&_UploadMedicalrecords.CallOpts, arg0, arg1)
}

// UserMedicalinformationrName is a free data retrieval call binding the contract method 0x0d121e4e.
//
// Solidity: function user_MedicalinformationrName(address , uint256 ) view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) UserMedicalinformationrName(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _UploadMedicalrecords.Contract.UserMedicalinformationrName(&_UploadMedicalrecords.CallOpts, arg0, arg1)
}

// UserViewMedicalExaminationReport is a free data retrieval call binding the contract method 0x603a78c2.
//
// Solidity: function user_ViewMedicalExaminationReport(string name) view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) UserViewMedicalExaminationReport(opts *bind.CallOpts, name string) (string, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "user_ViewMedicalExaminationReport", name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UserViewMedicalExaminationReport is a free data retrieval call binding the contract method 0x603a78c2.
//
// Solidity: function user_ViewMedicalExaminationReport(string name) view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) UserViewMedicalExaminationReport(name string) (string, error) {
	return _UploadMedicalrecords.Contract.UserViewMedicalExaminationReport(&_UploadMedicalrecords.CallOpts, name)
}

// UserViewMedicalExaminationReport is a free data retrieval call binding the contract method 0x603a78c2.
//
// Solidity: function user_ViewMedicalExaminationReport(string name) view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) UserViewMedicalExaminationReport(name string) (string, error) {
	return _UploadMedicalrecords.Contract.UserViewMedicalExaminationReport(&_UploadMedicalrecords.CallOpts, name)
}

// UserViewMedicalInformation is a free data retrieval call binding the contract method 0xd59409fe.
//
// Solidity: function user_ViewMedicalInformation(string name) view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) UserViewMedicalInformation(opts *bind.CallOpts, name string) (string, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "user_ViewMedicalInformation", name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UserViewMedicalInformation is a free data retrieval call binding the contract method 0xd59409fe.
//
// Solidity: function user_ViewMedicalInformation(string name) view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) UserViewMedicalInformation(name string) (string, error) {
	return _UploadMedicalrecords.Contract.UserViewMedicalInformation(&_UploadMedicalrecords.CallOpts, name)
}

// UserViewMedicalInformation is a free data retrieval call binding the contract method 0xd59409fe.
//
// Solidity: function user_ViewMedicalInformation(string name) view returns(string)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) UserViewMedicalInformation(name string) (string, error) {
	return _UploadMedicalrecords.Contract.UserViewMedicalInformation(&_UploadMedicalrecords.CallOpts, name)
}

// UserViewMedicalRecords is a free data retrieval call binding the contract method 0x3fe8ac8c.
//
// Solidity: function user_ViewMedicalRecords() view returns(string[], address soliciter)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) UserViewMedicalRecords(opts *bind.CallOpts) ([]string, common.Address, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "user_ViewMedicalRecords")

	if err != nil {
		return *new([]string), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// UserViewMedicalRecords is a free data retrieval call binding the contract method 0x3fe8ac8c.
//
// Solidity: function user_ViewMedicalRecords() view returns(string[], address soliciter)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) UserViewMedicalRecords() ([]string, common.Address, error) {
	return _UploadMedicalrecords.Contract.UserViewMedicalRecords(&_UploadMedicalrecords.CallOpts)
}

// UserViewMedicalRecords is a free data retrieval call binding the contract method 0x3fe8ac8c.
//
// Solidity: function user_ViewMedicalRecords() view returns(string[], address soliciter)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) UserViewMedicalRecords() ([]string, common.Address, error) {
	return _UploadMedicalrecords.Contract.UserViewMedicalRecords(&_UploadMedicalrecords.CallOpts)
}

// ViewMedicalExaminationReportCount is a free data retrieval call binding the contract method 0x03c8505d.
//
// Solidity: function viewMedicalExaminationReportCount() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) ViewMedicalExaminationReportCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "viewMedicalExaminationReportCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewMedicalExaminationReportCount is a free data retrieval call binding the contract method 0x03c8505d.
//
// Solidity: function viewMedicalExaminationReportCount() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) ViewMedicalExaminationReportCount() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.ViewMedicalExaminationReportCount(&_UploadMedicalrecords.CallOpts)
}

// ViewMedicalExaminationReportCount is a free data retrieval call binding the contract method 0x03c8505d.
//
// Solidity: function viewMedicalExaminationReportCount() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) ViewMedicalExaminationReportCount() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.ViewMedicalExaminationReportCount(&_UploadMedicalrecords.CallOpts)
}

// ViewMedicalinformationCount is a free data retrieval call binding the contract method 0x9d012e1c.
//
// Solidity: function viewMedicalinformationCount() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCaller) ViewMedicalinformationCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UploadMedicalrecords.contract.Call(opts, &out, "viewMedicalinformationCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewMedicalinformationCount is a free data retrieval call binding the contract method 0x9d012e1c.
//
// Solidity: function viewMedicalinformationCount() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) ViewMedicalinformationCount() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.ViewMedicalinformationCount(&_UploadMedicalrecords.CallOpts)
}

// ViewMedicalinformationCount is a free data retrieval call binding the contract method 0x9d012e1c.
//
// Solidity: function viewMedicalinformationCount() view returns(uint256)
func (_UploadMedicalrecords *UploadMedicalrecordsCallerSession) ViewMedicalinformationCount() (*big.Int, error) {
	return _UploadMedicalrecords.Contract.ViewMedicalinformationCount(&_UploadMedicalrecords.CallOpts)
}

// A11SmileGiveETH is a paid mutator transaction binding the contract method 0xae622b60.
//
// Solidity: function A11SmileGiveETH(address Startuse) payable returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) A11SmileGiveETH(opts *bind.TransactOpts, Startuse common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "A11SmileGiveETH", Startuse)
}

// A11SmileGiveETH is a paid mutator transaction binding the contract method 0xae622b60.
//
// Solidity: function A11SmileGiveETH(address Startuse) payable returns()
func (_UploadMedicalrecords *UploadMedicalrecordsSession) A11SmileGiveETH(Startuse common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.A11SmileGiveETH(&_UploadMedicalrecords.TransactOpts, Startuse)
}

// A11SmileGiveETH is a paid mutator transaction binding the contract method 0xae622b60.
//
// Solidity: function A11SmileGiveETH(address Startuse) payable returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) A11SmileGiveETH(Startuse common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.A11SmileGiveETH(&_UploadMedicalrecords.TransactOpts, Startuse)
}

// A11SmileSetErc is a paid mutator transaction binding the contract method 0xe7809062.
//
// Solidity: function A11Smile_setErc(address ercaddress_) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) A11SmileSetErc(opts *bind.TransactOpts, ercaddress_ common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "A11Smile_setErc", ercaddress_)
}

// A11SmileSetErc is a paid mutator transaction binding the contract method 0xe7809062.
//
// Solidity: function A11Smile_setErc(address ercaddress_) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsSession) A11SmileSetErc(ercaddress_ common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.A11SmileSetErc(&_UploadMedicalrecords.TransactOpts, ercaddress_)
}

// A11SmileSetErc is a paid mutator transaction binding the contract method 0xe7809062.
//
// Solidity: function A11Smile_setErc(address ercaddress_) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) A11SmileSetErc(ercaddress_ common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.A11SmileSetErc(&_UploadMedicalrecords.TransactOpts, ercaddress_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.Approve(&_UploadMedicalrecords.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.Approve(&_UploadMedicalrecords.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.DecreaseAllowance(&_UploadMedicalrecords.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.DecreaseAllowance(&_UploadMedicalrecords.TransactOpts, spender, subtractedValue)
}

// GainerAddMedicalInformation is a paid mutator transaction binding the contract method 0x0a3336f0.
//
// Solidity: function gainer_AddMedicalInformation(uint256 min, uint256 max, string MedicalName, string _Medicalrecordrequirements, string _Requirementdescription) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) GainerAddMedicalInformation(opts *bind.TransactOpts, min *big.Int, max *big.Int, MedicalName string, _Medicalrecordrequirements string, _Requirementdescription string) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "gainer_AddMedicalInformation", min, max, MedicalName, _Medicalrecordrequirements, _Requirementdescription)
}

// GainerAddMedicalInformation is a paid mutator transaction binding the contract method 0x0a3336f0.
//
// Solidity: function gainer_AddMedicalInformation(uint256 min, uint256 max, string MedicalName, string _Medicalrecordrequirements, string _Requirementdescription) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsSession) GainerAddMedicalInformation(min *big.Int, max *big.Int, MedicalName string, _Medicalrecordrequirements string, _Requirementdescription string) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.GainerAddMedicalInformation(&_UploadMedicalrecords.TransactOpts, min, max, MedicalName, _Medicalrecordrequirements, _Requirementdescription)
}

// GainerAddMedicalInformation is a paid mutator transaction binding the contract method 0x0a3336f0.
//
// Solidity: function gainer_AddMedicalInformation(uint256 min, uint256 max, string MedicalName, string _Medicalrecordrequirements, string _Requirementdescription) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) GainerAddMedicalInformation(min *big.Int, max *big.Int, MedicalName string, _Medicalrecordrequirements string, _Requirementdescription string) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.GainerAddMedicalInformation(&_UploadMedicalrecords.TransactOpts, min, max, MedicalName, _Medicalrecordrequirements, _Requirementdescription)
}

// GainerSetDoctor is a paid mutator transaction binding the contract method 0x5d268670.
//
// Solidity: function gainer_setDoctor(address soliciter) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) GainerSetDoctor(opts *bind.TransactOpts, soliciter common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "gainer_setDoctor", soliciter)
}

// GainerSetDoctor is a paid mutator transaction binding the contract method 0x5d268670.
//
// Solidity: function gainer_setDoctor(address soliciter) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsSession) GainerSetDoctor(soliciter common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.GainerSetDoctor(&_UploadMedicalrecords.TransactOpts, soliciter)
}

// GainerSetDoctor is a paid mutator transaction binding the contract method 0x5d268670.
//
// Solidity: function gainer_setDoctor(address soliciter) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) GainerSetDoctor(soliciter common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.GainerSetDoctor(&_UploadMedicalrecords.TransactOpts, soliciter)
}

// GainerTransferAS is a paid mutator transaction binding the contract method 0x3d199172.
//
// Solidity: function gainer_transfer_AS(address from1, address to1, uint256 quantity1) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) GainerTransferAS(opts *bind.TransactOpts, from1 common.Address, to1 common.Address, quantity1 *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "gainer_transfer_AS", from1, to1, quantity1)
}

// GainerTransferAS is a paid mutator transaction binding the contract method 0x3d199172.
//
// Solidity: function gainer_transfer_AS(address from1, address to1, uint256 quantity1) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsSession) GainerTransferAS(from1 common.Address, to1 common.Address, quantity1 *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.GainerTransferAS(&_UploadMedicalrecords.TransactOpts, from1, to1, quantity1)
}

// GainerTransferAS is a paid mutator transaction binding the contract method 0x3d199172.
//
// Solidity: function gainer_transfer_AS(address from1, address to1, uint256 quantity1) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) GainerTransferAS(from1 common.Address, to1 common.Address, quantity1 *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.GainerTransferAS(&_UploadMedicalrecords.TransactOpts, from1, to1, quantity1)
}

// GainerWhether is a paid mutator transaction binding the contract method 0xac5679d0.
//
// Solidity: function gainer_whether(address _soliciter, string PictureRoute, bool whether, uint256 ercnum_) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) GainerWhether(opts *bind.TransactOpts, _soliciter common.Address, PictureRoute string, whether bool, ercnum_ *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "gainer_whether", _soliciter, PictureRoute, whether, ercnum_)
}

// GainerWhether is a paid mutator transaction binding the contract method 0xac5679d0.
//
// Solidity: function gainer_whether(address _soliciter, string PictureRoute, bool whether, uint256 ercnum_) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsSession) GainerWhether(_soliciter common.Address, PictureRoute string, whether bool, ercnum_ *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.GainerWhether(&_UploadMedicalrecords.TransactOpts, _soliciter, PictureRoute, whether, ercnum_)
}

// GainerWhether is a paid mutator transaction binding the contract method 0xac5679d0.
//
// Solidity: function gainer_whether(address _soliciter, string PictureRoute, bool whether, uint256 ercnum_) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) GainerWhether(_soliciter common.Address, PictureRoute string, whether bool, ercnum_ *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.GainerWhether(&_UploadMedicalrecords.TransactOpts, _soliciter, PictureRoute, whether, ercnum_)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.IncreaseAllowance(&_UploadMedicalrecords.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.IncreaseAllowance(&_UploadMedicalrecords.TransactOpts, spender, addedValue)
}

// Mint1 is a paid mutator transaction binding the contract method 0x25cccb94.
//
// Solidity: function mint1(uint256 amount) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) Mint1(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "mint1", amount)
}

// Mint1 is a paid mutator transaction binding the contract method 0x25cccb94.
//
// Solidity: function mint1(uint256 amount) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsSession) Mint1(amount *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.Mint1(&_UploadMedicalrecords.TransactOpts, amount)
}

// Mint1 is a paid mutator transaction binding the contract method 0x25cccb94.
//
// Solidity: function mint1(uint256 amount) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) Mint1(amount *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.Mint1(&_UploadMedicalrecords.TransactOpts, amount)
}

// ToContract is a paid mutator transaction binding the contract method 0x61b21f4d.
//
// Solidity: function toContract() payable returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) ToContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "toContract")
}

// ToContract is a paid mutator transaction binding the contract method 0x61b21f4d.
//
// Solidity: function toContract() payable returns()
func (_UploadMedicalrecords *UploadMedicalrecordsSession) ToContract() (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.ToContract(&_UploadMedicalrecords.TransactOpts)
}

// ToContract is a paid mutator transaction binding the contract method 0x61b21f4d.
//
// Solidity: function toContract() payable returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) ToContract() (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.ToContract(&_UploadMedicalrecords.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.Transfer(&_UploadMedicalrecords.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.Transfer(&_UploadMedicalrecords.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.TransferFrom(&_UploadMedicalrecords.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.TransferFrom(&_UploadMedicalrecords.TransactOpts, from, to, amount)
}

// UserAddMedicalInformation is a paid mutator transaction binding the contract method 0xab547346.
//
// Solidity: function user_AddMedicalInformation(string Proute, address _soliciter) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) UserAddMedicalInformation(opts *bind.TransactOpts, Proute string, _soliciter common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "user_AddMedicalInformation", Proute, _soliciter)
}

// UserAddMedicalInformation is a paid mutator transaction binding the contract method 0xab547346.
//
// Solidity: function user_AddMedicalInformation(string Proute, address _soliciter) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsSession) UserAddMedicalInformation(Proute string, _soliciter common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.UserAddMedicalInformation(&_UploadMedicalrecords.TransactOpts, Proute, _soliciter)
}

// UserAddMedicalInformation is a paid mutator transaction binding the contract method 0xab547346.
//
// Solidity: function user_AddMedicalInformation(string Proute, address _soliciter) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) UserAddMedicalInformation(Proute string, _soliciter common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.UserAddMedicalInformation(&_UploadMedicalrecords.TransactOpts, Proute, _soliciter)
}

// UserAdduser is a paid mutator transaction binding the contract method 0x1dd9ec67.
//
// Solidity: function user_Adduser(address _user) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) UserAdduser(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "user_Adduser", _user)
}

// UserAdduser is a paid mutator transaction binding the contract method 0x1dd9ec67.
//
// Solidity: function user_Adduser(address _user) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsSession) UserAdduser(_user common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.UserAdduser(&_UploadMedicalrecords.TransactOpts, _user)
}

// UserAdduser is a paid mutator transaction binding the contract method 0x1dd9ec67.
//
// Solidity: function user_Adduser(address _user) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) UserAdduser(_user common.Address) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.UserAdduser(&_UploadMedicalrecords.TransactOpts, _user)
}

// UserUPMedicalExaminationReport is a paid mutator transaction binding the contract method 0x197e456a.
//
// Solidity: function user_UPMedicalExaminationReport(string name, string url) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) UserUPMedicalExaminationReport(opts *bind.TransactOpts, name string, url string) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "user_UPMedicalExaminationReport", name, url)
}

// UserUPMedicalExaminationReport is a paid mutator transaction binding the contract method 0x197e456a.
//
// Solidity: function user_UPMedicalExaminationReport(string name, string url) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsSession) UserUPMedicalExaminationReport(name string, url string) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.UserUPMedicalExaminationReport(&_UploadMedicalrecords.TransactOpts, name, url)
}

// UserUPMedicalExaminationReport is a paid mutator transaction binding the contract method 0x197e456a.
//
// Solidity: function user_UPMedicalExaminationReport(string name, string url) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) UserUPMedicalExaminationReport(name string, url string) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.UserUPMedicalExaminationReport(&_UploadMedicalrecords.TransactOpts, name, url)
}

// UserUPMedicalinformation is a paid mutator transaction binding the contract method 0x868e0bd4.
//
// Solidity: function user_UPMedicalinformation(string name, string url) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactor) UserUPMedicalinformation(opts *bind.TransactOpts, name string, url string) (*types.Transaction, error) {
	return _UploadMedicalrecords.contract.Transact(opts, "user_UPMedicalinformation", name, url)
}

// UserUPMedicalinformation is a paid mutator transaction binding the contract method 0x868e0bd4.
//
// Solidity: function user_UPMedicalinformation(string name, string url) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsSession) UserUPMedicalinformation(name string, url string) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.UserUPMedicalinformation(&_UploadMedicalrecords.TransactOpts, name, url)
}

// UserUPMedicalinformation is a paid mutator transaction binding the contract method 0x868e0bd4.
//
// Solidity: function user_UPMedicalinformation(string name, string url) returns()
func (_UploadMedicalrecords *UploadMedicalrecordsTransactorSession) UserUPMedicalinformation(name string, url string) (*types.Transaction, error) {
	return _UploadMedicalrecords.Contract.UserUPMedicalinformation(&_UploadMedicalrecords.TransactOpts, name, url)
}

// UploadMedicalrecordsA11GiveETHIterator is returned from FilterA11GiveETH and is used to iterate over the raw logs and unpacked data for A11GiveETH events raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsA11GiveETHIterator struct {
	Event *UploadMedicalrecordsA11GiveETH // Event containing the contract specifics and raw log

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
func (it *UploadMedicalrecordsA11GiveETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UploadMedicalrecordsA11GiveETH)
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
		it.Event = new(UploadMedicalrecordsA11GiveETH)
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
func (it *UploadMedicalrecordsA11GiveETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UploadMedicalrecordsA11GiveETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UploadMedicalrecordsA11GiveETH represents a A11GiveETH event raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsA11GiveETH struct {
	User1 common.Address
	User2 common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterA11GiveETH is a free log retrieval operation binding the contract event 0x41003739cb7adda22b5e7029b076372e289d38e49fcb784b70e3f2883dfb0265.
//
// Solidity: event A11GiveETH(address indexed user1, address indexed user2, uint256 indexed price)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) FilterA11GiveETH(opts *bind.FilterOpts, user1 []common.Address, user2 []common.Address, price []*big.Int) (*UploadMedicalrecordsA11GiveETHIterator, error) {

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

	logs, sub, err := _UploadMedicalrecords.contract.FilterLogs(opts, "A11GiveETH", user1Rule, user2Rule, priceRule)
	if err != nil {
		return nil, err
	}
	return &UploadMedicalrecordsA11GiveETHIterator{contract: _UploadMedicalrecords.contract, event: "A11GiveETH", logs: logs, sub: sub}, nil
}

// WatchA11GiveETH is a free log subscription operation binding the contract event 0x41003739cb7adda22b5e7029b076372e289d38e49fcb784b70e3f2883dfb0265.
//
// Solidity: event A11GiveETH(address indexed user1, address indexed user2, uint256 indexed price)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) WatchA11GiveETH(opts *bind.WatchOpts, sink chan<- *UploadMedicalrecordsA11GiveETH, user1 []common.Address, user2 []common.Address, price []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _UploadMedicalrecords.contract.WatchLogs(opts, "A11GiveETH", user1Rule, user2Rule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UploadMedicalrecordsA11GiveETH)
				if err := _UploadMedicalrecords.contract.UnpackLog(event, "A11GiveETH", log); err != nil {
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
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) ParseA11GiveETH(log types.Log) (*UploadMedicalrecordsA11GiveETH, error) {
	event := new(UploadMedicalrecordsA11GiveETH)
	if err := _UploadMedicalrecords.contract.UnpackLog(event, "A11GiveETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UploadMedicalrecordsAddtokensIterator is returned from FilterAddtokens and is used to iterate over the raw logs and unpacked data for Addtokens events raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsAddtokensIterator struct {
	Event *UploadMedicalrecordsAddtokens // Event containing the contract specifics and raw log

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
func (it *UploadMedicalrecordsAddtokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UploadMedicalrecordsAddtokens)
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
		it.Event = new(UploadMedicalrecordsAddtokens)
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
func (it *UploadMedicalrecordsAddtokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UploadMedicalrecordsAddtokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UploadMedicalrecordsAddtokens represents a Addtokens event raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsAddtokens struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddtokens is a free log retrieval operation binding the contract event 0x140232c0c4b8ad28248113c4185b3b3af0ba6f0589381f850b8731ad8af29db8.
//
// Solidity: event Addtokens(address indexed owner, uint256 indexed amount)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) FilterAddtokens(opts *bind.FilterOpts, owner []common.Address, amount []*big.Int) (*UploadMedicalrecordsAddtokensIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _UploadMedicalrecords.contract.FilterLogs(opts, "Addtokens", ownerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &UploadMedicalrecordsAddtokensIterator{contract: _UploadMedicalrecords.contract, event: "Addtokens", logs: logs, sub: sub}, nil
}

// WatchAddtokens is a free log subscription operation binding the contract event 0x140232c0c4b8ad28248113c4185b3b3af0ba6f0589381f850b8731ad8af29db8.
//
// Solidity: event Addtokens(address indexed owner, uint256 indexed amount)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) WatchAddtokens(opts *bind.WatchOpts, sink chan<- *UploadMedicalrecordsAddtokens, owner []common.Address, amount []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _UploadMedicalrecords.contract.WatchLogs(opts, "Addtokens", ownerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UploadMedicalrecordsAddtokens)
				if err := _UploadMedicalrecords.contract.UnpackLog(event, "Addtokens", log); err != nil {
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
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) ParseAddtokens(log types.Log) (*UploadMedicalrecordsAddtokens, error) {
	event := new(UploadMedicalrecordsAddtokens)
	if err := _UploadMedicalrecords.contract.UnpackLog(event, "Addtokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UploadMedicalrecordsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsApprovalIterator struct {
	Event *UploadMedicalrecordsApproval // Event containing the contract specifics and raw log

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
func (it *UploadMedicalrecordsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UploadMedicalrecordsApproval)
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
		it.Event = new(UploadMedicalrecordsApproval)
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
func (it *UploadMedicalrecordsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UploadMedicalrecordsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UploadMedicalrecordsApproval represents a Approval event raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UploadMedicalrecordsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UploadMedicalrecords.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UploadMedicalrecordsApprovalIterator{contract: _UploadMedicalrecords.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UploadMedicalrecordsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UploadMedicalrecords.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UploadMedicalrecordsApproval)
				if err := _UploadMedicalrecords.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) ParseApproval(log types.Log) (*UploadMedicalrecordsApproval, error) {
	event := new(UploadMedicalrecordsApproval)
	if err := _UploadMedicalrecords.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UploadMedicalrecordsEthgetAsIterator is returned from FilterEthgetAs and is used to iterate over the raw logs and unpacked data for EthgetAs events raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsEthgetAsIterator struct {
	Event *UploadMedicalrecordsEthgetAs // Event containing the contract specifics and raw log

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
func (it *UploadMedicalrecordsEthgetAsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UploadMedicalrecordsEthgetAs)
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
		it.Event = new(UploadMedicalrecordsEthgetAs)
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
func (it *UploadMedicalrecordsEthgetAsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UploadMedicalrecordsEthgetAsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UploadMedicalrecordsEthgetAs represents a EthgetAs event raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsEthgetAs struct {
	User1 common.Address
	User2 common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEthgetAs is a free log retrieval operation binding the contract event 0x4e28aa9ff438c656a77043806d0d6a21a098eda7102c9b52698d2a6f321fc3e6.
//
// Solidity: event EthgetAs(address indexed user1, address indexed user2, uint256 indexed price)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) FilterEthgetAs(opts *bind.FilterOpts, user1 []common.Address, user2 []common.Address, price []*big.Int) (*UploadMedicalrecordsEthgetAsIterator, error) {

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

	logs, sub, err := _UploadMedicalrecords.contract.FilterLogs(opts, "EthgetAs", user1Rule, user2Rule, priceRule)
	if err != nil {
		return nil, err
	}
	return &UploadMedicalrecordsEthgetAsIterator{contract: _UploadMedicalrecords.contract, event: "EthgetAs", logs: logs, sub: sub}, nil
}

// WatchEthgetAs is a free log subscription operation binding the contract event 0x4e28aa9ff438c656a77043806d0d6a21a098eda7102c9b52698d2a6f321fc3e6.
//
// Solidity: event EthgetAs(address indexed user1, address indexed user2, uint256 indexed price)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) WatchEthgetAs(opts *bind.WatchOpts, sink chan<- *UploadMedicalrecordsEthgetAs, user1 []common.Address, user2 []common.Address, price []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _UploadMedicalrecords.contract.WatchLogs(opts, "EthgetAs", user1Rule, user2Rule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UploadMedicalrecordsEthgetAs)
				if err := _UploadMedicalrecords.contract.UnpackLog(event, "EthgetAs", log); err != nil {
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
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) ParseEthgetAs(log types.Log) (*UploadMedicalrecordsEthgetAs, error) {
	event := new(UploadMedicalrecordsEthgetAs)
	if err := _UploadMedicalrecords.contract.UnpackLog(event, "EthgetAs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UploadMedicalrecordsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsTransferIterator struct {
	Event *UploadMedicalrecordsTransfer // Event containing the contract specifics and raw log

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
func (it *UploadMedicalrecordsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UploadMedicalrecordsTransfer)
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
		it.Event = new(UploadMedicalrecordsTransfer)
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
func (it *UploadMedicalrecordsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UploadMedicalrecordsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UploadMedicalrecordsTransfer represents a Transfer event raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UploadMedicalrecordsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UploadMedicalrecords.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UploadMedicalrecordsTransferIterator{contract: _UploadMedicalrecords.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UploadMedicalrecordsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UploadMedicalrecords.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UploadMedicalrecordsTransfer)
				if err := _UploadMedicalrecords.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) ParseTransfer(log types.Log) (*UploadMedicalrecordsTransfer, error) {
	event := new(UploadMedicalrecordsTransfer)
	if err := _UploadMedicalrecords.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UploadMedicalrecordsUploadmedicaldataIterator is returned from FilterUploadmedicaldata and is used to iterate over the raw logs and unpacked data for Uploadmedicaldata events raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsUploadmedicaldataIterator struct {
	Event *UploadMedicalrecordsUploadmedicaldata // Event containing the contract specifics and raw log

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
func (it *UploadMedicalrecordsUploadmedicaldataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UploadMedicalrecordsUploadmedicaldata)
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
		it.Event = new(UploadMedicalrecordsUploadmedicaldata)
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
func (it *UploadMedicalrecordsUploadmedicaldataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UploadMedicalrecordsUploadmedicaldataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UploadMedicalrecordsUploadmedicaldata represents a Uploadmedicaldata event raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsUploadmedicaldata struct {
	User      common.Address
	Route     common.Hash
	Soliciter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUploadmedicaldata is a free log retrieval operation binding the contract event 0x49868a6e303058cc0ee85fad003812b132b8ba154aafa643a8763b774a8b8495.
//
// Solidity: event Uploadmedicaldata(address indexed user, string indexed route, address indexed soliciter)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) FilterUploadmedicaldata(opts *bind.FilterOpts, user []common.Address, route []string, soliciter []common.Address) (*UploadMedicalrecordsUploadmedicaldataIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var routeRule []interface{}
	for _, routeItem := range route {
		routeRule = append(routeRule, routeItem)
	}
	var soliciterRule []interface{}
	for _, soliciterItem := range soliciter {
		soliciterRule = append(soliciterRule, soliciterItem)
	}

	logs, sub, err := _UploadMedicalrecords.contract.FilterLogs(opts, "Uploadmedicaldata", userRule, routeRule, soliciterRule)
	if err != nil {
		return nil, err
	}
	return &UploadMedicalrecordsUploadmedicaldataIterator{contract: _UploadMedicalrecords.contract, event: "Uploadmedicaldata", logs: logs, sub: sub}, nil
}

// WatchUploadmedicaldata is a free log subscription operation binding the contract event 0x49868a6e303058cc0ee85fad003812b132b8ba154aafa643a8763b774a8b8495.
//
// Solidity: event Uploadmedicaldata(address indexed user, string indexed route, address indexed soliciter)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) WatchUploadmedicaldata(opts *bind.WatchOpts, sink chan<- *UploadMedicalrecordsUploadmedicaldata, user []common.Address, route []string, soliciter []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var routeRule []interface{}
	for _, routeItem := range route {
		routeRule = append(routeRule, routeItem)
	}
	var soliciterRule []interface{}
	for _, soliciterItem := range soliciter {
		soliciterRule = append(soliciterRule, soliciterItem)
	}

	logs, sub, err := _UploadMedicalrecords.contract.WatchLogs(opts, "Uploadmedicaldata", userRule, routeRule, soliciterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UploadMedicalrecordsUploadmedicaldata)
				if err := _UploadMedicalrecords.contract.UnpackLog(event, "Uploadmedicaldata", log); err != nil {
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

// ParseUploadmedicaldata is a log parse operation binding the contract event 0x49868a6e303058cc0ee85fad003812b132b8ba154aafa643a8763b774a8b8495.
//
// Solidity: event Uploadmedicaldata(address indexed user, string indexed route, address indexed soliciter)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) ParseUploadmedicaldata(log types.Log) (*UploadMedicalrecordsUploadmedicaldata, error) {
	event := new(UploadMedicalrecordsUploadmedicaldata)
	if err := _UploadMedicalrecords.contract.UnpackLog(event, "Uploadmedicaldata", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UploadMedicalrecordsGainerUploadmedicaldataIterator is returned from FilterGainerUploadmedicaldata and is used to iterate over the raw logs and unpacked data for GainerUploadmedicaldata events raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsGainerUploadmedicaldataIterator struct {
	Event *UploadMedicalrecordsGainerUploadmedicaldata // Event containing the contract specifics and raw log

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
func (it *UploadMedicalrecordsGainerUploadmedicaldataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UploadMedicalrecordsGainerUploadmedicaldata)
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
		it.Event = new(UploadMedicalrecordsGainerUploadmedicaldata)
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
func (it *UploadMedicalrecordsGainerUploadmedicaldataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UploadMedicalrecordsGainerUploadmedicaldataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UploadMedicalrecordsGainerUploadmedicaldata represents a GainerUploadmedicaldata event raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsGainerUploadmedicaldata struct {
	Route     common.Hash
	Soliciter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGainerUploadmedicaldata is a free log retrieval operation binding the contract event 0x575ef43b443a501121681e81acef51358e5f96f8b3493a3b964389af363dbaf7.
//
// Solidity: event gainerUploadmedicaldata(string indexed route, address indexed soliciter)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) FilterGainerUploadmedicaldata(opts *bind.FilterOpts, route []string, soliciter []common.Address) (*UploadMedicalrecordsGainerUploadmedicaldataIterator, error) {

	var routeRule []interface{}
	for _, routeItem := range route {
		routeRule = append(routeRule, routeItem)
	}
	var soliciterRule []interface{}
	for _, soliciterItem := range soliciter {
		soliciterRule = append(soliciterRule, soliciterItem)
	}

	logs, sub, err := _UploadMedicalrecords.contract.FilterLogs(opts, "gainerUploadmedicaldata", routeRule, soliciterRule)
	if err != nil {
		return nil, err
	}
	return &UploadMedicalrecordsGainerUploadmedicaldataIterator{contract: _UploadMedicalrecords.contract, event: "gainerUploadmedicaldata", logs: logs, sub: sub}, nil
}

// WatchGainerUploadmedicaldata is a free log subscription operation binding the contract event 0x575ef43b443a501121681e81acef51358e5f96f8b3493a3b964389af363dbaf7.
//
// Solidity: event gainerUploadmedicaldata(string indexed route, address indexed soliciter)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) WatchGainerUploadmedicaldata(opts *bind.WatchOpts, sink chan<- *UploadMedicalrecordsGainerUploadmedicaldata, route []string, soliciter []common.Address) (event.Subscription, error) {

	var routeRule []interface{}
	for _, routeItem := range route {
		routeRule = append(routeRule, routeItem)
	}
	var soliciterRule []interface{}
	for _, soliciterItem := range soliciter {
		soliciterRule = append(soliciterRule, soliciterItem)
	}

	logs, sub, err := _UploadMedicalrecords.contract.WatchLogs(opts, "gainerUploadmedicaldata", routeRule, soliciterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UploadMedicalrecordsGainerUploadmedicaldata)
				if err := _UploadMedicalrecords.contract.UnpackLog(event, "gainerUploadmedicaldata", log); err != nil {
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

// ParseGainerUploadmedicaldata is a log parse operation binding the contract event 0x575ef43b443a501121681e81acef51358e5f96f8b3493a3b964389af363dbaf7.
//
// Solidity: event gainerUploadmedicaldata(string indexed route, address indexed soliciter)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) ParseGainerUploadmedicaldata(log types.Log) (*UploadMedicalrecordsGainerUploadmedicaldata, error) {
	event := new(UploadMedicalrecordsGainerUploadmedicaldata)
	if err := _UploadMedicalrecords.contract.UnpackLog(event, "gainerUploadmedicaldata", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UploadMedicalrecordsRewardIterator is returned from FilterReward and is used to iterate over the raw logs and unpacked data for Reward events raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsRewardIterator struct {
	Event *UploadMedicalrecordsReward // Event containing the contract specifics and raw log

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
func (it *UploadMedicalrecordsRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UploadMedicalrecordsReward)
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
		it.Event = new(UploadMedicalrecordsReward)
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
func (it *UploadMedicalrecordsRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UploadMedicalrecordsRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UploadMedicalrecordsReward represents a Reward event raised by the UploadMedicalrecords contract.
type UploadMedicalrecordsReward struct {
	Owner common.Address
	Route common.Hash
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReward is a free log retrieval operation binding the contract event 0xab58f406e20f71235aadf9324c46cdb688e9be29c543b558f4cc8601a41ea4bb.
//
// Solidity: event reward(address indexed owner, string indexed route, uint256 indexed price)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) FilterReward(opts *bind.FilterOpts, owner []common.Address, route []string, price []*big.Int) (*UploadMedicalrecordsRewardIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var routeRule []interface{}
	for _, routeItem := range route {
		routeRule = append(routeRule, routeItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _UploadMedicalrecords.contract.FilterLogs(opts, "reward", ownerRule, routeRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &UploadMedicalrecordsRewardIterator{contract: _UploadMedicalrecords.contract, event: "reward", logs: logs, sub: sub}, nil
}

// WatchReward is a free log subscription operation binding the contract event 0xab58f406e20f71235aadf9324c46cdb688e9be29c543b558f4cc8601a41ea4bb.
//
// Solidity: event reward(address indexed owner, string indexed route, uint256 indexed price)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) WatchReward(opts *bind.WatchOpts, sink chan<- *UploadMedicalrecordsReward, owner []common.Address, route []string, price []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var routeRule []interface{}
	for _, routeItem := range route {
		routeRule = append(routeRule, routeItem)
	}
	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _UploadMedicalrecords.contract.WatchLogs(opts, "reward", ownerRule, routeRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UploadMedicalrecordsReward)
				if err := _UploadMedicalrecords.contract.UnpackLog(event, "reward", log); err != nil {
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

// ParseReward is a log parse operation binding the contract event 0xab58f406e20f71235aadf9324c46cdb688e9be29c543b558f4cc8601a41ea4bb.
//
// Solidity: event reward(address indexed owner, string indexed route, uint256 indexed price)
func (_UploadMedicalrecords *UploadMedicalrecordsFilterer) ParseReward(log types.Log) (*UploadMedicalrecordsReward, error) {
	event := new(UploadMedicalrecordsReward)
	if err := _UploadMedicalrecords.contract.UnpackLog(event, "reward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
