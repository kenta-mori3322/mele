package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/melechain/mele/x/treasury/types"
)

const (
	QueryParameters = "parameters"
	QueryTreasury = "treasury"
	QueryOperators = "operators"
	QueryDisbursements = "disbursements"
	QueryAllDisbursements = "all_disbursements"
	QueryDisbursement = "disbursement"
	QueryBurns = "burns"
)

// NewQuerier creates a new querier for treasury clients.
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case QueryParameters:
			return queryParams(ctx, k, legacyQuerierCdc)
		case QueryTreasury:
			return queryTreasury(ctx, k, legacyQuerierCdc)
		case QueryOperators:
			return queryOperators(ctx, k, legacyQuerierCdc)
		case QueryDisbursement:
			return queryDisbursement(ctx, k, path[1:], legacyQuerierCdc)
		case QueryDisbursements:
			return queryDisbursements(ctx, k, legacyQuerierCdc)
		case QueryAllDisbursements:
			return queryAllDisbursements(ctx, k, legacyQuerierCdc)
		case QueryBurns:
			return queryBurns(ctx, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown treasury query endpoint")
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryTreasury(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetTreasury(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryOperators(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	operators := k.GetOperators(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, operators)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryDisbursement(ctx sdk.Context, k Keeper, path []string, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	found, disbursement := k.GetDisbursementReferenceInfo(ctx, path[0])

	if !found {
		return nil, types.ErrReferenceNotFound
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, disbursement)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryDisbursements(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	disbursements := k.GetDisbursements(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, disbursements)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryAllDisbursements(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	disbursements := k.GetAllDisbursements(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, disbursements)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryBurns(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	burns := k.GetBurns(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, burns)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
