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

// UploadMedicalrecords_gainerData is an auto generated low-level Go binding around an user-defined struct.
type UploadMedicalrecords_gainerData struct {
	M1 []string
	M2 []string
}

// UploadMedicalrecords_gainergainer_Solicit is an auto generated low-level Go binding around an user-defined struct.
type UploadMedicalrecords_gainergainer_Solicit struct {
	MedicalName string
	Min         *big.Int
	Max         *big.Int
}

// UploadMedicalrecords_gainergainer_upMedicalInformation is an auto generated low-level Go binding around an user-defined struct.
type UploadMedicalrecords_gainergainer_upMedicalInformation struct {
	MedicalName  string
	Department   string
	Min          *big.Int
	Max          *big.Int
	Addr         common.Address
	HospitalName string
	Account      *big.Int
	Exit         bool
}

// UploadMedicalrecords_gainergainer_upMedicalInformation1 is an auto generated low-level Go binding around an user-defined struct.
type UploadMedicalrecords_gainergainer_upMedicalInformation1 struct {
	MedicalName               string
	Min                       *big.Int
	Max                       *big.Int
	Medicalrecordrequirements string
	Requirementdescription    string
}

// UploadMedicalrecords_usercertificate is an auto generated low-level Go binding around an user-defined struct.
type UploadMedicalrecords_usercertificate struct {
	User     common.Address
	Time     *big.Int
	BlockNum *big.Int
	Serial   [32]byte
	M1       []string
	M2       []string
}

// UploadMedicalrecords_useruser_MedicalCertificateState is an auto generated low-level Go binding around an user-defined struct.
type UploadMedicalrecords_useruser_MedicalCertificateState struct {
	User         common.Address
	Soliciter    common.Address
	HospitalName string
	MedicalName  string
	Certificate  [32]byte
	Department   string
}

// GenMetaData contains all meta data concerning the Gen contract.
var GenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"account_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"Department\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"MedicalName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_Medicalrecordrequirements\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_Requirementdescription\",\"type\":\"string\"}],\"name\":\"gainer_AddMedicalInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"HospitalName\",\"type\":\"string\"}],\"name\":\"gainer_setDoctor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"CertificatePath\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"MediacalName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"whether\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ercnum_\",\"type\":\"uint256\"}],\"name\":\"gainer_whether\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_m1\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_m2\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"SetCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"Proute\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_soliciter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"MedicalName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Department\",\"type\":\"string\"}],\"name\":\"user_AddMedicalInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"user_Adduser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"user_UPMedicalExaminationReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"user_UPMedicalinformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AllUserSerialNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Ausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gainer_SeeOwnMedicalInformationed\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structUploadMedicalrecords_gainer.gainer_Solicit[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gainer_SeeOwnMedicalInformationing\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structUploadMedicalrecords_gainer.gainer_Solicit[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gainer_SeeuserUploadMedical\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"HospitalName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"certificate\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"Department\",\"type\":\"string\"}],\"internalType\":\"structUploadMedicalrecords_user.user_MedicalCertificateState[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Medicalinformation_\",\"type\":\"string\"}],\"name\":\"QueryMedicalinformation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ExaminationReportName_\",\"type\":\"string\"}],\"name\":\"QueryPhysicalExaminationReport\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"QueryUserSerialLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"HospitalAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"}],\"name\":\"SeeGainerMedicalInformations\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"Medicalrecordrequirements\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Requirementdescription\",\"type\":\"string\"}],\"internalType\":\"structUploadMedicalrecords_gainer.gainer_upMedicalInformation1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SeeGainerMedicalInformationsName\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Department\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"HospitalName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"account\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"}],\"internalType\":\"structUploadMedicalrecords_gainer.gainer_upMedicalInformation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"user_MedicalExaminationReportstrucrName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"user_MedicalinformationrName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"user_NSeeCertificateState\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"HospitalName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"certificate\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"Department\",\"type\":\"string\"}],\"internalType\":\"structUploadMedicalrecords_user.user_MedicalCertificateState[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"user_ViewMedicalExaminationReport\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"user_ViewMedicalInformation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"user_YSeeCertificateState\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"HospitalName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"certificate\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"Department\",\"type\":\"string\"}],\"internalType\":\"structUploadMedicalrecords_user.user_MedicalCertificateState[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userAmedicalInformation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"soliciters\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"UserSerial\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"UserSerialName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userWareHouse\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewMedicalExaminationReportCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewMedicalinformationCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ViewMedicalNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ViewUserAllMedicalExaminationReport\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ViewUserAllMedicalinformation\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_num\",\"type\":\"bytes32\"}],\"name\":\"ViewUserCertificate\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"serial\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"m1\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"m2\",\"type\":\"string[]\"}],\"internalType\":\"structUploadMedicalrecords_user.certificate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_medicalName\",\"type\":\"string\"}],\"name\":\"ViewWareHouseLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_medicalName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"ViewWareouse\",\"outputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"m1\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"m2\",\"type\":\"string[]\"}],\"internalType\":\"structUploadMedicalrecords_gainer.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// AllUserSerialNames is a free data retrieval call binding the contract method 0x7a3d7175.
//
// Solidity: function AllUserSerialNames() view returns(string[])
func (_Gen *GenCaller) AllUserSerialNames(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "AllUserSerialNames")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// AllUserSerialNames is a free data retrieval call binding the contract method 0x7a3d7175.
//
// Solidity: function AllUserSerialNames() view returns(string[])
func (_Gen *GenSession) AllUserSerialNames() ([]string, error) {
	return _Gen.Contract.AllUserSerialNames(&_Gen.CallOpts)
}

// AllUserSerialNames is a free data retrieval call binding the contract method 0x7a3d7175.
//
// Solidity: function AllUserSerialNames() view returns(string[])
func (_Gen *GenCallerSession) AllUserSerialNames() ([]string, error) {
	return _Gen.Contract.AllUserSerialNames(&_Gen.CallOpts)
}

// Ausers is a free data retrieval call binding the contract method 0x51cc2fc7.
//
// Solidity: function Ausers(address ) view returns(bool)
func (_Gen *GenCaller) Ausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "Ausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ausers is a free data retrieval call binding the contract method 0x51cc2fc7.
//
// Solidity: function Ausers(address ) view returns(bool)
func (_Gen *GenSession) Ausers(arg0 common.Address) (bool, error) {
	return _Gen.Contract.Ausers(&_Gen.CallOpts, arg0)
}

// Ausers is a free data retrieval call binding the contract method 0x51cc2fc7.
//
// Solidity: function Ausers(address ) view returns(bool)
func (_Gen *GenCallerSession) Ausers(arg0 common.Address) (bool, error) {
	return _Gen.Contract.Ausers(&_Gen.CallOpts, arg0)
}

// QueryMedicalinformation is a free data retrieval call binding the contract method 0x9ba614d5.
//
// Solidity: function QueryMedicalinformation(string Medicalinformation_) view returns(string)
func (_Gen *GenCaller) QueryMedicalinformation(opts *bind.CallOpts, Medicalinformation_ string) (string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "QueryMedicalinformation", Medicalinformation_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// QueryMedicalinformation is a free data retrieval call binding the contract method 0x9ba614d5.
//
// Solidity: function QueryMedicalinformation(string Medicalinformation_) view returns(string)
func (_Gen *GenSession) QueryMedicalinformation(Medicalinformation_ string) (string, error) {
	return _Gen.Contract.QueryMedicalinformation(&_Gen.CallOpts, Medicalinformation_)
}

// QueryMedicalinformation is a free data retrieval call binding the contract method 0x9ba614d5.
//
// Solidity: function QueryMedicalinformation(string Medicalinformation_) view returns(string)
func (_Gen *GenCallerSession) QueryMedicalinformation(Medicalinformation_ string) (string, error) {
	return _Gen.Contract.QueryMedicalinformation(&_Gen.CallOpts, Medicalinformation_)
}

// QueryPhysicalExaminationReport is a free data retrieval call binding the contract method 0xe6dec846.
//
// Solidity: function QueryPhysicalExaminationReport(string ExaminationReportName_) view returns(string)
func (_Gen *GenCaller) QueryPhysicalExaminationReport(opts *bind.CallOpts, ExaminationReportName_ string) (string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "QueryPhysicalExaminationReport", ExaminationReportName_)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// QueryPhysicalExaminationReport is a free data retrieval call binding the contract method 0xe6dec846.
//
// Solidity: function QueryPhysicalExaminationReport(string ExaminationReportName_) view returns(string)
func (_Gen *GenSession) QueryPhysicalExaminationReport(ExaminationReportName_ string) (string, error) {
	return _Gen.Contract.QueryPhysicalExaminationReport(&_Gen.CallOpts, ExaminationReportName_)
}

// QueryPhysicalExaminationReport is a free data retrieval call binding the contract method 0xe6dec846.
//
// Solidity: function QueryPhysicalExaminationReport(string ExaminationReportName_) view returns(string)
func (_Gen *GenCallerSession) QueryPhysicalExaminationReport(ExaminationReportName_ string) (string, error) {
	return _Gen.Contract.QueryPhysicalExaminationReport(&_Gen.CallOpts, ExaminationReportName_)
}

// QueryUserSerialLength is a free data retrieval call binding the contract method 0x10540f37.
//
// Solidity: function QueryUserSerialLength(address user) view returns(uint256)
func (_Gen *GenCaller) QueryUserSerialLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "QueryUserSerialLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryUserSerialLength is a free data retrieval call binding the contract method 0x10540f37.
//
// Solidity: function QueryUserSerialLength(address user) view returns(uint256)
func (_Gen *GenSession) QueryUserSerialLength(user common.Address) (*big.Int, error) {
	return _Gen.Contract.QueryUserSerialLength(&_Gen.CallOpts, user)
}

// QueryUserSerialLength is a free data retrieval call binding the contract method 0x10540f37.
//
// Solidity: function QueryUserSerialLength(address user) view returns(uint256)
func (_Gen *GenCallerSession) QueryUserSerialLength(user common.Address) (*big.Int, error) {
	return _Gen.Contract.QueryUserSerialLength(&_Gen.CallOpts, user)
}

// SeeGainerMedicalInformations is a free data retrieval call binding the contract method 0x2285b528.
//
// Solidity: function SeeGainerMedicalInformations(address HospitalAddress, string MedicalName) view returns((string,uint256,uint256,string,string))
func (_Gen *GenCaller) SeeGainerMedicalInformations(opts *bind.CallOpts, HospitalAddress common.Address, MedicalName string) (UploadMedicalrecords_gainergainer_upMedicalInformation1, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "SeeGainerMedicalInformations", HospitalAddress, MedicalName)

	if err != nil {
		return *new(UploadMedicalrecords_gainergainer_upMedicalInformation1), err
	}

	out0 := *abi.ConvertType(out[0], new(UploadMedicalrecords_gainergainer_upMedicalInformation1)).(*UploadMedicalrecords_gainergainer_upMedicalInformation1)

	return out0, err

}

// SeeGainerMedicalInformations is a free data retrieval call binding the contract method 0x2285b528.
//
// Solidity: function SeeGainerMedicalInformations(address HospitalAddress, string MedicalName) view returns((string,uint256,uint256,string,string))
func (_Gen *GenSession) SeeGainerMedicalInformations(HospitalAddress common.Address, MedicalName string) (UploadMedicalrecords_gainergainer_upMedicalInformation1, error) {
	return _Gen.Contract.SeeGainerMedicalInformations(&_Gen.CallOpts, HospitalAddress, MedicalName)
}

// SeeGainerMedicalInformations is a free data retrieval call binding the contract method 0x2285b528.
//
// Solidity: function SeeGainerMedicalInformations(address HospitalAddress, string MedicalName) view returns((string,uint256,uint256,string,string))
func (_Gen *GenCallerSession) SeeGainerMedicalInformations(HospitalAddress common.Address, MedicalName string) (UploadMedicalrecords_gainergainer_upMedicalInformation1, error) {
	return _Gen.Contract.SeeGainerMedicalInformations(&_Gen.CallOpts, HospitalAddress, MedicalName)
}

// SeeGainerMedicalInformationsName is a free data retrieval call binding the contract method 0xc506b536.
//
// Solidity: function SeeGainerMedicalInformationsName() view returns((string,string,uint256,uint256,address,string,uint256,bool)[])
func (_Gen *GenCaller) SeeGainerMedicalInformationsName(opts *bind.CallOpts) ([]UploadMedicalrecords_gainergainer_upMedicalInformation, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "SeeGainerMedicalInformationsName")

	if err != nil {
		return *new([]UploadMedicalrecords_gainergainer_upMedicalInformation), err
	}

	out0 := *abi.ConvertType(out[0], new([]UploadMedicalrecords_gainergainer_upMedicalInformation)).(*[]UploadMedicalrecords_gainergainer_upMedicalInformation)

	return out0, err

}

// SeeGainerMedicalInformationsName is a free data retrieval call binding the contract method 0xc506b536.
//
// Solidity: function SeeGainerMedicalInformationsName() view returns((string,string,uint256,uint256,address,string,uint256,bool)[])
func (_Gen *GenSession) SeeGainerMedicalInformationsName() ([]UploadMedicalrecords_gainergainer_upMedicalInformation, error) {
	return _Gen.Contract.SeeGainerMedicalInformationsName(&_Gen.CallOpts)
}

// SeeGainerMedicalInformationsName is a free data retrieval call binding the contract method 0xc506b536.
//
// Solidity: function SeeGainerMedicalInformationsName() view returns((string,string,uint256,uint256,address,string,uint256,bool)[])
func (_Gen *GenCallerSession) SeeGainerMedicalInformationsName() ([]UploadMedicalrecords_gainergainer_upMedicalInformation, error) {
	return _Gen.Contract.SeeGainerMedicalInformationsName(&_Gen.CallOpts)
}

// UserSerial is a free data retrieval call binding the contract method 0x9c30f930.
//
// Solidity: function UserSerial(address , uint256 ) view returns(bytes32)
func (_Gen *GenCaller) UserSerial(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "UserSerial", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UserSerial is a free data retrieval call binding the contract method 0x9c30f930.
//
// Solidity: function UserSerial(address , uint256 ) view returns(bytes32)
func (_Gen *GenSession) UserSerial(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Gen.Contract.UserSerial(&_Gen.CallOpts, arg0, arg1)
}

// UserSerial is a free data retrieval call binding the contract method 0x9c30f930.
//
// Solidity: function UserSerial(address , uint256 ) view returns(bytes32)
func (_Gen *GenCallerSession) UserSerial(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Gen.Contract.UserSerial(&_Gen.CallOpts, arg0, arg1)
}

// UserSerialName is a free data retrieval call binding the contract method 0x1a8b09ed.
//
// Solidity: function UserSerialName(address , string ) view returns(bytes32)
func (_Gen *GenCaller) UserSerialName(opts *bind.CallOpts, arg0 common.Address, arg1 string) ([32]byte, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "UserSerialName", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UserSerialName is a free data retrieval call binding the contract method 0x1a8b09ed.
//
// Solidity: function UserSerialName(address , string ) view returns(bytes32)
func (_Gen *GenSession) UserSerialName(arg0 common.Address, arg1 string) ([32]byte, error) {
	return _Gen.Contract.UserSerialName(&_Gen.CallOpts, arg0, arg1)
}

// UserSerialName is a free data retrieval call binding the contract method 0x1a8b09ed.
//
// Solidity: function UserSerialName(address , string ) view returns(bytes32)
func (_Gen *GenCallerSession) UserSerialName(arg0 common.Address, arg1 string) ([32]byte, error) {
	return _Gen.Contract.UserSerialName(&_Gen.CallOpts, arg0, arg1)
}

// ViewMedicalNames is a free data retrieval call binding the contract method 0x001e13e1.
//
// Solidity: function ViewMedicalNames() view returns(string[])
func (_Gen *GenCaller) ViewMedicalNames(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "ViewMedicalNames")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ViewMedicalNames is a free data retrieval call binding the contract method 0x001e13e1.
//
// Solidity: function ViewMedicalNames() view returns(string[])
func (_Gen *GenSession) ViewMedicalNames() ([]string, error) {
	return _Gen.Contract.ViewMedicalNames(&_Gen.CallOpts)
}

// ViewMedicalNames is a free data retrieval call binding the contract method 0x001e13e1.
//
// Solidity: function ViewMedicalNames() view returns(string[])
func (_Gen *GenCallerSession) ViewMedicalNames() ([]string, error) {
	return _Gen.Contract.ViewMedicalNames(&_Gen.CallOpts)
}

// ViewUserAllMedicalExaminationReport is a free data retrieval call binding the contract method 0x76889bff.
//
// Solidity: function ViewUserAllMedicalExaminationReport() view returns(string[])
func (_Gen *GenCaller) ViewUserAllMedicalExaminationReport(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "ViewUserAllMedicalExaminationReport")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ViewUserAllMedicalExaminationReport is a free data retrieval call binding the contract method 0x76889bff.
//
// Solidity: function ViewUserAllMedicalExaminationReport() view returns(string[])
func (_Gen *GenSession) ViewUserAllMedicalExaminationReport() ([]string, error) {
	return _Gen.Contract.ViewUserAllMedicalExaminationReport(&_Gen.CallOpts)
}

// ViewUserAllMedicalExaminationReport is a free data retrieval call binding the contract method 0x76889bff.
//
// Solidity: function ViewUserAllMedicalExaminationReport() view returns(string[])
func (_Gen *GenCallerSession) ViewUserAllMedicalExaminationReport() ([]string, error) {
	return _Gen.Contract.ViewUserAllMedicalExaminationReport(&_Gen.CallOpts)
}

// ViewUserAllMedicalinformation is a free data retrieval call binding the contract method 0xa511152e.
//
// Solidity: function ViewUserAllMedicalinformation() view returns(string[])
func (_Gen *GenCaller) ViewUserAllMedicalinformation(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "ViewUserAllMedicalinformation")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ViewUserAllMedicalinformation is a free data retrieval call binding the contract method 0xa511152e.
//
// Solidity: function ViewUserAllMedicalinformation() view returns(string[])
func (_Gen *GenSession) ViewUserAllMedicalinformation() ([]string, error) {
	return _Gen.Contract.ViewUserAllMedicalinformation(&_Gen.CallOpts)
}

// ViewUserAllMedicalinformation is a free data retrieval call binding the contract method 0xa511152e.
//
// Solidity: function ViewUserAllMedicalinformation() view returns(string[])
func (_Gen *GenCallerSession) ViewUserAllMedicalinformation() ([]string, error) {
	return _Gen.Contract.ViewUserAllMedicalinformation(&_Gen.CallOpts)
}

// ViewUserCertificate is a free data retrieval call binding the contract method 0x2a888eed.
//
// Solidity: function ViewUserCertificate(bytes32 _num) view returns((address,uint256,uint256,bytes32,string[],string[]))
func (_Gen *GenCaller) ViewUserCertificate(opts *bind.CallOpts, _num [32]byte) (UploadMedicalrecords_usercertificate, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "ViewUserCertificate", _num)

	if err != nil {
		return *new(UploadMedicalrecords_usercertificate), err
	}

	out0 := *abi.ConvertType(out[0], new(UploadMedicalrecords_usercertificate)).(*UploadMedicalrecords_usercertificate)

	return out0, err

}

// ViewUserCertificate is a free data retrieval call binding the contract method 0x2a888eed.
//
// Solidity: function ViewUserCertificate(bytes32 _num) view returns((address,uint256,uint256,bytes32,string[],string[]))
func (_Gen *GenSession) ViewUserCertificate(_num [32]byte) (UploadMedicalrecords_usercertificate, error) {
	return _Gen.Contract.ViewUserCertificate(&_Gen.CallOpts, _num)
}

// ViewUserCertificate is a free data retrieval call binding the contract method 0x2a888eed.
//
// Solidity: function ViewUserCertificate(bytes32 _num) view returns((address,uint256,uint256,bytes32,string[],string[]))
func (_Gen *GenCallerSession) ViewUserCertificate(_num [32]byte) (UploadMedicalrecords_usercertificate, error) {
	return _Gen.Contract.ViewUserCertificate(&_Gen.CallOpts, _num)
}

// ViewWareHouseLength is a free data retrieval call binding the contract method 0x49c4a44a.
//
// Solidity: function ViewWareHouseLength(string _medicalName) view returns(uint256)
func (_Gen *GenCaller) ViewWareHouseLength(opts *bind.CallOpts, _medicalName string) (*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "ViewWareHouseLength", _medicalName)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewWareHouseLength is a free data retrieval call binding the contract method 0x49c4a44a.
//
// Solidity: function ViewWareHouseLength(string _medicalName) view returns(uint256)
func (_Gen *GenSession) ViewWareHouseLength(_medicalName string) (*big.Int, error) {
	return _Gen.Contract.ViewWareHouseLength(&_Gen.CallOpts, _medicalName)
}

// ViewWareHouseLength is a free data retrieval call binding the contract method 0x49c4a44a.
//
// Solidity: function ViewWareHouseLength(string _medicalName) view returns(uint256)
func (_Gen *GenCallerSession) ViewWareHouseLength(_medicalName string) (*big.Int, error) {
	return _Gen.Contract.ViewWareHouseLength(&_Gen.CallOpts, _medicalName)
}

// ViewWareouse is a free data retrieval call binding the contract method 0xca09226a.
//
// Solidity: function ViewWareouse(string _medicalName, address _user) view returns((string[],string[]))
func (_Gen *GenCaller) ViewWareouse(opts *bind.CallOpts, _medicalName string, _user common.Address) (UploadMedicalrecords_gainerData, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "ViewWareouse", _medicalName, _user)

	if err != nil {
		return *new(UploadMedicalrecords_gainerData), err
	}

	out0 := *abi.ConvertType(out[0], new(UploadMedicalrecords_gainerData)).(*UploadMedicalrecords_gainerData)

	return out0, err

}

// ViewWareouse is a free data retrieval call binding the contract method 0xca09226a.
//
// Solidity: function ViewWareouse(string _medicalName, address _user) view returns((string[],string[]))
func (_Gen *GenSession) ViewWareouse(_medicalName string, _user common.Address) (UploadMedicalrecords_gainerData, error) {
	return _Gen.Contract.ViewWareouse(&_Gen.CallOpts, _medicalName, _user)
}

// ViewWareouse is a free data retrieval call binding the contract method 0xca09226a.
//
// Solidity: function ViewWareouse(string _medicalName, address _user) view returns((string[],string[]))
func (_Gen *GenCallerSession) ViewWareouse(_medicalName string, _user common.Address) (UploadMedicalrecords_gainerData, error) {
	return _Gen.Contract.ViewWareouse(&_Gen.CallOpts, _medicalName, _user)
}

// GainerSeeOwnMedicalInformationed is a free data retrieval call binding the contract method 0x2dee9e85.
//
// Solidity: function gainer_SeeOwnMedicalInformationed() view returns((string,uint256,uint256)[])
func (_Gen *GenCaller) GainerSeeOwnMedicalInformationed(opts *bind.CallOpts) ([]UploadMedicalrecords_gainergainer_Solicit, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "gainer_SeeOwnMedicalInformationed")

	if err != nil {
		return *new([]UploadMedicalrecords_gainergainer_Solicit), err
	}

	out0 := *abi.ConvertType(out[0], new([]UploadMedicalrecords_gainergainer_Solicit)).(*[]UploadMedicalrecords_gainergainer_Solicit)

	return out0, err

}

// GainerSeeOwnMedicalInformationed is a free data retrieval call binding the contract method 0x2dee9e85.
//
// Solidity: function gainer_SeeOwnMedicalInformationed() view returns((string,uint256,uint256)[])
func (_Gen *GenSession) GainerSeeOwnMedicalInformationed() ([]UploadMedicalrecords_gainergainer_Solicit, error) {
	return _Gen.Contract.GainerSeeOwnMedicalInformationed(&_Gen.CallOpts)
}

// GainerSeeOwnMedicalInformationed is a free data retrieval call binding the contract method 0x2dee9e85.
//
// Solidity: function gainer_SeeOwnMedicalInformationed() view returns((string,uint256,uint256)[])
func (_Gen *GenCallerSession) GainerSeeOwnMedicalInformationed() ([]UploadMedicalrecords_gainergainer_Solicit, error) {
	return _Gen.Contract.GainerSeeOwnMedicalInformationed(&_Gen.CallOpts)
}

// GainerSeeOwnMedicalInformationing is a free data retrieval call binding the contract method 0x41fb20e6.
//
// Solidity: function gainer_SeeOwnMedicalInformationing() view returns((string,uint256,uint256)[])
func (_Gen *GenCaller) GainerSeeOwnMedicalInformationing(opts *bind.CallOpts) ([]UploadMedicalrecords_gainergainer_Solicit, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "gainer_SeeOwnMedicalInformationing")

	if err != nil {
		return *new([]UploadMedicalrecords_gainergainer_Solicit), err
	}

	out0 := *abi.ConvertType(out[0], new([]UploadMedicalrecords_gainergainer_Solicit)).(*[]UploadMedicalrecords_gainergainer_Solicit)

	return out0, err

}

// GainerSeeOwnMedicalInformationing is a free data retrieval call binding the contract method 0x41fb20e6.
//
// Solidity: function gainer_SeeOwnMedicalInformationing() view returns((string,uint256,uint256)[])
func (_Gen *GenSession) GainerSeeOwnMedicalInformationing() ([]UploadMedicalrecords_gainergainer_Solicit, error) {
	return _Gen.Contract.GainerSeeOwnMedicalInformationing(&_Gen.CallOpts)
}

// GainerSeeOwnMedicalInformationing is a free data retrieval call binding the contract method 0x41fb20e6.
//
// Solidity: function gainer_SeeOwnMedicalInformationing() view returns((string,uint256,uint256)[])
func (_Gen *GenCallerSession) GainerSeeOwnMedicalInformationing() ([]UploadMedicalrecords_gainergainer_Solicit, error) {
	return _Gen.Contract.GainerSeeOwnMedicalInformationing(&_Gen.CallOpts)
}

// GainerSeeuserUploadMedical is a free data retrieval call binding the contract method 0x920b853e.
//
// Solidity: function gainer_SeeuserUploadMedical() view returns((address,address,string,string,bytes32,string)[])
func (_Gen *GenCaller) GainerSeeuserUploadMedical(opts *bind.CallOpts) ([]UploadMedicalrecords_useruser_MedicalCertificateState, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "gainer_SeeuserUploadMedical")

	if err != nil {
		return *new([]UploadMedicalrecords_useruser_MedicalCertificateState), err
	}

	out0 := *abi.ConvertType(out[0], new([]UploadMedicalrecords_useruser_MedicalCertificateState)).(*[]UploadMedicalrecords_useruser_MedicalCertificateState)

	return out0, err

}

// GainerSeeuserUploadMedical is a free data retrieval call binding the contract method 0x920b853e.
//
// Solidity: function gainer_SeeuserUploadMedical() view returns((address,address,string,string,bytes32,string)[])
func (_Gen *GenSession) GainerSeeuserUploadMedical() ([]UploadMedicalrecords_useruser_MedicalCertificateState, error) {
	return _Gen.Contract.GainerSeeuserUploadMedical(&_Gen.CallOpts)
}

// GainerSeeuserUploadMedical is a free data retrieval call binding the contract method 0x920b853e.
//
// Solidity: function gainer_SeeuserUploadMedical() view returns((address,address,string,string,bytes32,string)[])
func (_Gen *GenCallerSession) GainerSeeuserUploadMedical() ([]UploadMedicalrecords_useruser_MedicalCertificateState, error) {
	return _Gen.Contract.GainerSeeuserUploadMedical(&_Gen.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Gen *GenCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Gen *GenSession) LastUpdateTime() (*big.Int, error) {
	return _Gen.Contract.LastUpdateTime(&_Gen.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Gen *GenCallerSession) LastUpdateTime() (*big.Int, error) {
	return _Gen.Contract.LastUpdateTime(&_Gen.CallOpts)
}

// UserAmedicalInformation is a free data retrieval call binding the contract method 0xd112e898.
//
// Solidity: function userAmedicalInformation(address ) view returns(address user, address soliciters)
func (_Gen *GenCaller) UserAmedicalInformation(opts *bind.CallOpts, arg0 common.Address) (struct {
	User       common.Address
	Soliciters common.Address
}, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "userAmedicalInformation", arg0)

	outstruct := new(struct {
		User       common.Address
		Soliciters common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Soliciters = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// UserAmedicalInformation is a free data retrieval call binding the contract method 0xd112e898.
//
// Solidity: function userAmedicalInformation(address ) view returns(address user, address soliciters)
func (_Gen *GenSession) UserAmedicalInformation(arg0 common.Address) (struct {
	User       common.Address
	Soliciters common.Address
}, error) {
	return _Gen.Contract.UserAmedicalInformation(&_Gen.CallOpts, arg0)
}

// UserAmedicalInformation is a free data retrieval call binding the contract method 0xd112e898.
//
// Solidity: function userAmedicalInformation(address ) view returns(address user, address soliciters)
func (_Gen *GenCallerSession) UserAmedicalInformation(arg0 common.Address) (struct {
	User       common.Address
	Soliciters common.Address
}, error) {
	return _Gen.Contract.UserAmedicalInformation(&_Gen.CallOpts, arg0)
}

// UserWareHouse is a free data retrieval call binding the contract method 0x5cbfd4c0.
//
// Solidity: function userWareHouse(address , string , uint256 ) view returns(address)
func (_Gen *GenCaller) UserWareHouse(opts *bind.CallOpts, arg0 common.Address, arg1 string, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "userWareHouse", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserWareHouse is a free data retrieval call binding the contract method 0x5cbfd4c0.
//
// Solidity: function userWareHouse(address , string , uint256 ) view returns(address)
func (_Gen *GenSession) UserWareHouse(arg0 common.Address, arg1 string, arg2 *big.Int) (common.Address, error) {
	return _Gen.Contract.UserWareHouse(&_Gen.CallOpts, arg0, arg1, arg2)
}

// UserWareHouse is a free data retrieval call binding the contract method 0x5cbfd4c0.
//
// Solidity: function userWareHouse(address , string , uint256 ) view returns(address)
func (_Gen *GenCallerSession) UserWareHouse(arg0 common.Address, arg1 string, arg2 *big.Int) (common.Address, error) {
	return _Gen.Contract.UserWareHouse(&_Gen.CallOpts, arg0, arg1, arg2)
}

// UserMedicalExaminationReportstrucrName is a free data retrieval call binding the contract method 0xeae38791.
//
// Solidity: function user_MedicalExaminationReportstrucrName(address , uint256 ) view returns(string)
func (_Gen *GenCaller) UserMedicalExaminationReportstrucrName(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "user_MedicalExaminationReportstrucrName", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UserMedicalExaminationReportstrucrName is a free data retrieval call binding the contract method 0xeae38791.
//
// Solidity: function user_MedicalExaminationReportstrucrName(address , uint256 ) view returns(string)
func (_Gen *GenSession) UserMedicalExaminationReportstrucrName(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Gen.Contract.UserMedicalExaminationReportstrucrName(&_Gen.CallOpts, arg0, arg1)
}

// UserMedicalExaminationReportstrucrName is a free data retrieval call binding the contract method 0xeae38791.
//
// Solidity: function user_MedicalExaminationReportstrucrName(address , uint256 ) view returns(string)
func (_Gen *GenCallerSession) UserMedicalExaminationReportstrucrName(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Gen.Contract.UserMedicalExaminationReportstrucrName(&_Gen.CallOpts, arg0, arg1)
}

// UserMedicalinformationrName is a free data retrieval call binding the contract method 0x0d121e4e.
//
// Solidity: function user_MedicalinformationrName(address , uint256 ) view returns(string)
func (_Gen *GenCaller) UserMedicalinformationrName(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "user_MedicalinformationrName", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UserMedicalinformationrName is a free data retrieval call binding the contract method 0x0d121e4e.
//
// Solidity: function user_MedicalinformationrName(address , uint256 ) view returns(string)
func (_Gen *GenSession) UserMedicalinformationrName(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Gen.Contract.UserMedicalinformationrName(&_Gen.CallOpts, arg0, arg1)
}

// UserMedicalinformationrName is a free data retrieval call binding the contract method 0x0d121e4e.
//
// Solidity: function user_MedicalinformationrName(address , uint256 ) view returns(string)
func (_Gen *GenCallerSession) UserMedicalinformationrName(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Gen.Contract.UserMedicalinformationrName(&_Gen.CallOpts, arg0, arg1)
}

// UserNSeeCertificateState is a free data retrieval call binding the contract method 0x23a81a82.
//
// Solidity: function user_NSeeCertificateState() view returns((address,address,string,string,bytes32,string)[], uint256[])
func (_Gen *GenCaller) UserNSeeCertificateState(opts *bind.CallOpts) ([]UploadMedicalrecords_useruser_MedicalCertificateState, []*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "user_NSeeCertificateState")

	if err != nil {
		return *new([]UploadMedicalrecords_useruser_MedicalCertificateState), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]UploadMedicalrecords_useruser_MedicalCertificateState)).(*[]UploadMedicalrecords_useruser_MedicalCertificateState)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// UserNSeeCertificateState is a free data retrieval call binding the contract method 0x23a81a82.
//
// Solidity: function user_NSeeCertificateState() view returns((address,address,string,string,bytes32,string)[], uint256[])
func (_Gen *GenSession) UserNSeeCertificateState() ([]UploadMedicalrecords_useruser_MedicalCertificateState, []*big.Int, error) {
	return _Gen.Contract.UserNSeeCertificateState(&_Gen.CallOpts)
}

// UserNSeeCertificateState is a free data retrieval call binding the contract method 0x23a81a82.
//
// Solidity: function user_NSeeCertificateState() view returns((address,address,string,string,bytes32,string)[], uint256[])
func (_Gen *GenCallerSession) UserNSeeCertificateState() ([]UploadMedicalrecords_useruser_MedicalCertificateState, []*big.Int, error) {
	return _Gen.Contract.UserNSeeCertificateState(&_Gen.CallOpts)
}

// UserViewMedicalExaminationReport is a free data retrieval call binding the contract method 0x603a78c2.
//
// Solidity: function user_ViewMedicalExaminationReport(string name) view returns(string)
func (_Gen *GenCaller) UserViewMedicalExaminationReport(opts *bind.CallOpts, name string) (string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "user_ViewMedicalExaminationReport", name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UserViewMedicalExaminationReport is a free data retrieval call binding the contract method 0x603a78c2.
//
// Solidity: function user_ViewMedicalExaminationReport(string name) view returns(string)
func (_Gen *GenSession) UserViewMedicalExaminationReport(name string) (string, error) {
	return _Gen.Contract.UserViewMedicalExaminationReport(&_Gen.CallOpts, name)
}

// UserViewMedicalExaminationReport is a free data retrieval call binding the contract method 0x603a78c2.
//
// Solidity: function user_ViewMedicalExaminationReport(string name) view returns(string)
func (_Gen *GenCallerSession) UserViewMedicalExaminationReport(name string) (string, error) {
	return _Gen.Contract.UserViewMedicalExaminationReport(&_Gen.CallOpts, name)
}

// UserViewMedicalInformation is a free data retrieval call binding the contract method 0xd59409fe.
//
// Solidity: function user_ViewMedicalInformation(string name) view returns(string)
func (_Gen *GenCaller) UserViewMedicalInformation(opts *bind.CallOpts, name string) (string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "user_ViewMedicalInformation", name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UserViewMedicalInformation is a free data retrieval call binding the contract method 0xd59409fe.
//
// Solidity: function user_ViewMedicalInformation(string name) view returns(string)
func (_Gen *GenSession) UserViewMedicalInformation(name string) (string, error) {
	return _Gen.Contract.UserViewMedicalInformation(&_Gen.CallOpts, name)
}

// UserViewMedicalInformation is a free data retrieval call binding the contract method 0xd59409fe.
//
// Solidity: function user_ViewMedicalInformation(string name) view returns(string)
func (_Gen *GenCallerSession) UserViewMedicalInformation(name string) (string, error) {
	return _Gen.Contract.UserViewMedicalInformation(&_Gen.CallOpts, name)
}

// UserYSeeCertificateState is a free data retrieval call binding the contract method 0x26fa06d3.
//
// Solidity: function user_YSeeCertificateState() view returns((address,address,string,string,bytes32,string)[])
func (_Gen *GenCaller) UserYSeeCertificateState(opts *bind.CallOpts) ([]UploadMedicalrecords_useruser_MedicalCertificateState, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "user_YSeeCertificateState")

	if err != nil {
		return *new([]UploadMedicalrecords_useruser_MedicalCertificateState), err
	}

	out0 := *abi.ConvertType(out[0], new([]UploadMedicalrecords_useruser_MedicalCertificateState)).(*[]UploadMedicalrecords_useruser_MedicalCertificateState)

	return out0, err

}

// UserYSeeCertificateState is a free data retrieval call binding the contract method 0x26fa06d3.
//
// Solidity: function user_YSeeCertificateState() view returns((address,address,string,string,bytes32,string)[])
func (_Gen *GenSession) UserYSeeCertificateState() ([]UploadMedicalrecords_useruser_MedicalCertificateState, error) {
	return _Gen.Contract.UserYSeeCertificateState(&_Gen.CallOpts)
}

// UserYSeeCertificateState is a free data retrieval call binding the contract method 0x26fa06d3.
//
// Solidity: function user_YSeeCertificateState() view returns((address,address,string,string,bytes32,string)[])
func (_Gen *GenCallerSession) UserYSeeCertificateState() ([]UploadMedicalrecords_useruser_MedicalCertificateState, error) {
	return _Gen.Contract.UserYSeeCertificateState(&_Gen.CallOpts)
}

// ViewMedicalExaminationReportCount is a free data retrieval call binding the contract method 0x03c8505d.
//
// Solidity: function viewMedicalExaminationReportCount() view returns(uint256)
func (_Gen *GenCaller) ViewMedicalExaminationReportCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "viewMedicalExaminationReportCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewMedicalExaminationReportCount is a free data retrieval call binding the contract method 0x03c8505d.
//
// Solidity: function viewMedicalExaminationReportCount() view returns(uint256)
func (_Gen *GenSession) ViewMedicalExaminationReportCount() (*big.Int, error) {
	return _Gen.Contract.ViewMedicalExaminationReportCount(&_Gen.CallOpts)
}

// ViewMedicalExaminationReportCount is a free data retrieval call binding the contract method 0x03c8505d.
//
// Solidity: function viewMedicalExaminationReportCount() view returns(uint256)
func (_Gen *GenCallerSession) ViewMedicalExaminationReportCount() (*big.Int, error) {
	return _Gen.Contract.ViewMedicalExaminationReportCount(&_Gen.CallOpts)
}

// ViewMedicalinformationCount is a free data retrieval call binding the contract method 0x9d012e1c.
//
// Solidity: function viewMedicalinformationCount() view returns(uint256)
func (_Gen *GenCaller) ViewMedicalinformationCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "viewMedicalinformationCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewMedicalinformationCount is a free data retrieval call binding the contract method 0x9d012e1c.
//
// Solidity: function viewMedicalinformationCount() view returns(uint256)
func (_Gen *GenSession) ViewMedicalinformationCount() (*big.Int, error) {
	return _Gen.Contract.ViewMedicalinformationCount(&_Gen.CallOpts)
}

// ViewMedicalinformationCount is a free data retrieval call binding the contract method 0x9d012e1c.
//
// Solidity: function viewMedicalinformationCount() view returns(uint256)
func (_Gen *GenCallerSession) ViewMedicalinformationCount() (*big.Int, error) {
	return _Gen.Contract.ViewMedicalinformationCount(&_Gen.CallOpts)
}

// SetCertificate is a paid mutator transaction binding the contract method 0x03ab6eea.
//
// Solidity: function SetCertificate(string[] _m1, string[] _m2, string name) returns()
func (_Gen *GenTransactor) SetCertificate(opts *bind.TransactOpts, _m1 []string, _m2 []string, name string) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "SetCertificate", _m1, _m2, name)
}

// SetCertificate is a paid mutator transaction binding the contract method 0x03ab6eea.
//
// Solidity: function SetCertificate(string[] _m1, string[] _m2, string name) returns()
func (_Gen *GenSession) SetCertificate(_m1 []string, _m2 []string, name string) (*types.Transaction, error) {
	return _Gen.Contract.SetCertificate(&_Gen.TransactOpts, _m1, _m2, name)
}

// SetCertificate is a paid mutator transaction binding the contract method 0x03ab6eea.
//
// Solidity: function SetCertificate(string[] _m1, string[] _m2, string name) returns()
func (_Gen *GenTransactorSession) SetCertificate(_m1 []string, _m2 []string, name string) (*types.Transaction, error) {
	return _Gen.Contract.SetCertificate(&_Gen.TransactOpts, _m1, _m2, name)
}

// GainerAddMedicalInformation is a paid mutator transaction binding the contract method 0xd2e15d15.
//
// Solidity: function gainer_AddMedicalInformation(uint256 min_, uint256 max_, uint256 account_, string Department, string MedicalName_, string _Medicalrecordrequirements, string _Requirementdescription) returns()
func (_Gen *GenTransactor) GainerAddMedicalInformation(opts *bind.TransactOpts, min_ *big.Int, max_ *big.Int, account_ *big.Int, Department string, MedicalName_ string, _Medicalrecordrequirements string, _Requirementdescription string) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "gainer_AddMedicalInformation", min_, max_, account_, Department, MedicalName_, _Medicalrecordrequirements, _Requirementdescription)
}

// GainerAddMedicalInformation is a paid mutator transaction binding the contract method 0xd2e15d15.
//
// Solidity: function gainer_AddMedicalInformation(uint256 min_, uint256 max_, uint256 account_, string Department, string MedicalName_, string _Medicalrecordrequirements, string _Requirementdescription) returns()
func (_Gen *GenSession) GainerAddMedicalInformation(min_ *big.Int, max_ *big.Int, account_ *big.Int, Department string, MedicalName_ string, _Medicalrecordrequirements string, _Requirementdescription string) (*types.Transaction, error) {
	return _Gen.Contract.GainerAddMedicalInformation(&_Gen.TransactOpts, min_, max_, account_, Department, MedicalName_, _Medicalrecordrequirements, _Requirementdescription)
}

// GainerAddMedicalInformation is a paid mutator transaction binding the contract method 0xd2e15d15.
//
// Solidity: function gainer_AddMedicalInformation(uint256 min_, uint256 max_, uint256 account_, string Department, string MedicalName_, string _Medicalrecordrequirements, string _Requirementdescription) returns()
func (_Gen *GenTransactorSession) GainerAddMedicalInformation(min_ *big.Int, max_ *big.Int, account_ *big.Int, Department string, MedicalName_ string, _Medicalrecordrequirements string, _Requirementdescription string) (*types.Transaction, error) {
	return _Gen.Contract.GainerAddMedicalInformation(&_Gen.TransactOpts, min_, max_, account_, Department, MedicalName_, _Medicalrecordrequirements, _Requirementdescription)
}

// GainerSetDoctor is a paid mutator transaction binding the contract method 0x509ca131.
//
// Solidity: function gainer_setDoctor(address soliciter, string HospitalName) returns()
func (_Gen *GenTransactor) GainerSetDoctor(opts *bind.TransactOpts, soliciter common.Address, HospitalName string) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "gainer_setDoctor", soliciter, HospitalName)
}

// GainerSetDoctor is a paid mutator transaction binding the contract method 0x509ca131.
//
// Solidity: function gainer_setDoctor(address soliciter, string HospitalName) returns()
func (_Gen *GenSession) GainerSetDoctor(soliciter common.Address, HospitalName string) (*types.Transaction, error) {
	return _Gen.Contract.GainerSetDoctor(&_Gen.TransactOpts, soliciter, HospitalName)
}

// GainerSetDoctor is a paid mutator transaction binding the contract method 0x509ca131.
//
// Solidity: function gainer_setDoctor(address soliciter, string HospitalName) returns()
func (_Gen *GenTransactorSession) GainerSetDoctor(soliciter common.Address, HospitalName string) (*types.Transaction, error) {
	return _Gen.Contract.GainerSetDoctor(&_Gen.TransactOpts, soliciter, HospitalName)
}

// GainerWhether is a paid mutator transaction binding the contract method 0xc0d17054.
//
// Solidity: function gainer_whether(bytes32 CertificatePath, string MediacalName, bool whether, address user, uint256 ercnum_) returns()
func (_Gen *GenTransactor) GainerWhether(opts *bind.TransactOpts, CertificatePath [32]byte, MediacalName string, whether bool, user common.Address, ercnum_ *big.Int) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "gainer_whether", CertificatePath, MediacalName, whether, user, ercnum_)
}

// GainerWhether is a paid mutator transaction binding the contract method 0xc0d17054.
//
// Solidity: function gainer_whether(bytes32 CertificatePath, string MediacalName, bool whether, address user, uint256 ercnum_) returns()
func (_Gen *GenSession) GainerWhether(CertificatePath [32]byte, MediacalName string, whether bool, user common.Address, ercnum_ *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.GainerWhether(&_Gen.TransactOpts, CertificatePath, MediacalName, whether, user, ercnum_)
}

// GainerWhether is a paid mutator transaction binding the contract method 0xc0d17054.
//
// Solidity: function gainer_whether(bytes32 CertificatePath, string MediacalName, bool whether, address user, uint256 ercnum_) returns()
func (_Gen *GenTransactorSession) GainerWhether(CertificatePath [32]byte, MediacalName string, whether bool, user common.Address, ercnum_ *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.GainerWhether(&_Gen.TransactOpts, CertificatePath, MediacalName, whether, user, ercnum_)
}

// UserAddMedicalInformation is a paid mutator transaction binding the contract method 0x2892b8c9.
//
// Solidity: function user_AddMedicalInformation(bytes32 Proute, address _soliciter, string MedicalName_, string Department) returns()
func (_Gen *GenTransactor) UserAddMedicalInformation(opts *bind.TransactOpts, Proute [32]byte, _soliciter common.Address, MedicalName_ string, Department string) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "user_AddMedicalInformation", Proute, _soliciter, MedicalName_, Department)
}

// UserAddMedicalInformation is a paid mutator transaction binding the contract method 0x2892b8c9.
//
// Solidity: function user_AddMedicalInformation(bytes32 Proute, address _soliciter, string MedicalName_, string Department) returns()
func (_Gen *GenSession) UserAddMedicalInformation(Proute [32]byte, _soliciter common.Address, MedicalName_ string, Department string) (*types.Transaction, error) {
	return _Gen.Contract.UserAddMedicalInformation(&_Gen.TransactOpts, Proute, _soliciter, MedicalName_, Department)
}

// UserAddMedicalInformation is a paid mutator transaction binding the contract method 0x2892b8c9.
//
// Solidity: function user_AddMedicalInformation(bytes32 Proute, address _soliciter, string MedicalName_, string Department) returns()
func (_Gen *GenTransactorSession) UserAddMedicalInformation(Proute [32]byte, _soliciter common.Address, MedicalName_ string, Department string) (*types.Transaction, error) {
	return _Gen.Contract.UserAddMedicalInformation(&_Gen.TransactOpts, Proute, _soliciter, MedicalName_, Department)
}

// UserAdduser is a paid mutator transaction binding the contract method 0x1dd9ec67.
//
// Solidity: function user_Adduser(address _user) returns()
func (_Gen *GenTransactor) UserAdduser(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "user_Adduser", _user)
}

// UserAdduser is a paid mutator transaction binding the contract method 0x1dd9ec67.
//
// Solidity: function user_Adduser(address _user) returns()
func (_Gen *GenSession) UserAdduser(_user common.Address) (*types.Transaction, error) {
	return _Gen.Contract.UserAdduser(&_Gen.TransactOpts, _user)
}

// UserAdduser is a paid mutator transaction binding the contract method 0x1dd9ec67.
//
// Solidity: function user_Adduser(address _user) returns()
func (_Gen *GenTransactorSession) UserAdduser(_user common.Address) (*types.Transaction, error) {
	return _Gen.Contract.UserAdduser(&_Gen.TransactOpts, _user)
}

// UserUPMedicalExaminationReport is a paid mutator transaction binding the contract method 0x197e456a.
//
// Solidity: function user_UPMedicalExaminationReport(string name, string url) returns()
func (_Gen *GenTransactor) UserUPMedicalExaminationReport(opts *bind.TransactOpts, name string, url string) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "user_UPMedicalExaminationReport", name, url)
}

// UserUPMedicalExaminationReport is a paid mutator transaction binding the contract method 0x197e456a.
//
// Solidity: function user_UPMedicalExaminationReport(string name, string url) returns()
func (_Gen *GenSession) UserUPMedicalExaminationReport(name string, url string) (*types.Transaction, error) {
	return _Gen.Contract.UserUPMedicalExaminationReport(&_Gen.TransactOpts, name, url)
}

// UserUPMedicalExaminationReport is a paid mutator transaction binding the contract method 0x197e456a.
//
// Solidity: function user_UPMedicalExaminationReport(string name, string url) returns()
func (_Gen *GenTransactorSession) UserUPMedicalExaminationReport(name string, url string) (*types.Transaction, error) {
	return _Gen.Contract.UserUPMedicalExaminationReport(&_Gen.TransactOpts, name, url)
}

// UserUPMedicalinformation is a paid mutator transaction binding the contract method 0x868e0bd4.
//
// Solidity: function user_UPMedicalinformation(string name, string url) returns()
func (_Gen *GenTransactor) UserUPMedicalinformation(opts *bind.TransactOpts, name string, url string) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "user_UPMedicalinformation", name, url)
}

// UserUPMedicalinformation is a paid mutator transaction binding the contract method 0x868e0bd4.
//
// Solidity: function user_UPMedicalinformation(string name, string url) returns()
func (_Gen *GenSession) UserUPMedicalinformation(name string, url string) (*types.Transaction, error) {
	return _Gen.Contract.UserUPMedicalinformation(&_Gen.TransactOpts, name, url)
}

// UserUPMedicalinformation is a paid mutator transaction binding the contract method 0x868e0bd4.
//
// Solidity: function user_UPMedicalinformation(string name, string url) returns()
func (_Gen *GenTransactorSession) UserUPMedicalinformation(name string, url string) (*types.Transaction, error) {
	return _Gen.Contract.UserUPMedicalinformation(&_Gen.TransactOpts, name, url)
}
