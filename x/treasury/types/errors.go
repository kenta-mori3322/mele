package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNotManager = sdkerrors.Register(ModuleName, 101, "No Manager permissions to execute the action.")
	ErrNotOperator = sdkerrors.Register(ModuleName, 102, "No Operator permissions to execute the action.")
	ErrInvalidTime = sdkerrors.Register(ModuleName, 103, "Invalid time")
	ErrDuplicateReference = sdkerrors.Register(ModuleName, 104, "Reference already used")
	ErrDisbursementNotScheduled = sdkerrors.Register(ModuleName, 105, "Disbursement not scheduled")
	ErrBurnNotScheduled = sdkerrors.Register(ModuleName, 106, "Burn not scheduled")
	ErrReferenceNotFound = sdkerrors.Register(ModuleName, 107, "Reference not found")
	ErrMintDisabled = sdkerrors.Register(ModuleName, 108, "Minting/burning treasury funds is disabled.")
)