package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/treasury/types"
)

func (k Keeper) HandleBurn(ctx sdk.Context, operator sdk.AccAddress, amount sdk.Coins) error {
	if !k.IsOperator(ctx, operator) {
		return types.ErrNotOperator
	}

	scheduledFor := ctx.BlockTime()

	if amount.IsAnyGTE(k.DisbursementDelayThreshold(ctx)) {
		scheduledFor = scheduledFor.Add(k.DisbursementDelayDuration(ctx))
	}

	for k.HasBurnInQueue(ctx, scheduledFor) {
		scheduledFor = scheduledFor.Add(time.Millisecond)
	}

	burn := types.Burn{
		Operator:     operator.String(),
		Amount:       amount,
		ScheduledAt:  ctx.BlockTime().String(),
		ScheduledFor: scheduledFor.String(),
	}

	k.InsertBurnQueue(ctx, burn)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBurn,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
			sdk.NewAttribute(types.AttributeKeyScheduledFor, scheduledFor.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
			sdk.NewAttribute(types.AttributeKeySender, operator.String()),
		),
	})

	return nil
}

func (k Keeper) HandleApproveBurn(ctx sdk.Context, manager sdk.AccAddress, scheduledFor time.Time) error {
	if !k.IsManager(ctx, manager) {
		return types.ErrNotManager
	}

	if !k.HasBurnInQueue(ctx, scheduledFor) {
		return types.ErrBurnNotScheduled
	}

	burn := k.GetBurnFromQueue(ctx, scheduledFor)

	err := k.BurnFunds(ctx, manager, burn.Amount)
	if err != nil {
		k.Logger(ctx).Info(err.Error())
	}

	k.RemoveFromBurnQueue(ctx, scheduledFor)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeApproveBurn,
			sdk.NewAttribute(types.AttributeKeyScheduledFor, scheduledFor.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, burn.Amount.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
			sdk.NewAttribute(types.AttributeKeySender, manager.String()),
		),
	})

	return nil
}

func (k Keeper) HandleCancelBurn(ctx sdk.Context, manager sdk.AccAddress, scheduledFor time.Time) error {
	if !k.IsManager(ctx, manager) {
		return types.ErrNotManager
	}

	if !k.HasBurnInQueue(ctx, scheduledFor) {
		return types.ErrBurnNotScheduled
	}

	k.RemoveFromBurnQueue(ctx, scheduledFor)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCancelBurn,
			sdk.NewAttribute(types.AttributeKeyScheduledFor, scheduledFor.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
			sdk.NewAttribute(types.AttributeKeySender, manager.String()),
		),
	})

	return nil
}

func (k Keeper) BurnFunds(ctx sdk.Context, invoker sdk.AccAddress, amount sdk.Coins) error {
	if !invoker.Empty() && !k.IsOperator(ctx, invoker) && !k.IsManager(ctx, invoker) {
		return types.ErrNotOperator
	}

	if !amount.IsZero() {
		err := k.BankKeeper.BurnCoins(ctx, types.ModuleName, amount)
		if err != nil {
			return err
		}
	}

	treasury := k.GetTreasury(ctx)
	treasury.Burned = treasury.Burned.Add(amount...)
	k.SetTreasury(ctx, treasury)

	return nil
}

func (k Keeper) IterateScheduledBurnQueue(ctx sdk.Context, endTime time.Time, cb func(burn types.Burn) (stop bool)) {
	iterator := k.ScheduledBurnQueueIterator(ctx, endTime)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		var burn types.Burn
		k.cdc.MustUnmarshal(iterator.Value(), &burn)

		if cb(burn) {
			break
		}
	}
}

func (keeper Keeper) InsertBurnQueue(ctx sdk.Context, burn types.Burn) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshal(&burn)

	scheduledFor, _ := time.Parse(types.TimeFormat, burn.ScheduledFor)

	store.Set(types.BurnQueueKey(scheduledFor), bz)
}

func (keeper Keeper) RemoveFromBurnQueue(ctx sdk.Context, endTime time.Time) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(types.BurnQueueKey(endTime))
}

func (keeper Keeper) HasBurnInQueue(ctx sdk.Context, scheduledFor time.Time) bool {
	store := ctx.KVStore(keeper.storeKey)
	return store.Has(types.BurnQueueKey(scheduledFor))
}

func (keeper Keeper) GetBurnFromQueue(ctx sdk.Context, scheduledFor time.Time) types.Burn {
	store := ctx.KVStore(keeper.storeKey)

	var burn types.Burn
	bz := store.Get(types.BurnQueueKey(scheduledFor))

	keeper.cdc.MustUnmarshal(bz, &burn)

	return burn
}

func (keeper Keeper) ScheduledBurnQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	return store.Iterator(types.BurnQueueKeyPrefix, sdk.PrefixEndBytes(types.BurnQueueKey(endTime)))
}

func (k Keeper) IterateBurnQueue(ctx sdk.Context, cb func(burn types.Burn) (stop bool)) {
	iterator := k.BurnQueueIterator(ctx)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		var burn types.Burn
		k.cdc.MustUnmarshal(iterator.Value(), &burn)

		if cb(burn) {
			break
		}
	}
}

func (keeper Keeper) BurnQueueIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	return sdk.KVStorePrefixIterator(store, types.BurnQueueKeyPrefix)
}

func (k Keeper) GetBurns(ctx sdk.Context) []types.Burn {
	var burns []types.Burn
	k.IterateBurnQueue(ctx, func(burn types.Burn) (stop bool) {
		burns = append(burns, burn)
		return false
	})

	return burns
}
