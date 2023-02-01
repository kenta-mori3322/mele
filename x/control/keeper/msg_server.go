package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/control/types"
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

func (k msgServer) SubmitExecution(goCtx context.Context, msg *types.MsgSubmitExecution) (*types.MsgSubmitExecutionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	executor, err := sdk.AccAddressFromBech32(msg.Executor)
	if err != nil {
		return nil, err
	}

	execution, err := k.HandleSubmitExecution(ctx, msg.GetContent(), executor)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Executor),
		),
	)

	submitEvent := sdk.NewEvent(types.EventTypeSubmitExecution, sdk.NewAttribute(types.AttributeKeyExecutionType, msg.GetContent().ProposalType()))
	ctx.EventManager().EmitEvent(submitEvent)

	return &types.MsgSubmitExecutionResponse{
		ExecutionID: execution.Id,
	}, nil
}
