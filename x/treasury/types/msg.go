package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/melechain/mele/config"
	"github.com/tendermint/tendermint/types"
)

var _, _, _, _, _, _, _, _ sdk.Msg = &MsgAddOperator{},
	&MsgRemoveOperator{},
	&MsgDisburse{},
	&MsgApproveDisbursement{},
	&MsgCancelDisbursement{},
	&MsgBurn{},
	&MsgApproveBurn{},
	&MsgCancelBurn{}

func NewMsgAddOperator(sender sdk.AccAddress, operator sdk.AccAddress) *MsgAddOperator {
	return &MsgAddOperator{
		Sender:   sender.String(),
		Operator: operator.String(),
	}
}

func (msg MsgAddOperator) Route() string { return RouterKey }

func (msg MsgAddOperator) Type() string { return "add_operator" }

func (msg MsgAddOperator) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Operator == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Operator)
	}
	return nil
}

func (msg MsgAddOperator) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgAddOperator) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgRemoveOperator
func NewMsgRemoveOperator(sender sdk.AccAddress, operator sdk.AccAddress) *MsgRemoveOperator {
	return &MsgRemoveOperator{
		Sender:   sender.String(),
		Operator: operator.String(),
	}
}

func (msg MsgRemoveOperator) Route() string { return RouterKey }

func (msg MsgRemoveOperator) Type() string { return "remove_operator" }

func (msg MsgRemoveOperator) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Operator == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Operator)
	}
	return nil
}

func (msg MsgRemoveOperator) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgRemoveOperator) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgDisburse
func NewMsgDisburse(sender sdk.AccAddress, recipient sdk.AccAddress, amount sdk.Coins, reference string) *MsgDisburse {
	return &MsgDisburse{
		Operator:  sender.String(),
		Recipient: recipient.String(),
		Amount:    amount,
		Reference: reference,
	}
}

func (msg MsgDisburse) Route() string { return RouterKey }

func (msg MsgDisburse) Type() string { return "disburse" }

func (msg MsgDisburse) ValidateBasic() error {
	if msg.Operator == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Operator)
	}
	if msg.Recipient == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Operator)
	}
	if msg.Amount.AmountOf(config.DefaultDenomination).IsNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid amount.")
	}
	if msg.Amount.AmountOf(config.DefaultStableDenomination).IsNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid amount.")
	}
	if len(msg.Reference) > 255 {
		return sdkerrors.Wrapf(sdkerrors.ErrMemoTooLarge, "Reference too long")
	}
	return nil
}

func (msg MsgDisburse) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDisburse) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgCancelDisbursement
func NewMsgCancelDisbursement(manager sdk.AccAddress, recipient sdk.AccAddress, scheduledFor string) *MsgCancelDisbursement {
	return &MsgCancelDisbursement{
		Manager:      manager.String(),
		Recipient:    recipient.String(),
		ScheduledFor: scheduledFor,
	}
}

func (msg MsgCancelDisbursement) Route() string { return RouterKey }

func (msg MsgCancelDisbursement) Type() string { return "cancel_disbursement" }

func (msg MsgCancelDisbursement) ValidateBasic() error {
	if msg.Manager == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Manager)
	}
	if msg.Recipient == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Recipient)
	}
	scheduledFor, err := time.Parse(types.TimeFormat, msg.ScheduledFor)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidTime, msg.ScheduledFor)
	}
	if scheduledFor.IsZero() {
		return sdkerrors.Wrap(ErrInvalidTime, msg.ScheduledFor)
	}
	return nil
}

func (msg MsgCancelDisbursement) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCancelDisbursement) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgApproveDisbursement
func NewMsgApproveDisbursement(manager sdk.AccAddress, recipient sdk.AccAddress, scheduledFor string) *MsgApproveDisbursement {
	return &MsgApproveDisbursement{
		Manager:      manager.String(),
		Recipient:    recipient.String(),
		ScheduledFor: scheduledFor,
	}
}

func (msg MsgApproveDisbursement) Route() string { return RouterKey }

func (msg MsgApproveDisbursement) Type() string { return "approve_disbursement" }

func (msg MsgApproveDisbursement) ValidateBasic() error {
	if msg.Manager == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Manager)
	}
	if msg.Recipient == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Recipient)
	}
	scheduledFor, err := time.Parse(types.TimeFormat, msg.ScheduledFor)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidTime, msg.ScheduledFor)
	}
	if scheduledFor.IsZero() {
		return sdkerrors.Wrap(ErrInvalidTime, msg.ScheduledFor)
	}
	return nil
}

func (msg MsgApproveDisbursement) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgApproveDisbursement) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgBurn
func NewMsgBurn(sender sdk.AccAddress, amount sdk.Coins) *MsgBurn {
	return &MsgBurn{
		Operator: sender.String(),
		Amount:   amount,
	}
}

func (msg MsgBurn) Route() string { return RouterKey }

func (msg MsgBurn) Type() string { return "burn" }

func (msg MsgBurn) ValidateBasic() error {
	if msg.Operator == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Operator)
	}
	if !msg.Amount.AmountOf(config.DefaultDenomination).IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid amount.")
	}
	return nil
}

func (msg MsgBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgBurn) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgCancelBurn
func NewMsgCancelBurn(manager sdk.AccAddress, scheduledFor string) *MsgCancelBurn {
	return &MsgCancelBurn{
		Manager:      manager.String(),
		ScheduledFor: scheduledFor,
	}
}

func (msg MsgCancelBurn) Route() string { return RouterKey }

func (msg MsgCancelBurn) Type() string { return "cancel_burn" }

func (msg MsgCancelBurn) ValidateBasic() error {
	if msg.Manager == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Manager)
	}
	scheduledFor, err := time.Parse(types.TimeFormat, msg.ScheduledFor)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidTime, msg.ScheduledFor)
	}
	if scheduledFor.IsZero() {
		return sdkerrors.Wrap(ErrInvalidTime, msg.ScheduledFor)
	}
	return nil
}

func (msg MsgCancelBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCancelBurn) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MsgApproveBurn
func NewMsgApproveBurn(manager sdk.AccAddress, scheduledFor string) *MsgApproveBurn {
	return &MsgApproveBurn{
		Manager:      manager.String(),
		ScheduledFor: scheduledFor,
	}
}

func (msg MsgApproveBurn) Route() string { return RouterKey }

func (msg MsgApproveBurn) Type() string { return "approve_burn" }

func (msg MsgApproveBurn) ValidateBasic() error {
	if msg.Manager == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Manager)
	}
	scheduledFor, err := time.Parse(types.TimeFormat, msg.ScheduledFor)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidTime, msg.ScheduledFor)
	}
	if scheduledFor.IsZero() {
		return sdkerrors.Wrap(ErrInvalidTime, msg.ScheduledFor)
	}
	return nil
}

func (msg MsgApproveBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgApproveBurn) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Manager)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
