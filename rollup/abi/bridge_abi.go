package bridgeabi

import (
	"math/big"

	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
)

var (
	// ScrollChainABI holds information about ScrollChain's context and available invokable methods.
	ScrollChainABI *abi.ABI
	// L1MessageQueueABI holds information about L1MessageQueue contract's context and available invokable methods.
	L1MessageQueueABI *abi.ABI
	// L2GasPriceOracleABI holds information about L2GasPriceOracle's context and available invokable methods.
	L2GasPriceOracleABI *abi.ABI

	// L2ScrollMessengerABI holds information about L2ScrollMessenger's context and available invokable methods.
	L2ScrollMessengerABI *abi.ABI
	// L1GasPriceOracleABI holds information about L1GasPriceOracle's context and available invokable methods.
	L1GasPriceOracleABI *abi.ABI
	// L2MessageQueueABI holds information about L2MessageQueue contract's context and available invokable methods.
	L2MessageQueueABI *abi.ABI

	// L1CommitBatchEventSignature = keccak256("CommitBatch(uint256,bytes32)")
	L1CommitBatchEventSignature common.Hash
	// L1FinalizeBatchEventSignature = keccak256("FinalizeBatch(uint256,bytes32,bytes32,bytes32)")
	L1FinalizeBatchEventSignature common.Hash
	// L1QueueTransactionEventSignature = keccak256("QueueTransaction(address,address,uint256,uint64,uint256,bytes)")
	L1QueueTransactionEventSignature common.Hash

	// L2SentMessageEventSignature = keccak256("SentMessage(address,address,uint256,uint256,uint256,bytes,uint256,uint256)")
	L2SentMessageEventSignature common.Hash
	// L2RelayedMessageEventSignature = keccak256("RelayedMessage(bytes32)")
	L2RelayedMessageEventSignature common.Hash
	// L2FailedRelayedMessageEventSignature = keccak256("FailedRelayedMessage(bytes32)")
	L2FailedRelayedMessageEventSignature common.Hash

	// L2AppendMessageEventSignature = keccak256("AppendMessage(uint256,bytes32)")
	L2AppendMessageEventSignature common.Hash
)

func init() {
	ScrollChainABI, _ = ScrollChainTwoStepCommitMetaData.GetAbi()
	L1MessageQueueABI, _ = L1MessageQueueMetaData.GetAbi()
	L2GasPriceOracleABI, _ = L2GasPriceOracleMetaData.GetAbi()

	L2ScrollMessengerABI, _ = L2ScrollMessengerMetaData.GetAbi()
	L2MessageQueueABI, _ = L2MessageQueueMetaData.GetAbi()
	L1GasPriceOracleABI, _ = L1GasPriceOracleMetaData.GetAbi()

	L1CommitBatchEventSignature = ScrollChainABI.Events["CommitBatch"].ID
	L1FinalizeBatchEventSignature = ScrollChainABI.Events["FinalizeBatch"].ID

	L1QueueTransactionEventSignature = L1MessageQueueABI.Events["QueueTransaction"].ID

	L2SentMessageEventSignature = L2ScrollMessengerABI.Events["SentMessage"].ID
	L2RelayedMessageEventSignature = L2ScrollMessengerABI.Events["RelayedMessage"].ID
	L2FailedRelayedMessageEventSignature = L2ScrollMessengerABI.Events["FailedRelayedMessage"].ID

	L2AppendMessageEventSignature = L2MessageQueueABI.Events["AppendMessage"].ID
}

// Generated manually from abigen and only necessary events and mutable calls are kept.

// ScrollChainMetaData contains all meta data concerning the ScrollChain contract.
var ScrollChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"CommitBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"}],\"name\":\"FinalizeBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"RevertBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxNumL2TxInChunk\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxNumL2TxInChunk\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxNumL2TxInChunk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateProver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"UpdateVerifier\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_version\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_parentBatchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_chunks\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"_skippedL1MessageBitmap\",\"type\":\"bytes\"}],\"name\":\"commitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"committedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_withdrawRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_aggrProof\",\"type\":\"bytes\"}],\"name\":\"finalizeBatchWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"finalizedStateRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_stateRoot\",\"type\":\"bytes32\"}],\"name\":\"importGenesisBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageQueue\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxNumL2TxInChunk\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"}],\"name\":\"isBatchFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isProver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFinalizedBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"layer2ChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumL2TxInChunk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"revertBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxNumL2TxInChunk\",\"type\":\"uint256\"}],\"name\":\"updateMaxNumL2TxInChunk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"updateProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"updateSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

var ScrollChainTwoStepCommitMetaData = &bind.MetaData{
	ABI: `[
		{
		  "inputs": [
			{
			  "internalType": "uint64",
			  "name": "_chainId",
			  "type": "uint64"
			}
		  ],
		  "stateMutability": "nonpayable",
		  "type": "constructor"
		},
		{
		  "inputs": [],
		  "name": "CommittedProof",
		  "type": "error"
		},
		{
		  "inputs": [],
		  "name": "CommittedProofHash",
		  "type": "error"
		},
		{
		  "inputs": [],
		  "name": "CommittedTimeout",
		  "type": "error"
		},
		{
		  "inputs": [],
		  "name": "ErrCommitProof",
		  "type": "error"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes32",
			  "name": "",
			  "type": "bytes32"
			}
		  ],
		  "name": "ErrorBatchHash",
		  "type": "error"
		},
		{
		  "inputs": [],
		  "name": "InsufficientPledge",
		  "type": "error"
		},
		{
		  "inputs": [],
		  "name": "OnlyDeposit",
		  "type": "error"
		},
		{
		  "inputs": [],
		  "name": "SlotAdapterEmpty",
		  "type": "error"
		},
		{
		  "inputs": [],
		  "name": "SubmitProofEarly",
		  "type": "error"
		},
		{
		  "inputs": [],
		  "name": "SubmitProofTooLate",
		  "type": "error"
		},
		{
		  "inputs": [],
		  "name": "UnCommittedProofHash",
		  "type": "error"
		},
		{
		  "inputs": [],
		  "name": "ZeroAddress",
		  "type": "error"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": true,
			  "internalType": "uint256",
			  "name": "batchIndex",
			  "type": "uint256"
			},
			{
			  "indexed": true,
			  "internalType": "bytes32",
			  "name": "batchHash",
			  "type": "bytes32"
			}
		  ],
		  "name": "CommitBatch",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": true,
			  "internalType": "uint256",
			  "name": "batchIndex",
			  "type": "uint256"
			},
			{
			  "indexed": true,
			  "internalType": "bytes32",
			  "name": "batchHash",
			  "type": "bytes32"
			},
			{
			  "indexed": false,
			  "internalType": "bytes32",
			  "name": "stateRoot",
			  "type": "bytes32"
			},
			{
			  "indexed": false,
			  "internalType": "bytes32",
			  "name": "withdrawRoot",
			  "type": "bytes32"
			}
		  ],
		  "name": "FinalizeBatch",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "uint8",
			  "name": "version",
			  "type": "uint8"
			}
		  ],
		  "name": "Initialized",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "previousOwner",
			  "type": "address"
			},
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "newOwner",
			  "type": "address"
			}
		  ],
		  "name": "OwnershipTransferred",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "address",
			  "name": "account",
			  "type": "address"
			}
		  ],
		  "name": "Paused",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": true,
			  "internalType": "uint256",
			  "name": "batchIndex",
			  "type": "uint256"
			},
			{
			  "indexed": true,
			  "internalType": "bytes32",
			  "name": "batchHash",
			  "type": "bytes32"
			}
		  ],
		  "name": "RevertBatch",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "uint8",
			  "name": "newProofCommitEpoch",
			  "type": "uint8"
			}
		  ],
		  "name": "SetProofCommitEpoch",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "uint8",
			  "name": "newProofHashCommitEpoch",
			  "type": "uint8"
			}
		  ],
		  "name": "SetProofHashCommitEpoch",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "address",
			  "name": "_prover",
			  "type": "address"
			},
			{
			  "indexed": false,
			  "internalType": "uint256",
			  "name": "batchIndex",
			  "type": "uint256"
			},
			{
			  "indexed": false,
			  "internalType": "bytes32",
			  "name": "_proofHash",
			  "type": "bytes32"
			}
		  ],
		  "name": "SubmitProofHash",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "address",
			  "name": "account",
			  "type": "address"
			}
		  ],
		  "name": "Unpaused",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "uint256",
			  "name": "oldMaxNumTxInChunk",
			  "type": "uint256"
			},
			{
			  "indexed": false,
			  "internalType": "uint256",
			  "name": "newMaxNumTxInChunk",
			  "type": "uint256"
			}
		  ],
		  "name": "UpdateMaxNumTxInChunk",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "account",
			  "type": "address"
			},
			{
			  "indexed": false,
			  "internalType": "bool",
			  "name": "status",
			  "type": "bool"
			}
		  ],
		  "name": "UpdateProver",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "account",
			  "type": "address"
			},
			{
			  "indexed": false,
			  "internalType": "bool",
			  "name": "status",
			  "type": "bool"
			}
		  ],
		  "name": "UpdateSequencer",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "oldVerifier",
			  "type": "address"
			},
			{
			  "indexed": true,
			  "internalType": "address",
			  "name": "newVerifier",
			  "type": "address"
			}
		  ],
		  "name": "UpdateVerifier",
		  "type": "event"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "_account",
			  "type": "address"
			}
		  ],
		  "name": "addProver",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "_account",
			  "type": "address"
			}
		  ],
		  "name": "addSequencer",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "name": "blockCommitBatches",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint8",
			  "name": "_version",
			  "type": "uint8"
			},
			{
			  "internalType": "bytes",
			  "name": "_parentBatchHeader",
			  "type": "bytes"
			},
			{
			  "internalType": "bytes[]",
			  "name": "_chunks",
			  "type": "bytes[]"
			},
			{
			  "internalType": "bytes",
			  "name": "_skippedL1MessageBitmap",
			  "type": "bytes"
			}
		  ],
		  "name": "commitBatch",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "name": "committedBatchInfo",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "blockNumber",
			  "type": "uint256"
			},
			{
			  "internalType": "bool",
			  "name": "proofSubmitted",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "name": "committedBatches",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "",
			  "type": "bytes32"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_batchHeader",
			  "type": "bytes"
			},
			{
			  "internalType": "bytes32",
			  "name": "_prevStateRoot",
			  "type": "bytes32"
			},
			{
			  "internalType": "bytes32",
			  "name": "_postStateRoot",
			  "type": "bytes32"
			},
			{
			  "internalType": "bytes32",
			  "name": "_withdrawRoot",
			  "type": "bytes32"
			},
			{
			  "internalType": "bytes",
			  "name": "_aggrProof",
			  "type": "bytes"
			}
		  ],
		  "name": "finalizeBatchWithProof",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "name": "finalizedStateRoots",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "",
			  "type": "bytes32"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getBatchToProve",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "ideDeposit",
		  "outputs": [
			{
			  "internalType": "contract IDeposit",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_batchHeader",
			  "type": "bytes"
			},
			{
			  "internalType": "bytes32",
			  "name": "_stateRoot",
			  "type": "bytes32"
			}
		  ],
		  "name": "importGenesisBatch",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "incorrectProofHashPunishAmount",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "_messageQueue",
			  "type": "address"
			},
			{
			  "internalType": "address",
			  "name": "_verifier",
			  "type": "address"
			},
			{
			  "internalType": "uint256",
			  "name": "_maxNumTxInChunk",
			  "type": "uint256"
			}
		  ],
		  "name": "initialize",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "isAllLiquidated",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "_batchIndex",
			  "type": "uint256"
			}
		  ],
		  "name": "isBatchFinalized",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "batchIndex",
			  "type": "uint256"
			}
		  ],
		  "name": "isCommitProofAllowed",
		  "outputs": [
			{
			  "internalType": "uint8",
			  "name": "",
			  "type": "uint8"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "batchIndex",
			  "type": "uint256"
			}
		  ],
		  "name": "isCommitProofHashAllowed",
		  "outputs": [],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "name": "isProver",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "name": "isSequencer",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "lastFinalizedBatchIndex",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "layer2ChainId",
		  "outputs": [
			{
			  "internalType": "uint64",
			  "name": "",
			  "type": "uint64"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "maxNumTxInChunk",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "messageQueue",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "minDeposit",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "noProofPunishAmount",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "owner",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "paused",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "proofCommitEpoch",
		  "outputs": [
			{
			  "internalType": "uint8",
			  "name": "",
			  "type": "uint8"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "proofHashCommitEpoch",
		  "outputs": [
			{
			  "internalType": "uint8",
			  "name": "",
			  "type": "uint8"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "name": "proofNum",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			},
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "name": "proverCommitProofHash",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "proofHash",
			  "type": "bytes32"
			},
			{
			  "internalType": "uint256",
			  "name": "blockNumber",
			  "type": "uint256"
			},
			{
			  "internalType": "bool",
			  "name": "proof",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "name": "proverLastLiquidated",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			},
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "name": "proverLiquidation",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "prover",
			  "type": "address"
			},
			{
			  "internalType": "bool",
			  "name": "isSubmittedProofHash",
			  "type": "bool"
			},
			{
			  "internalType": "uint256",
			  "name": "submitHashBlockNumber",
			  "type": "uint256"
			},
			{
			  "internalType": "bool",
			  "name": "isSubmittedProof",
			  "type": "bool"
			},
			{
			  "internalType": "uint256",
			  "name": "submitProofBlockNumber",
			  "type": "uint256"
			},
			{
			  "internalType": "bool",
			  "name": "isLiquidated",
			  "type": "bool"
			},
			{
			  "internalType": "uint64",
			  "name": "finalNewBatch",
			  "type": "uint64"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			},
			{
			  "internalType": "bytes32",
			  "name": "",
			  "type": "bytes32"
			}
		  ],
		  "name": "proverPosition",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "_account",
			  "type": "address"
			}
		  ],
		  "name": "removeProver",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "_account",
			  "type": "address"
			}
		  ],
		  "name": "removeSequencer",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "renounceOwnership",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "_batchHeader",
			  "type": "bytes"
			},
			{
			  "internalType": "uint256",
			  "name": "_count",
			  "type": "uint256"
			}
		  ],
		  "name": "revertBatch",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "contract IDeposit",
			  "name": "_ideDeposit",
			  "type": "address"
			}
		  ],
		  "name": "setDeposit",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "_amount",
			  "type": "uint256"
			}
		  ],
		  "name": "setIncorrectProofPunishAmount",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "_amount",
			  "type": "uint256"
			}
		  ],
		  "name": "setMinDeposit",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "_amount",
			  "type": "uint256"
			}
		  ],
		  "name": "setNoProofPunishAmount",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bool",
			  "name": "_status",
			  "type": "bool"
			}
		  ],
		  "name": "setPause",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint8",
			  "name": "_newCommitEpoch",
			  "type": "uint8"
			}
		  ],
		  "name": "setProofCommitEpoch",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint8",
			  "name": "_newCommitEpoch",
			  "type": "uint8"
			}
		  ],
		  "name": "setProofHashCommitEpoch",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "contract ISlotAdapter",
			  "name": "_slotAdapter",
			  "type": "address"
			}
		  ],
		  "name": "setSlotAdapter",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "_account",
			  "type": "address"
			}
		  ],
		  "name": "settle",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "slotAdapter",
		  "outputs": [
			{
			  "internalType": "contract ISlotAdapter",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "batchIndex",
			  "type": "uint256"
			},
			{
			  "internalType": "bytes32",
			  "name": "_proofHash",
			  "type": "bytes32"
			}
		  ],
		  "name": "submitProofHash",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "newOwner",
			  "type": "address"
			}
		  ],
		  "name": "transferOwnership",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "_maxNumTxInChunk",
			  "type": "uint256"
			}
		  ],
		  "name": "updateMaxNumTxInChunk",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "_newVerifier",
			  "type": "address"
			}
		  ],
		  "name": "updateVerifier",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "verifier",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "name": "withdrawRoots",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "",
			  "type": "bytes32"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		}
	  ]`,
}

// L1ScrollMessengerMetaData contains all meta data concerning the L1ScrollMessenger contract.
var L1ScrollMessengerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldFeeVault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"name\":\"UpdateFeeVault\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageQueue\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isL1MessageRelayed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isL1MessageSent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isL2MessageExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"merkleProof\",\"type\":\"bytes\"}],\"internalType\":\"struct IL1ScrollMessenger.L2MessageProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"relayMessageWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_newGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"replayMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"name\":\"updateFeeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]\n",
}

// L1MessageQueueMetaData contains all meta data concerning the L1MessageQueue contract.
var L1MessageQueueMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skippedBitmap\",\"type\":\"uint256\"}],\"name\":\"DequeueTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"queueIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"QueueTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldGateway\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newGateway\",\"type\":\"address\"}],\"name\":\"UpdateEnforcedTxGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldGasOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newGasOracle\",\"type\":\"address\"}],\"name\":\"UpdateGasOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldMaxGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxGasLimit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"appendCrossDomainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"appendEnforcedTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"calculateIntrinsicGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"computeTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enforcedTxGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"estimateCrossDomainMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"}],\"name\":\"getCrossDomainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_scrollChain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_enforcedTxGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gasOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxGasLimit\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextCrossDomainMessageIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingQueueIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_skippedBitmap\",\"type\":\"uint256\"}],\"name\":\"popCrossDomainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scrollChain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGateway\",\"type\":\"address\"}],\"name\":\"updateEnforcedTxGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGasOracle\",\"type\":\"address\"}],\"name\":\"updateGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"updateMaxGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L2GasPriceOracleMetaData contains all meta data concerning the L2GasPriceOracle contract.
var L2GasPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txGasContractCreation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zeroGas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonZeroGas\",\"type\":\"uint256\"}],\"name\":\"IntrinsicParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l2BaseFee\",\"type\":\"uint256\"}],\"name\":\"L2BaseFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldWhitelist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newWhitelist\",\"type\":\"address\"}],\"name\":\"UpdateWhitelist\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"calculateIntrinsicGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"estimateCrossDomainMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_txGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_txGasContractCreation\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_zeroGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonZeroGas\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"intrinsicParams\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"txGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"txGasContractCreation\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zeroGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonZeroGas\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_txGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_txGasContractCreation\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_zeroGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonZeroGas\",\"type\":\"uint64\"}],\"name\":\"setIntrinsicParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BaseFee\",\"type\":\"uint256\"}],\"name\":\"setL2BaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newWhitelist\",\"type\":\"address\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"contract IWhitelist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]\n",
}

// L2ScrollMessengerMetaData contains all meta data concerning the L2ScrollMessenger contract.
var L2ScrollMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blockContainer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gasOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageQueue\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldFeeVault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"name\":\"UpdateFeeVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxFailedExecutionTimes\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxFailedExecutionTimes\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"blockContainer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeVault\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isL1MessageExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isL2MessageSent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"l1MessageFailedTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxFailedExecutionTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"relayMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"stateRootProof\",\"type\":\"bytes\"}],\"internalType\":\"struct IL2ScrollMessenger.L1MessageProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"retryMessageWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"name\":\"updateFeeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxFailedExecutionTimes\",\"type\":\"uint256\"}],\"name\":\"updateMaxFailedExecutionTimes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"verifyMessageExecutionStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"verifyMessageInclusionStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]\n",
}

// L1BlockContainerMetaData contains all meta data concerning the L1BlockContainer contract.
var L1BlockContainerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"name\":\"ImportBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldWhitelist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newWhitelist\",\"type\":\"address\"}],\"name\":\"UpdateWhitelist\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"}],\"name\":\"getBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"}],\"name\":\"getStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_blockHeaderRLP\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"_updateGasPriceOracle\",\"type\":\"bool\"}],\"name\":\"importBlockHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_startBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_startBlockHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_startBlockTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"_startBlockBaseFee\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"_startStateRoot\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"metadata\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"baseFee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"stateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newWhitelist\",\"type\":\"address\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"contract IWhitelist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]\n",
}

// L2MessageQueueMetaData contains all meta data concerning the L2MessageQueue contract.
var L2MessageQueueMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"AppendMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"appendMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"branches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextMessageIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"updateMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]\n",
}

// L1GasPriceOracleMetaData contains all meta data concerning the L1GasPriceOracle contract.
var L1GasPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1BaseFee\",\"type\":\"uint256\"}],\"name\":\"L1BaseFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"overhead\",\"type\":\"uint256\"}],\"name\":\"OverheadUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"ScalarUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldWhitelist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newWhitelist\",\"type\":\"address\"}],\"name\":\"UpdateWhitelist\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getL1Fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getL1GasUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1BaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"overhead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scalar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l1BaseFee\",\"type\":\"uint256\"}],\"name\":\"setL1BaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_overhead\",\"type\":\"uint256\"}],\"name\":\"setOverhead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_scalar\",\"type\":\"uint256\"}],\"name\":\"setScalar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newWhitelist\",\"type\":\"address\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"contract IWhitelist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]\n",
}

// IL1ScrollMessengerL2MessageProof is an auto generated low-level Go binding around an user-defined struct.
type IL1ScrollMessengerL2MessageProof struct {
	BatchIndex  *big.Int
	MerkleProof []byte
}

// IScrollChainBatch is an auto generated low-level Go binding around an user-defined struct.
type IScrollChainBatch struct {
	Blocks           []IScrollChainBlockContext
	PrevStateRoot    common.Hash
	NewStateRoot     common.Hash
	WithdrawTrieRoot common.Hash
	BatchIndex       uint64
	ParentBatchHash  common.Hash
	L2Transactions   []byte
}

// IScrollChainBlockContext is an auto generated low-level Go binding around an user-defined struct.
type IScrollChainBlockContext struct {
	BlockHash       common.Hash
	ParentHash      common.Hash
	BlockNumber     uint64
	Timestamp       uint64
	BaseFee         *big.Int
	GasLimit        uint64
	NumTransactions uint16
	NumL1Messages   uint16
}

// L1CommitBatchEvent represents a CommitBatch event raised by the ScrollChain contract.
type L1CommitBatchEvent struct {
	BatchIndex *big.Int
	BatchHash  common.Hash
}

// L1FinalizeBatchEvent represents a FinalizeBatch event raised by the ScrollChain contract.
type L1FinalizeBatchEvent struct {
	BatchIndex   *big.Int
	BatchHash    common.Hash
	StateRoot    common.Hash
	WithdrawRoot common.Hash
}

// L1RevertBatchEvent represents a RevertBatch event raised by the ScrollChain contract.
type L1RevertBatchEvent struct {
	BatchHash common.Hash
}

// L1QueueTransactionEvent represents a QueueTransaction event raised by the L1MessageQueue contract.
type L1QueueTransactionEvent struct {
	Sender     common.Address
	Target     common.Address
	Value      *big.Int
	QueueIndex uint64
	GasLimit   *big.Int
	Data       []byte
}

// L1SentMessageEvent represents a SentMessage event raised by the L1ScrollMessenger contract.
type L1SentMessageEvent struct {
	Sender       common.Address
	Target       common.Address
	Value        *big.Int
	MessageNonce *big.Int
	GasLimit     *big.Int
	Message      []byte
}

// L1FailedRelayedMessageEvent represents a FailedRelayedMessage event raised by the L1ScrollMessenger contract.
type L1FailedRelayedMessageEvent struct {
	MessageHash common.Hash
}

// L1RelayedMessageEvent represents a RelayedMessage event raised by the L1ScrollMessenger contract.
type L1RelayedMessageEvent struct {
	MessageHash common.Hash
}

// L2AppendMessageEvent represents a AppendMessage event raised by the L2MessageQueue contract.
type L2AppendMessageEvent struct {
	Index       *big.Int
	MessageHash common.Hash
}

// L2SentMessageEvent represents a SentMessage event raised by the L2ScrollMessenger contract.
type L2SentMessageEvent struct {
	Sender       common.Address
	Target       common.Address
	Value        *big.Int
	MessageNonce *big.Int
	GasLimit     *big.Int
	Message      []byte
}

// L2FailedRelayedMessageEvent represents a FailedRelayedMessage event raised by the L2ScrollMessenger contract.
type L2FailedRelayedMessageEvent struct {
	MessageHash common.Hash
}

// L2RelayedMessageEvent represents a RelayedMessage event raised by the L2ScrollMessenger contract.
type L2RelayedMessageEvent struct {
	MessageHash common.Hash
}
