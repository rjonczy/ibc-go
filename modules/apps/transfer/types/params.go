package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// DefaultSendEnabled enabled
	DefaultSendEnabled = true
	// DefaultReceiveEnabled enabled
	DefaultReceiveEnabled = true
)

var (
	// KeySendEnabled is store's key for SendEnabled Params
	KeySendEnabled = []byte("SendEnabled")
	// KeyReceiveEnabled is store's key for ReceiveEnabled Params
	KeyReceiveEnabled = []byte("ReceiveEnabled")
	// KeyProxyFee is store's key for ProxyFee Param
	KeyProxyFee = []byte("ProxyFee")
	// DefaultProxyFee 0.001% 0.00001
	DefaultProxyFee = sdk.NewDecWithPrec(1, 5)
)

// ParamKeyTable type declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new parameter configuration for the ibc transfer module
func NewParams(enableSend, enableReceive bool, proxyFee sdk.Dec) Params {
	return Params{
		SendEnabled:    enableSend,
		ReceiveEnabled: enableReceive,
		ProxyFee:       proxyFee,
	}
}

// DefaultParams is the default parameter configuration for the ibc-transfer module
func DefaultParams() Params {
	return NewParams(DefaultSendEnabled, DefaultReceiveEnabled, DefaultProxyFee)
}

// Validate all ibc-transfer module parameters
func (p Params) Validate() error {
	if err := validateEnabled(p.SendEnabled); err != nil {
		return err
	}

	return validateEnabled(p.ReceiveEnabled)
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySendEnabled, p.SendEnabled, validateEnabled),
		paramtypes.NewParamSetPair(KeyReceiveEnabled, p.ReceiveEnabled, validateEnabled),
		paramtypes.NewParamSetPair(KeyProxyFee, p.ProxyFee, validateProxyFee),
	}
}

func validateProxyFee(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("proxy fee cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("proxy fee too large: %s", v)
	}
	return nil
}

func validateEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
