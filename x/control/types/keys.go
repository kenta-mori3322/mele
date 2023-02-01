package types

import (
	"encoding/binary"
)

const (
	// ModuleName is the name of the module
	ModuleName = "control"

	// StoreKey is the store key string for gov
	StoreKey = ModuleName

	// RouterKey is the message route for gov
	RouterKey = ModuleName

	// QuerierRoute is the querier route for gov
	QuerierRoute = ModuleName

	// DefaultParamspace default name for parameter store
	DefaultParamspace = ModuleName
)

// Keys for control store
// Items are stored with the following key: values
//
// - 0x00<executionID_Bytes>: Execution
//
// - 0x01: nextExecutionID
var (
	ExecutionsKeyPrefix = []byte{0x00}
	ExecutionIDKey      = []byte{0x01}
)

func GetExecutionIDBytes(executionID uint64) (executionIDBz []byte) {
	executionIDBz = make([]byte, 8)
	binary.BigEndian.PutUint64(executionIDBz, executionID)
	return
}

func ExecutionKey(executionID uint64) []byte {
	return append(ExecutionsKeyPrefix, GetExecutionIDBytes(executionID)...)
}

func GetExecutionIDFromBytes(bz []byte) (executionID uint64) {
	return binary.BigEndian.Uint64(bz)
}