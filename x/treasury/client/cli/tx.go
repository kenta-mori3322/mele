package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	customutils "github.com/melechain/mele/custom"
	"github.com/melechain/mele/x/treasury/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	treasuryTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	treasuryTxCmd.AddCommand(
		GetCmdOperator(),
		GetCmdDisburse(),
		GetCmdApproveDisbursement(),
		GetCmdCancelDisbursement(),
		GetCmdBurn(),
		GetCmdApproveBurn(),
		GetCmdCancelBurn(),
	)

	return treasuryTxCmd
}

type newGenerateOrBroadcastFunc func(client.Context, *pflag.FlagSet, ...sdk.Msg) error

func newSplitAndApply(
	genOrBroadcastFn newGenerateOrBroadcastFunc, clientCtx client.Context,
	fs *pflag.FlagSet, msgs []sdk.Msg, chunkSize int,
) error {

	if chunkSize == 0 {
		return genOrBroadcastFn(clientCtx, fs, msgs...)
	}

	// split messages into slices of length chunkSize
	totalMessages := len(msgs)
	for i := 0; i < len(msgs); i += chunkSize {

		sliceEnd := i + chunkSize
		if sliceEnd > totalMessages {
			sliceEnd = totalMessages
		}

		msgChunk := msgs[i:sliceEnd]
		if err := genOrBroadcastFn(clientCtx, fs, msgChunk...); err != nil {
			return err
		}
	}

	return nil
}

func GetCmdDisburse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disburse [recipient] [amount] [reference]",
		Short: "Disburse treasury funds",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			recipient, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			// Parse the coins
			amount, err := customutils.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			reference := args[2]

			msgs := []sdk.Msg{types.NewMsgDisburse(clientCtx.GetFromAddress(), recipient, amount, reference)}

			for _, msg := range msgs {
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdApproveDisbursement() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-disbursement [recipient] [scheduled]",
		Short: "Approve scheduled distribution",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			recipient, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			scheduledFor := args[1]

			msgs := []sdk.Msg{types.NewMsgApproveDisbursement(clientCtx.GetFromAddress(), recipient, scheduledFor)}

			for _, msg := range msgs {
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdCancelDisbursement() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-disbursement [recipient] [scheduled]",
		Short: "Cancel scheduled distribution",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			recipient, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			scheduledFor := args[1]

			msgs := []sdk.Msg{types.NewMsgCancelDisbursement(clientCtx.GetFromAddress(), recipient, scheduledFor)}

			for _, msg := range msgs {
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [amount]",
		Short: "Burn treasury funds",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := customutils.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}

			msgs := []sdk.Msg{types.NewMsgBurn(clientCtx.GetFromAddress(), amount)}

			for _, msg := range msgs {
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdApproveBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-burn [scheduled]",
		Short: "Approve scheduled burn",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msgs := []sdk.Msg{types.NewMsgApproveBurn(clientCtx.GetFromAddress(), args[0])}

			for _, msg := range msgs {
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdCancelBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-burn [scheduled]",
		Short: "Cancel scheduled burn",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msgs := []sdk.Msg{types.NewMsgCancelBurn(clientCtx.GetFromAddress(), args[0])}

			for _, msg := range msgs {
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
