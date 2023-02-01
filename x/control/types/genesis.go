package types

import (
	"bytes"
)

func NewGenesisState(startingExecutionID uint64, params Params) *GenesisState {
	return &GenesisState{
		StartingExecutionId: startingExecutionID,
		Params:              params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		DefaultStartingExecutionID,
		DefaultParams(),
	)
}

func (data GenesisState) Equal(data2 GenesisState) bool {
	b1 := ModuleCdc.MustMarshal(&data)
	b2 := ModuleCdc.MustMarshal(&data2)
	return bytes.Equal(b1, b2)
}

func (data GenesisState) IsEmpty() bool {
	return data.Equal(GenesisState{})
}

func ValidateGenesis(data *GenesisState) error {
	return nil
}
