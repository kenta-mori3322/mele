package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/melechain/mele/x/fee/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	feeQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	feeQueryCmd.AddCommand(
		GetCmdQueryParams(queryRoute),
		GetCmdQueryFeeExcludedMessages(queryRoute),
	)

	return feeQueryCmd
}

func GetCmdQueryParams(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the current fee module parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			route := fmt.Sprintf("custom/%s/parameters", queryRoute)
			res, _, err := clientCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var params types.Params
			if err := clientCtx.LegacyAmino.UnmarshalJSON(res, &params); err != nil {
				return err
			}

			return clientCtx.PrintObjectLegacy(params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdQueryFeeExcludedMessages(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fee-excluded-messages",
		Short: "Query fee excluded messages",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			route := fmt.Sprintf("custom/%s/excluded-messages", queryRoute)
			res, _, err := clientCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var out []string
			if err := clientCtx.LegacyAmino.UnmarshalJSON(res, &out); err != nil {
				return err
			}

			return clientCtx.PrintObjectLegacy(out)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
