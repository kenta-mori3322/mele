package utils

import (
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/melechain/mele/x/treasury/types"
)

func ParseMintTreasurySupplyProposalWithDeposit(cdc codec.JSONCodec, proposalFile string) (types.MintTreasurySupplyProposalWithDeposit, error) {
	proposal := types.MintTreasurySupplyProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

func ParseBurnTreasurySupplyProposalWithDeposit(cdc codec.JSONCodec, proposalFile string) (types.BurnTreasurySupplyProposalWithDeposit, error) {
	proposal := types.BurnTreasurySupplyProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
