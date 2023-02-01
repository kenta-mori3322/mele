package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	controlclient "github.com/melechain/mele/x/control/client"
	"github.com/melechain/mele/x/treasury/client/cli"
)

var (
	MintTreasurySupplyProposalHandler = govclient.NewProposalHandler(cli.GetCmdMintTreasurySupplyProposal, nil)
	BurnTreasurySupplyProposalHandler = govclient.NewProposalHandler(cli.GetCmdBurnTreasurySupplyProposal, nil)

	MintTreasurySupplyExecutionHandler = controlclient.NewExecutionHandler(cli.GetCmdMintTreasurySupplyProposal)
	BurnTreasurySupplyExecutionHandler = controlclient.NewExecutionHandler(cli.GetCmdBurnTreasurySupplyProposal)
)
