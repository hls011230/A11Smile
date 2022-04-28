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

// gggGainer_upMedicalInformation is an auto generated low-level Go binding around an user-defined struct.
type gggGainer_upMedicalInformation struct {
	MedicalName  string
	Min          uint64
	Max          uint64
	Account      uint64
	HospitalName string
	Exit         bool
}

// gggGainer_upMedicalInformation1 is an auto generated low-level Go binding around an user-defined struct.
type gggGainer_upMedicalInformation1 struct {
	Min                       uint64
	Max                       uint64
	Medicalrecordrequirements string
	Requirementdescription    string
}

// GenMetaData contains all meta data concerning the Gen contract.
var GenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hospitalName_\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"min_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"account_\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"MedicalName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_Medicalrecordrequirements\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_Requirementdescription\",\"type\":\"string\"}],\"name\":\"gainer_AddMedicalInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bbb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gainer_SeeOwnMedicalInformationName\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"Neirong1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SeeGainerMedicalInformations\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"Medicalrecordrequirements\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Requirementdescription\",\"type\":\"string\"}],\"internalType\":\"structggg.Gainer_upMedicalInformation1[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SeeGainerMedicalInformationsName\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"account\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"HospitalName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"}],\"internalType\":\"structggg.Gainer_upMedicalInformation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"Shouye1\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GenABI is the input ABI used to generate the binding from.
// Deprecated: Use GenMetaData.ABI instead.
var GenABI = GenMetaData.ABI

// Gen is an auto generated Go binding around an Ethereum contract.
type Gen struct {
	GenCaller     // Read-only binding to the contract
	GenTransactor // Write-only binding to the contract
	GenFilterer   // Log filterer for contract events
}

// GenCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenSession struct {
	Contract     *Gen              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenCallerSession struct {
	Contract *GenCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenTransactorSession struct {
	Contract     *GenTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenRaw struct {
	Contract *Gen // Generic contract binding to access the raw methods on
}

// GenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenCallerRaw struct {
	Contract *GenCaller // Generic read-only contract binding to access the raw methods on
}

// GenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenTransactorRaw struct {
	Contract *GenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGen creates a new instance of Gen, bound to a specific deployed contract.
func NewGen(address common.Address, backend bind.ContractBackend) (*Gen, error) {
	contract, err := bindGen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gen{GenCaller: GenCaller{contract: contract}, GenTransactor: GenTransactor{contract: contract}, GenFilterer: GenFilterer{contract: contract}}, nil
}

// NewGenCaller creates a new read-only instance of Gen, bound to a specific deployed contract.
func NewGenCaller(address common.Address, caller bind.ContractCaller) (*GenCaller, error) {
	contract, err := bindGen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenCaller{contract: contract}, nil
}

// NewGenTransactor creates a new write-only instance of Gen, bound to a specific deployed contract.
func NewGenTransactor(address common.Address, transactor bind.ContractTransactor) (*GenTransactor, error) {
	contract, err := bindGen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenTransactor{contract: contract}, nil
}

// NewGenFilterer creates a new log filterer instance of Gen, bound to a specific deployed contract.
func NewGenFilterer(address common.Address, filterer bind.ContractFilterer) (*GenFilterer, error) {
	contract, err := bindGen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenFilterer{contract: contract}, nil
}

// bindGen binds a generic wrapper to an already deployed contract.
func bindGen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gen *GenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gen.Contract.GenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gen *GenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gen.Contract.GenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gen *GenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gen.Contract.GenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gen *GenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gen *GenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gen *GenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gen.Contract.contract.Transact(opts, method, params...)
}

// Neirong1 is a free data retrieval call binding the contract method 0xca5e20b6.
//
// Solidity: function Neirong1(address a, uint256 b) view returns(uint256, uint256, string, string)
func (_Gen *GenCaller) Neirong1(opts *bind.CallOpts, a common.Address, b *big.Int) (*big.Int, *big.Int, string, string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "Neirong1", a, b)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)

	return out0, out1, out2, out3, err

}

// Neirong1 is a free data retrieval call binding the contract method 0xca5e20b6.
//
// Solidity: function Neirong1(address a, uint256 b) view returns(uint256, uint256, string, string)
func (_Gen *GenSession) Neirong1(a common.Address, b *big.Int) (*big.Int, *big.Int, string, string, error) {
	return _Gen.Contract.Neirong1(&_Gen.CallOpts, a, b)
}

// Neirong1 is a free data retrieval call binding the contract method 0xca5e20b6.
//
// Solidity: function Neirong1(address a, uint256 b) view returns(uint256, uint256, string, string)
func (_Gen *GenCallerSession) Neirong1(a common.Address, b *big.Int) (*big.Int, *big.Int, string, string, error) {
	return _Gen.Contract.Neirong1(&_Gen.CallOpts, a, b)
}

// SeeGainerMedicalInformations is a free data retrieval call binding the contract method 0x96085081.
//
// Solidity: function SeeGainerMedicalInformations() view returns((uint64,uint64,string,string)[])
func (_Gen *GenCaller) SeeGainerMedicalInformations(opts *bind.CallOpts) ([]gggGainer_upMedicalInformation1, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "SeeGainerMedicalInformations")

	if err != nil {
		return *new([]gggGainer_upMedicalInformation1), err
	}

	out0 := *abi.ConvertType(out[0], new([]gggGainer_upMedicalInformation1)).(*[]gggGainer_upMedicalInformation1)

	return out0, err

}

// SeeGainerMedicalInformations is a free data retrieval call binding the contract method 0x96085081.
//
// Solidity: function SeeGainerMedicalInformations() view returns((uint64,uint64,string,string)[])
func (_Gen *GenSession) SeeGainerMedicalInformations() ([]gggGainer_upMedicalInformation1, error) {
	return _Gen.Contract.SeeGainerMedicalInformations(&_Gen.CallOpts)
}

// SeeGainerMedicalInformations is a free data retrieval call binding the contract method 0x96085081.
//
// Solidity: function SeeGainerMedicalInformations() view returns((uint64,uint64,string,string)[])
func (_Gen *GenCallerSession) SeeGainerMedicalInformations() ([]gggGainer_upMedicalInformation1, error) {
	return _Gen.Contract.SeeGainerMedicalInformations(&_Gen.CallOpts)
}

// SeeGainerMedicalInformationsName is a free data retrieval call binding the contract method 0xc506b536.
//
// Solidity: function SeeGainerMedicalInformationsName() view returns((string,uint64,uint64,uint64,string,bool)[])
func (_Gen *GenCaller) SeeGainerMedicalInformationsName(opts *bind.CallOpts) ([]gggGainer_upMedicalInformation, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "SeeGainerMedicalInformationsName")

	if err != nil {
		return *new([]gggGainer_upMedicalInformation), err
	}

	out0 := *abi.ConvertType(out[0], new([]gggGainer_upMedicalInformation)).(*[]gggGainer_upMedicalInformation)

	return out0, err

}

// SeeGainerMedicalInformationsName is a free data retrieval call binding the contract method 0xc506b536.
//
// Solidity: function SeeGainerMedicalInformationsName() view returns((string,uint64,uint64,uint64,string,bool)[])
func (_Gen *GenSession) SeeGainerMedicalInformationsName() ([]gggGainer_upMedicalInformation, error) {
	return _Gen.Contract.SeeGainerMedicalInformationsName(&_Gen.CallOpts)
}

// SeeGainerMedicalInformationsName is a free data retrieval call binding the contract method 0xc506b536.
//
// Solidity: function SeeGainerMedicalInformationsName() view returns((string,uint64,uint64,uint64,string,bool)[])
func (_Gen *GenCallerSession) SeeGainerMedicalInformationsName() ([]gggGainer_upMedicalInformation, error) {
	return _Gen.Contract.SeeGainerMedicalInformationsName(&_Gen.CallOpts)
}

// Shouye1 is a free data retrieval call binding the contract method 0x323c593b.
//
// Solidity: function Shouye1(address a, uint256 b) view returns(string, uint256, uint256, uint256, string, bool)
func (_Gen *GenCaller) Shouye1(opts *bind.CallOpts, a common.Address, b *big.Int) (string, *big.Int, *big.Int, *big.Int, string, bool, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "Shouye1", a, b)

	if err != nil {
		return *new(string), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// Shouye1 is a free data retrieval call binding the contract method 0x323c593b.
//
// Solidity: function Shouye1(address a, uint256 b) view returns(string, uint256, uint256, uint256, string, bool)
func (_Gen *GenSession) Shouye1(a common.Address, b *big.Int) (string, *big.Int, *big.Int, *big.Int, string, bool, error) {
	return _Gen.Contract.Shouye1(&_Gen.CallOpts, a, b)
}

// Shouye1 is a free data retrieval call binding the contract method 0x323c593b.
//
// Solidity: function Shouye1(address a, uint256 b) view returns(string, uint256, uint256, uint256, string, bool)
func (_Gen *GenCallerSession) Shouye1(a common.Address, b *big.Int) (string, *big.Int, *big.Int, *big.Int, string, bool, error) {
	return _Gen.Contract.Shouye1(&_Gen.CallOpts, a, b)
}

// Bbb is a free data retrieval call binding the contract method 0xd767aee0.
//
// Solidity: function bbb() view returns(uint256)
func (_Gen *GenCaller) Bbb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "bbb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bbb is a free data retrieval call binding the contract method 0xd767aee0.
//
// Solidity: function bbb() view returns(uint256)
func (_Gen *GenSession) Bbb() (*big.Int, error) {
	return _Gen.Contract.Bbb(&_Gen.CallOpts)
}

// Bbb is a free data retrieval call binding the contract method 0xd767aee0.
//
// Solidity: function bbb() view returns(uint256)
func (_Gen *GenCallerSession) Bbb() (*big.Int, error) {
	return _Gen.Contract.Bbb(&_Gen.CallOpts)
}

// GainerSeeOwnMedicalInformationName is a free data retrieval call binding the contract method 0x47f7acae.
//
// Solidity: function gainer_SeeOwnMedicalInformationName() view returns(string[])
func (_Gen *GenCaller) GainerSeeOwnMedicalInformationName(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "gainer_SeeOwnMedicalInformationName")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GainerSeeOwnMedicalInformationName is a free data retrieval call binding the contract method 0x47f7acae.
//
// Solidity: function gainer_SeeOwnMedicalInformationName() view returns(string[])
func (_Gen *GenSession) GainerSeeOwnMedicalInformationName() ([]string, error) {
	return _Gen.Contract.GainerSeeOwnMedicalInformationName(&_Gen.CallOpts)
}

// GainerSeeOwnMedicalInformationName is a free data retrieval call binding the contract method 0x47f7acae.
//
// Solidity: function gainer_SeeOwnMedicalInformationName() view returns(string[])
func (_Gen *GenCallerSession) GainerSeeOwnMedicalInformationName() ([]string, error) {
	return _Gen.Contract.GainerSeeOwnMedicalInformationName(&_Gen.CallOpts)
}

// GainerAddMedicalInformation is a paid mutator transaction binding the contract method 0xe4af7655.
//
// Solidity: function gainer_AddMedicalInformation(string hospitalName_, uint64 min_, uint64 max_, uint64 account_, string MedicalName_, string _Medicalrecordrequirements, string _Requirementdescription) returns()
func (_Gen *GenTransactor) GainerAddMedicalInformation(opts *bind.TransactOpts, hospitalName_ string, min_ uint64, max_ uint64, account_ uint64, MedicalName_ string, _Medicalrecordrequirements string, _Requirementdescription string) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "gainer_AddMedicalInformation", hospitalName_, min_, max_, account_, MedicalName_, _Medicalrecordrequirements, _Requirementdescription)
}

// GainerAddMedicalInformation is a paid mutator transaction binding the contract method 0xe4af7655.
//
// Solidity: function gainer_AddMedicalInformation(string hospitalName_, uint64 min_, uint64 max_, uint64 account_, string MedicalName_, string _Medicalrecordrequirements, string _Requirementdescription) returns()
func (_Gen *GenSession) GainerAddMedicalInformation(hospitalName_ string, min_ uint64, max_ uint64, account_ uint64, MedicalName_ string, _Medicalrecordrequirements string, _Requirementdescription string) (*types.Transaction, error) {
	return _Gen.Contract.GainerAddMedicalInformation(&_Gen.TransactOpts, hospitalName_, min_, max_, account_, MedicalName_, _Medicalrecordrequirements, _Requirementdescription)
}

// GainerAddMedicalInformation is a paid mutator transaction binding the contract method 0xe4af7655.
//
// Solidity: function gainer_AddMedicalInformation(string hospitalName_, uint64 min_, uint64 max_, uint64 account_, string MedicalName_, string _Medicalrecordrequirements, string _Requirementdescription) returns()
func (_Gen *GenTransactorSession) GainerAddMedicalInformation(hospitalName_ string, min_ uint64, max_ uint64, account_ uint64, MedicalName_ string, _Medicalrecordrequirements string, _Requirementdescription string) (*types.Transaction, error) {
	return _Gen.Contract.GainerAddMedicalInformation(&_Gen.TransactOpts, hospitalName_, min_, max_, account_, MedicalName_, _Medicalrecordrequirements, _Requirementdescription)
}
