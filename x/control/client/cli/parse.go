package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"

	govutils "github.com/cosmos/cosmos-sdk/x/gov/client/utils"
)

func parseSubmitProposalFlags() (*execution, error) {
	proposal := &execution{}
	proposalFile := viper.GetString(FlagExecution)

	if proposalFile == "" {
		proposal.Title = viper.GetString(FlagTitle)
		proposal.Description = viper.GetString(FlagDescription)
		proposal.Type = govutils.NormalizeProposalType(viper.GetString(flagProposalType))
		return proposal, nil
	}

	for _, flag := range ProposalFlags {
		if viper.GetString(flag) != "" {
			return nil, fmt.Errorf("--%s flag provided alongside --execution, which is a noop", flag)
		}
	}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, proposal)
	if err != nil {
		return nil, err
	}

	return proposal, nil
}
