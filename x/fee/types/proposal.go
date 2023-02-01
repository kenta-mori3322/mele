package types

import (
	"fmt"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeAddFeeExcludedMessage    = "AddFeeExcludedMessage"
	ProposalTypeRemoveFeeExcludedMessage = "RemoveFeeExcludedMessage"
)

var _ govtypes.Content = AddFeeExcludedMessageProposal{}
var _ govtypes.Content = RemoveFeeExcludedMessageProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeAddFeeExcludedMessage)
	govtypes.RegisterProposalTypeCodec(AddFeeExcludedMessageProposal{}, "fee/AddFeeExcludedMessageProposal")
	govtypes.RegisterProposalType(ProposalTypeRemoveFeeExcludedMessage)
	govtypes.RegisterProposalTypeCodec(RemoveFeeExcludedMessageProposal{}, "fee/RemoveFeeExcludedMessageProposal")
}

func NewAddFeeExcludedMessageProposal(title string, description string, messageType string) AddFeeExcludedMessageProposal {
	return AddFeeExcludedMessageProposal{title, description, messageType}
}

// nolint
func (p AddFeeExcludedMessageProposal) GetTitle() string       { return p.Title }
func (p AddFeeExcludedMessageProposal) GetDescription() string { return p.Description }
func (p AddFeeExcludedMessageProposal) ProposalRoute() string  { return RouterKey }
func (p AddFeeExcludedMessageProposal) ProposalType() string {
	return ProposalTypeAddFeeExcludedMessage
}
func (p AddFeeExcludedMessageProposal) ValidateBasic() error {
	return govtypes.ValidateAbstract(p)
}

func (p AddFeeExcludedMessageProposal) String() string {
	return fmt.Sprintf(`Add Fee-Excluded Message:
  Title:       %s
  Description: %s
  Message Type: %s
`, p.Title, p.Description, p.MessageType)
}

func NewRemoveFeeExcludedMessageProposal(title string, description string, messageType string) RemoveFeeExcludedMessageProposal {
	return RemoveFeeExcludedMessageProposal{title, description, messageType}
}

// nolint
func (p RemoveFeeExcludedMessageProposal) GetTitle() string       { return p.Title }
func (p RemoveFeeExcludedMessageProposal) GetDescription() string { return p.Description }
func (p RemoveFeeExcludedMessageProposal) ProposalRoute() string  { return RouterKey }
func (p RemoveFeeExcludedMessageProposal) ProposalType() string {
	return ProposalTypeRemoveFeeExcludedMessage
}
func (p RemoveFeeExcludedMessageProposal) ValidateBasic() error {
	return govtypes.ValidateAbstract(p)
}

func (p RemoveFeeExcludedMessageProposal) String() string {
	return fmt.Sprintf(`Remove Fee-Excluded Message:
  Title:       %s
  Description: %s
  Message Type: %s
`, p.Title, p.Description, p.MessageType)
}
