// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartcontracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
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

// ScrollChainMetaData contains all meta data concerning the ScrollChain contract.
var ScrollChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CommittedProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommittedProofHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommittedTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCommitProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ErrorBatchHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientPledge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SlotAdapterEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmitProofEarly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmitProofTooLate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnCommittedProofHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"submitProofHashNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"submitProofNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"CommitBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"}],\"name\":\"FinalizeBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"RevertBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newProofCommitEpoch\",\"type\":\"uint8\"}],\"name\":\"SetProofCommitEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newProofHashCommitEpoch\",\"type\":\"uint8\"}],\"name\":\"SetProofHashCommitEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_proofHash\",\"type\":\"bytes32\"}],\"name\":\"SubmitProofHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxNumTxInChunk\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxNumTxInChunk\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxNumTxInChunk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateProver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"UpdateVerifier\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blockCommitBatches\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_version\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_parentBatchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_chunks\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"_skippedL1MessageBitmap\",\"type\":\"bytes\"}],\"name\":\"commitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"committedBatchInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"proofSubmitted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"committedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_withdrawRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_aggrProof\",\"type\":\"bytes\"}],\"name\":\"finalizeBatchWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"finalizedStateRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchToProve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ideDeposit\",\"outputs\":[{\"internalType\":\"contractIDeposit\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_stateRoot\",\"type\":\"bytes32\"}],\"name\":\"importGenesisBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incorrectProofHashPunishAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageQueue\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxNumTxInChunk\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAllLiquidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"}],\"name\":\"isBatchFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"isCommitProofAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"isCommitProofHashAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isProver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFinalizedBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"layer2ChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumTxInChunk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noProofPunishAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proofCommitEpoch\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proofHashCommitEpoch\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proofNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proverCommitProofHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"proofHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"proof\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proverLastLiquidated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proverLiquidation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSubmittedProofHash\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submitHashBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSubmittedProof\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"submitProofBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLiquidated\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proverPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"revertBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDeposit\",\"name\":\"_ideDeposit\",\"type\":\"address\"}],\"name\":\"setDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setIncorrectProofPunishAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setMinDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setNoProofPunishAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_newCommitEpoch\",\"type\":\"uint8\"}],\"name\":\"setProofCommitEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_newCommitEpoch\",\"type\":\"uint8\"}],\"name\":\"setProofHashCommitEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISlotAdapter\",\"name\":\"_slotAdapter\",\"type\":\"address\"}],\"name\":\"setSlotAdapter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slotAdapter\",\"outputs\":[{\"internalType\":\"contractISlotAdapter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_proofHash\",\"type\":\"bytes32\"}],\"name\":\"submitProofHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxNumTxInChunk\",\"type\":\"uint256\"}],\"name\":\"updateMaxNumTxInChunk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ScrollChainABI is the input ABI used to generate the binding from.
// Deprecated: Use ScrollChainMetaData.ABI instead.
var ScrollChainABI = ScrollChainMetaData.ABI

// ScrollChain is an auto generated Go binding around an Ethereum contract.
type ScrollChain struct {
	ScrollChainCaller     // Read-only binding to the contract
	ScrollChainTransactor // Write-only binding to the contract
	ScrollChainFilterer   // Log filterer for contract events
}

// ScrollChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScrollChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrollChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScrollChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrollChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScrollChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrollChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScrollChainSession struct {
	Contract     *ScrollChain      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScrollChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScrollChainCallerSession struct {
	Contract *ScrollChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ScrollChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScrollChainTransactorSession struct {
	Contract     *ScrollChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ScrollChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScrollChainRaw struct {
	Contract *ScrollChain // Generic contract binding to access the raw methods on
}

// ScrollChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScrollChainCallerRaw struct {
	Contract *ScrollChainCaller // Generic read-only contract binding to access the raw methods on
}

// ScrollChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScrollChainTransactorRaw struct {
	Contract *ScrollChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScrollChain creates a new instance of ScrollChain, bound to a specific deployed contract.
func NewScrollChain(address common.Address, backend bind.ContractBackend) (*ScrollChain, error) {
	contract, err := bindScrollChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScrollChain{ScrollChainCaller: ScrollChainCaller{contract: contract}, ScrollChainTransactor: ScrollChainTransactor{contract: contract}, ScrollChainFilterer: ScrollChainFilterer{contract: contract}}, nil
}

// NewScrollChainCaller creates a new read-only instance of ScrollChain, bound to a specific deployed contract.
func NewScrollChainCaller(address common.Address, caller bind.ContractCaller) (*ScrollChainCaller, error) {
	contract, err := bindScrollChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollChainCaller{contract: contract}, nil
}

// NewScrollChainTransactor creates a new write-only instance of ScrollChain, bound to a specific deployed contract.
func NewScrollChainTransactor(address common.Address, transactor bind.ContractTransactor) (*ScrollChainTransactor, error) {
	contract, err := bindScrollChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollChainTransactor{contract: contract}, nil
}

// NewScrollChainFilterer creates a new log filterer instance of ScrollChain, bound to a specific deployed contract.
func NewScrollChainFilterer(address common.Address, filterer bind.ContractFilterer) (*ScrollChainFilterer, error) {
	contract, err := bindScrollChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScrollChainFilterer{contract: contract}, nil
}

// bindScrollChain binds a generic wrapper to an already deployed contract.
func bindScrollChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScrollChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScrollChain *ScrollChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollChain.Contract.ScrollChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScrollChain *ScrollChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollChain.Contract.ScrollChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScrollChain *ScrollChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollChain.Contract.ScrollChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScrollChain *ScrollChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScrollChain *ScrollChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScrollChain *ScrollChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollChain.Contract.contract.Transact(opts, method, params...)
}

// BlockCommitBatches is a free data retrieval call binding the contract method 0x97fb2863.
//
// Solidity: function blockCommitBatches(uint256 ) view returns(bool)
func (_ScrollChain *ScrollChainCaller) BlockCommitBatches(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "blockCommitBatches", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlockCommitBatches is a free data retrieval call binding the contract method 0x97fb2863.
//
// Solidity: function blockCommitBatches(uint256 ) view returns(bool)
func (_ScrollChain *ScrollChainSession) BlockCommitBatches(arg0 *big.Int) (bool, error) {
	return _ScrollChain.Contract.BlockCommitBatches(&_ScrollChain.CallOpts, arg0)
}

// BlockCommitBatches is a free data retrieval call binding the contract method 0x97fb2863.
//
// Solidity: function blockCommitBatches(uint256 ) view returns(bool)
func (_ScrollChain *ScrollChainCallerSession) BlockCommitBatches(arg0 *big.Int) (bool, error) {
	return _ScrollChain.Contract.BlockCommitBatches(&_ScrollChain.CallOpts, arg0)
}

// CommittedBatchInfo is a free data retrieval call binding the contract method 0x4a92882d.
//
// Solidity: function committedBatchInfo(uint256 ) view returns(uint256 blockNumber, bool proofSubmitted)
func (_ScrollChain *ScrollChainCaller) CommittedBatchInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockNumber    *big.Int
	ProofSubmitted bool
}, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "committedBatchInfo", arg0)

	outstruct := new(struct {
		BlockNumber    *big.Int
		ProofSubmitted bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ProofSubmitted = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// CommittedBatchInfo is a free data retrieval call binding the contract method 0x4a92882d.
//
// Solidity: function committedBatchInfo(uint256 ) view returns(uint256 blockNumber, bool proofSubmitted)
func (_ScrollChain *ScrollChainSession) CommittedBatchInfo(arg0 *big.Int) (struct {
	BlockNumber    *big.Int
	ProofSubmitted bool
}, error) {
	return _ScrollChain.Contract.CommittedBatchInfo(&_ScrollChain.CallOpts, arg0)
}

// CommittedBatchInfo is a free data retrieval call binding the contract method 0x4a92882d.
//
// Solidity: function committedBatchInfo(uint256 ) view returns(uint256 blockNumber, bool proofSubmitted)
func (_ScrollChain *ScrollChainCallerSession) CommittedBatchInfo(arg0 *big.Int) (struct {
	BlockNumber    *big.Int
	ProofSubmitted bool
}, error) {
	return _ScrollChain.Contract.CommittedBatchInfo(&_ScrollChain.CallOpts, arg0)
}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 ) view returns(bytes32)
func (_ScrollChain *ScrollChainCaller) CommittedBatches(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "committedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 ) view returns(bytes32)
func (_ScrollChain *ScrollChainSession) CommittedBatches(arg0 *big.Int) ([32]byte, error) {
	return _ScrollChain.Contract.CommittedBatches(&_ScrollChain.CallOpts, arg0)
}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 ) view returns(bytes32)
func (_ScrollChain *ScrollChainCallerSession) CommittedBatches(arg0 *big.Int) ([32]byte, error) {
	return _ScrollChain.Contract.CommittedBatches(&_ScrollChain.CallOpts, arg0)
}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 ) view returns(bytes32)
func (_ScrollChain *ScrollChainCaller) FinalizedStateRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "finalizedStateRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 ) view returns(bytes32)
func (_ScrollChain *ScrollChainSession) FinalizedStateRoots(arg0 *big.Int) ([32]byte, error) {
	return _ScrollChain.Contract.FinalizedStateRoots(&_ScrollChain.CallOpts, arg0)
}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 ) view returns(bytes32)
func (_ScrollChain *ScrollChainCallerSession) FinalizedStateRoots(arg0 *big.Int) ([32]byte, error) {
	return _ScrollChain.Contract.FinalizedStateRoots(&_ScrollChain.CallOpts, arg0)
}

// GetBatchToProve is a free data retrieval call binding the contract method 0x428ea953.
//
// Solidity: function getBatchToProve() view returns(uint256)
func (_ScrollChain *ScrollChainCaller) GetBatchToProve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "getBatchToProve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchToProve is a free data retrieval call binding the contract method 0x428ea953.
//
// Solidity: function getBatchToProve() view returns(uint256)
func (_ScrollChain *ScrollChainSession) GetBatchToProve() (*big.Int, error) {
	return _ScrollChain.Contract.GetBatchToProve(&_ScrollChain.CallOpts)
}

// GetBatchToProve is a free data retrieval call binding the contract method 0x428ea953.
//
// Solidity: function getBatchToProve() view returns(uint256)
func (_ScrollChain *ScrollChainCallerSession) GetBatchToProve() (*big.Int, error) {
	return _ScrollChain.Contract.GetBatchToProve(&_ScrollChain.CallOpts)
}

// IdeDeposit is a free data retrieval call binding the contract method 0x5c22ee9d.
//
// Solidity: function ideDeposit() view returns(address)
func (_ScrollChain *ScrollChainCaller) IdeDeposit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "ideDeposit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdeDeposit is a free data retrieval call binding the contract method 0x5c22ee9d.
//
// Solidity: function ideDeposit() view returns(address)
func (_ScrollChain *ScrollChainSession) IdeDeposit() (common.Address, error) {
	return _ScrollChain.Contract.IdeDeposit(&_ScrollChain.CallOpts)
}

// IdeDeposit is a free data retrieval call binding the contract method 0x5c22ee9d.
//
// Solidity: function ideDeposit() view returns(address)
func (_ScrollChain *ScrollChainCallerSession) IdeDeposit() (common.Address, error) {
	return _ScrollChain.Contract.IdeDeposit(&_ScrollChain.CallOpts)
}

// IncorrectProofHashPunishAmount is a free data retrieval call binding the contract method 0x430073dc.
//
// Solidity: function incorrectProofHashPunishAmount() view returns(uint256)
func (_ScrollChain *ScrollChainCaller) IncorrectProofHashPunishAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "incorrectProofHashPunishAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncorrectProofHashPunishAmount is a free data retrieval call binding the contract method 0x430073dc.
//
// Solidity: function incorrectProofHashPunishAmount() view returns(uint256)
func (_ScrollChain *ScrollChainSession) IncorrectProofHashPunishAmount() (*big.Int, error) {
	return _ScrollChain.Contract.IncorrectProofHashPunishAmount(&_ScrollChain.CallOpts)
}

// IncorrectProofHashPunishAmount is a free data retrieval call binding the contract method 0x430073dc.
//
// Solidity: function incorrectProofHashPunishAmount() view returns(uint256)
func (_ScrollChain *ScrollChainCallerSession) IncorrectProofHashPunishAmount() (*big.Int, error) {
	return _ScrollChain.Contract.IncorrectProofHashPunishAmount(&_ScrollChain.CallOpts)
}

// IsAllLiquidated is a free data retrieval call binding the contract method 0x265b44e1.
//
// Solidity: function isAllLiquidated() view returns(bool)
func (_ScrollChain *ScrollChainCaller) IsAllLiquidated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "isAllLiquidated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllLiquidated is a free data retrieval call binding the contract method 0x265b44e1.
//
// Solidity: function isAllLiquidated() view returns(bool)
func (_ScrollChain *ScrollChainSession) IsAllLiquidated() (bool, error) {
	return _ScrollChain.Contract.IsAllLiquidated(&_ScrollChain.CallOpts)
}

// IsAllLiquidated is a free data retrieval call binding the contract method 0x265b44e1.
//
// Solidity: function isAllLiquidated() view returns(bool)
func (_ScrollChain *ScrollChainCallerSession) IsAllLiquidated() (bool, error) {
	return _ScrollChain.Contract.IsAllLiquidated(&_ScrollChain.CallOpts)
}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 _batchIndex) view returns(bool)
func (_ScrollChain *ScrollChainCaller) IsBatchFinalized(opts *bind.CallOpts, _batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "isBatchFinalized", _batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 _batchIndex) view returns(bool)
func (_ScrollChain *ScrollChainSession) IsBatchFinalized(_batchIndex *big.Int) (bool, error) {
	return _ScrollChain.Contract.IsBatchFinalized(&_ScrollChain.CallOpts, _batchIndex)
}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 _batchIndex) view returns(bool)
func (_ScrollChain *ScrollChainCallerSession) IsBatchFinalized(_batchIndex *big.Int) (bool, error) {
	return _ScrollChain.Contract.IsBatchFinalized(&_ScrollChain.CallOpts, _batchIndex)
}

// IsCommitProofAllowed is a free data retrieval call binding the contract method 0x89da21cb.
//
// Solidity: function isCommitProofAllowed(uint256 batchIndex) view returns(bool)
func (_ScrollChain *ScrollChainCaller) IsCommitProofAllowed(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "isCommitProofAllowed", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCommitProofAllowed is a free data retrieval call binding the contract method 0x89da21cb.
//
// Solidity: function isCommitProofAllowed(uint256 batchIndex) view returns(bool)
func (_ScrollChain *ScrollChainSession) IsCommitProofAllowed(batchIndex *big.Int) (bool, error) {
	return _ScrollChain.Contract.IsCommitProofAllowed(&_ScrollChain.CallOpts, batchIndex)
}

// IsCommitProofAllowed is a free data retrieval call binding the contract method 0x89da21cb.
//
// Solidity: function isCommitProofAllowed(uint256 batchIndex) view returns(bool)
func (_ScrollChain *ScrollChainCallerSession) IsCommitProofAllowed(batchIndex *big.Int) (bool, error) {
	return _ScrollChain.Contract.IsCommitProofAllowed(&_ScrollChain.CallOpts, batchIndex)
}

// IsCommitProofHashAllowed is a free data retrieval call binding the contract method 0xd371c8f0.
//
// Solidity: function isCommitProofHashAllowed(uint256 batchIndex) view returns(bool)
func (_ScrollChain *ScrollChainCaller) IsCommitProofHashAllowed(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "isCommitProofHashAllowed", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCommitProofHashAllowed is a free data retrieval call binding the contract method 0xd371c8f0.
//
// Solidity: function isCommitProofHashAllowed(uint256 batchIndex) view returns(bool)
func (_ScrollChain *ScrollChainSession) IsCommitProofHashAllowed(batchIndex *big.Int) (bool, error) {
	return _ScrollChain.Contract.IsCommitProofHashAllowed(&_ScrollChain.CallOpts, batchIndex)
}

// IsCommitProofHashAllowed is a free data retrieval call binding the contract method 0xd371c8f0.
//
// Solidity: function isCommitProofHashAllowed(uint256 batchIndex) view returns(bool)
func (_ScrollChain *ScrollChainCallerSession) IsCommitProofHashAllowed(batchIndex *big.Int) (bool, error) {
	return _ScrollChain.Contract.IsCommitProofHashAllowed(&_ScrollChain.CallOpts, batchIndex)
}

// IsProver is a free data retrieval call binding the contract method 0x0a245924.
//
// Solidity: function isProver(address ) view returns(bool)
func (_ScrollChain *ScrollChainCaller) IsProver(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "isProver", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProver is a free data retrieval call binding the contract method 0x0a245924.
//
// Solidity: function isProver(address ) view returns(bool)
func (_ScrollChain *ScrollChainSession) IsProver(arg0 common.Address) (bool, error) {
	return _ScrollChain.Contract.IsProver(&_ScrollChain.CallOpts, arg0)
}

// IsProver is a free data retrieval call binding the contract method 0x0a245924.
//
// Solidity: function isProver(address ) view returns(bool)
func (_ScrollChain *ScrollChainCallerSession) IsProver(arg0 common.Address) (bool, error) {
	return _ScrollChain.Contract.IsProver(&_ScrollChain.CallOpts, arg0)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_ScrollChain *ScrollChainCaller) IsSequencer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "isSequencer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_ScrollChain *ScrollChainSession) IsSequencer(arg0 common.Address) (bool, error) {
	return _ScrollChain.Contract.IsSequencer(&_ScrollChain.CallOpts, arg0)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_ScrollChain *ScrollChainCallerSession) IsSequencer(arg0 common.Address) (bool, error) {
	return _ScrollChain.Contract.IsSequencer(&_ScrollChain.CallOpts, arg0)
}

// LastFinalizedBatchIndex is a free data retrieval call binding the contract method 0x059def61.
//
// Solidity: function lastFinalizedBatchIndex() view returns(uint256)
func (_ScrollChain *ScrollChainCaller) LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "lastFinalizedBatchIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFinalizedBatchIndex is a free data retrieval call binding the contract method 0x059def61.
//
// Solidity: function lastFinalizedBatchIndex() view returns(uint256)
func (_ScrollChain *ScrollChainSession) LastFinalizedBatchIndex() (*big.Int, error) {
	return _ScrollChain.Contract.LastFinalizedBatchIndex(&_ScrollChain.CallOpts)
}

// LastFinalizedBatchIndex is a free data retrieval call binding the contract method 0x059def61.
//
// Solidity: function lastFinalizedBatchIndex() view returns(uint256)
func (_ScrollChain *ScrollChainCallerSession) LastFinalizedBatchIndex() (*big.Int, error) {
	return _ScrollChain.Contract.LastFinalizedBatchIndex(&_ScrollChain.CallOpts)
}

// Layer2ChainId is a free data retrieval call binding the contract method 0x03c7f4af.
//
// Solidity: function layer2ChainId() view returns(uint64)
func (_ScrollChain *ScrollChainCaller) Layer2ChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "layer2ChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Layer2ChainId is a free data retrieval call binding the contract method 0x03c7f4af.
//
// Solidity: function layer2ChainId() view returns(uint64)
func (_ScrollChain *ScrollChainSession) Layer2ChainId() (uint64, error) {
	return _ScrollChain.Contract.Layer2ChainId(&_ScrollChain.CallOpts)
}

// Layer2ChainId is a free data retrieval call binding the contract method 0x03c7f4af.
//
// Solidity: function layer2ChainId() view returns(uint64)
func (_ScrollChain *ScrollChainCallerSession) Layer2ChainId() (uint64, error) {
	return _ScrollChain.Contract.Layer2ChainId(&_ScrollChain.CallOpts)
}

// MaxNumTxInChunk is a free data retrieval call binding the contract method 0xef6602ba.
//
// Solidity: function maxNumTxInChunk() view returns(uint256)
func (_ScrollChain *ScrollChainCaller) MaxNumTxInChunk(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "maxNumTxInChunk")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNumTxInChunk is a free data retrieval call binding the contract method 0xef6602ba.
//
// Solidity: function maxNumTxInChunk() view returns(uint256)
func (_ScrollChain *ScrollChainSession) MaxNumTxInChunk() (*big.Int, error) {
	return _ScrollChain.Contract.MaxNumTxInChunk(&_ScrollChain.CallOpts)
}

// MaxNumTxInChunk is a free data retrieval call binding the contract method 0xef6602ba.
//
// Solidity: function maxNumTxInChunk() view returns(uint256)
func (_ScrollChain *ScrollChainCallerSession) MaxNumTxInChunk() (*big.Int, error) {
	return _ScrollChain.Contract.MaxNumTxInChunk(&_ScrollChain.CallOpts)
}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_ScrollChain *ScrollChainCaller) MessageQueue(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "messageQueue")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_ScrollChain *ScrollChainSession) MessageQueue() (common.Address, error) {
	return _ScrollChain.Contract.MessageQueue(&_ScrollChain.CallOpts)
}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_ScrollChain *ScrollChainCallerSession) MessageQueue() (common.Address, error) {
	return _ScrollChain.Contract.MessageQueue(&_ScrollChain.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_ScrollChain *ScrollChainCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "minDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_ScrollChain *ScrollChainSession) MinDeposit() (*big.Int, error) {
	return _ScrollChain.Contract.MinDeposit(&_ScrollChain.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_ScrollChain *ScrollChainCallerSession) MinDeposit() (*big.Int, error) {
	return _ScrollChain.Contract.MinDeposit(&_ScrollChain.CallOpts)
}

// NoProofPunishAmount is a free data retrieval call binding the contract method 0xc809b9d0.
//
// Solidity: function noProofPunishAmount() view returns(uint256)
func (_ScrollChain *ScrollChainCaller) NoProofPunishAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "noProofPunishAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoProofPunishAmount is a free data retrieval call binding the contract method 0xc809b9d0.
//
// Solidity: function noProofPunishAmount() view returns(uint256)
func (_ScrollChain *ScrollChainSession) NoProofPunishAmount() (*big.Int, error) {
	return _ScrollChain.Contract.NoProofPunishAmount(&_ScrollChain.CallOpts)
}

// NoProofPunishAmount is a free data retrieval call binding the contract method 0xc809b9d0.
//
// Solidity: function noProofPunishAmount() view returns(uint256)
func (_ScrollChain *ScrollChainCallerSession) NoProofPunishAmount() (*big.Int, error) {
	return _ScrollChain.Contract.NoProofPunishAmount(&_ScrollChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ScrollChain *ScrollChainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ScrollChain *ScrollChainSession) Owner() (common.Address, error) {
	return _ScrollChain.Contract.Owner(&_ScrollChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ScrollChain *ScrollChainCallerSession) Owner() (common.Address, error) {
	return _ScrollChain.Contract.Owner(&_ScrollChain.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ScrollChain *ScrollChainCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ScrollChain *ScrollChainSession) Paused() (bool, error) {
	return _ScrollChain.Contract.Paused(&_ScrollChain.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ScrollChain *ScrollChainCallerSession) Paused() (bool, error) {
	return _ScrollChain.Contract.Paused(&_ScrollChain.CallOpts)
}

// ProofCommitEpoch is a free data retrieval call binding the contract method 0xc07863ee.
//
// Solidity: function proofCommitEpoch() view returns(uint8)
func (_ScrollChain *ScrollChainCaller) ProofCommitEpoch(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "proofCommitEpoch")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ProofCommitEpoch is a free data retrieval call binding the contract method 0xc07863ee.
//
// Solidity: function proofCommitEpoch() view returns(uint8)
func (_ScrollChain *ScrollChainSession) ProofCommitEpoch() (uint8, error) {
	return _ScrollChain.Contract.ProofCommitEpoch(&_ScrollChain.CallOpts)
}

// ProofCommitEpoch is a free data retrieval call binding the contract method 0xc07863ee.
//
// Solidity: function proofCommitEpoch() view returns(uint8)
func (_ScrollChain *ScrollChainCallerSession) ProofCommitEpoch() (uint8, error) {
	return _ScrollChain.Contract.ProofCommitEpoch(&_ScrollChain.CallOpts)
}

// ProofHashCommitEpoch is a free data retrieval call binding the contract method 0x76c53168.
//
// Solidity: function proofHashCommitEpoch() view returns(uint8)
func (_ScrollChain *ScrollChainCaller) ProofHashCommitEpoch(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "proofHashCommitEpoch")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ProofHashCommitEpoch is a free data retrieval call binding the contract method 0x76c53168.
//
// Solidity: function proofHashCommitEpoch() view returns(uint8)
func (_ScrollChain *ScrollChainSession) ProofHashCommitEpoch() (uint8, error) {
	return _ScrollChain.Contract.ProofHashCommitEpoch(&_ScrollChain.CallOpts)
}

// ProofHashCommitEpoch is a free data retrieval call binding the contract method 0x76c53168.
//
// Solidity: function proofHashCommitEpoch() view returns(uint8)
func (_ScrollChain *ScrollChainCallerSession) ProofHashCommitEpoch() (uint8, error) {
	return _ScrollChain.Contract.ProofHashCommitEpoch(&_ScrollChain.CallOpts)
}

// ProofNum is a free data retrieval call binding the contract method 0x21f5de8c.
//
// Solidity: function proofNum(address ) view returns(uint256)
func (_ScrollChain *ScrollChainCaller) ProofNum(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "proofNum", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProofNum is a free data retrieval call binding the contract method 0x21f5de8c.
//
// Solidity: function proofNum(address ) view returns(uint256)
func (_ScrollChain *ScrollChainSession) ProofNum(arg0 common.Address) (*big.Int, error) {
	return _ScrollChain.Contract.ProofNum(&_ScrollChain.CallOpts, arg0)
}

// ProofNum is a free data retrieval call binding the contract method 0x21f5de8c.
//
// Solidity: function proofNum(address ) view returns(uint256)
func (_ScrollChain *ScrollChainCallerSession) ProofNum(arg0 common.Address) (*big.Int, error) {
	return _ScrollChain.Contract.ProofNum(&_ScrollChain.CallOpts, arg0)
}

// ProverCommitProofHash is a free data retrieval call binding the contract method 0xb02a3a3a.
//
// Solidity: function proverCommitProofHash(uint256 , address ) view returns(bytes32 proofHash, uint256 blockNumber, bool proof)
func (_ScrollChain *ScrollChainCaller) ProverCommitProofHash(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	ProofHash   [32]byte
	BlockNumber *big.Int
	Proof       bool
}, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "proverCommitProofHash", arg0, arg1)

	outstruct := new(struct {
		ProofHash   [32]byte
		BlockNumber *big.Int
		Proof       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProofHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Proof = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// ProverCommitProofHash is a free data retrieval call binding the contract method 0xb02a3a3a.
//
// Solidity: function proverCommitProofHash(uint256 , address ) view returns(bytes32 proofHash, uint256 blockNumber, bool proof)
func (_ScrollChain *ScrollChainSession) ProverCommitProofHash(arg0 *big.Int, arg1 common.Address) (struct {
	ProofHash   [32]byte
	BlockNumber *big.Int
	Proof       bool
}, error) {
	return _ScrollChain.Contract.ProverCommitProofHash(&_ScrollChain.CallOpts, arg0, arg1)
}

// ProverCommitProofHash is a free data retrieval call binding the contract method 0xb02a3a3a.
//
// Solidity: function proverCommitProofHash(uint256 , address ) view returns(bytes32 proofHash, uint256 blockNumber, bool proof)
func (_ScrollChain *ScrollChainCallerSession) ProverCommitProofHash(arg0 *big.Int, arg1 common.Address) (struct {
	ProofHash   [32]byte
	BlockNumber *big.Int
	Proof       bool
}, error) {
	return _ScrollChain.Contract.ProverCommitProofHash(&_ScrollChain.CallOpts, arg0, arg1)
}

// ProverLastLiquidated is a free data retrieval call binding the contract method 0x2ee6046f.
//
// Solidity: function proverLastLiquidated(address ) view returns(uint256)
func (_ScrollChain *ScrollChainCaller) ProverLastLiquidated(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "proverLastLiquidated", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProverLastLiquidated is a free data retrieval call binding the contract method 0x2ee6046f.
//
// Solidity: function proverLastLiquidated(address ) view returns(uint256)
func (_ScrollChain *ScrollChainSession) ProverLastLiquidated(arg0 common.Address) (*big.Int, error) {
	return _ScrollChain.Contract.ProverLastLiquidated(&_ScrollChain.CallOpts, arg0)
}

// ProverLastLiquidated is a free data retrieval call binding the contract method 0x2ee6046f.
//
// Solidity: function proverLastLiquidated(address ) view returns(uint256)
func (_ScrollChain *ScrollChainCallerSession) ProverLastLiquidated(arg0 common.Address) (*big.Int, error) {
	return _ScrollChain.Contract.ProverLastLiquidated(&_ScrollChain.CallOpts, arg0)
}

// ProverLiquidation is a free data retrieval call binding the contract method 0x3ae93939.
//
// Solidity: function proverLiquidation(address , uint256 ) view returns(address prover, bool isSubmittedProofHash, uint256 submitHashBlockNumber, bool isSubmittedProof, uint256 submitProofBlockNumber, bool isLiquidated, uint64 finalNewBatch)
func (_ScrollChain *ScrollChainCaller) ProverLiquidation(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Prover                 common.Address
	IsSubmittedProofHash   bool
	SubmitHashBlockNumber  *big.Int
	IsSubmittedProof       bool
	SubmitProofBlockNumber *big.Int
	IsLiquidated           bool
	FinalNewBatch          uint64
}, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "proverLiquidation", arg0, arg1)

	outstruct := new(struct {
		Prover                 common.Address
		IsSubmittedProofHash   bool
		SubmitHashBlockNumber  *big.Int
		IsSubmittedProof       bool
		SubmitProofBlockNumber *big.Int
		IsLiquidated           bool
		FinalNewBatch          uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Prover = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsSubmittedProofHash = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.SubmitHashBlockNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsSubmittedProof = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.SubmitProofBlockNumber = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsLiquidated = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.FinalNewBatch = *abi.ConvertType(out[6], new(uint64)).(*uint64)

	return *outstruct, err

}

// ProverLiquidation is a free data retrieval call binding the contract method 0x3ae93939.
//
// Solidity: function proverLiquidation(address , uint256 ) view returns(address prover, bool isSubmittedProofHash, uint256 submitHashBlockNumber, bool isSubmittedProof, uint256 submitProofBlockNumber, bool isLiquidated, uint64 finalNewBatch)
func (_ScrollChain *ScrollChainSession) ProverLiquidation(arg0 common.Address, arg1 *big.Int) (struct {
	Prover                 common.Address
	IsSubmittedProofHash   bool
	SubmitHashBlockNumber  *big.Int
	IsSubmittedProof       bool
	SubmitProofBlockNumber *big.Int
	IsLiquidated           bool
	FinalNewBatch          uint64
}, error) {
	return _ScrollChain.Contract.ProverLiquidation(&_ScrollChain.CallOpts, arg0, arg1)
}

// ProverLiquidation is a free data retrieval call binding the contract method 0x3ae93939.
//
// Solidity: function proverLiquidation(address , uint256 ) view returns(address prover, bool isSubmittedProofHash, uint256 submitHashBlockNumber, bool isSubmittedProof, uint256 submitProofBlockNumber, bool isLiquidated, uint64 finalNewBatch)
func (_ScrollChain *ScrollChainCallerSession) ProverLiquidation(arg0 common.Address, arg1 *big.Int) (struct {
	Prover                 common.Address
	IsSubmittedProofHash   bool
	SubmitHashBlockNumber  *big.Int
	IsSubmittedProof       bool
	SubmitProofBlockNumber *big.Int
	IsLiquidated           bool
	FinalNewBatch          uint64
}, error) {
	return _ScrollChain.Contract.ProverLiquidation(&_ScrollChain.CallOpts, arg0, arg1)
}

// ProverPosition is a free data retrieval call binding the contract method 0x3df1bea9.
//
// Solidity: function proverPosition(address , bytes32 ) view returns(uint256)
func (_ScrollChain *ScrollChainCaller) ProverPosition(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "proverPosition", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProverPosition is a free data retrieval call binding the contract method 0x3df1bea9.
//
// Solidity: function proverPosition(address , bytes32 ) view returns(uint256)
func (_ScrollChain *ScrollChainSession) ProverPosition(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _ScrollChain.Contract.ProverPosition(&_ScrollChain.CallOpts, arg0, arg1)
}

// ProverPosition is a free data retrieval call binding the contract method 0x3df1bea9.
//
// Solidity: function proverPosition(address , bytes32 ) view returns(uint256)
func (_ScrollChain *ScrollChainCallerSession) ProverPosition(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _ScrollChain.Contract.ProverPosition(&_ScrollChain.CallOpts, arg0, arg1)
}

// SlotAdapter is a free data retrieval call binding the contract method 0x1c52e346.
//
// Solidity: function slotAdapter() view returns(address)
func (_ScrollChain *ScrollChainCaller) SlotAdapter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "slotAdapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlotAdapter is a free data retrieval call binding the contract method 0x1c52e346.
//
// Solidity: function slotAdapter() view returns(address)
func (_ScrollChain *ScrollChainSession) SlotAdapter() (common.Address, error) {
	return _ScrollChain.Contract.SlotAdapter(&_ScrollChain.CallOpts)
}

// SlotAdapter is a free data retrieval call binding the contract method 0x1c52e346.
//
// Solidity: function slotAdapter() view returns(address)
func (_ScrollChain *ScrollChainCallerSession) SlotAdapter() (common.Address, error) {
	return _ScrollChain.Contract.SlotAdapter(&_ScrollChain.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ScrollChain *ScrollChainCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ScrollChain *ScrollChainSession) Verifier() (common.Address, error) {
	return _ScrollChain.Contract.Verifier(&_ScrollChain.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ScrollChain *ScrollChainCallerSession) Verifier() (common.Address, error) {
	return _ScrollChain.Contract.Verifier(&_ScrollChain.CallOpts)
}

// WithdrawRoots is a free data retrieval call binding the contract method 0xea5f084f.
//
// Solidity: function withdrawRoots(uint256 ) view returns(bytes32)
func (_ScrollChain *ScrollChainCaller) WithdrawRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ScrollChain.contract.Call(opts, &out, "withdrawRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawRoots is a free data retrieval call binding the contract method 0xea5f084f.
//
// Solidity: function withdrawRoots(uint256 ) view returns(bytes32)
func (_ScrollChain *ScrollChainSession) WithdrawRoots(arg0 *big.Int) ([32]byte, error) {
	return _ScrollChain.Contract.WithdrawRoots(&_ScrollChain.CallOpts, arg0)
}

// WithdrawRoots is a free data retrieval call binding the contract method 0xea5f084f.
//
// Solidity: function withdrawRoots(uint256 ) view returns(bytes32)
func (_ScrollChain *ScrollChainCallerSession) WithdrawRoots(arg0 *big.Int) ([32]byte, error) {
	return _ScrollChain.Contract.WithdrawRoots(&_ScrollChain.CallOpts, arg0)
}

// AddProver is a paid mutator transaction binding the contract method 0x1d49e457.
//
// Solidity: function addProver(address _account) returns()
func (_ScrollChain *ScrollChainTransactor) AddProver(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "addProver", _account)
}

// AddProver is a paid mutator transaction binding the contract method 0x1d49e457.
//
// Solidity: function addProver(address _account) returns()
func (_ScrollChain *ScrollChainSession) AddProver(_account common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.AddProver(&_ScrollChain.TransactOpts, _account)
}

// AddProver is a paid mutator transaction binding the contract method 0x1d49e457.
//
// Solidity: function addProver(address _account) returns()
func (_ScrollChain *ScrollChainTransactorSession) AddProver(_account common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.AddProver(&_ScrollChain.TransactOpts, _account)
}

// AddSequencer is a paid mutator transaction binding the contract method 0x8a336231.
//
// Solidity: function addSequencer(address _account) returns()
func (_ScrollChain *ScrollChainTransactor) AddSequencer(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "addSequencer", _account)
}

// AddSequencer is a paid mutator transaction binding the contract method 0x8a336231.
//
// Solidity: function addSequencer(address _account) returns()
func (_ScrollChain *ScrollChainSession) AddSequencer(_account common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.AddSequencer(&_ScrollChain.TransactOpts, _account)
}

// AddSequencer is a paid mutator transaction binding the contract method 0x8a336231.
//
// Solidity: function addSequencer(address _account) returns()
func (_ScrollChain *ScrollChainTransactorSession) AddSequencer(_account common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.AddSequencer(&_ScrollChain.TransactOpts, _account)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x1325aca0.
//
// Solidity: function commitBatch(uint8 _version, bytes _parentBatchHeader, bytes[] _chunks, bytes _skippedL1MessageBitmap) returns()
func (_ScrollChain *ScrollChainTransactor) CommitBatch(opts *bind.TransactOpts, _version uint8, _parentBatchHeader []byte, _chunks [][]byte, _skippedL1MessageBitmap []byte) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "commitBatch", _version, _parentBatchHeader, _chunks, _skippedL1MessageBitmap)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x1325aca0.
//
// Solidity: function commitBatch(uint8 _version, bytes _parentBatchHeader, bytes[] _chunks, bytes _skippedL1MessageBitmap) returns()
func (_ScrollChain *ScrollChainSession) CommitBatch(_version uint8, _parentBatchHeader []byte, _chunks [][]byte, _skippedL1MessageBitmap []byte) (*types.Transaction, error) {
	return _ScrollChain.Contract.CommitBatch(&_ScrollChain.TransactOpts, _version, _parentBatchHeader, _chunks, _skippedL1MessageBitmap)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x1325aca0.
//
// Solidity: function commitBatch(uint8 _version, bytes _parentBatchHeader, bytes[] _chunks, bytes _skippedL1MessageBitmap) returns()
func (_ScrollChain *ScrollChainTransactorSession) CommitBatch(_version uint8, _parentBatchHeader []byte, _chunks [][]byte, _skippedL1MessageBitmap []byte) (*types.Transaction, error) {
	return _ScrollChain.Contract.CommitBatch(&_ScrollChain.TransactOpts, _version, _parentBatchHeader, _chunks, _skippedL1MessageBitmap)
}

// FinalizeBatchWithProof is a paid mutator transaction binding the contract method 0x31fa742d.
//
// Solidity: function finalizeBatchWithProof(bytes _batchHeader, bytes32 _prevStateRoot, bytes32 _postStateRoot, bytes32 _withdrawRoot, bytes _aggrProof) returns()
func (_ScrollChain *ScrollChainTransactor) FinalizeBatchWithProof(opts *bind.TransactOpts, _batchHeader []byte, _prevStateRoot [32]byte, _postStateRoot [32]byte, _withdrawRoot [32]byte, _aggrProof []byte) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "finalizeBatchWithProof", _batchHeader, _prevStateRoot, _postStateRoot, _withdrawRoot, _aggrProof)
}

// FinalizeBatchWithProof is a paid mutator transaction binding the contract method 0x31fa742d.
//
// Solidity: function finalizeBatchWithProof(bytes _batchHeader, bytes32 _prevStateRoot, bytes32 _postStateRoot, bytes32 _withdrawRoot, bytes _aggrProof) returns()
func (_ScrollChain *ScrollChainSession) FinalizeBatchWithProof(_batchHeader []byte, _prevStateRoot [32]byte, _postStateRoot [32]byte, _withdrawRoot [32]byte, _aggrProof []byte) (*types.Transaction, error) {
	return _ScrollChain.Contract.FinalizeBatchWithProof(&_ScrollChain.TransactOpts, _batchHeader, _prevStateRoot, _postStateRoot, _withdrawRoot, _aggrProof)
}

// FinalizeBatchWithProof is a paid mutator transaction binding the contract method 0x31fa742d.
//
// Solidity: function finalizeBatchWithProof(bytes _batchHeader, bytes32 _prevStateRoot, bytes32 _postStateRoot, bytes32 _withdrawRoot, bytes _aggrProof) returns()
func (_ScrollChain *ScrollChainTransactorSession) FinalizeBatchWithProof(_batchHeader []byte, _prevStateRoot [32]byte, _postStateRoot [32]byte, _withdrawRoot [32]byte, _aggrProof []byte) (*types.Transaction, error) {
	return _ScrollChain.Contract.FinalizeBatchWithProof(&_ScrollChain.TransactOpts, _batchHeader, _prevStateRoot, _postStateRoot, _withdrawRoot, _aggrProof)
}

// ImportGenesisBatch is a paid mutator transaction binding the contract method 0x3fdeecb2.
//
// Solidity: function importGenesisBatch(bytes _batchHeader, bytes32 _stateRoot) returns()
func (_ScrollChain *ScrollChainTransactor) ImportGenesisBatch(opts *bind.TransactOpts, _batchHeader []byte, _stateRoot [32]byte) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "importGenesisBatch", _batchHeader, _stateRoot)
}

// ImportGenesisBatch is a paid mutator transaction binding the contract method 0x3fdeecb2.
//
// Solidity: function importGenesisBatch(bytes _batchHeader, bytes32 _stateRoot) returns()
func (_ScrollChain *ScrollChainSession) ImportGenesisBatch(_batchHeader []byte, _stateRoot [32]byte) (*types.Transaction, error) {
	return _ScrollChain.Contract.ImportGenesisBatch(&_ScrollChain.TransactOpts, _batchHeader, _stateRoot)
}

// ImportGenesisBatch is a paid mutator transaction binding the contract method 0x3fdeecb2.
//
// Solidity: function importGenesisBatch(bytes _batchHeader, bytes32 _stateRoot) returns()
func (_ScrollChain *ScrollChainTransactorSession) ImportGenesisBatch(_batchHeader []byte, _stateRoot [32]byte) (*types.Transaction, error) {
	return _ScrollChain.Contract.ImportGenesisBatch(&_ScrollChain.TransactOpts, _batchHeader, _stateRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _messageQueue, address _verifier, uint256 _maxNumTxInChunk) returns()
func (_ScrollChain *ScrollChainTransactor) Initialize(opts *bind.TransactOpts, _messageQueue common.Address, _verifier common.Address, _maxNumTxInChunk *big.Int) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "initialize", _messageQueue, _verifier, _maxNumTxInChunk)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _messageQueue, address _verifier, uint256 _maxNumTxInChunk) returns()
func (_ScrollChain *ScrollChainSession) Initialize(_messageQueue common.Address, _verifier common.Address, _maxNumTxInChunk *big.Int) (*types.Transaction, error) {
	return _ScrollChain.Contract.Initialize(&_ScrollChain.TransactOpts, _messageQueue, _verifier, _maxNumTxInChunk)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _messageQueue, address _verifier, uint256 _maxNumTxInChunk) returns()
func (_ScrollChain *ScrollChainTransactorSession) Initialize(_messageQueue common.Address, _verifier common.Address, _maxNumTxInChunk *big.Int) (*types.Transaction, error) {
	return _ScrollChain.Contract.Initialize(&_ScrollChain.TransactOpts, _messageQueue, _verifier, _maxNumTxInChunk)
}

// RemoveProver is a paid mutator transaction binding the contract method 0xb571d3dd.
//
// Solidity: function removeProver(address _account) returns()
func (_ScrollChain *ScrollChainTransactor) RemoveProver(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "removeProver", _account)
}

// RemoveProver is a paid mutator transaction binding the contract method 0xb571d3dd.
//
// Solidity: function removeProver(address _account) returns()
func (_ScrollChain *ScrollChainSession) RemoveProver(_account common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.RemoveProver(&_ScrollChain.TransactOpts, _account)
}

// RemoveProver is a paid mutator transaction binding the contract method 0xb571d3dd.
//
// Solidity: function removeProver(address _account) returns()
func (_ScrollChain *ScrollChainTransactorSession) RemoveProver(_account common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.RemoveProver(&_ScrollChain.TransactOpts, _account)
}

// RemoveSequencer is a paid mutator transaction binding the contract method 0x6989ca7c.
//
// Solidity: function removeSequencer(address _account) returns()
func (_ScrollChain *ScrollChainTransactor) RemoveSequencer(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "removeSequencer", _account)
}

// RemoveSequencer is a paid mutator transaction binding the contract method 0x6989ca7c.
//
// Solidity: function removeSequencer(address _account) returns()
func (_ScrollChain *ScrollChainSession) RemoveSequencer(_account common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.RemoveSequencer(&_ScrollChain.TransactOpts, _account)
}

// RemoveSequencer is a paid mutator transaction binding the contract method 0x6989ca7c.
//
// Solidity: function removeSequencer(address _account) returns()
func (_ScrollChain *ScrollChainTransactorSession) RemoveSequencer(_account common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.RemoveSequencer(&_ScrollChain.TransactOpts, _account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ScrollChain *ScrollChainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ScrollChain *ScrollChainSession) RenounceOwnership() (*types.Transaction, error) {
	return _ScrollChain.Contract.RenounceOwnership(&_ScrollChain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ScrollChain *ScrollChainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ScrollChain.Contract.RenounceOwnership(&_ScrollChain.TransactOpts)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes _batchHeader, uint256 _count) returns()
func (_ScrollChain *ScrollChainTransactor) RevertBatch(opts *bind.TransactOpts, _batchHeader []byte, _count *big.Int) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "revertBatch", _batchHeader, _count)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes _batchHeader, uint256 _count) returns()
func (_ScrollChain *ScrollChainSession) RevertBatch(_batchHeader []byte, _count *big.Int) (*types.Transaction, error) {
	return _ScrollChain.Contract.RevertBatch(&_ScrollChain.TransactOpts, _batchHeader, _count)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes _batchHeader, uint256 _count) returns()
func (_ScrollChain *ScrollChainTransactorSession) RevertBatch(_batchHeader []byte, _count *big.Int) (*types.Transaction, error) {
	return _ScrollChain.Contract.RevertBatch(&_ScrollChain.TransactOpts, _batchHeader, _count)
}

// SetDeposit is a paid mutator transaction binding the contract method 0x3b80938e.
//
// Solidity: function setDeposit(address _ideDeposit) returns()
func (_ScrollChain *ScrollChainTransactor) SetDeposit(opts *bind.TransactOpts, _ideDeposit common.Address) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "setDeposit", _ideDeposit)
}

// SetDeposit is a paid mutator transaction binding the contract method 0x3b80938e.
//
// Solidity: function setDeposit(address _ideDeposit) returns()
func (_ScrollChain *ScrollChainSession) SetDeposit(_ideDeposit common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetDeposit(&_ScrollChain.TransactOpts, _ideDeposit)
}

// SetDeposit is a paid mutator transaction binding the contract method 0x3b80938e.
//
// Solidity: function setDeposit(address _ideDeposit) returns()
func (_ScrollChain *ScrollChainTransactorSession) SetDeposit(_ideDeposit common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetDeposit(&_ScrollChain.TransactOpts, _ideDeposit)
}

// SetIncorrectProofPunishAmount is a paid mutator transaction binding the contract method 0x5bb79603.
//
// Solidity: function setIncorrectProofPunishAmount(uint256 _amount) returns()
func (_ScrollChain *ScrollChainTransactor) SetIncorrectProofPunishAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "setIncorrectProofPunishAmount", _amount)
}

// SetIncorrectProofPunishAmount is a paid mutator transaction binding the contract method 0x5bb79603.
//
// Solidity: function setIncorrectProofPunishAmount(uint256 _amount) returns()
func (_ScrollChain *ScrollChainSession) SetIncorrectProofPunishAmount(_amount *big.Int) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetIncorrectProofPunishAmount(&_ScrollChain.TransactOpts, _amount)
}

// SetIncorrectProofPunishAmount is a paid mutator transaction binding the contract method 0x5bb79603.
//
// Solidity: function setIncorrectProofPunishAmount(uint256 _amount) returns()
func (_ScrollChain *ScrollChainTransactorSession) SetIncorrectProofPunishAmount(_amount *big.Int) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetIncorrectProofPunishAmount(&_ScrollChain.TransactOpts, _amount)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 _amount) returns()
func (_ScrollChain *ScrollChainTransactor) SetMinDeposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "setMinDeposit", _amount)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 _amount) returns()
func (_ScrollChain *ScrollChainSession) SetMinDeposit(_amount *big.Int) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetMinDeposit(&_ScrollChain.TransactOpts, _amount)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 _amount) returns()
func (_ScrollChain *ScrollChainTransactorSession) SetMinDeposit(_amount *big.Int) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetMinDeposit(&_ScrollChain.TransactOpts, _amount)
}

// SetNoProofPunishAmount is a paid mutator transaction binding the contract method 0x8db05052.
//
// Solidity: function setNoProofPunishAmount(uint256 _amount) returns()
func (_ScrollChain *ScrollChainTransactor) SetNoProofPunishAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "setNoProofPunishAmount", _amount)
}

// SetNoProofPunishAmount is a paid mutator transaction binding the contract method 0x8db05052.
//
// Solidity: function setNoProofPunishAmount(uint256 _amount) returns()
func (_ScrollChain *ScrollChainSession) SetNoProofPunishAmount(_amount *big.Int) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetNoProofPunishAmount(&_ScrollChain.TransactOpts, _amount)
}

// SetNoProofPunishAmount is a paid mutator transaction binding the contract method 0x8db05052.
//
// Solidity: function setNoProofPunishAmount(uint256 _amount) returns()
func (_ScrollChain *ScrollChainTransactorSession) SetNoProofPunishAmount(_amount *big.Int) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetNoProofPunishAmount(&_ScrollChain.TransactOpts, _amount)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_ScrollChain *ScrollChainTransactor) SetPause(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "setPause", _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_ScrollChain *ScrollChainSession) SetPause(_status bool) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetPause(&_ScrollChain.TransactOpts, _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_ScrollChain *ScrollChainTransactorSession) SetPause(_status bool) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetPause(&_ScrollChain.TransactOpts, _status)
}

// SetProofCommitEpoch is a paid mutator transaction binding the contract method 0x4799a9a8.
//
// Solidity: function setProofCommitEpoch(uint8 _newCommitEpoch) returns()
func (_ScrollChain *ScrollChainTransactor) SetProofCommitEpoch(opts *bind.TransactOpts, _newCommitEpoch uint8) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "setProofCommitEpoch", _newCommitEpoch)
}

// SetProofCommitEpoch is a paid mutator transaction binding the contract method 0x4799a9a8.
//
// Solidity: function setProofCommitEpoch(uint8 _newCommitEpoch) returns()
func (_ScrollChain *ScrollChainSession) SetProofCommitEpoch(_newCommitEpoch uint8) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetProofCommitEpoch(&_ScrollChain.TransactOpts, _newCommitEpoch)
}

// SetProofCommitEpoch is a paid mutator transaction binding the contract method 0x4799a9a8.
//
// Solidity: function setProofCommitEpoch(uint8 _newCommitEpoch) returns()
func (_ScrollChain *ScrollChainTransactorSession) SetProofCommitEpoch(_newCommitEpoch uint8) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetProofCommitEpoch(&_ScrollChain.TransactOpts, _newCommitEpoch)
}

// SetProofHashCommitEpoch is a paid mutator transaction binding the contract method 0x8f2aa42a.
//
// Solidity: function setProofHashCommitEpoch(uint8 _newCommitEpoch) returns()
func (_ScrollChain *ScrollChainTransactor) SetProofHashCommitEpoch(opts *bind.TransactOpts, _newCommitEpoch uint8) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "setProofHashCommitEpoch", _newCommitEpoch)
}

// SetProofHashCommitEpoch is a paid mutator transaction binding the contract method 0x8f2aa42a.
//
// Solidity: function setProofHashCommitEpoch(uint8 _newCommitEpoch) returns()
func (_ScrollChain *ScrollChainSession) SetProofHashCommitEpoch(_newCommitEpoch uint8) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetProofHashCommitEpoch(&_ScrollChain.TransactOpts, _newCommitEpoch)
}

// SetProofHashCommitEpoch is a paid mutator transaction binding the contract method 0x8f2aa42a.
//
// Solidity: function setProofHashCommitEpoch(uint8 _newCommitEpoch) returns()
func (_ScrollChain *ScrollChainTransactorSession) SetProofHashCommitEpoch(_newCommitEpoch uint8) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetProofHashCommitEpoch(&_ScrollChain.TransactOpts, _newCommitEpoch)
}

// SetSlotAdapter is a paid mutator transaction binding the contract method 0xa62125c9.
//
// Solidity: function setSlotAdapter(address _slotAdapter) returns()
func (_ScrollChain *ScrollChainTransactor) SetSlotAdapter(opts *bind.TransactOpts, _slotAdapter common.Address) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "setSlotAdapter", _slotAdapter)
}

// SetSlotAdapter is a paid mutator transaction binding the contract method 0xa62125c9.
//
// Solidity: function setSlotAdapter(address _slotAdapter) returns()
func (_ScrollChain *ScrollChainSession) SetSlotAdapter(_slotAdapter common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetSlotAdapter(&_ScrollChain.TransactOpts, _slotAdapter)
}

// SetSlotAdapter is a paid mutator transaction binding the contract method 0xa62125c9.
//
// Solidity: function setSlotAdapter(address _slotAdapter) returns()
func (_ScrollChain *ScrollChainTransactorSession) SetSlotAdapter(_slotAdapter common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.SetSlotAdapter(&_ScrollChain.TransactOpts, _slotAdapter)
}

// Settle is a paid mutator transaction binding the contract method 0x6a256b29.
//
// Solidity: function settle(address _account) returns()
func (_ScrollChain *ScrollChainTransactor) Settle(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "settle", _account)
}

// Settle is a paid mutator transaction binding the contract method 0x6a256b29.
//
// Solidity: function settle(address _account) returns()
func (_ScrollChain *ScrollChainSession) Settle(_account common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.Settle(&_ScrollChain.TransactOpts, _account)
}

// Settle is a paid mutator transaction binding the contract method 0x6a256b29.
//
// Solidity: function settle(address _account) returns()
func (_ScrollChain *ScrollChainTransactorSession) Settle(_account common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.Settle(&_ScrollChain.TransactOpts, _account)
}

// SubmitProofHash is a paid mutator transaction binding the contract method 0xcd32c210.
//
// Solidity: function submitProofHash(uint256 batchIndex, bytes32 _proofHash) returns()
func (_ScrollChain *ScrollChainTransactor) SubmitProofHash(opts *bind.TransactOpts, batchIndex *big.Int, _proofHash [32]byte) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "submitProofHash", batchIndex, _proofHash)
}

// SubmitProofHash is a paid mutator transaction binding the contract method 0xcd32c210.
//
// Solidity: function submitProofHash(uint256 batchIndex, bytes32 _proofHash) returns()
func (_ScrollChain *ScrollChainSession) SubmitProofHash(batchIndex *big.Int, _proofHash [32]byte) (*types.Transaction, error) {
	return _ScrollChain.Contract.SubmitProofHash(&_ScrollChain.TransactOpts, batchIndex, _proofHash)
}

// SubmitProofHash is a paid mutator transaction binding the contract method 0xcd32c210.
//
// Solidity: function submitProofHash(uint256 batchIndex, bytes32 _proofHash) returns()
func (_ScrollChain *ScrollChainTransactorSession) SubmitProofHash(batchIndex *big.Int, _proofHash [32]byte) (*types.Transaction, error) {
	return _ScrollChain.Contract.SubmitProofHash(&_ScrollChain.TransactOpts, batchIndex, _proofHash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ScrollChain *ScrollChainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ScrollChain *ScrollChainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.TransferOwnership(&_ScrollChain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ScrollChain *ScrollChainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.TransferOwnership(&_ScrollChain.TransactOpts, newOwner)
}

// UpdateMaxNumTxInChunk is a paid mutator transaction binding the contract method 0x1e228302.
//
// Solidity: function updateMaxNumTxInChunk(uint256 _maxNumTxInChunk) returns()
func (_ScrollChain *ScrollChainTransactor) UpdateMaxNumTxInChunk(opts *bind.TransactOpts, _maxNumTxInChunk *big.Int) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "updateMaxNumTxInChunk", _maxNumTxInChunk)
}

// UpdateMaxNumTxInChunk is a paid mutator transaction binding the contract method 0x1e228302.
//
// Solidity: function updateMaxNumTxInChunk(uint256 _maxNumTxInChunk) returns()
func (_ScrollChain *ScrollChainSession) UpdateMaxNumTxInChunk(_maxNumTxInChunk *big.Int) (*types.Transaction, error) {
	return _ScrollChain.Contract.UpdateMaxNumTxInChunk(&_ScrollChain.TransactOpts, _maxNumTxInChunk)
}

// UpdateMaxNumTxInChunk is a paid mutator transaction binding the contract method 0x1e228302.
//
// Solidity: function updateMaxNumTxInChunk(uint256 _maxNumTxInChunk) returns()
func (_ScrollChain *ScrollChainTransactorSession) UpdateMaxNumTxInChunk(_maxNumTxInChunk *big.Int) (*types.Transaction, error) {
	return _ScrollChain.Contract.UpdateMaxNumTxInChunk(&_ScrollChain.TransactOpts, _maxNumTxInChunk)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address _newVerifier) returns()
func (_ScrollChain *ScrollChainTransactor) UpdateVerifier(opts *bind.TransactOpts, _newVerifier common.Address) (*types.Transaction, error) {
	return _ScrollChain.contract.Transact(opts, "updateVerifier", _newVerifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address _newVerifier) returns()
func (_ScrollChain *ScrollChainSession) UpdateVerifier(_newVerifier common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.UpdateVerifier(&_ScrollChain.TransactOpts, _newVerifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address _newVerifier) returns()
func (_ScrollChain *ScrollChainTransactorSession) UpdateVerifier(_newVerifier common.Address) (*types.Transaction, error) {
	return _ScrollChain.Contract.UpdateVerifier(&_ScrollChain.TransactOpts, _newVerifier)
}

// ScrollChainCommitBatchIterator is returned from FilterCommitBatch and is used to iterate over the raw logs and unpacked data for CommitBatch events raised by the ScrollChain contract.
type ScrollChainCommitBatchIterator struct {
	Event *ScrollChainCommitBatch // Event containing the contract specifics and raw log

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
func (it *ScrollChainCommitBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainCommitBatch)
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
		it.Event = new(ScrollChainCommitBatch)
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
func (it *ScrollChainCommitBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainCommitBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainCommitBatch represents a CommitBatch event raised by the ScrollChain contract.
type ScrollChainCommitBatch struct {
	BatchIndex *big.Int
	BatchHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitBatch is a free log retrieval operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_ScrollChain *ScrollChainFilterer) FilterCommitBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*ScrollChainCommitBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "CommitBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &ScrollChainCommitBatchIterator{contract: _ScrollChain.contract, event: "CommitBatch", logs: logs, sub: sub}, nil
}

// WatchCommitBatch is a free log subscription operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_ScrollChain *ScrollChainFilterer) WatchCommitBatch(opts *bind.WatchOpts, sink chan<- *ScrollChainCommitBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "CommitBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainCommitBatch)
				if err := _ScrollChain.contract.UnpackLog(event, "CommitBatch", log); err != nil {
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

// ParseCommitBatch is a log parse operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_ScrollChain *ScrollChainFilterer) ParseCommitBatch(log types.Log) (*ScrollChainCommitBatch, error) {
	event := new(ScrollChainCommitBatch)
	if err := _ScrollChain.contract.UnpackLog(event, "CommitBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainFinalizeBatchIterator is returned from FilterFinalizeBatch and is used to iterate over the raw logs and unpacked data for FinalizeBatch events raised by the ScrollChain contract.
type ScrollChainFinalizeBatchIterator struct {
	Event *ScrollChainFinalizeBatch // Event containing the contract specifics and raw log

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
func (it *ScrollChainFinalizeBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainFinalizeBatch)
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
		it.Event = new(ScrollChainFinalizeBatch)
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
func (it *ScrollChainFinalizeBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainFinalizeBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainFinalizeBatch represents a FinalizeBatch event raised by the ScrollChain contract.
type ScrollChainFinalizeBatch struct {
	BatchIndex   *big.Int
	BatchHash    [32]byte
	StateRoot    [32]byte
	WithdrawRoot [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBatch is a free log retrieval operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_ScrollChain *ScrollChainFilterer) FilterFinalizeBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*ScrollChainFinalizeBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "FinalizeBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &ScrollChainFinalizeBatchIterator{contract: _ScrollChain.contract, event: "FinalizeBatch", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatch is a free log subscription operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_ScrollChain *ScrollChainFilterer) WatchFinalizeBatch(opts *bind.WatchOpts, sink chan<- *ScrollChainFinalizeBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "FinalizeBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainFinalizeBatch)
				if err := _ScrollChain.contract.UnpackLog(event, "FinalizeBatch", log); err != nil {
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

// ParseFinalizeBatch is a log parse operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_ScrollChain *ScrollChainFilterer) ParseFinalizeBatch(log types.Log) (*ScrollChainFinalizeBatch, error) {
	event := new(ScrollChainFinalizeBatch)
	if err := _ScrollChain.contract.UnpackLog(event, "FinalizeBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ScrollChain contract.
type ScrollChainInitializedIterator struct {
	Event *ScrollChainInitialized // Event containing the contract specifics and raw log

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
func (it *ScrollChainInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainInitialized)
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
		it.Event = new(ScrollChainInitialized)
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
func (it *ScrollChainInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainInitialized represents a Initialized event raised by the ScrollChain contract.
type ScrollChainInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ScrollChain *ScrollChainFilterer) FilterInitialized(opts *bind.FilterOpts) (*ScrollChainInitializedIterator, error) {

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ScrollChainInitializedIterator{contract: _ScrollChain.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ScrollChain *ScrollChainFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ScrollChainInitialized) (event.Subscription, error) {

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainInitialized)
				if err := _ScrollChain.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ScrollChain *ScrollChainFilterer) ParseInitialized(log types.Log) (*ScrollChainInitialized, error) {
	event := new(ScrollChainInitialized)
	if err := _ScrollChain.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ScrollChain contract.
type ScrollChainOwnershipTransferredIterator struct {
	Event *ScrollChainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ScrollChainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainOwnershipTransferred)
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
		it.Event = new(ScrollChainOwnershipTransferred)
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
func (it *ScrollChainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainOwnershipTransferred represents a OwnershipTransferred event raised by the ScrollChain contract.
type ScrollChainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ScrollChain *ScrollChainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ScrollChainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ScrollChainOwnershipTransferredIterator{contract: _ScrollChain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ScrollChain *ScrollChainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollChainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainOwnershipTransferred)
				if err := _ScrollChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ScrollChain *ScrollChainFilterer) ParseOwnershipTransferred(log types.Log) (*ScrollChainOwnershipTransferred, error) {
	event := new(ScrollChainOwnershipTransferred)
	if err := _ScrollChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ScrollChain contract.
type ScrollChainPausedIterator struct {
	Event *ScrollChainPaused // Event containing the contract specifics and raw log

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
func (it *ScrollChainPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainPaused)
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
		it.Event = new(ScrollChainPaused)
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
func (it *ScrollChainPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainPaused represents a Paused event raised by the ScrollChain contract.
type ScrollChainPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ScrollChain *ScrollChainFilterer) FilterPaused(opts *bind.FilterOpts) (*ScrollChainPausedIterator, error) {

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ScrollChainPausedIterator{contract: _ScrollChain.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ScrollChain *ScrollChainFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ScrollChainPaused) (event.Subscription, error) {

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainPaused)
				if err := _ScrollChain.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ScrollChain *ScrollChainFilterer) ParsePaused(log types.Log) (*ScrollChainPaused, error) {
	event := new(ScrollChainPaused)
	if err := _ScrollChain.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainRevertBatchIterator is returned from FilterRevertBatch and is used to iterate over the raw logs and unpacked data for RevertBatch events raised by the ScrollChain contract.
type ScrollChainRevertBatchIterator struct {
	Event *ScrollChainRevertBatch // Event containing the contract specifics and raw log

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
func (it *ScrollChainRevertBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainRevertBatch)
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
		it.Event = new(ScrollChainRevertBatch)
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
func (it *ScrollChainRevertBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainRevertBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainRevertBatch represents a RevertBatch event raised by the ScrollChain contract.
type ScrollChainRevertBatch struct {
	BatchIndex *big.Int
	BatchHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevertBatch is a free log retrieval operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_ScrollChain *ScrollChainFilterer) FilterRevertBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*ScrollChainRevertBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "RevertBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &ScrollChainRevertBatchIterator{contract: _ScrollChain.contract, event: "RevertBatch", logs: logs, sub: sub}, nil
}

// WatchRevertBatch is a free log subscription operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_ScrollChain *ScrollChainFilterer) WatchRevertBatch(opts *bind.WatchOpts, sink chan<- *ScrollChainRevertBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "RevertBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainRevertBatch)
				if err := _ScrollChain.contract.UnpackLog(event, "RevertBatch", log); err != nil {
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

// ParseRevertBatch is a log parse operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_ScrollChain *ScrollChainFilterer) ParseRevertBatch(log types.Log) (*ScrollChainRevertBatch, error) {
	event := new(ScrollChainRevertBatch)
	if err := _ScrollChain.contract.UnpackLog(event, "RevertBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainSetProofCommitEpochIterator is returned from FilterSetProofCommitEpoch and is used to iterate over the raw logs and unpacked data for SetProofCommitEpoch events raised by the ScrollChain contract.
type ScrollChainSetProofCommitEpochIterator struct {
	Event *ScrollChainSetProofCommitEpoch // Event containing the contract specifics and raw log

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
func (it *ScrollChainSetProofCommitEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainSetProofCommitEpoch)
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
		it.Event = new(ScrollChainSetProofCommitEpoch)
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
func (it *ScrollChainSetProofCommitEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainSetProofCommitEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainSetProofCommitEpoch represents a SetProofCommitEpoch event raised by the ScrollChain contract.
type ScrollChainSetProofCommitEpoch struct {
	NewProofCommitEpoch uint8
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetProofCommitEpoch is a free log retrieval operation binding the contract event 0xeb7b81984df2c6f86796b8f38c1b19e8f6b326f861d5188594b6cbc02dd2b5db.
//
// Solidity: event SetProofCommitEpoch(uint8 newProofCommitEpoch)
func (_ScrollChain *ScrollChainFilterer) FilterSetProofCommitEpoch(opts *bind.FilterOpts) (*ScrollChainSetProofCommitEpochIterator, error) {

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "SetProofCommitEpoch")
	if err != nil {
		return nil, err
	}
	return &ScrollChainSetProofCommitEpochIterator{contract: _ScrollChain.contract, event: "SetProofCommitEpoch", logs: logs, sub: sub}, nil
}

// WatchSetProofCommitEpoch is a free log subscription operation binding the contract event 0xeb7b81984df2c6f86796b8f38c1b19e8f6b326f861d5188594b6cbc02dd2b5db.
//
// Solidity: event SetProofCommitEpoch(uint8 newProofCommitEpoch)
func (_ScrollChain *ScrollChainFilterer) WatchSetProofCommitEpoch(opts *bind.WatchOpts, sink chan<- *ScrollChainSetProofCommitEpoch) (event.Subscription, error) {

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "SetProofCommitEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainSetProofCommitEpoch)
				if err := _ScrollChain.contract.UnpackLog(event, "SetProofCommitEpoch", log); err != nil {
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

// ParseSetProofCommitEpoch is a log parse operation binding the contract event 0xeb7b81984df2c6f86796b8f38c1b19e8f6b326f861d5188594b6cbc02dd2b5db.
//
// Solidity: event SetProofCommitEpoch(uint8 newProofCommitEpoch)
func (_ScrollChain *ScrollChainFilterer) ParseSetProofCommitEpoch(log types.Log) (*ScrollChainSetProofCommitEpoch, error) {
	event := new(ScrollChainSetProofCommitEpoch)
	if err := _ScrollChain.contract.UnpackLog(event, "SetProofCommitEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainSetProofHashCommitEpochIterator is returned from FilterSetProofHashCommitEpoch and is used to iterate over the raw logs and unpacked data for SetProofHashCommitEpoch events raised by the ScrollChain contract.
type ScrollChainSetProofHashCommitEpochIterator struct {
	Event *ScrollChainSetProofHashCommitEpoch // Event containing the contract specifics and raw log

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
func (it *ScrollChainSetProofHashCommitEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainSetProofHashCommitEpoch)
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
		it.Event = new(ScrollChainSetProofHashCommitEpoch)
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
func (it *ScrollChainSetProofHashCommitEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainSetProofHashCommitEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainSetProofHashCommitEpoch represents a SetProofHashCommitEpoch event raised by the ScrollChain contract.
type ScrollChainSetProofHashCommitEpoch struct {
	NewProofHashCommitEpoch uint8
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSetProofHashCommitEpoch is a free log retrieval operation binding the contract event 0x3baed86e39e55c44af3b5e06d60eb89cf23a7b1b7ebb7f0b8a272f64bce1165a.
//
// Solidity: event SetProofHashCommitEpoch(uint8 newProofHashCommitEpoch)
func (_ScrollChain *ScrollChainFilterer) FilterSetProofHashCommitEpoch(opts *bind.FilterOpts) (*ScrollChainSetProofHashCommitEpochIterator, error) {

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "SetProofHashCommitEpoch")
	if err != nil {
		return nil, err
	}
	return &ScrollChainSetProofHashCommitEpochIterator{contract: _ScrollChain.contract, event: "SetProofHashCommitEpoch", logs: logs, sub: sub}, nil
}

// WatchSetProofHashCommitEpoch is a free log subscription operation binding the contract event 0x3baed86e39e55c44af3b5e06d60eb89cf23a7b1b7ebb7f0b8a272f64bce1165a.
//
// Solidity: event SetProofHashCommitEpoch(uint8 newProofHashCommitEpoch)
func (_ScrollChain *ScrollChainFilterer) WatchSetProofHashCommitEpoch(opts *bind.WatchOpts, sink chan<- *ScrollChainSetProofHashCommitEpoch) (event.Subscription, error) {

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "SetProofHashCommitEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainSetProofHashCommitEpoch)
				if err := _ScrollChain.contract.UnpackLog(event, "SetProofHashCommitEpoch", log); err != nil {
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

// ParseSetProofHashCommitEpoch is a log parse operation binding the contract event 0x3baed86e39e55c44af3b5e06d60eb89cf23a7b1b7ebb7f0b8a272f64bce1165a.
//
// Solidity: event SetProofHashCommitEpoch(uint8 newProofHashCommitEpoch)
func (_ScrollChain *ScrollChainFilterer) ParseSetProofHashCommitEpoch(log types.Log) (*ScrollChainSetProofHashCommitEpoch, error) {
	event := new(ScrollChainSetProofHashCommitEpoch)
	if err := _ScrollChain.contract.UnpackLog(event, "SetProofHashCommitEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainSubmitProofHashIterator is returned from FilterSubmitProofHash and is used to iterate over the raw logs and unpacked data for SubmitProofHash events raised by the ScrollChain contract.
type ScrollChainSubmitProofHashIterator struct {
	Event *ScrollChainSubmitProofHash // Event containing the contract specifics and raw log

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
func (it *ScrollChainSubmitProofHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainSubmitProofHash)
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
		it.Event = new(ScrollChainSubmitProofHash)
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
func (it *ScrollChainSubmitProofHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainSubmitProofHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainSubmitProofHash represents a SubmitProofHash event raised by the ScrollChain contract.
type ScrollChainSubmitProofHash struct {
	Prover     common.Address
	BatchIndex *big.Int
	ProofHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubmitProofHash is a free log retrieval operation binding the contract event 0xab87afef6fb753dd9c1b9e846c8370a6427c3db49fef6eddfb4da22af00d2423.
//
// Solidity: event SubmitProofHash(address _prover, uint256 batchIndex, bytes32 _proofHash)
func (_ScrollChain *ScrollChainFilterer) FilterSubmitProofHash(opts *bind.FilterOpts) (*ScrollChainSubmitProofHashIterator, error) {

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "SubmitProofHash")
	if err != nil {
		return nil, err
	}
	return &ScrollChainSubmitProofHashIterator{contract: _ScrollChain.contract, event: "SubmitProofHash", logs: logs, sub: sub}, nil
}

// WatchSubmitProofHash is a free log subscription operation binding the contract event 0xab87afef6fb753dd9c1b9e846c8370a6427c3db49fef6eddfb4da22af00d2423.
//
// Solidity: event SubmitProofHash(address _prover, uint256 batchIndex, bytes32 _proofHash)
func (_ScrollChain *ScrollChainFilterer) WatchSubmitProofHash(opts *bind.WatchOpts, sink chan<- *ScrollChainSubmitProofHash) (event.Subscription, error) {

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "SubmitProofHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainSubmitProofHash)
				if err := _ScrollChain.contract.UnpackLog(event, "SubmitProofHash", log); err != nil {
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

// ParseSubmitProofHash is a log parse operation binding the contract event 0xab87afef6fb753dd9c1b9e846c8370a6427c3db49fef6eddfb4da22af00d2423.
//
// Solidity: event SubmitProofHash(address _prover, uint256 batchIndex, bytes32 _proofHash)
func (_ScrollChain *ScrollChainFilterer) ParseSubmitProofHash(log types.Log) (*ScrollChainSubmitProofHash, error) {
	event := new(ScrollChainSubmitProofHash)
	if err := _ScrollChain.contract.UnpackLog(event, "SubmitProofHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ScrollChain contract.
type ScrollChainUnpausedIterator struct {
	Event *ScrollChainUnpaused // Event containing the contract specifics and raw log

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
func (it *ScrollChainUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainUnpaused)
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
		it.Event = new(ScrollChainUnpaused)
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
func (it *ScrollChainUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainUnpaused represents a Unpaused event raised by the ScrollChain contract.
type ScrollChainUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ScrollChain *ScrollChainFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ScrollChainUnpausedIterator, error) {

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ScrollChainUnpausedIterator{contract: _ScrollChain.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ScrollChain *ScrollChainFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ScrollChainUnpaused) (event.Subscription, error) {

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainUnpaused)
				if err := _ScrollChain.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ScrollChain *ScrollChainFilterer) ParseUnpaused(log types.Log) (*ScrollChainUnpaused, error) {
	event := new(ScrollChainUnpaused)
	if err := _ScrollChain.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainUpdateMaxNumTxInChunkIterator is returned from FilterUpdateMaxNumTxInChunk and is used to iterate over the raw logs and unpacked data for UpdateMaxNumTxInChunk events raised by the ScrollChain contract.
type ScrollChainUpdateMaxNumTxInChunkIterator struct {
	Event *ScrollChainUpdateMaxNumTxInChunk // Event containing the contract specifics and raw log

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
func (it *ScrollChainUpdateMaxNumTxInChunkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainUpdateMaxNumTxInChunk)
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
		it.Event = new(ScrollChainUpdateMaxNumTxInChunk)
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
func (it *ScrollChainUpdateMaxNumTxInChunkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainUpdateMaxNumTxInChunkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainUpdateMaxNumTxInChunk represents a UpdateMaxNumTxInChunk event raised by the ScrollChain contract.
type ScrollChainUpdateMaxNumTxInChunk struct {
	OldMaxNumTxInChunk *big.Int
	NewMaxNumTxInChunk *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxNumTxInChunk is a free log retrieval operation binding the contract event 0x6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a.
//
// Solidity: event UpdateMaxNumTxInChunk(uint256 oldMaxNumTxInChunk, uint256 newMaxNumTxInChunk)
func (_ScrollChain *ScrollChainFilterer) FilterUpdateMaxNumTxInChunk(opts *bind.FilterOpts) (*ScrollChainUpdateMaxNumTxInChunkIterator, error) {

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "UpdateMaxNumTxInChunk")
	if err != nil {
		return nil, err
	}
	return &ScrollChainUpdateMaxNumTxInChunkIterator{contract: _ScrollChain.contract, event: "UpdateMaxNumTxInChunk", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxNumTxInChunk is a free log subscription operation binding the contract event 0x6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a.
//
// Solidity: event UpdateMaxNumTxInChunk(uint256 oldMaxNumTxInChunk, uint256 newMaxNumTxInChunk)
func (_ScrollChain *ScrollChainFilterer) WatchUpdateMaxNumTxInChunk(opts *bind.WatchOpts, sink chan<- *ScrollChainUpdateMaxNumTxInChunk) (event.Subscription, error) {

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "UpdateMaxNumTxInChunk")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainUpdateMaxNumTxInChunk)
				if err := _ScrollChain.contract.UnpackLog(event, "UpdateMaxNumTxInChunk", log); err != nil {
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

// ParseUpdateMaxNumTxInChunk is a log parse operation binding the contract event 0x6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a.
//
// Solidity: event UpdateMaxNumTxInChunk(uint256 oldMaxNumTxInChunk, uint256 newMaxNumTxInChunk)
func (_ScrollChain *ScrollChainFilterer) ParseUpdateMaxNumTxInChunk(log types.Log) (*ScrollChainUpdateMaxNumTxInChunk, error) {
	event := new(ScrollChainUpdateMaxNumTxInChunk)
	if err := _ScrollChain.contract.UnpackLog(event, "UpdateMaxNumTxInChunk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainUpdateProverIterator is returned from FilterUpdateProver and is used to iterate over the raw logs and unpacked data for UpdateProver events raised by the ScrollChain contract.
type ScrollChainUpdateProverIterator struct {
	Event *ScrollChainUpdateProver // Event containing the contract specifics and raw log

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
func (it *ScrollChainUpdateProverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainUpdateProver)
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
		it.Event = new(ScrollChainUpdateProver)
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
func (it *ScrollChainUpdateProverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainUpdateProverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainUpdateProver represents a UpdateProver event raised by the ScrollChain contract.
type ScrollChainUpdateProver struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateProver is a free log retrieval operation binding the contract event 0x967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e.
//
// Solidity: event UpdateProver(address indexed account, bool status)
func (_ScrollChain *ScrollChainFilterer) FilterUpdateProver(opts *bind.FilterOpts, account []common.Address) (*ScrollChainUpdateProverIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "UpdateProver", accountRule)
	if err != nil {
		return nil, err
	}
	return &ScrollChainUpdateProverIterator{contract: _ScrollChain.contract, event: "UpdateProver", logs: logs, sub: sub}, nil
}

// WatchUpdateProver is a free log subscription operation binding the contract event 0x967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e.
//
// Solidity: event UpdateProver(address indexed account, bool status)
func (_ScrollChain *ScrollChainFilterer) WatchUpdateProver(opts *bind.WatchOpts, sink chan<- *ScrollChainUpdateProver, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "UpdateProver", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainUpdateProver)
				if err := _ScrollChain.contract.UnpackLog(event, "UpdateProver", log); err != nil {
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

// ParseUpdateProver is a log parse operation binding the contract event 0x967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e.
//
// Solidity: event UpdateProver(address indexed account, bool status)
func (_ScrollChain *ScrollChainFilterer) ParseUpdateProver(log types.Log) (*ScrollChainUpdateProver, error) {
	event := new(ScrollChainUpdateProver)
	if err := _ScrollChain.contract.UnpackLog(event, "UpdateProver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainUpdateSequencerIterator is returned from FilterUpdateSequencer and is used to iterate over the raw logs and unpacked data for UpdateSequencer events raised by the ScrollChain contract.
type ScrollChainUpdateSequencerIterator struct {
	Event *ScrollChainUpdateSequencer // Event containing the contract specifics and raw log

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
func (it *ScrollChainUpdateSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainUpdateSequencer)
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
		it.Event = new(ScrollChainUpdateSequencer)
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
func (it *ScrollChainUpdateSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainUpdateSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainUpdateSequencer represents a UpdateSequencer event raised by the ScrollChain contract.
type ScrollChainUpdateSequencer struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateSequencer is a free log retrieval operation binding the contract event 0x631cb110fbe6a87fba5414d6b2cff02264480535cd1f5abdbc4fa638bc0b5692.
//
// Solidity: event UpdateSequencer(address indexed account, bool status)
func (_ScrollChain *ScrollChainFilterer) FilterUpdateSequencer(opts *bind.FilterOpts, account []common.Address) (*ScrollChainUpdateSequencerIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "UpdateSequencer", accountRule)
	if err != nil {
		return nil, err
	}
	return &ScrollChainUpdateSequencerIterator{contract: _ScrollChain.contract, event: "UpdateSequencer", logs: logs, sub: sub}, nil
}

// WatchUpdateSequencer is a free log subscription operation binding the contract event 0x631cb110fbe6a87fba5414d6b2cff02264480535cd1f5abdbc4fa638bc0b5692.
//
// Solidity: event UpdateSequencer(address indexed account, bool status)
func (_ScrollChain *ScrollChainFilterer) WatchUpdateSequencer(opts *bind.WatchOpts, sink chan<- *ScrollChainUpdateSequencer, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "UpdateSequencer", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainUpdateSequencer)
				if err := _ScrollChain.contract.UnpackLog(event, "UpdateSequencer", log); err != nil {
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

// ParseUpdateSequencer is a log parse operation binding the contract event 0x631cb110fbe6a87fba5414d6b2cff02264480535cd1f5abdbc4fa638bc0b5692.
//
// Solidity: event UpdateSequencer(address indexed account, bool status)
func (_ScrollChain *ScrollChainFilterer) ParseUpdateSequencer(log types.Log) (*ScrollChainUpdateSequencer, error) {
	event := new(ScrollChainUpdateSequencer)
	if err := _ScrollChain.contract.UnpackLog(event, "UpdateSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollChainUpdateVerifierIterator is returned from FilterUpdateVerifier and is used to iterate over the raw logs and unpacked data for UpdateVerifier events raised by the ScrollChain contract.
type ScrollChainUpdateVerifierIterator struct {
	Event *ScrollChainUpdateVerifier // Event containing the contract specifics and raw log

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
func (it *ScrollChainUpdateVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollChainUpdateVerifier)
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
		it.Event = new(ScrollChainUpdateVerifier)
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
func (it *ScrollChainUpdateVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollChainUpdateVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollChainUpdateVerifier represents a UpdateVerifier event raised by the ScrollChain contract.
type ScrollChainUpdateVerifier struct {
	OldVerifier common.Address
	NewVerifier common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateVerifier is a free log retrieval operation binding the contract event 0x728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96.
//
// Solidity: event UpdateVerifier(address indexed oldVerifier, address indexed newVerifier)
func (_ScrollChain *ScrollChainFilterer) FilterUpdateVerifier(opts *bind.FilterOpts, oldVerifier []common.Address, newVerifier []common.Address) (*ScrollChainUpdateVerifierIterator, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _ScrollChain.contract.FilterLogs(opts, "UpdateVerifier", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return &ScrollChainUpdateVerifierIterator{contract: _ScrollChain.contract, event: "UpdateVerifier", logs: logs, sub: sub}, nil
}

// WatchUpdateVerifier is a free log subscription operation binding the contract event 0x728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96.
//
// Solidity: event UpdateVerifier(address indexed oldVerifier, address indexed newVerifier)
func (_ScrollChain *ScrollChainFilterer) WatchUpdateVerifier(opts *bind.WatchOpts, sink chan<- *ScrollChainUpdateVerifier, oldVerifier []common.Address, newVerifier []common.Address) (event.Subscription, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _ScrollChain.contract.WatchLogs(opts, "UpdateVerifier", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollChainUpdateVerifier)
				if err := _ScrollChain.contract.UnpackLog(event, "UpdateVerifier", log); err != nil {
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

// ParseUpdateVerifier is a log parse operation binding the contract event 0x728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96.
//
// Solidity: event UpdateVerifier(address indexed oldVerifier, address indexed newVerifier)
func (_ScrollChain *ScrollChainFilterer) ParseUpdateVerifier(log types.Log) (*ScrollChainUpdateVerifier, error) {
	event := new(ScrollChainUpdateVerifier)
	if err := _ScrollChain.contract.UnpackLog(event, "UpdateVerifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
