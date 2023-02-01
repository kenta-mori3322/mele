package fee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/melechain/mele/config"
	"github.com/melechain/mele/x/fee/keeper"
	"github.com/melechain/mele/x/fee/types"
)

type FeeDecorator struct {
	feeKeeper          keeper.Keeper
	bankKeeper         bankkeeper.Keeper
	accountKeeper      authkeeper.AccountKeeper
	feeCollectorModule string
}

func NewFeeDecorator(fk keeper.Keeper, bk bankkeeper.Keeper, ak authkeeper.AccountKeeper, feeCollectorModule string) FeeDecorator {
	if addr := ak.GetModuleAddress(feeCollectorModule); addr == nil {
		panic("the fee collector module account has not been set")
	}

	return FeeDecorator{
		feeKeeper:          fk,
		bankKeeper:         bk,
		accountKeeper:      ak,
		feeCollectorModule: feeCollectorModule,
	}
}

func (d FeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	msgs := feeTx.GetMsgs()
	feePayer := feeTx.FeePayer()

	txFees := sdk.NewCoins()
	systemFees := sdk.NewCoins()

	for _, msg := range msgs {
		msgFee := sdk.NewCoins()

		switch msg := msg.(type) {
		case *banktypes.MsgSend:
			msgFee = msgFee.Add(msg.Amount...)

		case *banktypes.MsgMultiSend:
			for _, input := range msg.Inputs {
				msgFee = msgFee.Add(input.Coins...)
			}
		}

		txFees = txFees.Add(msgFee...)

		if !d.feeKeeper.IsMessageFeeExcluded(ctx, msg) {
			systemFees = systemFees.Add(d.CalculateSystemFee(ctx, msgFee)...)
		}
	}

	totalFees := txFees.Add(systemFees...)

	if !d.bankKeeper.HasBalance(ctx, feePayer, totalFees[0]) {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; required: %s", totalFees)
	}

	err = d.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		feePayer,
		d.feeCollectorModule,
		systemFees, // we deduct only system fees in the ante handler!
	)

	if err != nil {
		return ctx, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.AttributeValueModule,
			sdk.NewAttribute(types.AttributeKeyTotalFee, totalFees.String()),
			sdk.NewAttribute(types.AttributeKeySystemFee, systemFees.String()),
		),
	)

	return next(ctx, tx, simulate)
}

func (d FeeDecorator) CalculateSystemFee(ctx sdk.Context, txFee sdk.Coins) sdk.Coins {
	feePercentage := d.feeKeeper.FeePercentage(ctx)
	systemFeeInt := txFee.AmountOf(config.DefaultDenomination).ToDec().Mul(feePercentage).TruncateInt()

	if txFee.AmountOf(config.DefaultStableDenomination).GT(sdk.ZeroInt()) {
		// MELG transfer
		// supply := d.bankKeeper.GetSupply(ctx)
		supply := d.bankKeeper.GetSupply(ctx, config.DefaultStableDenomination) // cosmos-sdk 0.44.0 updates

		// totalUmelg := supply.GetTotal().AmountOf(config.DefaultStableDenomination).ToDec()
		totalUmelg := supply.Amount.ToDec() // cosmos-sdk 0.44.0 updates

		totalPercentage := txFee.AmountOf(config.DefaultStableDenomination).ToDec().Quo(totalUmelg)

		totalUmelgFee := totalPercentage.Mul(d.feeKeeper.MelgFeePercentage(ctx)).Mul(txFee.AmountOf(config.DefaultStableDenomination).ToDec())

		melgPriceInMelc := d.feeKeeper.MelgPrice(ctx)
		totalUmelcFee := totalUmelgFee.Mul(melgPriceInMelc)

		systemFeeInt = systemFeeInt.Add(totalUmelcFee.TruncateInt())
	}

	systemFee := sdk.NewCoins(sdk.NewCoin(config.DefaultDenomination, systemFeeInt))
	minimumFee := d.feeKeeper.MinimumFee(ctx)
	maximumFee := d.feeKeeper.MaximumFee(ctx)

	if systemFee.IsAllLT(minimumFee) {
		return minimumFee
	}

	if systemFee.IsAllGT(maximumFee) {
		return maximumFee
	}

	return systemFee
}
