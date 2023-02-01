package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/melechain/mele/x/treasury/types"


	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetCmdOperator() *cobra.Command {
	operatorTxCmd := &cobra.Command{
		Use:                        "operator",
		Short:                      fmt.Sprintf("%s transactions subcommands", "operator"),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	operatorTxCmd.AddCommand(
		GetCmdAddOperator(),
		GetCmdRemoveOperator(),
	)

	return operatorTxCmd
}

func GetCmdAddOperator() *cobra.Command {
	return &cobra.Command{
		Use:   "add [address]",
		Short: "Add a Treasury Distribution Operator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			operator, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msgs := []sdk.Msg{types.NewMsgAddOperator(clientCtx.GetFromAddress(), operator)}

			for _, msg := range msgs {
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}
}

func GetCmdRemoveOperator() *cobra.Command {
	return &cobra.Command{
		Use:   "remove [address]",
		Short: "Remove a Treasury Distribution Operator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			operator, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msgs := []sdk.Msg{types.NewMsgRemoveOperator(clientCtx.GetFromAddress(), operator)}

			for _, msg := range msgs {
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}
}
