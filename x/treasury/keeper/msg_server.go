package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/treasury/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the distribution MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) AddOperator(goCtx context.Context, msg *types.MsgAddOperator) (*types.MsgAddOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	operator, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return nil, err
	}

	err = k.HandleAddOperator(ctx, sender, operator)
	if err != nil {
		return nil, err
	}

	return &types.MsgAddOperatorResponse{}, nil
}

func (k msgServer) RemoveOperator(goCtx context.Context, msg *types.MsgRemoveOperator) (*types.MsgRemoveOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	operator, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return nil, err
	}

	err = k.HandleRemoveOperator(ctx, sender, operator)
	if err != nil {
		return nil, err
	}

	return &types.MsgRemoveOperatorResponse{}, nil
}

func (k msgServer) Disburse(goCtx context.Context, msg *types.MsgDisburse) (*types.MsgDisburseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	operator, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return nil, err
	}

	err = k.HandleDisburse(ctx, operator, recipient, msg.Amount, msg.Reference)
	if err != nil {
		return nil, err
	}

	return &types.MsgDisburseResponse{}, nil
}

func (k msgServer) ApproveDisbursement(goCtx context.Context, msg *types.MsgApproveDisbursement) (*types.MsgApproveDisbursementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	manager, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return nil, err
	}

	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	scheduledFor, err := time.Parse(types.TimeFormat, msg.ScheduledFor)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidTime, msg.ScheduledFor)
	}

	err = k.HandleApproveDisbursement(ctx, manager, recipient, scheduledFor)
	if err != nil {
		return nil, err
	}

	return &types.MsgApproveDisbursementResponse{}, nil
}

func (k msgServer) CancelDisbursement(goCtx context.Context, msg *types.MsgCancelDisbursement) (*types.MsgCancelDisbursementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	manager, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return nil, err
	}

	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	scheduledFor, err := time.Parse(types.TimeFormat, msg.ScheduledFor)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidTime, msg.ScheduledFor)
	}

	err = k.HandleCancelDisbursement(ctx, manager, recipient, scheduledFor)
	if err != nil {
		return nil, err
	}

	return &types.MsgCancelDisbursementResponse{}, nil
}

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	operator, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return nil, err
	}

	err = k.HandleBurn(ctx, operator, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgBurnResponse{}, nil
}

func (k msgServer) ApproveBurn(goCtx context.Context, msg *types.MsgApproveBurn) (*types.MsgApproveBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	manager, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return nil, err
	}

	scheduledFor, err := time.Parse(types.TimeFormat, msg.ScheduledFor)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidTime, msg.ScheduledFor)
	}

	err = k.HandleApproveBurn(ctx, manager, scheduledFor)
	if err != nil {
		return nil, err
	}

	return &types.MsgApproveBurnResponse{}, nil
}

func (k msgServer) CancelBurn(goCtx context.Context, msg *types.MsgCancelBurn) (*types.MsgCancelBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	manager, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		return nil, err
	}

	scheduledFor, err := time.Parse(types.TimeFormat, msg.ScheduledFor)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidTime, msg.ScheduledFor)
	}

	err = k.HandleCancelBurn(ctx, manager, scheduledFor)
	if err != nil {
		return nil, err
	}

	return &types.MsgCancelBurnResponse{}, nil
}

func MintTreasurySupplyProposal(ctx sdk.Context, k Keeper, p *types.MintTreasurySupplyProposal) error {
	err := k.MintCoins(ctx, p.Amount)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMintTreasurySupply,
			sdk.NewAttribute(types.AttributeKeyTitle, p.Title),
			sdk.NewAttribute(types.AttributeKeyDescription, p.Description),
			sdk.NewAttribute(types.AttributeKeyAmount, p.Amount.String()),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
		),
	)

	return nil
}

func BurnTreasurySupplyProposal(ctx sdk.Context, k Keeper, p *types.BurnTreasurySupplyProposal) error {
	err := k.BurnCoins(ctx, p.Amount)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBurnTreasurySupply,
			sdk.NewAttribute(types.AttributeKeyTitle, p.Title),
			sdk.NewAttribute(types.AttributeKeyDescription, p.Description),
			sdk.NewAttribute(types.AttributeKeyAmount, p.Amount.String()),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
		),
	)

	return nil
}
