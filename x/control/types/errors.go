package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrUnknownExecution         = sdkerrors.Register(ModuleName, 1, "unknown execution")
	ErrInvalidExecutionContent  = sdkerrors.Register(ModuleName, 2, "invalid execution content")
	ErrInvalidProposalType      = sdkerrors.Register(ModuleName, 3, "invalid execution type")
	ErrInvalidGenesis           = sdkerrors.Register(ModuleName, 4, "invalid genesis state")
	ErrNoExecutionHandlerExists = sdkerrors.Register(ModuleName, 5, "no handler exists for execution type")
	ErrControlNotEnabled 	    = sdkerrors.Register(ModuleName, 6, "control module disabled")
	ErrNotManager		        = sdkerrors.Register(ModuleName, 7, "not a manager")
)
