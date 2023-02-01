package types

func NewGenesisState(params Params, feeExcludedMessages []string) GenesisState {
	return GenesisState{
		Params: params,
		FeeExcludedMessages: feeExcludedMessages,
	}
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
		FeeExcludedMessages: DefaultFeeExcludedMessages,
	}
}

func ValidateGenesis(data *GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}

	return nil
}