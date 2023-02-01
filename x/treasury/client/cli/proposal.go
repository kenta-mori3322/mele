package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	customutils "github.com/melechain/mele/custom"
	treasuryutils "github.com/melechain/mele/x/treasury/client/utils"
	"github.com/melechain/mele/x/treasury/types"
	"github.com/spf13/cobra"
)

func GetCmdMintTreasurySupplyProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-treasury-supply [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Mint treasury supply proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := treasuryutils.ParseMintTreasurySupplyProposalWithDeposit(clientCtx.JSONCodec, args[0])
			if err != nil {
				return err
			}

			amount, err := customutils.ParseCoinsNormalized(proposal.Amount)
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
			content := types.NewMintTreasurySupplyProposal(proposal.Title, proposal.Description, amount)

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

func GetCmdBurnTreasurySupplyProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-treasury-supply [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Burn treasury supply proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := treasuryutils.ParseMintTreasurySupplyProposalWithDeposit(clientCtx.JSONCodec, args[0])
			if err != nil {
				return err
			}

			amount, err := customutils.ParseCoinsNormalized(proposal.Amount)
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
			content := types.NewMintTreasurySupplyProposal(proposal.Title, proposal.Description, amount)

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
