package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/fee/types"
)

// FeePercentage
func (k Keeper) FeePercentage(ctx sdk.Context) (res sdk.Dec) {
	var s string
	k.paramspace.Get(ctx, types.KeyFeePercentage, &s)

	res, _ = sdk.NewDecFromStr(s)

	return
}

// MinimumFee
func (k Keeper) MinimumFee(ctx sdk.Context) (res sdk.Coins) {
	k.paramspace.Get(ctx, types.KeyMinimumFee, &res)
	return
}

// Maximum Fee
func (k Keeper) MaximumFee(ctx sdk.Context) (res sdk.Coins) {
	k.paramspace.Get(ctx, types.KeyMaximumFee, &res)
	return
}

// MelgPrice
func (k Keeper) MelgPrice(ctx sdk.Context) (res sdk.Dec) {
	var s string
	k.paramspace.Get(ctx, types.KeyMelgPrice, &s)

	res, _ = sdk.NewDecFromStr(s)

	return
}

// MelgFeePercentage
func (k Keeper) MelgFeePercentage(ctx sdk.Context) (res sdk.Dec) {
	var s string
	k.paramspace.Get(ctx, types.KeyMelgFeePercentage, &s)

	res, _ = sdk.NewDecFromStr(s)

	return
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}
