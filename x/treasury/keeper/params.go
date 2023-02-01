package keeper

import	(
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/treasury/types"
	"time"
)

func (k Keeper) Managers(ctx sdk.Context) (res []string) {
	k.paramspace.Get(ctx, types.KeyManagers, &res)
	return
}

func (k Keeper) IsMintEnabled(ctx sdk.Context) (res bool) {
	k.paramspace.Get(ctx, types.KeyMintEnabled, &res)
	return
}

func (k Keeper) DisbursementDelayThreshold(ctx sdk.Context) (res sdk.Coins) {
	k.paramspace.Get(ctx, types.KeyDisbursementDelayThresholdAmount, &res)
	return
}

func (k Keeper) DisbursementDelayDuration(ctx sdk.Context) (res time.Duration) {
	var s string
	k.paramspace.Get(ctx, types.KeyDisbursementDelayDuration, &s)

	res, _ = time.ParseDuration(s)

	return
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}