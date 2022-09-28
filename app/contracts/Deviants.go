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

// DeviantsParams is an auto generated low-level Go binding around an user-defined struct.
type DeviantsParams struct {
	Name                 string
	Symbol               string
	BaseURI              string
	MaxTotalSupply       *big.Int
	BulkBuyLimit         *big.Int
	PublicPrice          *big.Int
	DiscountPrice        *big.Int
	DaoAddress           common.Address
	PolymorphsV1Contract common.Address
	PolymorphsV2Contract common.Address
	FacesContract        common.Address
	LobstersContract     common.Address
	RoyaltyFee           *big.Int
}

// DeviantsMetaData contains all meta data concerning the Deviants contract.
var DeviantsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bulkBuyLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_publicPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_discountPrice\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_daoAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polymorphsV1Contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polymorphsV2Contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_facesContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lobstersContract\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_royaltyFee\",\"type\":\"uint96\"}],\"internalType\":\"structDeviants.Params\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newReciver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"value\",\"type\":\"uint96\"}],\"name\":\"RoyaltiesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"usedDiscount\",\"type\":\"bool\"}],\"name\":\"TokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bulkBuyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"daoMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"discountMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"discountPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"discountsUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergeWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facesContract\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"geneOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gene\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lobstersContract\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polymorphsV1Contract\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polymorphsV2Contract\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newLimit\",\"type\":\"uint256\"}],\"name\":\"setBulkBuyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"setDiscountPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_value\",\"type\":\"uint96\"}],\"name\":\"setRoyalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DeviantsABI is the input ABI used to generate the binding from.
// Deprecated: Use DeviantsMetaData.ABI instead.
var DeviantsABI = DeviantsMetaData.ABI

// Deviants is an auto generated Go binding around an Ethereum contract.
type Deviants struct {
	DeviantsCaller     // Read-only binding to the contract
	DeviantsTransactor // Write-only binding to the contract
	DeviantsFilterer   // Log filterer for contract events
}

// DeviantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeviantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeviantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeviantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeviantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeviantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeviantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeviantsSession struct {
	Contract     *Deviants         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeviantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeviantsCallerSession struct {
	Contract *DeviantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DeviantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeviantsTransactorSession struct {
	Contract     *DeviantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DeviantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeviantsRaw struct {
	Contract *Deviants // Generic contract binding to access the raw methods on
}

// DeviantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeviantsCallerRaw struct {
	Contract *DeviantsCaller // Generic read-only contract binding to access the raw methods on
}

// DeviantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeviantsTransactorRaw struct {
	Contract *DeviantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeviants creates a new instance of Deviants, bound to a specific deployed contract.
func NewDeviants(address common.Address, backend bind.ContractBackend) (*Deviants, error) {
	contract, err := bindDeviants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deviants{DeviantsCaller: DeviantsCaller{contract: contract}, DeviantsTransactor: DeviantsTransactor{contract: contract}, DeviantsFilterer: DeviantsFilterer{contract: contract}}, nil
}

// NewDeviantsCaller creates a new read-only instance of Deviants, bound to a specific deployed contract.
func NewDeviantsCaller(address common.Address, caller bind.ContractCaller) (*DeviantsCaller, error) {
	contract, err := bindDeviants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeviantsCaller{contract: contract}, nil
}

// NewDeviantsTransactor creates a new write-only instance of Deviants, bound to a specific deployed contract.
func NewDeviantsTransactor(address common.Address, transactor bind.ContractTransactor) (*DeviantsTransactor, error) {
	contract, err := bindDeviants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeviantsTransactor{contract: contract}, nil
}

// NewDeviantsFilterer creates a new log filterer instance of Deviants, bound to a specific deployed contract.
func NewDeviantsFilterer(address common.Address, filterer bind.ContractFilterer) (*DeviantsFilterer, error) {
	contract, err := bindDeviants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeviantsFilterer{contract: contract}, nil
}

// bindDeviants binds a generic wrapper to an already deployed contract.
func bindDeviants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeviantsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deviants *DeviantsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deviants.Contract.DeviantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deviants *DeviantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deviants.Contract.DeviantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deviants *DeviantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deviants.Contract.DeviantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deviants *DeviantsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deviants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deviants *DeviantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deviants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deviants *DeviantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deviants.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Deviants *DeviantsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Deviants *DeviantsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Deviants.Contract.BalanceOf(&_Deviants.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Deviants *DeviantsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Deviants.Contract.BalanceOf(&_Deviants.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Deviants *DeviantsCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Deviants *DeviantsSession) BaseURI() (string, error) {
	return _Deviants.Contract.BaseURI(&_Deviants.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Deviants *DeviantsCallerSession) BaseURI() (string, error) {
	return _Deviants.Contract.BaseURI(&_Deviants.CallOpts)
}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint256)
func (_Deviants *DeviantsCaller) BulkBuyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "bulkBuyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint256)
func (_Deviants *DeviantsSession) BulkBuyLimit() (*big.Int, error) {
	return _Deviants.Contract.BulkBuyLimit(&_Deviants.CallOpts)
}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint256)
func (_Deviants *DeviantsCallerSession) BulkBuyLimit() (*big.Int, error) {
	return _Deviants.Contract.BulkBuyLimit(&_Deviants.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_Deviants *DeviantsCaller) DaoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "daoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_Deviants *DeviantsSession) DaoAddress() (common.Address, error) {
	return _Deviants.Contract.DaoAddress(&_Deviants.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_Deviants *DeviantsCallerSession) DaoAddress() (common.Address, error) {
	return _Deviants.Contract.DaoAddress(&_Deviants.CallOpts)
}

// DiscountPrice is a free data retrieval call binding the contract method 0x84ad8e8f.
//
// Solidity: function discountPrice() view returns(uint256)
func (_Deviants *DeviantsCaller) DiscountPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "discountPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DiscountPrice is a free data retrieval call binding the contract method 0x84ad8e8f.
//
// Solidity: function discountPrice() view returns(uint256)
func (_Deviants *DeviantsSession) DiscountPrice() (*big.Int, error) {
	return _Deviants.Contract.DiscountPrice(&_Deviants.CallOpts)
}

// DiscountPrice is a free data retrieval call binding the contract method 0x84ad8e8f.
//
// Solidity: function discountPrice() view returns(uint256)
func (_Deviants *DeviantsCallerSession) DiscountPrice() (*big.Int, error) {
	return _Deviants.Contract.DiscountPrice(&_Deviants.CallOpts)
}

// DiscountsUsed is a free data retrieval call binding the contract method 0x76963e58.
//
// Solidity: function discountsUsed(address ) view returns(uint256)
func (_Deviants *DeviantsCaller) DiscountsUsed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "discountsUsed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DiscountsUsed is a free data retrieval call binding the contract method 0x76963e58.
//
// Solidity: function discountsUsed(address ) view returns(uint256)
func (_Deviants *DeviantsSession) DiscountsUsed(arg0 common.Address) (*big.Int, error) {
	return _Deviants.Contract.DiscountsUsed(&_Deviants.CallOpts, arg0)
}

// DiscountsUsed is a free data retrieval call binding the contract method 0x76963e58.
//
// Solidity: function discountsUsed(address ) view returns(uint256)
func (_Deviants *DeviantsCallerSession) DiscountsUsed(arg0 common.Address) (*big.Int, error) {
	return _Deviants.Contract.DiscountsUsed(&_Deviants.CallOpts, arg0)
}

// FacesContract is a free data retrieval call binding the contract method 0x19c05f6d.
//
// Solidity: function facesContract() view returns(address)
func (_Deviants *DeviantsCaller) FacesContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "facesContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacesContract is a free data retrieval call binding the contract method 0x19c05f6d.
//
// Solidity: function facesContract() view returns(address)
func (_Deviants *DeviantsSession) FacesContract() (common.Address, error) {
	return _Deviants.Contract.FacesContract(&_Deviants.CallOpts)
}

// FacesContract is a free data retrieval call binding the contract method 0x19c05f6d.
//
// Solidity: function facesContract() view returns(address)
func (_Deviants *DeviantsCallerSession) FacesContract() (common.Address, error) {
	return _Deviants.Contract.FacesContract(&_Deviants.CallOpts)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_Deviants *DeviantsCaller) GeneOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "geneOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_Deviants *DeviantsSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _Deviants.Contract.GeneOf(&_Deviants.CallOpts, tokenId)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_Deviants *DeviantsCallerSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _Deviants.Contract.GeneOf(&_Deviants.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Deviants *DeviantsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Deviants *DeviantsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Deviants.Contract.GetApproved(&_Deviants.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Deviants *DeviantsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Deviants.Contract.GetApproved(&_Deviants.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Deviants *DeviantsCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Deviants *DeviantsSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Deviants.Contract.IsApprovedForAll(&_Deviants.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Deviants *DeviantsCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Deviants.Contract.IsApprovedForAll(&_Deviants.CallOpts, owner, operator)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256)
func (_Deviants *DeviantsCaller) LastTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "lastTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256)
func (_Deviants *DeviantsSession) LastTokenId() (*big.Int, error) {
	return _Deviants.Contract.LastTokenId(&_Deviants.CallOpts)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256)
func (_Deviants *DeviantsCallerSession) LastTokenId() (*big.Int, error) {
	return _Deviants.Contract.LastTokenId(&_Deviants.CallOpts)
}

// LobstersContract is a free data retrieval call binding the contract method 0x38a10ee3.
//
// Solidity: function lobstersContract() view returns(address)
func (_Deviants *DeviantsCaller) LobstersContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "lobstersContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LobstersContract is a free data retrieval call binding the contract method 0x38a10ee3.
//
// Solidity: function lobstersContract() view returns(address)
func (_Deviants *DeviantsSession) LobstersContract() (common.Address, error) {
	return _Deviants.Contract.LobstersContract(&_Deviants.CallOpts)
}

// LobstersContract is a free data retrieval call binding the contract method 0x38a10ee3.
//
// Solidity: function lobstersContract() view returns(address)
func (_Deviants *DeviantsCallerSession) LobstersContract() (common.Address, error) {
	return _Deviants.Contract.LobstersContract(&_Deviants.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Deviants *DeviantsCaller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "maxTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Deviants *DeviantsSession) MaxTotalSupply() (*big.Int, error) {
	return _Deviants.Contract.MaxTotalSupply(&_Deviants.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Deviants *DeviantsCallerSession) MaxTotalSupply() (*big.Int, error) {
	return _Deviants.Contract.MaxTotalSupply(&_Deviants.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Deviants *DeviantsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Deviants *DeviantsSession) Name() (string, error) {
	return _Deviants.Contract.Name(&_Deviants.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Deviants *DeviantsCallerSession) Name() (string, error) {
	return _Deviants.Contract.Name(&_Deviants.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Deviants *DeviantsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Deviants *DeviantsSession) Owner() (common.Address, error) {
	return _Deviants.Contract.Owner(&_Deviants.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Deviants *DeviantsCallerSession) Owner() (common.Address, error) {
	return _Deviants.Contract.Owner(&_Deviants.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Deviants *DeviantsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Deviants *DeviantsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Deviants.Contract.OwnerOf(&_Deviants.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Deviants *DeviantsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Deviants.Contract.OwnerOf(&_Deviants.CallOpts, tokenId)
}

// PolymorphsV1Contract is a free data retrieval call binding the contract method 0x8ddeb63c.
//
// Solidity: function polymorphsV1Contract() view returns(address)
func (_Deviants *DeviantsCaller) PolymorphsV1Contract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "polymorphsV1Contract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolymorphsV1Contract is a free data retrieval call binding the contract method 0x8ddeb63c.
//
// Solidity: function polymorphsV1Contract() view returns(address)
func (_Deviants *DeviantsSession) PolymorphsV1Contract() (common.Address, error) {
	return _Deviants.Contract.PolymorphsV1Contract(&_Deviants.CallOpts)
}

// PolymorphsV1Contract is a free data retrieval call binding the contract method 0x8ddeb63c.
//
// Solidity: function polymorphsV1Contract() view returns(address)
func (_Deviants *DeviantsCallerSession) PolymorphsV1Contract() (common.Address, error) {
	return _Deviants.Contract.PolymorphsV1Contract(&_Deviants.CallOpts)
}

// PolymorphsV2Contract is a free data retrieval call binding the contract method 0xde443ddf.
//
// Solidity: function polymorphsV2Contract() view returns(address)
func (_Deviants *DeviantsCaller) PolymorphsV2Contract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "polymorphsV2Contract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolymorphsV2Contract is a free data retrieval call binding the contract method 0xde443ddf.
//
// Solidity: function polymorphsV2Contract() view returns(address)
func (_Deviants *DeviantsSession) PolymorphsV2Contract() (common.Address, error) {
	return _Deviants.Contract.PolymorphsV2Contract(&_Deviants.CallOpts)
}

// PolymorphsV2Contract is a free data retrieval call binding the contract method 0xde443ddf.
//
// Solidity: function polymorphsV2Contract() view returns(address)
func (_Deviants *DeviantsCallerSession) PolymorphsV2Contract() (common.Address, error) {
	return _Deviants.Contract.PolymorphsV2Contract(&_Deviants.CallOpts)
}

// PublicPrice is a free data retrieval call binding the contract method 0xa945bf80.
//
// Solidity: function publicPrice() view returns(uint256)
func (_Deviants *DeviantsCaller) PublicPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "publicPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicPrice is a free data retrieval call binding the contract method 0xa945bf80.
//
// Solidity: function publicPrice() view returns(uint256)
func (_Deviants *DeviantsSession) PublicPrice() (*big.Int, error) {
	return _Deviants.Contract.PublicPrice(&_Deviants.CallOpts)
}

// PublicPrice is a free data retrieval call binding the contract method 0xa945bf80.
//
// Solidity: function publicPrice() view returns(uint256)
func (_Deviants *DeviantsCallerSession) PublicPrice() (*big.Int, error) {
	return _Deviants.Contract.PublicPrice(&_Deviants.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_Deviants *DeviantsCaller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_Deviants *DeviantsSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Deviants.Contract.RoyaltyInfo(&_Deviants.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_Deviants *DeviantsCallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Deviants.Contract.RoyaltyInfo(&_Deviants.CallOpts, _tokenId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Deviants *DeviantsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Deviants *DeviantsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Deviants.Contract.SupportsInterface(&_Deviants.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Deviants *DeviantsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Deviants.Contract.SupportsInterface(&_Deviants.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Deviants *DeviantsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Deviants *DeviantsSession) Symbol() (string, error) {
	return _Deviants.Contract.Symbol(&_Deviants.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Deviants *DeviantsCallerSession) Symbol() (string, error) {
	return _Deviants.Contract.Symbol(&_Deviants.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Deviants *DeviantsCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Deviants *DeviantsSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Deviants.Contract.TokenURI(&_Deviants.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Deviants *DeviantsCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Deviants.Contract.TokenURI(&_Deviants.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Deviants *DeviantsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deviants.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Deviants *DeviantsSession) TotalSupply() (*big.Int, error) {
	return _Deviants.Contract.TotalSupply(&_Deviants.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Deviants *DeviantsCallerSession) TotalSupply() (*big.Int, error) {
	return _Deviants.Contract.TotalSupply(&_Deviants.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Deviants *DeviantsTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Deviants *DeviantsSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.Approve(&_Deviants.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Deviants *DeviantsTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.Approve(&_Deviants.TransactOpts, to, tokenId)
}

// DaoMint is a paid mutator transaction binding the contract method 0x62759f6c.
//
// Solidity: function daoMint(uint256 _amount) returns()
func (_Deviants *DeviantsTransactor) DaoMint(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "daoMint", _amount)
}

// DaoMint is a paid mutator transaction binding the contract method 0x62759f6c.
//
// Solidity: function daoMint(uint256 _amount) returns()
func (_Deviants *DeviantsSession) DaoMint(_amount *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.DaoMint(&_Deviants.TransactOpts, _amount)
}

// DaoMint is a paid mutator transaction binding the contract method 0x62759f6c.
//
// Solidity: function daoMint(uint256 _amount) returns()
func (_Deviants *DeviantsTransactorSession) DaoMint(_amount *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.DaoMint(&_Deviants.TransactOpts, _amount)
}

// DiscountMint is a paid mutator transaction binding the contract method 0x41de890e.
//
// Solidity: function discountMint(uint256 _amount) payable returns()
func (_Deviants *DeviantsTransactor) DiscountMint(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "discountMint", _amount)
}

// DiscountMint is a paid mutator transaction binding the contract method 0x41de890e.
//
// Solidity: function discountMint(uint256 _amount) payable returns()
func (_Deviants *DeviantsSession) DiscountMint(_amount *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.DiscountMint(&_Deviants.TransactOpts, _amount)
}

// DiscountMint is a paid mutator transaction binding the contract method 0x41de890e.
//
// Solidity: function discountMint(uint256 _amount) payable returns()
func (_Deviants *DeviantsTransactorSession) DiscountMint(_amount *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.DiscountMint(&_Deviants.TransactOpts, _amount)
}

// EmergeWithdraw is a paid mutator transaction binding the contract method 0xb13529d5.
//
// Solidity: function emergeWithdraw() returns()
func (_Deviants *DeviantsTransactor) EmergeWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "emergeWithdraw")
}

// EmergeWithdraw is a paid mutator transaction binding the contract method 0xb13529d5.
//
// Solidity: function emergeWithdraw() returns()
func (_Deviants *DeviantsSession) EmergeWithdraw() (*types.Transaction, error) {
	return _Deviants.Contract.EmergeWithdraw(&_Deviants.TransactOpts)
}

// EmergeWithdraw is a paid mutator transaction binding the contract method 0xb13529d5.
//
// Solidity: function emergeWithdraw() returns()
func (_Deviants *DeviantsTransactorSession) EmergeWithdraw() (*types.Transaction, error) {
	return _Deviants.Contract.EmergeWithdraw(&_Deviants.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _amount) payable returns()
func (_Deviants *DeviantsTransactor) Mint(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "mint", _amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _amount) payable returns()
func (_Deviants *DeviantsSession) Mint(_amount *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.Mint(&_Deviants.TransactOpts, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _amount) payable returns()
func (_Deviants *DeviantsTransactorSession) Mint(_amount *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.Mint(&_Deviants.TransactOpts, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Deviants *DeviantsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Deviants *DeviantsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Deviants.Contract.RenounceOwnership(&_Deviants.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Deviants *DeviantsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Deviants.Contract.RenounceOwnership(&_Deviants.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Deviants *DeviantsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Deviants *DeviantsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.SafeTransferFrom(&_Deviants.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Deviants *DeviantsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.SafeTransferFrom(&_Deviants.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Deviants *DeviantsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Deviants *DeviantsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Deviants.Contract.SafeTransferFrom0(&_Deviants.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Deviants *DeviantsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Deviants.Contract.SafeTransferFrom0(&_Deviants.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Deviants *DeviantsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Deviants *DeviantsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Deviants.Contract.SetApprovalForAll(&_Deviants.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Deviants *DeviantsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Deviants.Contract.SetApprovalForAll(&_Deviants.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Deviants *DeviantsTransactor) SetBaseURI(opts *bind.TransactOpts, _baseURI string) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "setBaseURI", _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Deviants *DeviantsSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _Deviants.Contract.SetBaseURI(&_Deviants.TransactOpts, _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Deviants *DeviantsTransactorSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _Deviants.Contract.SetBaseURI(&_Deviants.TransactOpts, _baseURI)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x5e468dfd.
//
// Solidity: function setBulkBuyLimit(uint256 _newLimit) returns()
func (_Deviants *DeviantsTransactor) SetBulkBuyLimit(opts *bind.TransactOpts, _newLimit *big.Int) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "setBulkBuyLimit", _newLimit)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x5e468dfd.
//
// Solidity: function setBulkBuyLimit(uint256 _newLimit) returns()
func (_Deviants *DeviantsSession) SetBulkBuyLimit(_newLimit *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.SetBulkBuyLimit(&_Deviants.TransactOpts, _newLimit)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x5e468dfd.
//
// Solidity: function setBulkBuyLimit(uint256 _newLimit) returns()
func (_Deviants *DeviantsTransactorSession) SetBulkBuyLimit(_newLimit *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.SetBulkBuyLimit(&_Deviants.TransactOpts, _newLimit)
}

// SetDiscountPrice is a paid mutator transaction binding the contract method 0x8b78c116.
//
// Solidity: function setDiscountPrice(uint256 _newPrice) returns()
func (_Deviants *DeviantsTransactor) SetDiscountPrice(opts *bind.TransactOpts, _newPrice *big.Int) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "setDiscountPrice", _newPrice)
}

// SetDiscountPrice is a paid mutator transaction binding the contract method 0x8b78c116.
//
// Solidity: function setDiscountPrice(uint256 _newPrice) returns()
func (_Deviants *DeviantsSession) SetDiscountPrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.SetDiscountPrice(&_Deviants.TransactOpts, _newPrice)
}

// SetDiscountPrice is a paid mutator transaction binding the contract method 0x8b78c116.
//
// Solidity: function setDiscountPrice(uint256 _newPrice) returns()
func (_Deviants *DeviantsTransactorSession) SetDiscountPrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.SetDiscountPrice(&_Deviants.TransactOpts, _newPrice)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_Deviants *DeviantsTransactor) SetMaxSupply(opts *bind.TransactOpts, _maxSupply *big.Int) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "setMaxSupply", _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_Deviants *DeviantsSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.SetMaxSupply(&_Deviants.TransactOpts, _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_Deviants *DeviantsTransactorSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.SetMaxSupply(&_Deviants.TransactOpts, _maxSupply)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _newPrice) returns()
func (_Deviants *DeviantsTransactor) SetMintPrice(opts *bind.TransactOpts, _newPrice *big.Int) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "setMintPrice", _newPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _newPrice) returns()
func (_Deviants *DeviantsSession) SetMintPrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.SetMintPrice(&_Deviants.TransactOpts, _newPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _newPrice) returns()
func (_Deviants *DeviantsTransactorSession) SetMintPrice(_newPrice *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.SetMintPrice(&_Deviants.TransactOpts, _newPrice)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0xc21b471b.
//
// Solidity: function setRoyalties(address _recipient, uint96 _value) returns()
func (_Deviants *DeviantsTransactor) SetRoyalties(opts *bind.TransactOpts, _recipient common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "setRoyalties", _recipient, _value)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0xc21b471b.
//
// Solidity: function setRoyalties(address _recipient, uint96 _value) returns()
func (_Deviants *DeviantsSession) SetRoyalties(_recipient common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.SetRoyalties(&_Deviants.TransactOpts, _recipient, _value)
}

// SetRoyalties is a paid mutator transaction binding the contract method 0xc21b471b.
//
// Solidity: function setRoyalties(address _recipient, uint96 _value) returns()
func (_Deviants *DeviantsTransactorSession) SetRoyalties(_recipient common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.SetRoyalties(&_Deviants.TransactOpts, _recipient, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Deviants *DeviantsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Deviants *DeviantsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.TransferFrom(&_Deviants.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Deviants *DeviantsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Deviants.Contract.TransferFrom(&_Deviants.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Deviants *DeviantsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Deviants.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Deviants *DeviantsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Deviants.Contract.TransferOwnership(&_Deviants.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Deviants *DeviantsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Deviants.Contract.TransferOwnership(&_Deviants.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Deviants *DeviantsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deviants.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Deviants *DeviantsSession) Receive() (*types.Transaction, error) {
	return _Deviants.Contract.Receive(&_Deviants.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Deviants *DeviantsTransactorSession) Receive() (*types.Transaction, error) {
	return _Deviants.Contract.Receive(&_Deviants.TransactOpts)
}

// DeviantsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Deviants contract.
type DeviantsApprovalIterator struct {
	Event *DeviantsApproval // Event containing the contract specifics and raw log

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
func (it *DeviantsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeviantsApproval)
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
		it.Event = new(DeviantsApproval)
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
func (it *DeviantsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeviantsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeviantsApproval represents a Approval event raised by the Deviants contract.
type DeviantsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Deviants *DeviantsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DeviantsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Deviants.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DeviantsApprovalIterator{contract: _Deviants.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Deviants *DeviantsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DeviantsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Deviants.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeviantsApproval)
				if err := _Deviants.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Deviants *DeviantsFilterer) ParseApproval(log types.Log) (*DeviantsApproval, error) {
	event := new(DeviantsApproval)
	if err := _Deviants.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeviantsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Deviants contract.
type DeviantsApprovalForAllIterator struct {
	Event *DeviantsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DeviantsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeviantsApprovalForAll)
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
		it.Event = new(DeviantsApprovalForAll)
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
func (it *DeviantsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeviantsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeviantsApprovalForAll represents a ApprovalForAll event raised by the Deviants contract.
type DeviantsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Deviants *DeviantsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DeviantsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Deviants.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DeviantsApprovalForAllIterator{contract: _Deviants.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Deviants *DeviantsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DeviantsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Deviants.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeviantsApprovalForAll)
				if err := _Deviants.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Deviants *DeviantsFilterer) ParseApprovalForAll(log types.Log) (*DeviantsApprovalForAll, error) {
	event := new(DeviantsApprovalForAll)
	if err := _Deviants.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeviantsConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Deviants contract.
type DeviantsConsecutiveTransferIterator struct {
	Event *DeviantsConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *DeviantsConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeviantsConsecutiveTransfer)
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
		it.Event = new(DeviantsConsecutiveTransfer)
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
func (it *DeviantsConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeviantsConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeviantsConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Deviants contract.
type DeviantsConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Deviants *DeviantsFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*DeviantsConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Deviants.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DeviantsConsecutiveTransferIterator{contract: _Deviants.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Deviants *DeviantsFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *DeviantsConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Deviants.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeviantsConsecutiveTransfer)
				if err := _Deviants.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Deviants *DeviantsFilterer) ParseConsecutiveTransfer(log types.Log) (*DeviantsConsecutiveTransfer, error) {
	event := new(DeviantsConsecutiveTransfer)
	if err := _Deviants.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeviantsMaxSupplyChangedIterator is returned from FilterMaxSupplyChanged and is used to iterate over the raw logs and unpacked data for MaxSupplyChanged events raised by the Deviants contract.
type DeviantsMaxSupplyChangedIterator struct {
	Event *DeviantsMaxSupplyChanged // Event containing the contract specifics and raw log

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
func (it *DeviantsMaxSupplyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeviantsMaxSupplyChanged)
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
		it.Event = new(DeviantsMaxSupplyChanged)
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
func (it *DeviantsMaxSupplyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeviantsMaxSupplyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeviantsMaxSupplyChanged represents a MaxSupplyChanged event raised by the Deviants contract.
type DeviantsMaxSupplyChanged struct {
	NewMaxSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyChanged is a free log retrieval operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_Deviants *DeviantsFilterer) FilterMaxSupplyChanged(opts *bind.FilterOpts) (*DeviantsMaxSupplyChangedIterator, error) {

	logs, sub, err := _Deviants.contract.FilterLogs(opts, "MaxSupplyChanged")
	if err != nil {
		return nil, err
	}
	return &DeviantsMaxSupplyChangedIterator{contract: _Deviants.contract, event: "MaxSupplyChanged", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyChanged is a free log subscription operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_Deviants *DeviantsFilterer) WatchMaxSupplyChanged(opts *bind.WatchOpts, sink chan<- *DeviantsMaxSupplyChanged) (event.Subscription, error) {

	logs, sub, err := _Deviants.contract.WatchLogs(opts, "MaxSupplyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeviantsMaxSupplyChanged)
				if err := _Deviants.contract.UnpackLog(event, "MaxSupplyChanged", log); err != nil {
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

// ParseMaxSupplyChanged is a log parse operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_Deviants *DeviantsFilterer) ParseMaxSupplyChanged(log types.Log) (*DeviantsMaxSupplyChanged, error) {
	event := new(DeviantsMaxSupplyChanged)
	if err := _Deviants.contract.UnpackLog(event, "MaxSupplyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeviantsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Deviants contract.
type DeviantsOwnershipTransferredIterator struct {
	Event *DeviantsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DeviantsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeviantsOwnershipTransferred)
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
		it.Event = new(DeviantsOwnershipTransferred)
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
func (it *DeviantsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeviantsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeviantsOwnershipTransferred represents a OwnershipTransferred event raised by the Deviants contract.
type DeviantsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Deviants *DeviantsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DeviantsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Deviants.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DeviantsOwnershipTransferredIterator{contract: _Deviants.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Deviants *DeviantsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DeviantsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Deviants.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeviantsOwnershipTransferred)
				if err := _Deviants.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Deviants *DeviantsFilterer) ParseOwnershipTransferred(log types.Log) (*DeviantsOwnershipTransferred, error) {
	event := new(DeviantsOwnershipTransferred)
	if err := _Deviants.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeviantsRoyaltiesChangedIterator is returned from FilterRoyaltiesChanged and is used to iterate over the raw logs and unpacked data for RoyaltiesChanged events raised by the Deviants contract.
type DeviantsRoyaltiesChangedIterator struct {
	Event *DeviantsRoyaltiesChanged // Event containing the contract specifics and raw log

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
func (it *DeviantsRoyaltiesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeviantsRoyaltiesChanged)
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
		it.Event = new(DeviantsRoyaltiesChanged)
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
func (it *DeviantsRoyaltiesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeviantsRoyaltiesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeviantsRoyaltiesChanged represents a RoyaltiesChanged event raised by the Deviants contract.
type DeviantsRoyaltiesChanged struct {
	NewReciver common.Address
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoyaltiesChanged is a free log retrieval operation binding the contract event 0x111989d7f0e41872bbd4cf4d94f4125711bfde242cfd6caca6658064ad6c23d1.
//
// Solidity: event RoyaltiesChanged(address newReciver, uint96 value)
func (_Deviants *DeviantsFilterer) FilterRoyaltiesChanged(opts *bind.FilterOpts) (*DeviantsRoyaltiesChangedIterator, error) {

	logs, sub, err := _Deviants.contract.FilterLogs(opts, "RoyaltiesChanged")
	if err != nil {
		return nil, err
	}
	return &DeviantsRoyaltiesChangedIterator{contract: _Deviants.contract, event: "RoyaltiesChanged", logs: logs, sub: sub}, nil
}

// WatchRoyaltiesChanged is a free log subscription operation binding the contract event 0x111989d7f0e41872bbd4cf4d94f4125711bfde242cfd6caca6658064ad6c23d1.
//
// Solidity: event RoyaltiesChanged(address newReciver, uint96 value)
func (_Deviants *DeviantsFilterer) WatchRoyaltiesChanged(opts *bind.WatchOpts, sink chan<- *DeviantsRoyaltiesChanged) (event.Subscription, error) {

	logs, sub, err := _Deviants.contract.WatchLogs(opts, "RoyaltiesChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeviantsRoyaltiesChanged)
				if err := _Deviants.contract.UnpackLog(event, "RoyaltiesChanged", log); err != nil {
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

// ParseRoyaltiesChanged is a log parse operation binding the contract event 0x111989d7f0e41872bbd4cf4d94f4125711bfde242cfd6caca6658064ad6c23d1.
//
// Solidity: event RoyaltiesChanged(address newReciver, uint96 value)
func (_Deviants *DeviantsFilterer) ParseRoyaltiesChanged(log types.Log) (*DeviantsRoyaltiesChanged, error) {
	event := new(DeviantsRoyaltiesChanged)
	if err := _Deviants.contract.UnpackLog(event, "RoyaltiesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeviantsTokenMintedIterator is returned from FilterTokenMinted and is used to iterate over the raw logs and unpacked data for TokenMinted events raised by the Deviants contract.
type DeviantsTokenMintedIterator struct {
	Event *DeviantsTokenMinted // Event containing the contract specifics and raw log

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
func (it *DeviantsTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeviantsTokenMinted)
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
		it.Event = new(DeviantsTokenMinted)
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
func (it *DeviantsTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeviantsTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeviantsTokenMinted represents a TokenMinted event raised by the Deviants contract.
type DeviantsTokenMinted struct {
	Id           *big.Int
	To           common.Address
	UsedDiscount bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenMinted is a free log retrieval operation binding the contract event 0x5b4944f193638a626a944a372995a6f80b4b85736a75a124dd90fc3948249a4b.
//
// Solidity: event TokenMinted(uint256 id, address indexed to, bool indexed usedDiscount)
func (_Deviants *DeviantsFilterer) FilterTokenMinted(opts *bind.FilterOpts, to []common.Address, usedDiscount []bool) (*DeviantsTokenMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var usedDiscountRule []interface{}
	for _, usedDiscountItem := range usedDiscount {
		usedDiscountRule = append(usedDiscountRule, usedDiscountItem)
	}

	logs, sub, err := _Deviants.contract.FilterLogs(opts, "TokenMinted", toRule, usedDiscountRule)
	if err != nil {
		return nil, err
	}
	return &DeviantsTokenMintedIterator{contract: _Deviants.contract, event: "TokenMinted", logs: logs, sub: sub}, nil
}

// WatchTokenMinted is a free log subscription operation binding the contract event 0x5b4944f193638a626a944a372995a6f80b4b85736a75a124dd90fc3948249a4b.
//
// Solidity: event TokenMinted(uint256 id, address indexed to, bool indexed usedDiscount)
func (_Deviants *DeviantsFilterer) WatchTokenMinted(opts *bind.WatchOpts, sink chan<- *DeviantsTokenMinted, to []common.Address, usedDiscount []bool) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var usedDiscountRule []interface{}
	for _, usedDiscountItem := range usedDiscount {
		usedDiscountRule = append(usedDiscountRule, usedDiscountItem)
	}

	logs, sub, err := _Deviants.contract.WatchLogs(opts, "TokenMinted", toRule, usedDiscountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeviantsTokenMinted)
				if err := _Deviants.contract.UnpackLog(event, "TokenMinted", log); err != nil {
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

// ParseTokenMinted is a log parse operation binding the contract event 0x5b4944f193638a626a944a372995a6f80b4b85736a75a124dd90fc3948249a4b.
//
// Solidity: event TokenMinted(uint256 id, address indexed to, bool indexed usedDiscount)
func (_Deviants *DeviantsFilterer) ParseTokenMinted(log types.Log) (*DeviantsTokenMinted, error) {
	event := new(DeviantsTokenMinted)
	if err := _Deviants.contract.UnpackLog(event, "TokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeviantsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Deviants contract.
type DeviantsTransferIterator struct {
	Event *DeviantsTransfer // Event containing the contract specifics and raw log

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
func (it *DeviantsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeviantsTransfer)
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
		it.Event = new(DeviantsTransfer)
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
func (it *DeviantsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeviantsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeviantsTransfer represents a Transfer event raised by the Deviants contract.
type DeviantsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Deviants *DeviantsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DeviantsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Deviants.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DeviantsTransferIterator{contract: _Deviants.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Deviants *DeviantsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DeviantsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Deviants.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeviantsTransfer)
				if err := _Deviants.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Deviants *DeviantsFilterer) ParseTransfer(log types.Log) (*DeviantsTransfer, error) {
	event := new(DeviantsTransfer)
	if err := _Deviants.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
