package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/melechain/mele/x/control/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group gov queries under a subcommand
	govQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the control module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	govQueryCmd.AddCommand(
		GetCmdQueryExecution(queryRoute),
		GetCmdQueryExecutions(queryRoute),
		GetCmdQueryParams(queryRoute))

	return govQueryCmd
}

func GetCmdQueryExecution(queryRoute string) *cobra.Command {
	return &cobra.Command{
		Use:   "execution [execution-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query details of a single execution",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			// validate that the execution id is a uint
			executionID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("execution-id %s not a valid uint, please input a valid execution-id", args[0])
			}

			params := types.NewQueryExecutionParams(executionID)
			bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/execution", queryRoute), bz)
			if err != nil {
				return err
			}

			var execution types.Execution
			if err := clientCtx.LegacyAmino.UnmarshalJSON(res, &execution); err != nil {
				return err
			}

			return clientCtx.PrintObjectLegacy(execution) // nolint:errcheck
		},
	}
}

func GetCmdQueryExecutions(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "executions",
		Short: "Query executions",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/executions", queryRoute), nil)
			if err != nil {
				return err
			}

			var executions []types.Execution
			if err := clientCtx.LegacyAmino.UnmarshalJSON(res, &executions); err != nil {
				return err
			}

			if len(executions) == 0 {
				return fmt.Errorf("no executions found")
			}

			return clientCtx.PrintObjectLegacy(executions) // nolint:errcheck
		},
	}

	return cmd
}

func GetCmdQueryParams(queryRoute string) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Query the parameters of the control process",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/params", queryRoute), nil)
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
}
