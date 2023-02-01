package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/gogo/protobuf/proto"
)

const DefaultStartingExecutionID uint64 = 1

func NewExecution(content govtypes.Content, id uint64, submitTime string, executor string) (Execution, error) {
	execution := Execution{
		Id:         id,
		Executor:   executor,
		SubmitTime: submitTime,
	}

	msg, ok := content.(proto.Message)
	if !ok {
		return Execution{}, fmt.Errorf("%T does not implement proto.Message", content)
	}

	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return Execution{}, err
	}

	execution.Content = any

	return execution, nil
}

// GetContent returns the proposal Content
func (p Execution) GetContent() govtypes.Content {
	content, ok := p.Content.GetCachedValue().(govtypes.Content)
	if !ok {
		return nil
	}
	return content
}

func (p Execution) ProposalType() string {
	content := p.GetContent()
	if content == nil {
		return ""
	}
	return content.ProposalType()
}

func (p Execution) ProposalRoute() string {
	content := p.GetContent()
	if content == nil {
		return ""
	}
	return content.ProposalRoute()
}

func (p Execution) GetTitle() string {
	content := p.GetContent()
	if content == nil {
		return ""
	}
	return content.GetTitle()
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Execution) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var content govtypes.Content
	return unpacker.UnpackAny(p.Content, &content)
}

// Proposals is an array of proposal
type Executions []Execution

var _ types.UnpackInterfacesMessage = Executions{}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Executions) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, x := range p {
		err := x.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}
