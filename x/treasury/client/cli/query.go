package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/melechain/mele/x/treasury/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group treasury queries under a subcommand
	treasuryQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	treasuryQueryCmd.AddCommand(
		GetCmdQueryParams(queryRoute),
		GetCmdQueryTreasury(queryRoute),
		GetCmdOperators(queryRoute),
		GetCmdDisbursement(queryRoute),
		GetCmdDisbursements(queryRoute),
		GetCmdAllDisbursements(queryRoute),
		GetCmdBurns(queryRoute),
	)

	return treasuryQueryCmd
}

func GetCmdQueryParams(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the current treasury module parameters",
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

func GetCmdQueryTreasury(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "treasury",
		Short: "Query the current treasury",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			route := fmt.Sprintf("custom/%s/treasury", queryRoute)
			res, _, err := clientCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var params types.Treasury
			if err := clientCtx.LegacyAmino.UnmarshalJSON(res, &params); err != nil {
				return err
			}

			return clientCtx.PrintObjectLegacy(params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdOperators(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "operators",
		Short: "Query Treasury Distribution Operators",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/operators", queryRoute), nil)
			if err != nil {
				fmt.Printf("Could not resolve operators\n")
				return nil
			}

			var out types.QueryResOperators
			clientCtx.LegacyAmino.MustUnmarshalJSON(res, &out)
			return clientCtx.PrintObjectLegacy(out)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdDisbursements(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disbursements",
		Short: "Query Treasury Scheduled Disbursements",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/disbursements", queryRoute), nil)
			if err != nil {
				fmt.Printf("Could not resolve disbursements {%s}\n", err)
				return nil
			}

			var out types.QueryResDisbursements
			clientCtx.LegacyAmino.MustUnmarshalJSON(res, &out)
			return clientCtx.PrintObjectLegacy(out)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdDisbursement(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disbursement [reference]",
		Short: "Query Treasury Disbursement by Reference",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/disbursement/%s", queryRoute, args[0]), nil)
			if err != nil {
				fmt.Printf("Could not resolve disbursements\n")
				return nil
			}

			var out types.Disbursement
			clientCtx.LegacyAmino.MustUnmarshalJSON(res, &out)
			return clientCtx.PrintObjectLegacy(out)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdAllDisbursements(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-disbursements",
		Short: "Query All Treasury Disbursements",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/all_disbursements", queryRoute), nil)
			if err != nil {
				fmt.Printf("Could not resolve disbursements\n")
				return nil
			}

			var out types.QueryResDisbursements
			clientCtx.LegacyAmino.MustUnmarshalJSON(res, &out)
			return clientCtx.PrintObjectLegacy(out)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdBurns(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burns",
		Short: "Query Treasury Scheduled Burns",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/burns", queryRoute), nil)
			if err != nil {
				fmt.Printf("Could not resolve burns\n")
				return nil
			}

			var out types.QueryResBurns
			clientCtx.LegacyAmino.MustUnmarshalJSON(res, &out)
			return clientCtx.PrintObjectLegacy(out)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
