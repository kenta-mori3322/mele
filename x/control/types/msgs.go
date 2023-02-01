package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgSubmitExecution = "submit_execution"
)

var _ sdk.Msg = &MsgSubmitExecution{}

func NewMsgSubmitExecution(content govtypes.Content, executor sdk.AccAddress) (*MsgSubmitExecution, error) {
	m := &MsgSubmitExecution{
		Executor: executor.String(),
	}
	err := m.SetContent(content)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Route implements Msg
func (msg MsgSubmitExecution) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgSubmitExecution) Type() string { return TypeMsgSubmitExecution }

// ValidateBasic implements Msg
func (msg MsgSubmitExecution) ValidateBasic() error {
	if msg.Content == nil {
		return sdkerrors.Wrap(ErrInvalidExecutionContent, "missing content")
	}
	if msg.Executor == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Executor)
	}
	content := msg.GetContent()

	if !govtypes.IsValidProposalType(content.ProposalType()) {
		return sdkerrors.Wrap(ErrInvalidProposalType, content.ProposalType())
	}

	return content.ValidateBasic()
}

// String implements the Stringer interface
func (msg MsgSubmitExecution) String() string {
	return fmt.Sprintf(`Submit Executor Message:
  Content:         %s
`, msg.Content.String())
}

// GetSignBytes implements Msg
func (msg MsgSubmitExecution) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements Msg
func (msg MsgSubmitExecution) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Executor)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgSubmitExecution) SetContent(content govtypes.Content) error {
	msg, ok := content.(proto.Message)
	if !ok {
		return fmt.Errorf("can't proto marshal %T", msg)
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return err
	}
	m.Content = any
	return nil
}

func (m *MsgSubmitExecution) GetContent() govtypes.Content {
	content, ok := m.Content.GetCachedValue().(govtypes.Content)
	if !ok {
		return nil
	}
	return content
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m MsgSubmitExecution) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var content govtypes.Content
	return unpacker.UnpackAny(m.Content, &content)
}
