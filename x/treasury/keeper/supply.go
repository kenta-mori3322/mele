package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/treasury/types"
)

func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		return nil
	}

	if ! k.IsMintEnabled(ctx) {
		return types.ErrMintDisabled
	}

	return k.BankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

func (k Keeper) BurnCoins(ctx sdk.Context, coins sdk.Coins) error {
	if coins.Empty() {
		return nil
	}

	if ! k.IsMintEnabled(ctx) {
		return types.ErrMintDisabled
	}

	return k.BankKeeper.BurnCoins(ctx, types.ModuleName, coins)
}
