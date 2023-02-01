package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultParamspace = ModuleName

	DefaultDisbursementDelayDuration        = time.Second * 30
	DefaultDisbursementDelayThresholdAmount = 1000 // melc
)

var (
	KeyManagers                         = []byte("Managers")
	KeyDisbursementDelayThresholdAmount = []byte("DisbursementDelayThresholdAmount")
	KeyDisbursementDelayDuration        = []byte("DisbursementDelayDuration")
	KeyMintEnabled                      = []byte("MintEnabled")

	DefaultManagerAddress = "mele12rmu8657nunpgnd5ufwqphnwlspzcwl29ejqpa"
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p Params) String() string {
	return fmt.Sprintf(`
	Managers: %s
	Disbursement Delay Threshold Amount: %s
	Disbursement Delay Duration: %s
	Mint enabled: %t
	`, p.Managers, p.DisbursementDelayThresholdAmount, p.DisbursementDelayDuration, p.MintEnabled)
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyManagers, &p.Managers, validateManager),
		paramtypes.NewParamSetPair(KeyDisbursementDelayThresholdAmount, &p.DisbursementDelayThresholdAmount, validateCoins),
		paramtypes.NewParamSetPair(KeyDisbursementDelayDuration, &p.DisbursementDelayDuration, validateDuration),
		paramtypes.NewParamSetPair(KeyMintEnabled, &p.MintEnabled, validateMintEnabled),
	}
}

func DefaultParams() Params {
	amount, _ := sdk.ConvertCoin(sdk.NewInt64Coin("melc", DefaultDisbursementDelayThresholdAmount), "umelc")
	duration := DefaultDisbursementDelayDuration.String()

	return Params{
		[]string{DefaultManagerAddress},
		sdk.NewCoins(amount),
		duration,
		true,
	}
}

func (p Params) Validate() error {
	if err := validateManager(p.Managers); err != nil {
		return err
	}

	if err := validateCoins(p.DisbursementDelayThresholdAmount); err != nil {
		return err
	}

	if err := validateDuration(p.DisbursementDelayDuration); err != nil {
		return err
	}

	if err := validateMintEnabled(p.MintEnabled); err != nil {
		return err
	}

	return nil
}

func validateManager(i interface{}) error {
	addresses, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if len(addresses) <= 0 || len(addresses) > 2 {
		return fmt.Errorf("only 1 or 2 Managers can be specified: %d", len(addresses))
	}

	return nil
}

func validateDuration(i interface{}) error {
	v, ok := i.(string)

	duration, err := time.ParseDuration(v)

	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if duration <= 0 {
		return fmt.Errorf("duration must be positive: %s", v)
	}

	return nil
}

func validateMintEnabled(i interface{}) error {
	_, ok := i.(bool)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateCoins(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return fmt.Errorf("invalid amount: %s", v)
	}

	return nil
}
