package keeper

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/treasury/types"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.BinaryCodec
	paramspace    paramtypes.Subspace
	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper
}

func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, paramspace paramtypes.Subspace, accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper) Keeper {
	// set KeyTable if it has not already been set
	if !paramspace.HasKeyTable() {
		paramspace = paramspace.WithKeyTable(types.ParamKeyTable())
	}

	keeper := Keeper{
		storeKey:      key,
		cdc:           cdc,
		paramspace:    paramspace,
		AccountKeeper: accountKeeper,
		BankKeeper:    bankKeeper,
	}
	return keeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetTreasury(ctx sdk.Context) (minter types.Treasury) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.TreasuryKey)
	if b == nil {
		panic("stored treasury should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &minter)
	return
}

func (k Keeper) SetTreasury(ctx sdk.Context, minter types.Treasury) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&minter)
	store.Set(types.TreasuryKey, b)
}
