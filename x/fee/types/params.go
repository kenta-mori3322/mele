package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/melechain/mele/config"
)

var (
	DefaultFeePercentage 		= sdk.NewDecWithPrec(25, 4)
	DefaultMinimumFee 			= sdk.NewCoins(sdk.NewInt64Coin(config.DefaultDenomination, 200)) // 200umelc
	DefaultMaximumFee           = sdk.NewCoins(sdk.NewInt64Coin(config.DefaultDenomination, 1000000000)) // 1 melc
	DefaultMelgPrice 			= sdk.NewDecWithPrec(2677, 3)
	DefaultMelgFeePercentage 	= sdk.NewDecWithPrec(1, 2)

	DefaultFeeExcludedMessages []string

	KeyFeePercentage 			= []byte("FeePercentage")
	KeyMinimumFee				= []byte("MinimumFee")
	KeyMaximumFee				= []byte("MaximumFee")
	KeyMelgPrice				= []byte("MelgPrice")
	KeyMelgFeePercentage 		= []byte("MelgFeePercentage")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(feePercentage sdk.Dec, minimumFee sdk.Coins, maximumFee sdk.Coins, melgPrice sdk.Dec, melgFeePercentage sdk.Dec) Params {
	return Params{
		FeePercentage: feePercentage.String(),
		MinimumFee: minimumFee,
		MaximumFee: maximumFee,
		MelgPrice: melgPrice.String(),
		MelgFeePercentage: melgFeePercentage.String(),
	}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	return fmt.Sprintf(`Params:
`, )
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyFeePercentage, &p.FeePercentage, validateFeePercentage),
		paramtypes.NewParamSetPair(KeyMinimumFee, &p.MinimumFee, validateFee),
		paramtypes.NewParamSetPair(KeyMaximumFee, &p.MaximumFee, validateFee),
		paramtypes.NewParamSetPair(KeyMelgPrice, &p.MelgPrice, validateMelgPrice),
		paramtypes.NewParamSetPair(KeyMelgFeePercentage, &p.MelgFeePercentage, validateFeePercentage),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(
		DefaultFeePercentage,
		DefaultMinimumFee,
		DefaultMaximumFee,
		DefaultMelgPrice,
		DefaultMelgFeePercentage,
	)
}

func (p Params) Validate() error {
	if err := validateFeePercentage(p.FeePercentage); err != nil {
		return err
	}

	if err := validateFee(p.MinimumFee); err != nil {
		return err
	}

	if err := validateFee(p.MaximumFee); err != nil {
		return err
	}

	if err := validateMelgPrice(p.MelgPrice); err != nil {
		return err
	}

	if err := validateFeePercentage(p.MelgFeePercentage); err != nil {
		return err
	}

	return nil
}

func validateFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if ! v.IsValid() {
		return fmt.Errorf("invalid name info fee: %s", v)
	}
	if ! v.AmountOf(config.DefaultDenomination).IsPositive() {
		return fmt.Errorf("invalid fee denomination. expected: %s", config.DefaultDenomination)
	}
	return nil
}

func validateFeePercentage(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	parsed, e := sdk.NewDecFromStr(v)
	if e != nil {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if parsed.GT(sdk.OneDec()) {
		return fmt.Errorf("fee percentage must be less than 100%: %s", v)
	}
	if parsed.IsNegative() {
		return fmt.Errorf("fee percentage must be positive: %s", v)
	}

	return nil
}

func validateMelgPrice(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	parsed, e := sdk.NewDecFromStr(v)
	if e != nil {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if parsed.IsNegative() {
		return fmt.Errorf("fee percentage must be positive: %s", v)
	}

	return nil
}