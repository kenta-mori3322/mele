package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	controlclient "github.com/melechain/mele/x/control/client"
	"github.com/melechain/mele/x/fee/client/cli"
)

var (
	AddFeeExcludedMessageProposalHandler = govclient.NewProposalHandler(cli.GetCmdAddFeeExcludedMessageProposal, nil)
	RemoveFeeExcludedMessageProposalHandler = govclient.NewProposalHandler(cli.GetCmdRemoveFeeExcludedMessageProposal, nil)

	AddFeeExcludedMessageExecutionHandler = controlclient.NewExecutionHandler(cli.GetCmdAddFeeExcludedMessageProposal)
	RemoveFeeExcludedMessageExecutionHandler = controlclient.NewExecutionHandler(cli.GetCmdRemoveFeeExcludedMessageProposal)
)
