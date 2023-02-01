package types

const (
	EventTypeDisburse    		= "disburse"
	EventTypeBurn               = "burn"
	EventTypeBlockDisburse    	= "block_disburse"
	EventTypeBlockBurn          = "block_burn"
	EventTypeAddOperator 		= "add_operator"
	EventTypeRemoveOperator 	= "remove_operator"
	EventTypeApproveDisbursement = "approve_disbursement"
	EventTypeCancelDisbursement	= "cancel_disbursement"

	EventTypeMintTreasurySupply	= "MintTreasurySupply"
	EventTypeBurnTreasurySupply	= "BurnTreasurySupply"

	EventTypeApproveBurn            = "approve_burn"
	EventTypeCancelBurn	            = "cancel_burn"

	AttributeKeyInvoker				= "invoker"
	AttributeKeySender				= "sender"
	AttributeKeyRecipient  			= "recipient"
	AttributeKeyAmount				= "amount"
	AttributeKeyReference			= "reference"
	AttributeKeyScheduledFor 		= "scheduledFor"
	AttributeKeyOperator			= "operator"
	AttributeKeyTitle				= "title"
	AttributeKeyDescription			= "description"

	AttributeValueModule = ModuleName
)