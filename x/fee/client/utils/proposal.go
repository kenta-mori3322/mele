package utils

import (
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/melechain/mele/x/fee/types"
)

func ParseAddFeeExcludedMessageProposalWithDeposit(cdc codec.JSONCodec, proposalFile string) (types.AddFeeExcludedMessageProposalWithDeposit, error) {
	proposal := types.AddFeeExcludedMessageProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

func ParseRemoveFeeExcludedMessageProposalWithDeposit(cdc codec.JSONCodec, proposalFile string) (types.RemoveFeeExcludedMessageProposalWithDeposit, error) {
	proposal := types.RemoveFeeExcludedMessageProposalWithDeposit{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
