package types

import (
	"fmt"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyEnabled                         = []byte("Enabled")
	KeyManagers                        = []byte("Managers")

	DefaultManagerAddress = "mele12rmu8657nunpgnd5ufwqphnwlspzcwl29ejqpa"
)

func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(enabled bool, managers []string) Params {
	return Params{
		Enabled: enabled,
		Managers: managers,
	}
}

func (p Params) String() string {
	return fmt.Sprintf(`
	Enabled: %t
	Managers: %s
	`, p.Enabled, p.Managers)
}

func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(KeyEnabled, &p.Enabled, ValidateEnabled),
		params.NewParamSetPair(KeyManagers, &p.Managers, ValidateManager),
	}
}

func DefaultParams() Params {
	return NewParams(true, []string{DefaultManagerAddress})
}

func ValidateEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func ValidateManager(i interface{}) error {
	addresses, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if len(addresses) <= 0 || len(addresses) > 2 {
		return fmt.Errorf("only 1 or 2 Managers can be specified: %d", len(addresses))
	}

	return nil
}

func (p Params) Validate() error {
	if err := ValidateEnabled(p.Enabled); err != nil {
		return err
	}

	if err := ValidateManager(p.Managers); err != nil {
		return err
	}

	return nil
}
