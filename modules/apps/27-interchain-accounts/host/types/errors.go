package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ErrHostSubModuleDisabled ICA Host sentinel errors
var (
	ErrHostSubModuleDisabled = sdkerrors.Register(SubModuleName, 2, "host submodule is disabled")
)
