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

// UploadMedicalrecords_gainergainer_Solicit is an auto generated low-level Go binding around an user-defined struct.
type UploadMedicalrecords_gainergainer_Solicit struct {
	MedicalName string
	Min         *big.Int
	Max         *big.Int
}

// UploadMedicalrecords_gainergainer_upMedicalInformation is an auto generated low-level Go binding around an user-defined struct.
type UploadMedicalrecords_gainergainer_upMedicalInformation struct {
	MedicalName  string
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
	State        bool
	Soliciter    common.Address
	HospitalName string
	MedicalName  string
	Certificate  [32]byte
	Erum         *big.Int
}

// GenMetaData contains all meta data concerning the Gen contract.
var GenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"A11GiveETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ercaddress_\",\"type\":\"address\"}],\"name\":\"A11Smile_setErc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"Startuse\",\"type\":\"address\"}],\"name\":\"A11SmileGiveETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Addtokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"EthgetAs\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"AddEth\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"RedEth\",\"type\":\"address\"}],\"name\":\"EthGetAs\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"account_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"MedicalName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_Medicalrecordrequirements\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_Requirementdescription\",\"type\":\"string\"}],\"name\":\"gainer_AddMedicalInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"HospitalName\",\"type\":\"string\"}],\"name\":\"gainer_setDoctor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity1\",\"type\":\"uint256\"}],\"name\":\"gainer_transfer_AS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"CertificatePath\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"MediacalName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"whether\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ercnum_\",\"type\":\"uint256\"}],\"name\":\"gainer_whether\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"route\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"}],\"name\":\"Uploadmedicaldata\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"route\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"}],\"name\":\"gainerUploadmedicaldata\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"route\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"reward\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_m1\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_m2\",\"type\":\"string[]\"}],\"name\":\"SetCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"Proute\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_soliciter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"MedicalName_\",\"type\":\"string\"}],\"name\":\"user_AddMedicalInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"user_Adduser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"user_UPMedicalExaminationReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"user_UPMedicalinformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"A11Smile1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Adoctor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"HospitalName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"examineTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"}],\"name\":\"gainer_isDoctor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gainer_SeeOwnMedicalInformationed\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structUploadMedicalrecords_gainer.gainer_Solicit[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gainer_SeeOwnMedicalInformationing\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"internalType\":\"structUploadMedicalrecords_gainer.gainer_Solicit[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gainer_SeeuserUploadMedical\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"HospitalName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"certificate\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"erum\",\"type\":\"uint256\"}],\"internalType\":\"structUploadMedicalrecords_user.user_MedicalCertificateState[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Medicalinformation_\",\"type\":\"string\"}],\"name\":\"QueryMedicalinformation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ExaminationReportName_\",\"type\":\"string\"}],\"name\":\"QueryPhysicalExaminationReport\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"QueryUserSerialLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"HospitalAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"}],\"name\":\"SeeGainerMedicalInformations\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"Medicalrecordrequirements\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Requirementdescription\",\"type\":\"string\"}],\"internalType\":\"structUploadMedicalrecords_gainer.gainer_upMedicalInformation1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SeeGainerMedicalInformationsName\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"HospitalName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"account\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"}],\"internalType\":\"structUploadMedicalrecords_gainer.gainer_upMedicalInformation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"user_IsUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"user_MedicalExaminationReportstrucrName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"user_MedicalinformationrName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"user_NSeeCertificateState\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"HospitalName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"certificate\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"erum\",\"type\":\"uint256\"}],\"internalType\":\"structUploadMedicalrecords_user.user_MedicalCertificateState[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"user_ViewMedicalExaminationReport\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"user_ViewMedicalInformation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"user_YSeeCertificateState\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"soliciter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"HospitalName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"MedicalName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"certificate\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"erum\",\"type\":\"uint256\"}],\"internalType\":\"structUploadMedicalrecords_user.user_MedicalCertificateState[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userAmedicalInformation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"UserSerial\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewMedicalExaminationReportCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewMedicalinformationCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ViewUserAllMedicalExaminationReport\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ViewUserAllMedicalinformation\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_num\",\"type\":\"bytes32\"}],\"name\":\"ViewUserCertificate\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"serial\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"m1\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"m2\",\"type\":\"string[]\"}],\"internalType\":\"structUploadMedicalrecords_user.certificate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// A11Smile1 is a free data retrieval call binding the contract method 0x8f403c7d.
//
// Solidity: function A11Smile1() view returns(address)
func (_Gen *GenCaller) A11Smile1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "A11Smile1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// A11Smile1 is a free data retrieval call binding the contract method 0x8f403c7d.
//
// Solidity: function A11Smile1() view returns(address)
func (_Gen *GenSession) A11Smile1() (common.Address, error) {
	return _Gen.Contract.A11Smile1(&_Gen.CallOpts)
}

// A11Smile1 is a free data retrieval call binding the contract method 0x8f403c7d.
//
// Solidity: function A11Smile1() view returns(address)
func (_Gen *GenCallerSession) A11Smile1() (common.Address, error) {
	return _Gen.Contract.A11Smile1(&_Gen.CallOpts)
}

// Adoctor is a free data retrieval call binding the contract method 0x3b5fed16.
//
// Solidity: function Adoctor(address ) view returns(address soliciter, string HospitalName, bool exist)
func (_Gen *GenCaller) Adoctor(opts *bind.CallOpts, arg0 common.Address) (struct {
	Soliciter    common.Address
	HospitalName string
	Exist        bool
}, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "Adoctor", arg0)

	outstruct := new(struct {
		Soliciter    common.Address
		HospitalName string
		Exist        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Soliciter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.HospitalName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Exist = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Adoctor is a free data retrieval call binding the contract method 0x3b5fed16.
//
// Solidity: function Adoctor(address ) view returns(address soliciter, string HospitalName, bool exist)
func (_Gen *GenSession) Adoctor(arg0 common.Address) (struct {
	Soliciter    common.Address
	HospitalName string
	Exist        bool
}, error) {
	return _Gen.Contract.Adoctor(&_Gen.CallOpts, arg0)
}

// Adoctor is a free data retrieval call binding the contract method 0x3b5fed16.
//
// Solidity: function Adoctor(address ) view returns(address soliciter, string HospitalName, bool exist)
func (_Gen *GenCallerSession) Adoctor(arg0 common.Address) (struct {
	Soliciter    common.Address
	HospitalName string
	Exist        bool
}, error) {
	return _Gen.Contract.Adoctor(&_Gen.CallOpts, arg0)
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
// Solidity: function SeeGainerMedicalInformationsName() view returns((string,uint256,uint256,address,string,uint256,bool)[])
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
// Solidity: function SeeGainerMedicalInformationsName() view returns((string,uint256,uint256,address,string,uint256,bool)[])
func (_Gen *GenSession) SeeGainerMedicalInformationsName() ([]UploadMedicalrecords_gainergainer_upMedicalInformation, error) {
	return _Gen.Contract.SeeGainerMedicalInformationsName(&_Gen.CallOpts)
}

// SeeGainerMedicalInformationsName is a free data retrieval call binding the contract method 0xc506b536.
//
// Solidity: function SeeGainerMedicalInformationsName() view returns((string,uint256,uint256,address,string,uint256,bool)[])
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

// ViewUserCertificate is a free data retrieval call binding the contract method 0x8ff8f77e.
//
// Solidity: function ViewUserCertificate(address _user, bytes32 _num) view returns((address,uint256,uint256,bytes32,string[],string[]))
func (_Gen *GenCaller) ViewUserCertificate(opts *bind.CallOpts, _user common.Address, _num [32]byte) (UploadMedicalrecords_usercertificate, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "ViewUserCertificate", _user, _num)

	if err != nil {
		return *new(UploadMedicalrecords_usercertificate), err
	}

	out0 := *abi.ConvertType(out[0], new(UploadMedicalrecords_usercertificate)).(*UploadMedicalrecords_usercertificate)

	return out0, err

}

// ViewUserCertificate is a free data retrieval call binding the contract method 0x8ff8f77e.
//
// Solidity: function ViewUserCertificate(address _user, bytes32 _num) view returns((address,uint256,uint256,bytes32,string[],string[]))
func (_Gen *GenSession) ViewUserCertificate(_user common.Address, _num [32]byte) (UploadMedicalrecords_usercertificate, error) {
	return _Gen.Contract.ViewUserCertificate(&_Gen.CallOpts, _user, _num)
}

// ViewUserCertificate is a free data retrieval call binding the contract method 0x8ff8f77e.
//
// Solidity: function ViewUserCertificate(address _user, bytes32 _num) view returns((address,uint256,uint256,bytes32,string[],string[]))
func (_Gen *GenCallerSession) ViewUserCertificate(_user common.Address, _num [32]byte) (UploadMedicalrecords_usercertificate, error) {
	return _Gen.Contract.ViewUserCertificate(&_Gen.CallOpts, _user, _num)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Gen *GenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Gen *GenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Gen.Contract.Allowance(&_Gen.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Gen *GenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Gen.Contract.Allowance(&_Gen.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gen *GenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gen *GenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Gen.Contract.BalanceOf(&_Gen.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gen *GenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Gen.Contract.BalanceOf(&_Gen.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Gen *GenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Gen *GenSession) Decimals() (uint8, error) {
	return _Gen.Contract.Decimals(&_Gen.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Gen *GenCallerSession) Decimals() (uint8, error) {
	return _Gen.Contract.Decimals(&_Gen.CallOpts)
}

// ExamineTime is a free data retrieval call binding the contract method 0x495c184e.
//
// Solidity: function examineTime() view returns(uint256)
func (_Gen *GenCaller) ExamineTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "examineTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExamineTime is a free data retrieval call binding the contract method 0x495c184e.
//
// Solidity: function examineTime() view returns(uint256)
func (_Gen *GenSession) ExamineTime() (*big.Int, error) {
	return _Gen.Contract.ExamineTime(&_Gen.CallOpts)
}

// ExamineTime is a free data retrieval call binding the contract method 0x495c184e.
//
// Solidity: function examineTime() view returns(uint256)
func (_Gen *GenCallerSession) ExamineTime() (*big.Int, error) {
	return _Gen.Contract.ExamineTime(&_Gen.CallOpts)
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
// Solidity: function gainer_SeeuserUploadMedical() view returns((address,bool,address,string,string,bytes32,uint256)[])
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
// Solidity: function gainer_SeeuserUploadMedical() view returns((address,bool,address,string,string,bytes32,uint256)[])
func (_Gen *GenSession) GainerSeeuserUploadMedical() ([]UploadMedicalrecords_useruser_MedicalCertificateState, error) {
	return _Gen.Contract.GainerSeeuserUploadMedical(&_Gen.CallOpts)
}

// GainerSeeuserUploadMedical is a free data retrieval call binding the contract method 0x920b853e.
//
// Solidity: function gainer_SeeuserUploadMedical() view returns((address,bool,address,string,string,bytes32,uint256)[])
func (_Gen *GenCallerSession) GainerSeeuserUploadMedical() ([]UploadMedicalrecords_useruser_MedicalCertificateState, error) {
	return _Gen.Contract.GainerSeeuserUploadMedical(&_Gen.CallOpts)
}

// GainerIsDoctor is a free data retrieval call binding the contract method 0x93172bf7.
//
// Solidity: function gainer_isDoctor(address soliciter) view returns(bool)
func (_Gen *GenCaller) GainerIsDoctor(opts *bind.CallOpts, soliciter common.Address) (bool, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "gainer_isDoctor", soliciter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GainerIsDoctor is a free data retrieval call binding the contract method 0x93172bf7.
//
// Solidity: function gainer_isDoctor(address soliciter) view returns(bool)
func (_Gen *GenSession) GainerIsDoctor(soliciter common.Address) (bool, error) {
	return _Gen.Contract.GainerIsDoctor(&_Gen.CallOpts, soliciter)
}

// GainerIsDoctor is a free data retrieval call binding the contract method 0x93172bf7.
//
// Solidity: function gainer_isDoctor(address soliciter) view returns(bool)
func (_Gen *GenCallerSession) GainerIsDoctor(soliciter common.Address) (bool, error) {
	return _Gen.Contract.GainerIsDoctor(&_Gen.CallOpts, soliciter)
}

// GetUserAS is a free data retrieval call binding the contract method 0xda0b61f7.
//
// Solidity: function getUserAS() view returns(uint256)
func (_Gen *GenCaller) GetUserAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "getUserAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAS is a free data retrieval call binding the contract method 0xda0b61f7.
//
// Solidity: function getUserAS() view returns(uint256)
func (_Gen *GenSession) GetUserAS() (*big.Int, error) {
	return _Gen.Contract.GetUserAS(&_Gen.CallOpts)
}

// GetUserAS is a free data retrieval call binding the contract method 0xda0b61f7.
//
// Solidity: function getUserAS() view returns(uint256)
func (_Gen *GenCallerSession) GetUserAS() (*big.Int, error) {
	return _Gen.Contract.GetUserAS(&_Gen.CallOpts)
}

// GetUserETH is a free data retrieval call binding the contract method 0xbc7edaf2.
//
// Solidity: function getUserETH() view returns(uint256)
func (_Gen *GenCaller) GetUserETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "getUserETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserETH is a free data retrieval call binding the contract method 0xbc7edaf2.
//
// Solidity: function getUserETH() view returns(uint256)
func (_Gen *GenSession) GetUserETH() (*big.Int, error) {
	return _Gen.Contract.GetUserETH(&_Gen.CallOpts)
}

// GetUserETH is a free data retrieval call binding the contract method 0xbc7edaf2.
//
// Solidity: function getUserETH() view returns(uint256)
func (_Gen *GenCallerSession) GetUserETH() (*big.Int, error) {
	return _Gen.Contract.GetUserETH(&_Gen.CallOpts)
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

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gen *GenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gen *GenSession) Name() (string, error) {
	return _Gen.Contract.Name(&_Gen.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gen *GenCallerSession) Name() (string, error) {
	return _Gen.Contract.Name(&_Gen.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gen *GenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gen *GenSession) Symbol() (string, error) {
	return _Gen.Contract.Symbol(&_Gen.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gen *GenCallerSession) Symbol() (string, error) {
	return _Gen.Contract.Symbol(&_Gen.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gen *GenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gen *GenSession) TotalSupply() (*big.Int, error) {
	return _Gen.Contract.TotalSupply(&_Gen.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gen *GenCallerSession) TotalSupply() (*big.Int, error) {
	return _Gen.Contract.TotalSupply(&_Gen.CallOpts)
}

// UserAmedicalInformation is a free data retrieval call binding the contract method 0xd112e898.
//
// Solidity: function userAmedicalInformation(address ) view returns(address user)
func (_Gen *GenCaller) UserAmedicalInformation(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "userAmedicalInformation", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserAmedicalInformation is a free data retrieval call binding the contract method 0xd112e898.
//
// Solidity: function userAmedicalInformation(address ) view returns(address user)
func (_Gen *GenSession) UserAmedicalInformation(arg0 common.Address) (common.Address, error) {
	return _Gen.Contract.UserAmedicalInformation(&_Gen.CallOpts, arg0)
}

// UserAmedicalInformation is a free data retrieval call binding the contract method 0xd112e898.
//
// Solidity: function userAmedicalInformation(address ) view returns(address user)
func (_Gen *GenCallerSession) UserAmedicalInformation(arg0 common.Address) (common.Address, error) {
	return _Gen.Contract.UserAmedicalInformation(&_Gen.CallOpts, arg0)
}

// UserIsUser is a free data retrieval call binding the contract method 0xd5f03d3c.
//
// Solidity: function user_IsUser(address _user) view returns(bool)
func (_Gen *GenCaller) UserIsUser(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "user_IsUser", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserIsUser is a free data retrieval call binding the contract method 0xd5f03d3c.
//
// Solidity: function user_IsUser(address _user) view returns(bool)
func (_Gen *GenSession) UserIsUser(_user common.Address) (bool, error) {
	return _Gen.Contract.UserIsUser(&_Gen.CallOpts, _user)
}

// UserIsUser is a free data retrieval call binding the contract method 0xd5f03d3c.
//
// Solidity: function user_IsUser(address _user) view returns(bool)
func (_Gen *GenCallerSession) UserIsUser(_user common.Address) (bool, error) {
	return _Gen.Contract.UserIsUser(&_Gen.CallOpts, _user)
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
// Solidity: function user_NSeeCertificateState() view returns((address,bool,address,string,string,bytes32,uint256)[])
func (_Gen *GenCaller) UserNSeeCertificateState(opts *bind.CallOpts) ([]UploadMedicalrecords_useruser_MedicalCertificateState, error) {
	var out []interface{}
	err := _Gen.contract.Call(opts, &out, "user_NSeeCertificateState")

	if err != nil {
		return *new([]UploadMedicalrecords_useruser_MedicalCertificateState), err
	}

	out0 := *abi.ConvertType(out[0], new([]UploadMedicalrecords_useruser_MedicalCertificateState)).(*[]UploadMedicalrecords_useruser_MedicalCertificateState)

	return out0, err

}

// UserNSeeCertificateState is a free data retrieval call binding the contract method 0x23a81a82.
//
// Solidity: function user_NSeeCertificateState() view returns((address,bool,address,string,string,bytes32,uint256)[])
func (_Gen *GenSession) UserNSeeCertificateState() ([]UploadMedicalrecords_useruser_MedicalCertificateState, error) {
	return _Gen.Contract.UserNSeeCertificateState(&_Gen.CallOpts)
}

// UserNSeeCertificateState is a free data retrieval call binding the contract method 0x23a81a82.
//
// Solidity: function user_NSeeCertificateState() view returns((address,bool,address,string,string,bytes32,uint256)[])
func (_Gen *GenCallerSession) UserNSeeCertificateState() ([]UploadMedicalrecords_useruser_MedicalCertificateState, error) {
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
// Solidity: function user_YSeeCertificateState() view returns((address,bool,address,string,string,bytes32,uint256)[])
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
// Solidity: function user_YSeeCertificateState() view returns((address,bool,address,string,string,bytes32,uint256)[])
func (_Gen *GenSession) UserYSeeCertificateState() ([]UploadMedicalrecords_useruser_MedicalCertificateState, error) {
	return _Gen.Contract.UserYSeeCertificateState(&_Gen.CallOpts)
}

// UserYSeeCertificateState is a free data retrieval call binding the contract method 0x26fa06d3.
//
// Solidity: function user_YSeeCertificateState() view returns((address,bool,address,string,string,bytes32,uint256)[])
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

// A11SmileGiveETH is a paid mutator transaction binding the contract method 0xae622b60.
//
// Solidity: function A11SmileGiveETH(address Startuse) payable returns()
func (_Gen *GenTransactor) A11SmileGiveETH(opts *bind.TransactOpts, Startuse common.Address) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "A11SmileGiveETH", Startuse)
}

// A11SmileGiveETH is a paid mutator transaction binding the contract method 0xae622b60.
//
// Solidity: function A11SmileGiveETH(address Startuse) payable returns()
func (_Gen *GenSession) A11SmileGiveETH(Startuse common.Address) (*types.Transaction, error) {
	return _Gen.Contract.A11SmileGiveETH(&_Gen.TransactOpts, Startuse)
}

// A11SmileGiveETH is a paid mutator transaction binding the contract method 0xae622b60.
//
// Solidity: function A11SmileGiveETH(address Startuse) payable returns()
func (_Gen *GenTransactorSession) A11SmileGiveETH(Startuse common.Address) (*types.Transaction, error) {
	return _Gen.Contract.A11SmileGiveETH(&_Gen.TransactOpts, Startuse)
}

// A11SmileSetErc is a paid mutator transaction binding the contract method 0xe7809062.
//
// Solidity: function A11Smile_setErc(address ercaddress_) returns()
func (_Gen *GenTransactor) A11SmileSetErc(opts *bind.TransactOpts, ercaddress_ common.Address) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "A11Smile_setErc", ercaddress_)
}

// A11SmileSetErc is a paid mutator transaction binding the contract method 0xe7809062.
//
// Solidity: function A11Smile_setErc(address ercaddress_) returns()
func (_Gen *GenSession) A11SmileSetErc(ercaddress_ common.Address) (*types.Transaction, error) {
	return _Gen.Contract.A11SmileSetErc(&_Gen.TransactOpts, ercaddress_)
}

// A11SmileSetErc is a paid mutator transaction binding the contract method 0xe7809062.
//
// Solidity: function A11Smile_setErc(address ercaddress_) returns()
func (_Gen *GenTransactorSession) A11SmileSetErc(ercaddress_ common.Address) (*types.Transaction, error) {
	return _Gen.Contract.A11SmileSetErc(&_Gen.TransactOpts, ercaddress_)
}

// EthGetAs is a paid mutator transaction binding the contract method 0xba3e8a34.
//
// Solidity: function EthGetAs(address AddEth, address RedEth) payable returns()
func (_Gen *GenTransactor) EthGetAs(opts *bind.TransactOpts, AddEth common.Address, RedEth common.Address) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "EthGetAs", AddEth, RedEth)
}

// EthGetAs is a paid mutator transaction binding the contract method 0xba3e8a34.
//
// Solidity: function EthGetAs(address AddEth, address RedEth) payable returns()
func (_Gen *GenSession) EthGetAs(AddEth common.Address, RedEth common.Address) (*types.Transaction, error) {
	return _Gen.Contract.EthGetAs(&_Gen.TransactOpts, AddEth, RedEth)
}

// EthGetAs is a paid mutator transaction binding the contract method 0xba3e8a34.
//
// Solidity: function EthGetAs(address AddEth, address RedEth) payable returns()
func (_Gen *GenTransactorSession) EthGetAs(AddEth common.Address, RedEth common.Address) (*types.Transaction, error) {
	return _Gen.Contract.EthGetAs(&_Gen.TransactOpts, AddEth, RedEth)
}

// SetCertificate is a paid mutator transaction binding the contract method 0x29ccabdd.
//
// Solidity: function SetCertificate(string[] _m1, string[] _m2) returns()
func (_Gen *GenTransactor) SetCertificate(opts *bind.TransactOpts, _m1 []string, _m2 []string) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "SetCertificate", _m1, _m2)
}

// SetCertificate is a paid mutator transaction binding the contract method 0x29ccabdd.
//
// Solidity: function SetCertificate(string[] _m1, string[] _m2) returns()
func (_Gen *GenSession) SetCertificate(_m1 []string, _m2 []string) (*types.Transaction, error) {
	return _Gen.Contract.SetCertificate(&_Gen.TransactOpts, _m1, _m2)
}

// SetCertificate is a paid mutator transaction binding the contract method 0x29ccabdd.
//
// Solidity: function SetCertificate(string[] _m1, string[] _m2) returns()
func (_Gen *GenTransactorSession) SetCertificate(_m1 []string, _m2 []string) (*types.Transaction, error) {
	return _Gen.Contract.SetCertificate(&_Gen.TransactOpts, _m1, _m2)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Gen *GenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Gen *GenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.Approve(&_Gen.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Gen *GenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.Approve(&_Gen.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Gen *GenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Gen *GenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.DecreaseAllowance(&_Gen.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Gen *GenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.DecreaseAllowance(&_Gen.TransactOpts, spender, subtractedValue)
}

// GainerAddMedicalInformation is a paid mutator transaction binding the contract method 0x68bae494.
//
// Solidity: function gainer_AddMedicalInformation(uint256 min_, uint256 max_, uint256 account_, string MedicalName_, string _Medicalrecordrequirements, string _Requirementdescription) returns()
func (_Gen *GenTransactor) GainerAddMedicalInformation(opts *bind.TransactOpts, min_ *big.Int, max_ *big.Int, account_ *big.Int, MedicalName_ string, _Medicalrecordrequirements string, _Requirementdescription string) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "gainer_AddMedicalInformation", min_, max_, account_, MedicalName_, _Medicalrecordrequirements, _Requirementdescription)
}

// GainerAddMedicalInformation is a paid mutator transaction binding the contract method 0x68bae494.
//
// Solidity: function gainer_AddMedicalInformation(uint256 min_, uint256 max_, uint256 account_, string MedicalName_, string _Medicalrecordrequirements, string _Requirementdescription) returns()
func (_Gen *GenSession) GainerAddMedicalInformation(min_ *big.Int, max_ *big.Int, account_ *big.Int, MedicalName_ string, _Medicalrecordrequirements string, _Requirementdescription string) (*types.Transaction, error) {
	return _Gen.Contract.GainerAddMedicalInformation(&_Gen.TransactOpts, min_, max_, account_, MedicalName_, _Medicalrecordrequirements, _Requirementdescription)
}

// GainerAddMedicalInformation is a paid mutator transaction binding the contract method 0x68bae494.
//
// Solidity: function gainer_AddMedicalInformation(uint256 min_, uint256 max_, uint256 account_, string MedicalName_, string _Medicalrecordrequirements, string _Requirementdescription) returns()
func (_Gen *GenTransactorSession) GainerAddMedicalInformation(min_ *big.Int, max_ *big.Int, account_ *big.Int, MedicalName_ string, _Medicalrecordrequirements string, _Requirementdescription string) (*types.Transaction, error) {
	return _Gen.Contract.GainerAddMedicalInformation(&_Gen.TransactOpts, min_, max_, account_, MedicalName_, _Medicalrecordrequirements, _Requirementdescription)
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

// GainerTransferAS is a paid mutator transaction binding the contract method 0x3d199172.
//
// Solidity: function gainer_transfer_AS(address from1, address to1, uint256 quantity1) returns()
func (_Gen *GenTransactor) GainerTransferAS(opts *bind.TransactOpts, from1 common.Address, to1 common.Address, quantity1 *big.Int) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "gainer_transfer_AS", from1, to1, quantity1)
}

// GainerTransferAS is a paid mutator transaction binding the contract method 0x3d199172.
//
// Solidity: function gainer_transfer_AS(address from1, address to1, uint256 quantity1) returns()
func (_Gen *GenSession) GainerTransferAS(from1 common.Address, to1 common.Address, quantity1 *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.GainerTransferAS(&_Gen.TransactOpts, from1, to1, quantity1)
}

// GainerTransferAS is a paid mutator transaction binding the contract method 0x3d199172.
//
// Solidity: function gainer_transfer_AS(address from1, address to1, uint256 quantity1) returns()
func (_Gen *GenTransactorSession) GainerTransferAS(from1 common.Address, to1 common.Address, quantity1 *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.GainerTransferAS(&_Gen.TransactOpts, from1, to1, quantity1)
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

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Gen *GenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Gen *GenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.IncreaseAllowance(&_Gen.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Gen *GenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.IncreaseAllowance(&_Gen.TransactOpts, spender, addedValue)
}

// Mint1 is a paid mutator transaction binding the contract method 0x25cccb94.
//
// Solidity: function mint1(uint256 amount) returns()
func (_Gen *GenTransactor) Mint1(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "mint1", amount)
}

// Mint1 is a paid mutator transaction binding the contract method 0x25cccb94.
//
// Solidity: function mint1(uint256 amount) returns()
func (_Gen *GenSession) Mint1(amount *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.Mint1(&_Gen.TransactOpts, amount)
}

// Mint1 is a paid mutator transaction binding the contract method 0x25cccb94.
//
// Solidity: function mint1(uint256 amount) returns()
func (_Gen *GenTransactorSession) Mint1(amount *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.Mint1(&_Gen.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Gen *GenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Gen *GenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.Transfer(&_Gen.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Gen *GenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.Transfer(&_Gen.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Gen *GenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Gen *GenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.TransferFrom(&_Gen.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Gen *GenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gen.Contract.TransferFrom(&_Gen.TransactOpts, sender, recipient, amount)
}

// UserAddMedicalInformation is a paid mutator transaction binding the contract method 0x753772dc.
//
// Solidity: function user_AddMedicalInformation(bytes32 Proute, address _soliciter, string MedicalName_) returns()
func (_Gen *GenTransactor) UserAddMedicalInformation(opts *bind.TransactOpts, Proute [32]byte, _soliciter common.Address, MedicalName_ string) (*types.Transaction, error) {
	return _Gen.contract.Transact(opts, "user_AddMedicalInformation", Proute, _soliciter, MedicalName_)
}

// UserAddMedicalInformation is a paid mutator transaction binding the contract method 0x753772dc.
//
// Solidity: function user_AddMedicalInformation(bytes32 Proute, address _soliciter, string MedicalName_) returns()
func (_Gen *GenSession) UserAddMedicalInformation(Proute [32]byte, _soliciter common.Address, MedicalName_ string) (*types.Transaction, error) {
	return _Gen.Contract.UserAddMedicalInformation(&_Gen.TransactOpts, Proute, _soliciter, MedicalName_)
}

// UserAddMedicalInformation is a paid mutator transaction binding the contract method 0x753772dc.
//
// Solidity: function user_AddMedicalInformation(bytes32 Proute, address _soliciter, string MedicalName_) returns()
func (_Gen *GenTransactorSession) UserAddMedicalInformation(Proute [32]byte, _soliciter common.Address, MedicalName_ string) (*types.Transaction, error) {
	return _Gen.Contract.UserAddMedicalInformation(&_Gen.TransactOpts, Proute, _soliciter, MedicalName_)
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

// GenA11GiveETHIterator is returned from FilterA11GiveETH and is used to iterate over the raw logs and unpacked data for A11GiveETH events raised by the Gen contract.
type GenA11GiveETHIterator struct {
	Event *GenA11GiveETH // Event containing the contract specifics and raw log

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
func (it *GenA11GiveETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenA11GiveETH)
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
		it.Event = new(GenA11GiveETH)
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
func (it *GenA11GiveETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenA11GiveETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenA11GiveETH represents a A11GiveETH event raised by the Gen contract.
type GenA11GiveETH struct {
	User1 common.Address
	User2 common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterA11GiveETH is a free log retrieval operation binding the contract event 0x41003739cb7adda22b5e7029b076372e289d38e49fcb784b70e3f2883dfb0265.
//
// Solidity: event A11GiveETH(address indexed user1, address indexed user2, uint256 indexed price)
func (_Gen *GenFilterer) FilterA11GiveETH(opts *bind.FilterOpts, user1 []common.Address, user2 []common.Address, price []*big.Int) (*GenA11GiveETHIterator, error) {

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

	logs, sub, err := _Gen.contract.FilterLogs(opts, "A11GiveETH", user1Rule, user2Rule, priceRule)
	if err != nil {
		return nil, err
	}
	return &GenA11GiveETHIterator{contract: _Gen.contract, event: "A11GiveETH", logs: logs, sub: sub}, nil
}

// WatchA11GiveETH is a free log subscription operation binding the contract event 0x41003739cb7adda22b5e7029b076372e289d38e49fcb784b70e3f2883dfb0265.
//
// Solidity: event A11GiveETH(address indexed user1, address indexed user2, uint256 indexed price)
func (_Gen *GenFilterer) WatchA11GiveETH(opts *bind.WatchOpts, sink chan<- *GenA11GiveETH, user1 []common.Address, user2 []common.Address, price []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Gen.contract.WatchLogs(opts, "A11GiveETH", user1Rule, user2Rule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenA11GiveETH)
				if err := _Gen.contract.UnpackLog(event, "A11GiveETH", log); err != nil {
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
func (_Gen *GenFilterer) ParseA11GiveETH(log types.Log) (*GenA11GiveETH, error) {
	event := new(GenA11GiveETH)
	if err := _Gen.contract.UnpackLog(event, "A11GiveETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenAddtokensIterator is returned from FilterAddtokens and is used to iterate over the raw logs and unpacked data for Addtokens events raised by the Gen contract.
type GenAddtokensIterator struct {
	Event *GenAddtokens // Event containing the contract specifics and raw log

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
func (it *GenAddtokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenAddtokens)
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
		it.Event = new(GenAddtokens)
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
func (it *GenAddtokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenAddtokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenAddtokens represents a Addtokens event raised by the Gen contract.
type GenAddtokens struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddtokens is a free log retrieval operation binding the contract event 0x140232c0c4b8ad28248113c4185b3b3af0ba6f0589381f850b8731ad8af29db8.
//
// Solidity: event Addtokens(address indexed owner, uint256 indexed amount)
func (_Gen *GenFilterer) FilterAddtokens(opts *bind.FilterOpts, owner []common.Address, amount []*big.Int) (*GenAddtokensIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Gen.contract.FilterLogs(opts, "Addtokens", ownerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &GenAddtokensIterator{contract: _Gen.contract, event: "Addtokens", logs: logs, sub: sub}, nil
}

// WatchAddtokens is a free log subscription operation binding the contract event 0x140232c0c4b8ad28248113c4185b3b3af0ba6f0589381f850b8731ad8af29db8.
//
// Solidity: event Addtokens(address indexed owner, uint256 indexed amount)
func (_Gen *GenFilterer) WatchAddtokens(opts *bind.WatchOpts, sink chan<- *GenAddtokens, owner []common.Address, amount []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Gen.contract.WatchLogs(opts, "Addtokens", ownerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenAddtokens)
				if err := _Gen.contract.UnpackLog(event, "Addtokens", log); err != nil {
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
func (_Gen *GenFilterer) ParseAddtokens(log types.Log) (*GenAddtokens, error) {
	event := new(GenAddtokens)
	if err := _Gen.contract.UnpackLog(event, "Addtokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Gen contract.
type GenApprovalIterator struct {
	Event *GenApproval // Event containing the contract specifics and raw log

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
func (it *GenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenApproval)
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
		it.Event = new(GenApproval)
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
func (it *GenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenApproval represents a Approval event raised by the Gen contract.
type GenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Gen *GenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Gen.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GenApprovalIterator{contract: _Gen.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Gen *GenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Gen.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenApproval)
				if err := _Gen.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Gen *GenFilterer) ParseApproval(log types.Log) (*GenApproval, error) {
	event := new(GenApproval)
	if err := _Gen.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenEthgetAsIterator is returned from FilterEthgetAs and is used to iterate over the raw logs and unpacked data for EthgetAs events raised by the Gen contract.
type GenEthgetAsIterator struct {
	Event *GenEthgetAs // Event containing the contract specifics and raw log

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
func (it *GenEthgetAsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenEthgetAs)
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
		it.Event = new(GenEthgetAs)
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
func (it *GenEthgetAsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenEthgetAsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenEthgetAs represents a EthgetAs event raised by the Gen contract.
type GenEthgetAs struct {
	User1 common.Address
	User2 common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEthgetAs is a free log retrieval operation binding the contract event 0x4e28aa9ff438c656a77043806d0d6a21a098eda7102c9b52698d2a6f321fc3e6.
//
// Solidity: event EthgetAs(address indexed user1, address indexed user2, uint256 indexed price)
func (_Gen *GenFilterer) FilterEthgetAs(opts *bind.FilterOpts, user1 []common.Address, user2 []common.Address, price []*big.Int) (*GenEthgetAsIterator, error) {

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

	logs, sub, err := _Gen.contract.FilterLogs(opts, "EthgetAs", user1Rule, user2Rule, priceRule)
	if err != nil {
		return nil, err
	}
	return &GenEthgetAsIterator{contract: _Gen.contract, event: "EthgetAs", logs: logs, sub: sub}, nil
}

// WatchEthgetAs is a free log subscription operation binding the contract event 0x4e28aa9ff438c656a77043806d0d6a21a098eda7102c9b52698d2a6f321fc3e6.
//
// Solidity: event EthgetAs(address indexed user1, address indexed user2, uint256 indexed price)
func (_Gen *GenFilterer) WatchEthgetAs(opts *bind.WatchOpts, sink chan<- *GenEthgetAs, user1 []common.Address, user2 []common.Address, price []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Gen.contract.WatchLogs(opts, "EthgetAs", user1Rule, user2Rule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenEthgetAs)
				if err := _Gen.contract.UnpackLog(event, "EthgetAs", log); err != nil {
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
func (_Gen *GenFilterer) ParseEthgetAs(log types.Log) (*GenEthgetAs, error) {
	event := new(GenEthgetAs)
	if err := _Gen.contract.UnpackLog(event, "EthgetAs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Gen contract.
type GenTransferIterator struct {
	Event *GenTransfer // Event containing the contract specifics and raw log

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
func (it *GenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenTransfer)
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
		it.Event = new(GenTransfer)
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
func (it *GenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenTransfer represents a Transfer event raised by the Gen contract.
type GenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Gen *GenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Gen.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GenTransferIterator{contract: _Gen.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Gen *GenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Gen.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenTransfer)
				if err := _Gen.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Gen *GenFilterer) ParseTransfer(log types.Log) (*GenTransfer, error) {
	event := new(GenTransfer)
	if err := _Gen.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenUploadmedicaldataIterator is returned from FilterUploadmedicaldata and is used to iterate over the raw logs and unpacked data for Uploadmedicaldata events raised by the Gen contract.
type GenUploadmedicaldataIterator struct {
	Event *GenUploadmedicaldata // Event containing the contract specifics and raw log

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
func (it *GenUploadmedicaldataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenUploadmedicaldata)
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
		it.Event = new(GenUploadmedicaldata)
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
func (it *GenUploadmedicaldataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenUploadmedicaldataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenUploadmedicaldata represents a Uploadmedicaldata event raised by the Gen contract.
type GenUploadmedicaldata struct {
	User      common.Address
	Route     common.Hash
	Soliciter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUploadmedicaldata is a free log retrieval operation binding the contract event 0x49868a6e303058cc0ee85fad003812b132b8ba154aafa643a8763b774a8b8495.
//
// Solidity: event Uploadmedicaldata(address indexed user, string indexed route, address indexed soliciter)
func (_Gen *GenFilterer) FilterUploadmedicaldata(opts *bind.FilterOpts, user []common.Address, route []string, soliciter []common.Address) (*GenUploadmedicaldataIterator, error) {

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

	logs, sub, err := _Gen.contract.FilterLogs(opts, "Uploadmedicaldata", userRule, routeRule, soliciterRule)
	if err != nil {
		return nil, err
	}
	return &GenUploadmedicaldataIterator{contract: _Gen.contract, event: "Uploadmedicaldata", logs: logs, sub: sub}, nil
}

// WatchUploadmedicaldata is a free log subscription operation binding the contract event 0x49868a6e303058cc0ee85fad003812b132b8ba154aafa643a8763b774a8b8495.
//
// Solidity: event Uploadmedicaldata(address indexed user, string indexed route, address indexed soliciter)
func (_Gen *GenFilterer) WatchUploadmedicaldata(opts *bind.WatchOpts, sink chan<- *GenUploadmedicaldata, user []common.Address, route []string, soliciter []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Gen.contract.WatchLogs(opts, "Uploadmedicaldata", userRule, routeRule, soliciterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenUploadmedicaldata)
				if err := _Gen.contract.UnpackLog(event, "Uploadmedicaldata", log); err != nil {
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
func (_Gen *GenFilterer) ParseUploadmedicaldata(log types.Log) (*GenUploadmedicaldata, error) {
	event := new(GenUploadmedicaldata)
	if err := _Gen.contract.UnpackLog(event, "Uploadmedicaldata", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenGainerUploadmedicaldataIterator is returned from FilterGainerUploadmedicaldata and is used to iterate over the raw logs and unpacked data for GainerUploadmedicaldata events raised by the Gen contract.
type GenGainerUploadmedicaldataIterator struct {
	Event *GenGainerUploadmedicaldata // Event containing the contract specifics and raw log

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
func (it *GenGainerUploadmedicaldataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenGainerUploadmedicaldata)
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
		it.Event = new(GenGainerUploadmedicaldata)
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
func (it *GenGainerUploadmedicaldataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenGainerUploadmedicaldataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenGainerUploadmedicaldata represents a GainerUploadmedicaldata event raised by the Gen contract.
type GenGainerUploadmedicaldata struct {
	Route     common.Hash
	Soliciter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGainerUploadmedicaldata is a free log retrieval operation binding the contract event 0x575ef43b443a501121681e81acef51358e5f96f8b3493a3b964389af363dbaf7.
//
// Solidity: event gainerUploadmedicaldata(string indexed route, address indexed soliciter)
func (_Gen *GenFilterer) FilterGainerUploadmedicaldata(opts *bind.FilterOpts, route []string, soliciter []common.Address) (*GenGainerUploadmedicaldataIterator, error) {

	var routeRule []interface{}
	for _, routeItem := range route {
		routeRule = append(routeRule, routeItem)
	}
	var soliciterRule []interface{}
	for _, soliciterItem := range soliciter {
		soliciterRule = append(soliciterRule, soliciterItem)
	}

	logs, sub, err := _Gen.contract.FilterLogs(opts, "gainerUploadmedicaldata", routeRule, soliciterRule)
	if err != nil {
		return nil, err
	}
	return &GenGainerUploadmedicaldataIterator{contract: _Gen.contract, event: "gainerUploadmedicaldata", logs: logs, sub: sub}, nil
}

// WatchGainerUploadmedicaldata is a free log subscription operation binding the contract event 0x575ef43b443a501121681e81acef51358e5f96f8b3493a3b964389af363dbaf7.
//
// Solidity: event gainerUploadmedicaldata(string indexed route, address indexed soliciter)
func (_Gen *GenFilterer) WatchGainerUploadmedicaldata(opts *bind.WatchOpts, sink chan<- *GenGainerUploadmedicaldata, route []string, soliciter []common.Address) (event.Subscription, error) {

	var routeRule []interface{}
	for _, routeItem := range route {
		routeRule = append(routeRule, routeItem)
	}
	var soliciterRule []interface{}
	for _, soliciterItem := range soliciter {
		soliciterRule = append(soliciterRule, soliciterItem)
	}

	logs, sub, err := _Gen.contract.WatchLogs(opts, "gainerUploadmedicaldata", routeRule, soliciterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenGainerUploadmedicaldata)
				if err := _Gen.contract.UnpackLog(event, "gainerUploadmedicaldata", log); err != nil {
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
func (_Gen *GenFilterer) ParseGainerUploadmedicaldata(log types.Log) (*GenGainerUploadmedicaldata, error) {
	event := new(GenGainerUploadmedicaldata)
	if err := _Gen.contract.UnpackLog(event, "gainerUploadmedicaldata", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GenRewardIterator is returned from FilterReward and is used to iterate over the raw logs and unpacked data for Reward events raised by the Gen contract.
type GenRewardIterator struct {
	Event *GenReward // Event containing the contract specifics and raw log

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
func (it *GenRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GenReward)
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
		it.Event = new(GenReward)
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
func (it *GenRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GenRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GenReward represents a Reward event raised by the Gen contract.
type GenReward struct {
	Owner common.Address
	Route common.Hash
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReward is a free log retrieval operation binding the contract event 0xab58f406e20f71235aadf9324c46cdb688e9be29c543b558f4cc8601a41ea4bb.
//
// Solidity: event reward(address indexed owner, string indexed route, uint256 indexed price)
func (_Gen *GenFilterer) FilterReward(opts *bind.FilterOpts, owner []common.Address, route []string, price []*big.Int) (*GenRewardIterator, error) {

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

	logs, sub, err := _Gen.contract.FilterLogs(opts, "reward", ownerRule, routeRule, priceRule)
	if err != nil {
		return nil, err
	}
	return &GenRewardIterator{contract: _Gen.contract, event: "reward", logs: logs, sub: sub}, nil
}

// WatchReward is a free log subscription operation binding the contract event 0xab58f406e20f71235aadf9324c46cdb688e9be29c543b558f4cc8601a41ea4bb.
//
// Solidity: event reward(address indexed owner, string indexed route, uint256 indexed price)
func (_Gen *GenFilterer) WatchReward(opts *bind.WatchOpts, sink chan<- *GenReward, owner []common.Address, route []string, price []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Gen.contract.WatchLogs(opts, "reward", ownerRule, routeRule, priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GenReward)
				if err := _Gen.contract.UnpackLog(event, "reward", log); err != nil {
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
func (_Gen *GenFilterer) ParseReward(log types.Log) (*GenReward, error) {
	event := new(GenReward)
	if err := _Gen.contract.UnpackLog(event, "reward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
