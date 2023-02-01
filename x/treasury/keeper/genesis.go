package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/treasury/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) {
	for _, record := range data.DistributionOperators {
		address, err := sdk.AccAddressFromBech32(record)

		if err != nil {
			panic(err)
		}

		k.AddOperator(ctx, address)
	}

	k.SetParams(ctx, data.Params)

	if data.Treasury.MintGenesisSupply {
		targetSupply := data.Treasury.TargetSupply

		err := k.MintCoins(ctx, targetSupply)
		if err != nil {
			panic(err)
		}

		data.Treasury.MintGenesisSupply = false
	}

	k.SetTreasury(ctx, data.Treasury)

	for _, disbursement := range data.DisbursementQueue {
		k.InsertDisbursementQueue(ctx, disbursement)
	}

	for _, burn := range data.BurnQueue {
		k.InsertBurnQueue(ctx, burn)
	}

	for _, disbursementReference := range data.DisbursementReferences {
		k.SetDisbursementReferenceInfo(ctx, disbursementReference)
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) (data *types.GenesisState) {
	treasury := k.GetTreasury(ctx)
	params := k.GetParams(ctx)
	operators := k.GetOperators(ctx)
	disbursements := k.GetDisbursements(ctx)
	burns := k.GetBurns(ctx)

	var disbursementReferences []types.Disbursement
	k.IterateDisbursementReferences(ctx, func(disbursement types.Disbursement) (stop bool) {
		disbursementReferences = append(disbursementReferences, disbursement)

		return false
	})

	return &types.GenesisState{
		Treasury: treasury, Params: params, DistributionOperators: operators, DisbursementQueue: disbursements, BurnQueue: burns, DisbursementReferences: disbursementReferences,
	}
}
