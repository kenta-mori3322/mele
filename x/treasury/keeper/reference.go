package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/treasury/types"
)

func (k Keeper) SetDisbursementReferenceInfo(ctx sdk.Context, disbursement types.Disbursement) {
	store := ctx.KVStore(k.storeKey)

	reference := strings.ToLower(disbursement.Reference)

	store.Set(types.GetDisbursementReferenceKey(reference), k.cdc.MustMarshal(&disbursement))
}

func (k Keeper) RemoveDisbursementReferenceInfo(ctx sdk.Context, reference string) {
	store := ctx.KVStore(k.storeKey)

	store.Delete(types.GetDisbursementReferenceKey(reference))
}

func (k Keeper) GetDisbursementReferenceInfo(ctx sdk.Context, reference string) (bool, types.Disbursement) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetDisbursementReferenceKey(reference))
	if bz == nil {
		return false, types.Disbursement{}
	}

	var disbursement types.Disbursement
	k.cdc.MustUnmarshal(bz, &disbursement)

	return true, disbursement
}

func (k Keeper) IsDisbursementReferenceSet(ctx sdk.Context, reference string) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetDisbursementReferenceKey(reference))
}

func (k Keeper) GetDisbursementReferenceIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.GetDisbursementReferenceIteratorKey())
}

func (k Keeper) IterateDisbursementReferences(ctx sdk.Context, cb func(disbursement types.Disbursement) (stop bool)) {
	iterator := k.GetDisbursementReferenceIterator(ctx)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var disbursement types.Disbursement
		k.cdc.MustUnmarshal(iterator.Value(), &disbursement)

		if cb(disbursement) {
			break
		}
	}
}

func (k Keeper) GetAllDisbursements(ctx sdk.Context) []types.Disbursement {
	var disbursements []types.Disbursement
	k.IterateDisbursementReferences(ctx, func(disbursement types.Disbursement) (stop bool) {
		disbursements = append(disbursements, disbursement)
		return false
	})

	return disbursements
}
