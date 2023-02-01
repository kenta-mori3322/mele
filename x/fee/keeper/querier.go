package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryParameters = "parameters"
	QueryFeeExcludedMessages = "excluded-messages"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
			case QueryParameters:
				return queryParams(ctx, k, legacyQuerierCdc)

			case QueryFeeExcludedMessages:
				return queryFeeExcludedMessages(ctx, k, legacyQuerierCdc)

			default:
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown distribution query endpoint")
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

func queryFeeExcludedMessages(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	feeExcludedMessages := k.GetFeeExcludedMessages(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, feeExcludedMessages)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}