// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Auction6ABI is the input ABI used to generate the binding from.
const Auction6ABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_increment\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"FirstBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entry\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"increment\",\"type\":\"uint256\"}],\"name\":\"IncreasedBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lastAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"lPos\",\"type\":\"uint16\"}],\"name\":\"ToBeContinued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"VerificationFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Verified\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"sequence\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"reset\",\"type\":\"bool\"}],\"name\":\"newWinningSequence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfEntries\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pos\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"verifySequence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"sequence\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"winningSequence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Auction6Bin is the compiled bytecode used for deploying new contracts.
var Auction6Bin = "0x608060405234801561001057600080fd5b50604051610f01380380610f018339818101604052608081101561003357600080fd5b50805160208201516040830151606090930151600192909255600255600391909155600455610e9a806100676000396000f3fe60806040526004361061007b5760003560e01c8063ab6a070c1161004e578063ab6a070c1461018d578063ba414fa61461023f578063bbb82d8914610268578063c56551b61461027d5761007b565b80631998aeef146100805780631ea6f2111461008a578063259ae5161461013c578063825f787214610166575b600080fd5b6100886102a9565b005b34801561009657600080fd5b50610088600480360360408110156100ad57600080fd5b8101906020810181356401000000008111156100c857600080fd5b8201836020820111156100da57600080fd5b803590602001918460208302840111640100000000831117156100fc57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505091359250610515915050565b34801561014857600080fd5b506100886004803603602081101561015f57600080fd5b50356105f5565b34801561017257600080fd5b5061017b610932565b60408051918252519081900360200190f35b34801561019957600080fd5b50610088600480360360408110156101b057600080fd5b8101906020810181356401000000008111156101cb57600080fd5b8201836020820111156101dd57600080fd5b803590602001918460208302840111640100000000831117156101ff57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505050503515159050610938565b34801561024b57600080fd5b50610254610d60565b604080519115158252519081900360200190f35b34801561027457600080fd5b50610254610d69565b34801561028957600080fd5b50610292610d77565b6040805161ffff9092168252519081900360200190f35b60015442101580156102bd57506002544211155b61030e576040805162461bcd60e51b815260206004820152601d60248201527f41756374696f6e206e6f742063757272656e746c792072756e6e696e67000000604482015290519081900360640190fd5b336000908152602081905260409020546001600160501b031680156103ce5760045434101561037a576040805162461bcd60e51b81526020600482015260136024820152721a5b98dc995b595b9d081d1bdbc81cdb585b1b606a1b604482015290519081900360640190fd5b600554604080519182523360208301526001600160501b03831682820152346060830152517ff2bed04e2470b19089271262d4996abf365f47ff4834443ecc2010a703503ac89181900360800190a161046f565b600354341015610425576040805162461bcd60e51b815260206004820152601f60248201527f62696420646f6573206e6f74206d656574207265736572766520707269636500604482015290519081900360640190fd5b600680546001019055600554604080519182523360208301523482820152517f427ec838c38270fde4dcc8c81f02ae764f9fad491853c05f09eaba09a48bbb989181900360600190a15b610477610d88565b50604080516060810182526001600160501b03349390930183168152600580546001810190915563ffffffff908116602080840191825260008486018181523382529181905294909420925183549151945161ffff16600160701b0261ffff60701b1995909316600160501b0263ffffffff60501b199190961669ffffffffffffffffffff1990921691909117169390931791909116919091179055565b600854610100900460ff161561056e576040805162461bcd60e51b81526020600482015260196024820152781cd95c5d595b98d948185b1c9958591e481d995c9a599a5959603a1b604482015290519081900360640190fd5b6006548251146105c5576040805162461bcd60e51b815260206004820152601960248201527f73657175656e6365206c656e67746820696e636f727265637400000000000000604482015290519081900360640190fd5b81516105d8906009906020850190610da8565b5060006007556008805460ff191690556105f1816105f5565b5050565b600854600090610100900460ff1615610651576040805162461bcd60e51b81526020600482015260196024820152781cd95c5d595b98d948185b1c9958591e481d995c9a599a5959603a1b604482015290519081900360640190fd5b60085460ff16156106935760405162461bcd60e51b8152600401808060200182810382526022815260200180610e436022913960400191505060405180910390fd5b600654600754106106eb576040805162461bcd60e51b815260206004820152601860248201527f746869732073686f756c64206e657665722068617070656e0000000000000000604482015290519081900360640190fd5b6007546106f6576001015b6106fe610d88565b600080600960018560075401038154811061071557fe5b6000918252602080832091909101546001600160a01b031683528281019390935260409182019020815160608101835290546001600160501b0381168252600160501b810463ffffffff1693820193909352600160701b90920461ffff169082015260075490915061078b576007805460010190555b815b838110156108f85760065460075414156107cf576008805461ff001916610100179055604051600080516020610e2383398151915290600090a150505061092f565b6107d7610d88565b60078054600181019091556009805460009283929181106107f457fe5b6000918252602080832091909101546001600160a01b031683528281019390935260409182019020815160608101835290546001600160501b03808216808452600160501b830463ffffffff1695840195909552600160701b90910461ffff169282019290925285519093501610806108995750805183516001600160501b0390811691161480156108995750806020015163ffffffff16836020015163ffffffff16115b156108ee576008805460ff19166001179055600754604080516000199092018252517f32629d580208e19f97e5752eef849e102f803999c88aa7f75e12b1744eecd5a79181900360200190a15050505061092f565b915060010161078d565b50600654600754141561092c576008805461ff001916610100179055604051600080516020610e2383398151915290600090a15b50505b50565b60065481565b6000811561098c578260008151811061094d57fe5b60209081029190910101516008805461ffff60b01b196001600160a01b03909316620100000262010000600160b01b0319909116179190911690555060015b600854600160b01b900461ffff166109a2610d88565b506008546201000090046001600160a01b031660009081526020818152604091829020825160608101845290546001600160501b0381168252600160501b810463ffffffff1692820192909252600160701b90910461ffff1691810191909152825b8551811015610c66576001600654038361ffff161415610a6c576008805461ff001916610100179055604051600080516020610e2383398151915290600090a150506008805461ffff909216600160b01b0261ffff60b01b19909216919091179055506105f1565b610a74610d88565b600080888481518110610a8357fe5b6020908102919091018101516001600160a01b03168252818101929092526040908101600020815160608101835290546001600160501b03808216808452600160501b830463ffffffff1695840195909552600160701b90910461ffff16928201929092528551909350161080610b265750805183516001600160501b039081169116148015610b265750806020015163ffffffff16836020015163ffffffff16115b80610b39575080516001600160501b0316155b15610bad576008805460ff191660011790556040805161ffff8616815290517f32629d580208e19f97e5752eef849e102f803999c88aa7f75e12b1744eecd5a79181900360200190a150506008805461ffff909316600160b01b0261ffff60b01b1990931692909217909155506105f19050565b61ffff600190940193841660408201528651819060009081908a9086908110610bd257fe5b6020908102919091018101516001600160a01b031682528181019290925260409081016000208351815493850151949092015169ffffffffffffffffffff199093166001600160501b039092169190911763ffffffff60501b1916600160501b63ffffffff909416939093029290921761ffff60701b1916600160701b61ffff909216919091021790559150600101610a04565b506001600654038261ffff161415610cbc576008805461010061ff00199091161761ffff60b01b1916600160b01b61ffff851602179055604051600080516020610e2383398151915290600090a15050506105f1565b84600186510381518110610ccc57fe5b6020908102919091018101516008805461ffff8616600160b01b810261ffff60b01b196001600160a01b039586166201000090810262010000600160b01b031990951694909417161792839055604080519290930490931681529283019190915280517f5d970fe28f8d9fc1642508f654459830ceeb41371644f300106bb92dc7c28b9e9281900390910190a15050505050565b60085460ff1681565b600854610100900460ff1681565b600854600160b01b900461ffff1681565b604080516060810182526000808252602082018190529181019190915290565b828054828255906000526020600020908101928215610dfd579160200282015b82811115610dfd57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610dc8565b50610e09929150610e0d565b5090565b5b80821115610e095760008155600101610e0e56fe2e5e1405bcc465f1fe8b5964e79910ef9548bb882e1bae02ca901e855a7a4ff573657175656e636520616c7265616479206661696c65642e2054727920616761696ea26469706673582212207809fa163ba03d71680f28fae67f95b5b53bb70147e434c152a2d848e0d1820b64736f6c63430007030033"

// DeployAuction6 deploys a new Ethereum contract, binding an instance of Auction6 to it.
func DeployAuction6(auth *bind.TransactOpts, backend bind.ContractBackend, _start *big.Int, _end *big.Int, _reserve *big.Int, _increment *big.Int) (common.Address, *types.Transaction, *Auction6, error) {
	parsed, err := abi.JSON(strings.NewReader(Auction6ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Auction6Bin), backend, _start, _end, _reserve, _increment)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Auction6{Auction6Caller: Auction6Caller{contract: contract}, Auction6Transactor: Auction6Transactor{contract: contract}, Auction6Filterer: Auction6Filterer{contract: contract}}, nil
}

// Auction6 is an auto generated Go binding around an Ethereum contract.
type Auction6 struct {
	Auction6Caller     // Read-only binding to the contract
	Auction6Transactor // Write-only binding to the contract
	Auction6Filterer   // Log filterer for contract events
}

// Auction6Caller is an auto generated read-only Go binding around an Ethereum contract.
type Auction6Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Auction6Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Auction6Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Auction6Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Auction6Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Auction6Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Auction6Session struct {
	Contract     *Auction6         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Auction6CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Auction6CallerSession struct {
	Contract *Auction6Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Auction6TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Auction6TransactorSession struct {
	Contract     *Auction6Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Auction6Raw is an auto generated low-level Go binding around an Ethereum contract.
type Auction6Raw struct {
	Contract *Auction6 // Generic contract binding to access the raw methods on
}

// Auction6CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Auction6CallerRaw struct {
	Contract *Auction6Caller // Generic read-only contract binding to access the raw methods on
}

// Auction6TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Auction6TransactorRaw struct {
	Contract *Auction6Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAuction6 creates a new instance of Auction6, bound to a specific deployed contract.
func NewAuction6(address common.Address, backend bind.ContractBackend) (*Auction6, error) {
	contract, err := bindAuction6(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auction6{Auction6Caller: Auction6Caller{contract: contract}, Auction6Transactor: Auction6Transactor{contract: contract}, Auction6Filterer: Auction6Filterer{contract: contract}}, nil
}

// NewAuction6Caller creates a new read-only instance of Auction6, bound to a specific deployed contract.
func NewAuction6Caller(address common.Address, caller bind.ContractCaller) (*Auction6Caller, error) {
	contract, err := bindAuction6(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Auction6Caller{contract: contract}, nil
}

// NewAuction6Transactor creates a new write-only instance of Auction6, bound to a specific deployed contract.
func NewAuction6Transactor(address common.Address, transactor bind.ContractTransactor) (*Auction6Transactor, error) {
	contract, err := bindAuction6(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Auction6Transactor{contract: contract}, nil
}

// NewAuction6Filterer creates a new log filterer instance of Auction6, bound to a specific deployed contract.
func NewAuction6Filterer(address common.Address, filterer bind.ContractFilterer) (*Auction6Filterer, error) {
	contract, err := bindAuction6(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Auction6Filterer{contract: contract}, nil
}

// bindAuction6 binds a generic wrapper to an already deployed contract.
func bindAuction6(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Auction6ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction6 *Auction6Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction6.Contract.Auction6Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction6 *Auction6Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction6.Contract.Auction6Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction6 *Auction6Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction6.Contract.Auction6Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction6 *Auction6CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction6.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction6 *Auction6TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction6.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction6 *Auction6TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction6.Contract.contract.Transact(opts, method, params...)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_Auction6 *Auction6Caller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Auction6.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_Auction6 *Auction6Session) Failed() (bool, error) {
	return _Auction6.Contract.Failed(&_Auction6.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_Auction6 *Auction6CallerSession) Failed() (bool, error) {
	return _Auction6.Contract.Failed(&_Auction6.CallOpts)
}

// NumberOfEntries is a free data retrieval call binding the contract method 0x825f7872.
//
// Solidity: function numberOfEntries() view returns(uint256)
func (_Auction6 *Auction6Caller) NumberOfEntries(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction6.contract.Call(opts, &out, "numberOfEntries")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfEntries is a free data retrieval call binding the contract method 0x825f7872.
//
// Solidity: function numberOfEntries() view returns(uint256)
func (_Auction6 *Auction6Session) NumberOfEntries() (*big.Int, error) {
	return _Auction6.Contract.NumberOfEntries(&_Auction6.CallOpts)
}

// NumberOfEntries is a free data retrieval call binding the contract method 0x825f7872.
//
// Solidity: function numberOfEntries() view returns(uint256)
func (_Auction6 *Auction6CallerSession) NumberOfEntries() (*big.Int, error) {
	return _Auction6.Contract.NumberOfEntries(&_Auction6.CallOpts)
}

// Pos is a free data retrieval call binding the contract method 0xc56551b6.
//
// Solidity: function pos() view returns(uint16)
func (_Auction6 *Auction6Caller) Pos(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Auction6.contract.Call(opts, &out, "pos")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Pos is a free data retrieval call binding the contract method 0xc56551b6.
//
// Solidity: function pos() view returns(uint16)
func (_Auction6 *Auction6Session) Pos() (uint16, error) {
	return _Auction6.Contract.Pos(&_Auction6.CallOpts)
}

// Pos is a free data retrieval call binding the contract method 0xc56551b6.
//
// Solidity: function pos() view returns(uint16)
func (_Auction6 *Auction6CallerSession) Pos() (uint16, error) {
	return _Auction6.Contract.Pos(&_Auction6.CallOpts)
}

// Verified is a free data retrieval call binding the contract method 0xbbb82d89.
//
// Solidity: function verified() view returns(bool)
func (_Auction6 *Auction6Caller) Verified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Auction6.contract.Call(opts, &out, "verified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verified is a free data retrieval call binding the contract method 0xbbb82d89.
//
// Solidity: function verified() view returns(bool)
func (_Auction6 *Auction6Session) Verified() (bool, error) {
	return _Auction6.Contract.Verified(&_Auction6.CallOpts)
}

// Verified is a free data retrieval call binding the contract method 0xbbb82d89.
//
// Solidity: function verified() view returns(bool)
func (_Auction6 *Auction6CallerSession) Verified() (bool, error) {
	return _Auction6.Contract.Verified(&_Auction6.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction6 *Auction6Transactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction6.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction6 *Auction6Session) Bid() (*types.Transaction, error) {
	return _Auction6.Contract.Bid(&_Auction6.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction6 *Auction6TransactorSession) Bid() (*types.Transaction, error) {
	return _Auction6.Contract.Bid(&_Auction6.TransactOpts)
}

// NewWinningSequence is a paid mutator transaction binding the contract method 0xab6a070c.
//
// Solidity: function newWinningSequence(address[] sequence, bool reset) returns()
func (_Auction6 *Auction6Transactor) NewWinningSequence(opts *bind.TransactOpts, sequence []common.Address, reset bool) (*types.Transaction, error) {
	return _Auction6.contract.Transact(opts, "newWinningSequence", sequence, reset)
}

// NewWinningSequence is a paid mutator transaction binding the contract method 0xab6a070c.
//
// Solidity: function newWinningSequence(address[] sequence, bool reset) returns()
func (_Auction6 *Auction6Session) NewWinningSequence(sequence []common.Address, reset bool) (*types.Transaction, error) {
	return _Auction6.Contract.NewWinningSequence(&_Auction6.TransactOpts, sequence, reset)
}

// NewWinningSequence is a paid mutator transaction binding the contract method 0xab6a070c.
//
// Solidity: function newWinningSequence(address[] sequence, bool reset) returns()
func (_Auction6 *Auction6TransactorSession) NewWinningSequence(sequence []common.Address, reset bool) (*types.Transaction, error) {
	return _Auction6.Contract.NewWinningSequence(&_Auction6.TransactOpts, sequence, reset)
}

// VerifySequence is a paid mutator transaction binding the contract method 0x259ae516.
//
// Solidity: function verifySequence(uint256 count) returns()
func (_Auction6 *Auction6Transactor) VerifySequence(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _Auction6.contract.Transact(opts, "verifySequence", count)
}

// VerifySequence is a paid mutator transaction binding the contract method 0x259ae516.
//
// Solidity: function verifySequence(uint256 count) returns()
func (_Auction6 *Auction6Session) VerifySequence(count *big.Int) (*types.Transaction, error) {
	return _Auction6.Contract.VerifySequence(&_Auction6.TransactOpts, count)
}

// VerifySequence is a paid mutator transaction binding the contract method 0x259ae516.
//
// Solidity: function verifySequence(uint256 count) returns()
func (_Auction6 *Auction6TransactorSession) VerifySequence(count *big.Int) (*types.Transaction, error) {
	return _Auction6.Contract.VerifySequence(&_Auction6.TransactOpts, count)
}

// WinningSequence is a paid mutator transaction binding the contract method 0x1ea6f211.
//
// Solidity: function winningSequence(address[] sequence, uint256 count) returns()
func (_Auction6 *Auction6Transactor) WinningSequence(opts *bind.TransactOpts, sequence []common.Address, count *big.Int) (*types.Transaction, error) {
	return _Auction6.contract.Transact(opts, "winningSequence", sequence, count)
}

// WinningSequence is a paid mutator transaction binding the contract method 0x1ea6f211.
//
// Solidity: function winningSequence(address[] sequence, uint256 count) returns()
func (_Auction6 *Auction6Session) WinningSequence(sequence []common.Address, count *big.Int) (*types.Transaction, error) {
	return _Auction6.Contract.WinningSequence(&_Auction6.TransactOpts, sequence, count)
}

// WinningSequence is a paid mutator transaction binding the contract method 0x1ea6f211.
//
// Solidity: function winningSequence(address[] sequence, uint256 count) returns()
func (_Auction6 *Auction6TransactorSession) WinningSequence(sequence []common.Address, count *big.Int) (*types.Transaction, error) {
	return _Auction6.Contract.WinningSequence(&_Auction6.TransactOpts, sequence, count)
}

// Auction6FirstBidIterator is returned from FilterFirstBid and is used to iterate over the raw logs and unpacked data for FirstBid events raised by the Auction6 contract.
type Auction6FirstBidIterator struct {
	Event *Auction6FirstBid // Event containing the contract specifics and raw log

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
func (it *Auction6FirstBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Auction6FirstBid)
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
		it.Event = new(Auction6FirstBid)
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
func (it *Auction6FirstBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Auction6FirstBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Auction6FirstBid represents a FirstBid event raised by the Auction6 contract.
type Auction6FirstBid struct {
	Entry  *big.Int
	Bidder common.Address
	Bid    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFirstBid is a free log retrieval operation binding the contract event 0x427ec838c38270fde4dcc8c81f02ae764f9fad491853c05f09eaba09a48bbb98.
//
// Solidity: event FirstBid(uint256 entry, address bidder, uint256 bid)
func (_Auction6 *Auction6Filterer) FilterFirstBid(opts *bind.FilterOpts) (*Auction6FirstBidIterator, error) {

	logs, sub, err := _Auction6.contract.FilterLogs(opts, "FirstBid")
	if err != nil {
		return nil, err
	}
	return &Auction6FirstBidIterator{contract: _Auction6.contract, event: "FirstBid", logs: logs, sub: sub}, nil
}

// WatchFirstBid is a free log subscription operation binding the contract event 0x427ec838c38270fde4dcc8c81f02ae764f9fad491853c05f09eaba09a48bbb98.
//
// Solidity: event FirstBid(uint256 entry, address bidder, uint256 bid)
func (_Auction6 *Auction6Filterer) WatchFirstBid(opts *bind.WatchOpts, sink chan<- *Auction6FirstBid) (event.Subscription, error) {

	logs, sub, err := _Auction6.contract.WatchLogs(opts, "FirstBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Auction6FirstBid)
				if err := _Auction6.contract.UnpackLog(event, "FirstBid", log); err != nil {
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

// ParseFirstBid is a log parse operation binding the contract event 0x427ec838c38270fde4dcc8c81f02ae764f9fad491853c05f09eaba09a48bbb98.
//
// Solidity: event FirstBid(uint256 entry, address bidder, uint256 bid)
func (_Auction6 *Auction6Filterer) ParseFirstBid(log types.Log) (*Auction6FirstBid, error) {
	event := new(Auction6FirstBid)
	if err := _Auction6.contract.UnpackLog(event, "FirstBid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Auction6IncreasedBidIterator is returned from FilterIncreasedBid and is used to iterate over the raw logs and unpacked data for IncreasedBid events raised by the Auction6 contract.
type Auction6IncreasedBidIterator struct {
	Event *Auction6IncreasedBid // Event containing the contract specifics and raw log

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
func (it *Auction6IncreasedBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Auction6IncreasedBid)
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
		it.Event = new(Auction6IncreasedBid)
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
func (it *Auction6IncreasedBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Auction6IncreasedBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Auction6IncreasedBid represents a IncreasedBid event raised by the Auction6 contract.
type Auction6IncreasedBid struct {
	Entry     *big.Int
	Bidder    common.Address
	OldBid    *big.Int
	Increment *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreasedBid is a free log retrieval operation binding the contract event 0xf2bed04e2470b19089271262d4996abf365f47ff4834443ecc2010a703503ac8.
//
// Solidity: event IncreasedBid(uint256 entry, address bidder, uint256 oldBid, uint256 increment)
func (_Auction6 *Auction6Filterer) FilterIncreasedBid(opts *bind.FilterOpts) (*Auction6IncreasedBidIterator, error) {

	logs, sub, err := _Auction6.contract.FilterLogs(opts, "IncreasedBid")
	if err != nil {
		return nil, err
	}
	return &Auction6IncreasedBidIterator{contract: _Auction6.contract, event: "IncreasedBid", logs: logs, sub: sub}, nil
}

// WatchIncreasedBid is a free log subscription operation binding the contract event 0xf2bed04e2470b19089271262d4996abf365f47ff4834443ecc2010a703503ac8.
//
// Solidity: event IncreasedBid(uint256 entry, address bidder, uint256 oldBid, uint256 increment)
func (_Auction6 *Auction6Filterer) WatchIncreasedBid(opts *bind.WatchOpts, sink chan<- *Auction6IncreasedBid) (event.Subscription, error) {

	logs, sub, err := _Auction6.contract.WatchLogs(opts, "IncreasedBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Auction6IncreasedBid)
				if err := _Auction6.contract.UnpackLog(event, "IncreasedBid", log); err != nil {
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

// ParseIncreasedBid is a log parse operation binding the contract event 0xf2bed04e2470b19089271262d4996abf365f47ff4834443ecc2010a703503ac8.
//
// Solidity: event IncreasedBid(uint256 entry, address bidder, uint256 oldBid, uint256 increment)
func (_Auction6 *Auction6Filterer) ParseIncreasedBid(log types.Log) (*Auction6IncreasedBid, error) {
	event := new(Auction6IncreasedBid)
	if err := _Auction6.contract.UnpackLog(event, "IncreasedBid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Auction6ToBeContinuedIterator is returned from FilterToBeContinued and is used to iterate over the raw logs and unpacked data for ToBeContinued events raised by the Auction6 contract.
type Auction6ToBeContinuedIterator struct {
	Event *Auction6ToBeContinued // Event containing the contract specifics and raw log

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
func (it *Auction6ToBeContinuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Auction6ToBeContinued)
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
		it.Event = new(Auction6ToBeContinued)
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
func (it *Auction6ToBeContinuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Auction6ToBeContinuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Auction6ToBeContinued represents a ToBeContinued event raised by the Auction6 contract.
type Auction6ToBeContinued struct {
	LastAddress common.Address
	LPos        uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterToBeContinued is a free log retrieval operation binding the contract event 0x5d970fe28f8d9fc1642508f654459830ceeb41371644f300106bb92dc7c28b9e.
//
// Solidity: event ToBeContinued(address lastAddress, uint16 lPos)
func (_Auction6 *Auction6Filterer) FilterToBeContinued(opts *bind.FilterOpts) (*Auction6ToBeContinuedIterator, error) {

	logs, sub, err := _Auction6.contract.FilterLogs(opts, "ToBeContinued")
	if err != nil {
		return nil, err
	}
	return &Auction6ToBeContinuedIterator{contract: _Auction6.contract, event: "ToBeContinued", logs: logs, sub: sub}, nil
}

// WatchToBeContinued is a free log subscription operation binding the contract event 0x5d970fe28f8d9fc1642508f654459830ceeb41371644f300106bb92dc7c28b9e.
//
// Solidity: event ToBeContinued(address lastAddress, uint16 lPos)
func (_Auction6 *Auction6Filterer) WatchToBeContinued(opts *bind.WatchOpts, sink chan<- *Auction6ToBeContinued) (event.Subscription, error) {

	logs, sub, err := _Auction6.contract.WatchLogs(opts, "ToBeContinued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Auction6ToBeContinued)
				if err := _Auction6.contract.UnpackLog(event, "ToBeContinued", log); err != nil {
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

// ParseToBeContinued is a log parse operation binding the contract event 0x5d970fe28f8d9fc1642508f654459830ceeb41371644f300106bb92dc7c28b9e.
//
// Solidity: event ToBeContinued(address lastAddress, uint16 lPos)
func (_Auction6 *Auction6Filterer) ParseToBeContinued(log types.Log) (*Auction6ToBeContinued, error) {
	event := new(Auction6ToBeContinued)
	if err := _Auction6.contract.UnpackLog(event, "ToBeContinued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Auction6VerificationFailedIterator is returned from FilterVerificationFailed and is used to iterate over the raw logs and unpacked data for VerificationFailed events raised by the Auction6 contract.
type Auction6VerificationFailedIterator struct {
	Event *Auction6VerificationFailed // Event containing the contract specifics and raw log

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
func (it *Auction6VerificationFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Auction6VerificationFailed)
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
		it.Event = new(Auction6VerificationFailed)
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
func (it *Auction6VerificationFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Auction6VerificationFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Auction6VerificationFailed represents a VerificationFailed event raised by the Auction6 contract.
type Auction6VerificationFailed struct {
	Position *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerificationFailed is a free log retrieval operation binding the contract event 0x32629d580208e19f97e5752eef849e102f803999c88aa7f75e12b1744eecd5a7.
//
// Solidity: event VerificationFailed(uint256 position)
func (_Auction6 *Auction6Filterer) FilterVerificationFailed(opts *bind.FilterOpts) (*Auction6VerificationFailedIterator, error) {

	logs, sub, err := _Auction6.contract.FilterLogs(opts, "VerificationFailed")
	if err != nil {
		return nil, err
	}
	return &Auction6VerificationFailedIterator{contract: _Auction6.contract, event: "VerificationFailed", logs: logs, sub: sub}, nil
}

// WatchVerificationFailed is a free log subscription operation binding the contract event 0x32629d580208e19f97e5752eef849e102f803999c88aa7f75e12b1744eecd5a7.
//
// Solidity: event VerificationFailed(uint256 position)
func (_Auction6 *Auction6Filterer) WatchVerificationFailed(opts *bind.WatchOpts, sink chan<- *Auction6VerificationFailed) (event.Subscription, error) {

	logs, sub, err := _Auction6.contract.WatchLogs(opts, "VerificationFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Auction6VerificationFailed)
				if err := _Auction6.contract.UnpackLog(event, "VerificationFailed", log); err != nil {
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

// ParseVerificationFailed is a log parse operation binding the contract event 0x32629d580208e19f97e5752eef849e102f803999c88aa7f75e12b1744eecd5a7.
//
// Solidity: event VerificationFailed(uint256 position)
func (_Auction6 *Auction6Filterer) ParseVerificationFailed(log types.Log) (*Auction6VerificationFailed, error) {
	event := new(Auction6VerificationFailed)
	if err := _Auction6.contract.UnpackLog(event, "VerificationFailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Auction6VerifiedIterator is returned from FilterVerified and is used to iterate over the raw logs and unpacked data for Verified events raised by the Auction6 contract.
type Auction6VerifiedIterator struct {
	Event *Auction6Verified // Event containing the contract specifics and raw log

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
func (it *Auction6VerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Auction6Verified)
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
		it.Event = new(Auction6Verified)
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
func (it *Auction6VerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Auction6VerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Auction6Verified represents a Verified event raised by the Auction6 contract.
type Auction6Verified struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVerified is a free log retrieval operation binding the contract event 0x2e5e1405bcc465f1fe8b5964e79910ef9548bb882e1bae02ca901e855a7a4ff5.
//
// Solidity: event Verified()
func (_Auction6 *Auction6Filterer) FilterVerified(opts *bind.FilterOpts) (*Auction6VerifiedIterator, error) {

	logs, sub, err := _Auction6.contract.FilterLogs(opts, "Verified")
	if err != nil {
		return nil, err
	}
	return &Auction6VerifiedIterator{contract: _Auction6.contract, event: "Verified", logs: logs, sub: sub}, nil
}

// WatchVerified is a free log subscription operation binding the contract event 0x2e5e1405bcc465f1fe8b5964e79910ef9548bb882e1bae02ca901e855a7a4ff5.
//
// Solidity: event Verified()
func (_Auction6 *Auction6Filterer) WatchVerified(opts *bind.WatchOpts, sink chan<- *Auction6Verified) (event.Subscription, error) {

	logs, sub, err := _Auction6.contract.WatchLogs(opts, "Verified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Auction6Verified)
				if err := _Auction6.contract.UnpackLog(event, "Verified", log); err != nil {
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

// ParseVerified is a log parse operation binding the contract event 0x2e5e1405bcc465f1fe8b5964e79910ef9548bb882e1bae02ca901e855a7a4ff5.
//
// Solidity: event Verified()
func (_Auction6 *Auction6Filterer) ParseVerified(log types.Log) (*Auction6Verified, error) {
	event := new(Auction6Verified)
	if err := _Auction6.contract.UnpackLog(event, "Verified", log); err != nil {
		return nil, err
	}
	return event, nil
}
