package common

import (
	"math/big"
)

const (
	RewardMasterPercent        = 40
	RewardVoterPercent         = 50
	RewardFoundationPercent    = 10
	HexSignMethod              = "e341eaa4"
	HexSetSecret               = "34d38600"
	HexSetOpening              = "e11f5ba2"
	EpocBlockSecret            = 896
	EpocBlockOpening           = 898
	EpocBlockRandomize         = 900
	MaxMasternodes             = 150
	LimitPenaltyEpoch          = 4
	BlocksPerYear              = uint64(10512000)
	LimitThresholdNonceInQueue = 10
	DefaultMinGasPrice         = 250000000
	MergeSignRange             = 15
	RangeReturnSigner          = 150
	MinimunMinerBlockPerEpoch  = 1
)

// spend 3 years for prepairing these features
var TIP2019Block = big.NewInt(31536000)
var TIPSigning = big.NewInt(31536000)
var TIPRandomize = big.NewInt(31536000)
var BlackListHFNumber = uint64(0)
var IsTestnet bool = false
var StoreRewardFolder string
var RollbackHash Hash
var MinGasPrice = big.NewInt(DefaultMinGasPrice)
var TRC21IssuerSMCTestNet = HexToAddress("0x7081C72c9DC44686C7B7EAB1d338EA137Fa9f0D3")
var TRC21IssuerSMC = HexToAddress("0x8c0faeb5C6bEd2129b8674F262Fd45c4e9468bee")
var TRC21GasPriceBefore = big.NewInt(0)
var TRC21GasPrice = big.NewInt(0)
var Blacklist = map[Address]bool{
	// HexToAddress("0x5248bfb72fd4f234e062d3e9bb76f08643004fcd"): true,
}
var TIPTRC21Fee = big.NewInt(0)
