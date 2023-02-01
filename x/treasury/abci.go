package treasury

import (
	"time"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/treasury/keeper"
	"github.com/melechain/mele/x/treasury/types"
)

// BeginBlocker sets the proposer for determining distribution during endblock
// and distribute rewards for the previous block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {

}

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.IterateScheduledDisbursementQueue(ctx, ctx.BlockTime(), func(disbursement types.Disbursement) (stop bool) {
		operator, _ := sdk.AccAddressFromBech32(disbursement.Operator)
		recipient, _ := sdk.AccAddressFromBech32(disbursement.Recipient)

		err := k.DisburseFunds(ctx, operator, recipient, disbursement.Amount)
		if err != nil {
			k.Logger(ctx).Info(err.Error())
		}

		scheduledFor, _ := time.Parse(types.TimeFormat, disbursement.ScheduledFor)
		k.RemoveFromDisbursementQueue(ctx, recipient.String(), scheduledFor)

		disbursement.Executed = true
		k.SetDisbursementReferenceInfo(ctx, disbursement)

		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.EventTypeBlockDisburse,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
				sdk.NewAttribute(types.AttributeKeyInvoker, disbursement.Operator),
				sdk.NewAttribute(types.AttributeKeyRecipient, disbursement.Recipient),
				sdk.NewAttribute(types.AttributeKeyAmount, disbursement.Amount.String()),
				sdk.NewAttribute(types.AttributeKeyReference, disbursement.Reference),
			),
		})

		return false
	})

	k.IterateScheduledBurnQueue(ctx, ctx.BlockTime(), func(burn types.Burn) (stop bool) {
		operator, _ := sdk.AccAddressFromBech32(burn.Operator)

		err := k.BurnFunds(ctx, operator, burn.Amount)
		if err != nil {
			k.Logger(ctx).Info(err.Error())
		}

		scheduledFor, _ := time.Parse(types.TimeFormat, burn.ScheduledFor)
		k.RemoveFromBurnQueue(ctx, scheduledFor)

		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.EventTypeBlockBurn,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
				sdk.NewAttribute(types.AttributeKeyInvoker, burn.Operator),
				sdk.NewAttribute(types.AttributeKeyAmount, burn.Amount.String()),
			),
		})

		return false
	})
}