package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// RegisterCodec registers concrete types on codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAddOperator{}, "treasury/AddOperator", nil)
	cdc.RegisterConcrete(&MsgRemoveOperator{}, "treasury/RemoveOperator", nil)

	cdc.RegisterConcrete(&MsgDisburse{}, "treasury/Disburse", nil)
	cdc.RegisterConcrete(&MsgApproveDisbursement{}, "treasury/ApproveDisbursement", nil)
	cdc.RegisterConcrete(&MsgCancelDisbursement{}, "treasury/CancelDisbursement", nil)

	cdc.RegisterConcrete(&MsgBurn{}, "treasury/Burn", nil)
	cdc.RegisterConcrete(&MsgApproveBurn{}, "treasury/ApproveBurn", nil)
	cdc.RegisterConcrete(&MsgCancelBurn{}, "treasury/CancelBurn", nil)

	cdc.RegisterConcrete(&MintTreasurySupplyProposal{}, "treasury/MintTreasurySupplyProposal", nil)
	cdc.RegisterConcrete(&BurnTreasurySupplyProposal{}, "treasury/BurnTreasurySupplyProposal", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgAddOperator{},
		&MsgRemoveOperator{},
		&MsgDisburse{},
		&MsgApproveDisbursement{},
		&MsgCancelDisbursement{},
		&MsgBurn{},
		&MsgApproveBurn{},
		&MsgCancelBurn{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&MintTreasurySupplyProposal{},
		&BurnTreasurySupplyProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
