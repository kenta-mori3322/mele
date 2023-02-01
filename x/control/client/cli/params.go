package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	controltypes "github.com/melechain/mele/x/control/types"
	"github.com/spf13/cobra"

	paramscutils "github.com/cosmos/cosmos-sdk/x/params/client/utils"
)

// NewSubmitParamChangeProposalTxCmd returns a CLI command handler for creating
// a parameter change proposal governance transaction.
func GetCmdSubmitParamChangeProposal() *cobra.Command {
	return &cobra.Command{
		Use:   "param-change [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a parameter change proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := paramscutils.ParseParamChangeProposalJSON(clientCtx.LegacyAmino, args[0])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := paramproposal.NewParameterChangeProposal(
				proposal.Title, proposal.Description, proposal.Changes.ToParamChanges(),
			)

			msg, err := controltypes.NewMsgSubmitExecution(content, from)
			if err != nil {
				return err
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}
