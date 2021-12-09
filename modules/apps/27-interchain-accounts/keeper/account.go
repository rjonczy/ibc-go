package keeper

import (
	// authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	// channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	// host "github.com/cosmos/ibc-go/modules/core/24-host"
)

// Determine account's address that will be created.
func GenerateAddress(identifier string) []byte {
	return tmhash.SumTruncated(append([]byte(identifier)))
}
