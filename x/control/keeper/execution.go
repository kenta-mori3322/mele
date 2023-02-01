package keeper

import (
	"fmt"

	"github.com/melechain/mele/x/control/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (keeper Keeper) HandleSubmitExecution(ctx sdk.Context, content govtypes.Content, executor sdk.AccAddress) (types.Execution, error) {
	if !keeper.GetParams(ctx).Enabled {
		return types.Execution{}, sdkerrors.Wrap(types.ErrControlNotEnabled, "Control module is disabled.")
	}

	if !keeper.IsManager(ctx, executor) {
		return types.Execution{}, sdkerrors.Wrap(types.ErrNotManager, "Executor is not a manager.")
	}

	if !keeper.router.HasRoute(content.ProposalRoute()) {
		return types.Execution{}, sdkerrors.Wrap(types.ErrNoExecutionHandlerExists, content.ProposalRoute())
	}

	handler := keeper.router.GetRoute(content.ProposalRoute())
	if err := handler(ctx, content); err != nil {
		return types.Execution{}, sdkerrors.Wrap(types.ErrInvalidExecutionContent, err.Error())
	}

	executionID, err := keeper.GetExecutionID(ctx)
	if err != nil {
		return types.Execution{}, err
	}

	submitTime := ctx.BlockHeader().Time

	execution, err := types.NewExecution(content, executionID, submitTime.String(), executor.String())

	keeper.SetExecution(ctx, execution)
	keeper.SetExecutionID(ctx, executionID+1)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSubmitExecution,
			sdk.NewAttribute(types.AttributeKeyExecutionID, fmt.Sprintf("%d", executionID)),
		),
	)

	return execution, nil
}

func (keeper Keeper) GetExecution(ctx sdk.Context, executionID uint64) (execution types.Execution, ok bool) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(types.ExecutionKey(executionID))
	if bz == nil {
		return
	}
	keeper.cdc.MustUnmarshalLengthPrefixed(bz, &execution)
	return execution, true
}

func (keeper Keeper) SetExecution(ctx sdk.Context, execution types.Execution) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalLengthPrefixed(&execution)
	store.Set(types.ExecutionKey(execution.Id), bz)
}

func (keeper Keeper) DeleteExecution(ctx sdk.Context, executionID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	_, ok := keeper.GetExecution(ctx, executionID)
	if !ok {
		panic(fmt.Sprintf("couldn't find execution with id#%d", executionID))
	}
	store.Delete(types.ExecutionKey(executionID))
}

func (keeper Keeper) IterateExecutions(ctx sdk.Context, cb func(execution types.Execution) (stop bool)) {
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ExecutionsKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var execution types.Execution
		keeper.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &execution)

		if cb(execution) {
			break
		}
	}
}

func (keeper Keeper) GetExecutions(ctx sdk.Context) (executions types.Executions) {
	keeper.IterateExecutions(ctx, func(execution types.Execution) bool {
		executions = append(executions, execution)
		return false
	})
	return
}

func (keeper Keeper) GetExecutionID(ctx sdk.Context) (executionID uint64, err error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(types.ExecutionIDKey)
	if bz == nil {
		return 0, sdkerrors.Wrap(types.ErrInvalidGenesis, "initial execution ID hasn't been set")
	}

	executionID = types.GetExecutionIDFromBytes(bz)
	return executionID, nil
}

func (keeper Keeper) SetExecutionID(ctx sdk.Context, executionID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	store.Set(types.ExecutionIDKey, types.GetExecutionIDBytes(executionID))
}
