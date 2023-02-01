package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/control/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) {
	k.SetExecutionID(ctx, data.StartingExecutionId)

	for _, execution := range data.Executions {
		k.SetExecution(ctx, execution)
	}

	k.SetParams(ctx, data.Params)
}

func (k Keeper) ExportGenesis(ctx sdk.Context) (data *types.GenesisState) {
	startingExecutionID, _ := k.GetExecutionID(ctx)
	executions := k.GetExecutions(ctx)
	params := k.GetParams(ctx)

	return &types.GenesisState {
		StartingExecutionId: startingExecutionID,
		Executions:          executions,
		Params: params,
	}
}
