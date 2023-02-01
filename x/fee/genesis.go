package fee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/fee/keeper"
	"github.com/melechain/mele/x/fee/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, data *types.GenesisState) {
	k.SetParams(ctx, data.Params)

	for _, record := range data.FeeExcludedMessages {
		k.SetFeeExcludedMessage(ctx, record)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)

	return &types.GenesisState{
		Params: params,
		FeeExcludedMessages: k.GetFeeExcludedMessages(ctx),
	}
}