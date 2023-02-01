package keeper

import	(
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/control/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

func (k Keeper) Managers(ctx sdk.Context) (res []string) {
	k.paramSpace.Get(ctx, types.KeyManagers, &res)
	return
}

func (k Keeper) IsManager(ctx sdk.Context, address sdk.AccAddress) bool {
	managers := k.Managers(ctx)

	for _, manager := range managers {
		if manager == address.String() {
			return true
		}
	}

	return false
}