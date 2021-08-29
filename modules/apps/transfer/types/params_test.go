package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestValidateParams(t *testing.T) {
	require.NoError(t, DefaultParams().Validate())
	require.NoError(t, NewParams(true, false, sdk.NewDecWithPrec(1, 5)).Validate())
}
