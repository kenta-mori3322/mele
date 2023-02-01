package types


// NewGenesisState creates a new GenesisState object
func NewGenesisState(treasury Treasury, params Params, operators []string, disbursements []Disbursement, burns []Burn, references []Disbursement) GenesisState {
	return GenesisState{
		Treasury: 	treasury,
		Params: 	params,
		DistributionOperators: 	operators,
		DisbursementQueue: disbursements,
		BurnQueue: burns,
		DisbursementReferences: references,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Treasury: 	DefaultInitialTreasury(),
		Params: 	DefaultParams(),
		DistributionOperators: 	DefaultOperators(),
		DisbursementQueue: []Disbursement{},
		BurnQueue: []Burn{},
		DisbursementReferences: []Disbursement{},
	}
}

// ValidateGenesis validates the treasury genesis parameters
func ValidateGenesis(data *GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}

	return nil
}