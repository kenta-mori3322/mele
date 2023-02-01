package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/melechain/mele/x/control/types"
)

// Executor flags
const (
	FlagTitle        = "title"
	FlagDescription  = "description"
	flagProposalType = "type"
	FlagExecution    = "execution"
)

type execution struct {
	Title       string
	Description string
	Type        string
}

var ProposalFlags = []string{
	FlagTitle,
	FlagDescription,
	flagProposalType,
}

func GetTxCmd(storeKey string, pcmds []*cobra.Command) *cobra.Command {
	govTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Control transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmdSubmitProp := GetCmdSubmitExecution()
	for _, pcmd := range pcmds {
		flags.AddTxFlagsToCmd(pcmd)
		cmdSubmitProp.AddCommand(pcmd)
	}
	cmdSubmitProp.AddCommand(GetCmdSubmitParamChangeProposal())

	govTxCmd.AddCommand(
		cmdSubmitProp,
	)

	return govTxCmd
}

// GetCmdSubmitExecution implements submitting a execution transaction command.
func GetCmdSubmitExecution() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-execution",
		Short: "Submit an execution",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			execution, err := parseSubmitProposalFlags()
			if err != nil {
				return err
			}

			content := govtypes.ContentFromProposalType(execution.Title, execution.Description, execution.Type)

			ex, err := types.NewMsgSubmitExecution(content, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			msgs := []sdk.Msg{ex}

			for _, msg := range msgs {
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	cmd.Flags().String(FlagTitle, "", "title of execution")
	cmd.Flags().String(FlagDescription, "", "description of execution")
	cmd.Flags().String(flagProposalType, "", "proposalType of execution, types: text/parameter_change/software_upgrade")
	cmd.Flags().String(FlagExecution, "", "execution file path (if this path is given, other execution flags are ignored)")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
