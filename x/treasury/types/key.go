package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

var (
	TreasuryKey = []byte{0x00}

	OperatorKeyPrefix          = []byte{0x10}
	DisbursementQueueKeyPrefix = []byte{0x12}
	BurnQueueKeyPrefix         = []byte{0x13}

	DisbursementReferenceKeyPrefix = []byte{0x15}

	StatusPresent = []byte{0x01}
)

var lenTime = len(sdk.FormatTimeBytes(time.Now()))

const (
	// ModuleName is the name of the module
	ModuleName  					= "treasury"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName

	TimeFormat = "2006-01-02T15:04:05.99999999999Z"
)


func GetOperatorKey(address sdk.AccAddress) []byte {
	return append(OperatorKeyPrefix, address...)
}

func GetOperatorIteratorKey() []byte {
	return OperatorKeyPrefix
}

func SplitOperatorKey(key []byte) sdk.AccAddress {
	return key[1:]
}

func DisbursementByTimeKey(endTime time.Time) []byte {
	return append(DisbursementQueueKeyPrefix, sdk.FormatTimeBytes(endTime)...)
}

func DisbursementQueueKey(address string, endTime time.Time) []byte {
	return append(DisbursementByTimeKey(endTime), address...)
}

func BurnQueueKey(endTime time.Time) []byte {
	return append(BurnQueueKeyPrefix, sdk.FormatTimeBytes(endTime)...)
}

func GetDisbursementReferenceKey(reference string) []byte {
	return append(DisbursementReferenceKeyPrefix, []byte(reference)...)
}

func GetDisbursementReferenceIteratorKey() []byte {
	return DisbursementReferenceKeyPrefix
}

func SplitDisbursementReferenceKey(key []byte) (string) {
	return string(key[1:])
}