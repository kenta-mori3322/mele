package types

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeMintTreasurySupply = "MintTreasurySupply"
	ProposalTypeBurnTreasurySupply = "BurnTreasurySupply"
)

var _ govtypes.Content = MintTreasurySupplyProposal{}
var _ govtypes.Content = BurnTreasurySupplyProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeMintTreasurySupply)
	govtypes.RegisterProposalTypeCodec(MintTreasurySupplyProposal{}, "treasury/MintTreasurySupplyProposal")

	govtypes.RegisterProposalType(ProposalTypeBurnTreasurySupply)
	govtypes.RegisterProposalTypeCodec(BurnTreasurySupplyProposal{}, "treasury/BurnTreasurySupplyProposal")
}

// NewMintTreasurySupplyProposal creates a new mint treasury supply proposal.
func NewMintTreasurySupplyProposal(title, description string, amount sdk.Coins) MintTreasurySupplyProposal {
	return MintTreasurySupplyProposal{title, description, amount}
}

func (mtsp MintTreasurySupplyProposal) GetTitle() string { return mtsp.Title }

func (mtsp MintTreasurySupplyProposal) GetDescription() string { return mtsp.Description }

func (mtsp MintTreasurySupplyProposal) ProposalRoute() string { return RouterKey }

func (mtsp MintTreasurySupplyProposal) ProposalType() string { return ProposalTypeMintTreasurySupply }

func (mtsp MintTreasurySupplyProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(mtsp)
	if err != nil {
		return err
	}
	if !mtsp.Amount.IsValid() {
		return errors.ErrInvalidCoins
	}

	return nil
}

func (mtsp MintTreasurySupplyProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Mint Treasury Supply Proposal:
  Title:       %s
  Description: %s
  Amount:      %s
`, mtsp.Title, mtsp.Description, mtsp.Amount))
	return b.String()
}

// NewBurnTreasurySupplyProposal creates a new burn treasury supply proposal.
func NewBurnTreasurySupplyProposal(title, description string, amount sdk.Coins) BurnTreasurySupplyProposal {
	return BurnTreasurySupplyProposal{title, description, amount}
}

func (btsp BurnTreasurySupplyProposal) GetTitle() string { return btsp.Title }

func (btsp BurnTreasurySupplyProposal) GetDescription() string { return btsp.Description }

func (btsp BurnTreasurySupplyProposal) ProposalRoute() string { return RouterKey }

func (btsp BurnTreasurySupplyProposal) ProposalType() string { return ProposalTypeBurnTreasurySupply }

func (btsp BurnTreasurySupplyProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(btsp)
	if err != nil {
		return err
	}
	if !btsp.Amount.IsValid() {
		return errors.ErrInvalidCoins
	}

	return nil
}

func (btsp BurnTreasurySupplyProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Burn Treasury Supply Proposal:
  Title:       %s
  Description: %s
  Amount:      %s
`, btsp.Title, btsp.Description, btsp.Amount))
	return b.String()
}
