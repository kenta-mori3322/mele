package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/melechain/mele/x/control/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	paramSpace params.Subspace

	storeKey sdk.StoreKey

	cdc codec.BinaryCodec

	router govtypes.Router
}

func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace params.Subspace, rtr govtypes.Router,
) Keeper {
	return Keeper{
		storeKey:   key,
		paramSpace: paramSpace.WithKeyTable(types.ParamKeyTable()),
		cdc:        cdc,
		router:     rtr,
	}
}

func (keeper Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (keeper Keeper) Router() govtypes.Router {
	return keeper.router
}
