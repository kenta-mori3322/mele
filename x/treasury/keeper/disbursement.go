package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/treasury/types"
)

func (k Keeper) HandleDisburse(ctx sdk.Context, operator sdk.AccAddress, recipient sdk.AccAddress, amount sdk.Coins, reference string) error {
	if !k.IsOperator(ctx, operator) {
		return types.ErrNotOperator
	}

	if k.IsDisbursementReferenceSet(ctx, reference) {
		return types.ErrDuplicateReference
	}

	scheduledFor := ctx.BlockTime()

	if amount.IsAnyGTE(k.DisbursementDelayThreshold(ctx)) {
		scheduledFor = scheduledFor.Add(k.DisbursementDelayDuration(ctx))
	}

	for k.HasDisbursementInQueue(ctx, recipient.String(), scheduledFor) {
		scheduledFor = scheduledFor.Add(time.Millisecond)
	}

	disbursement := types.Disbursement{
		Operator:     operator.String(),
		Recipient:    recipient.String(),
		Amount:       amount,
		ScheduledAt:  ctx.BlockTime().String(),
		ScheduledFor: scheduledFor.String(),
		Reference:    reference,
	}

	k.InsertDisbursementQueue(ctx, disbursement)
	k.SetDisbursementReferenceInfo(ctx, disbursement)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDisburse,
			sdk.NewAttribute(types.AttributeKeyScheduledFor, scheduledFor.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyReference, reference),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
			sdk.NewAttribute(types.AttributeKeySender, operator.String()),
		),
	})

	return nil
}

func (k Keeper) HandleApproveDisbursement(ctx sdk.Context, manager sdk.AccAddress, recipient sdk.AccAddress, scheduledFor time.Time) error {
	if !k.IsManager(ctx, manager) {
		return types.ErrNotManager
	}

	if !k.HasDisbursementInQueue(ctx, recipient.String(), scheduledFor) {
		return types.ErrDisbursementNotScheduled
	}

	disbursement := k.GetDisbursementFromQueue(ctx, recipient.String(), scheduledFor)

	err := k.DisburseFunds(ctx, manager, recipient, disbursement.Amount)
	if err != nil {
		k.Logger(ctx).Info(err.Error())
	}

	k.RemoveFromDisbursementQueue(ctx, recipient.String(), scheduledFor)

	disbursement.Executed = true
	k.SetDisbursementReferenceInfo(ctx, disbursement)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeApproveDisbursement,
			sdk.NewAttribute(types.AttributeKeyScheduledFor, scheduledFor.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, disbursement.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyReference, disbursement.Reference),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
			sdk.NewAttribute(types.AttributeKeySender, manager.String()),
		),
	})

	return nil
}

func (k Keeper) HandleCancelDisbursement(ctx sdk.Context, manager sdk.AccAddress, recipient sdk.AccAddress, scheduledFor time.Time) error {
	if !k.IsManager(ctx, manager) {
		return types.ErrNotManager
	}

	if !k.HasDisbursementInQueue(ctx, recipient.String(), scheduledFor) {
		return types.ErrDisbursementNotScheduled
	}

	k.RemoveFromDisbursementQueue(ctx, recipient.String(), scheduledFor)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCancelDisbursement,
			sdk.NewAttribute(types.AttributeKeyScheduledFor, scheduledFor.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
			sdk.NewAttribute(types.AttributeKeySender, manager.String()),
		),
	})

	return nil
}

func (k Keeper) DisburseFunds(ctx sdk.Context, invoker sdk.AccAddress, recipient sdk.AccAddress, amount sdk.Coins) error {
	if !invoker.Empty() && !k.IsOperator(ctx, invoker) && !k.IsManager(ctx, invoker) {
		return types.ErrNotOperator
	}

	if !amount.IsZero() {
		err := k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, amount)
		if err != nil {
			return err
		}
	}

	treasury := k.GetTreasury(ctx)
	treasury.Distributed = treasury.Distributed.Add(amount...)
	k.SetTreasury(ctx, treasury)

	return nil
}

func (k Keeper) IterateScheduledDisbursementQueue(ctx sdk.Context, endTime time.Time, cb func(disbursement types.Disbursement) (stop bool)) {
	iterator := k.ScheduledDisbursementQueueIterator(ctx, endTime)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		var disbursement types.Disbursement
		k.cdc.MustUnmarshal(iterator.Value(), &disbursement)

		if cb(disbursement) {
			break
		}
	}
}

func (keeper Keeper) InsertDisbursementQueue(ctx sdk.Context, disbursement types.Disbursement) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshal(&disbursement)
	scheduledFor, _ := time.Parse(types.TimeFormat, disbursement.ScheduledFor)
	store.Set(types.DisbursementQueueKey(disbursement.Recipient, scheduledFor), bz)
}

func (keeper Keeper) RemoveFromDisbursementQueue(ctx sdk.Context, address string, endTime time.Time) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(types.DisbursementQueueKey(address, endTime))
}

func (keeper Keeper) HasDisbursementInQueue(ctx sdk.Context, recipient string, scheduledFor time.Time) bool {
	store := ctx.KVStore(keeper.storeKey)
	return store.Has(types.DisbursementQueueKey(recipient, scheduledFor))
}

func (keeper Keeper) GetDisbursementFromQueue(ctx sdk.Context, recipient string, scheduledFor time.Time) types.Disbursement {
	store := ctx.KVStore(keeper.storeKey)

	var disbursement types.Disbursement
	bz := store.Get(types.DisbursementQueueKey(recipient, scheduledFor))

	keeper.cdc.MustUnmarshal(bz, &disbursement)

	return disbursement
}

func (keeper Keeper) ScheduledDisbursementQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	return store.Iterator(types.DisbursementQueueKeyPrefix, sdk.PrefixEndBytes(types.DisbursementByTimeKey(endTime)))
}

func (k Keeper) IterateDisbursementQueue(ctx sdk.Context, cb func(disbursement types.Disbursement) (stop bool)) {
	iterator := k.DisbursementQueueIterator(ctx)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		var disbursement types.Disbursement
		k.cdc.MustUnmarshal(iterator.Value(), &disbursement)

		if cb(disbursement) {
			break
		}
	}
}

func (keeper Keeper) DisbursementQueueIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	return sdk.KVStorePrefixIterator(store, types.DisbursementQueueKeyPrefix)
}

func (k Keeper) GetDisbursements(ctx sdk.Context) []types.Disbursement {
	var disbursements []types.Disbursement
	k.IterateDisbursementQueue(ctx, func(disbursement types.Disbursement) (stop bool) {
		disbursements = append(disbursements, disbursement)
		return false
	})

	return disbursements
}
