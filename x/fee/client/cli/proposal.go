package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	customutils "github.com/melechain/mele/custom"
	feeutils "github.com/melechain/mele/x/fee/client/utils"
	"github.com/melechain/mele/x/fee/types"
	"github.com/spf13/cobra"
)

func GetCmdAddFeeExcludedMessageProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-fee-excluded-message [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Add Fee Excluded Message Proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := feeutils.ParseAddFeeExcludedMessageProposalWithDeposit(clientCtx.JSONCodec, args[0])
			if err != nil {
				return err
			}

			deposit, err := customutils.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			if err != nil {
				return err
			}
			content := types.NewAddFeeExcludedMessageProposal(proposal.Title, proposal.Description, proposal.MessageType)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}

func GetCmdRemoveFeeExcludedMessageProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-fee-excluded-message [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Remove Fee Excluded Message Proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := feeutils.ParseRemoveFeeExcludedMessageProposalWithDeposit(clientCtx.JSONCodec, args[0])
			if err != nil {
				return err
			}

			deposit, err := customutils.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			if err != nil {
				return err
			}
			content := types.NewRemoveFeeExcludedMessageProposal(proposal.Title, proposal.Description, proposal.MessageType)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}
