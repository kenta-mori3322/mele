package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

type QueryResOperators []sdk.AccAddress

func (n QueryResOperators) String() string {
	var operators []string

	for _, operator := range n {
		operators = append(operators, operator.String())
	}

	return strings.Join(operators, "\n")
}

type QueryResDisbursements []Disbursement

func (n QueryResDisbursements) String() string {
	var disbursements []string

	for _, disbursement := range n {
		disbursements = append(disbursements, disbursement.String())
	}

	return strings.Join(disbursements, "\n")
}

type QueryResBurns []Burn

func (n QueryResBurns) String() string {
	var burns []string

	for _, burn := range n {
		burns = append(burns, burn.String())
	}

	return strings.Join(burns, "\n")
}

type QueryResPrice sdk.Coins

func (n QueryResPrice) String() string {
	return n.String()
}