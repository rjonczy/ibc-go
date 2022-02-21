package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	commitmenttypes "github.com/cosmos/ibc-go/v3/modules/core/23-commitment/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
)

var (
	chainID             = "gaiamainnet"
	connectionID        = "connection-0"
	clientID            = "clientidone"
	connectionID2       = "connectionidtwo"
	clientID2           = "clientidtwo"
	invalidConnectionID = "(invalidConnectionID)"
	clientHeight        = clienttypes.NewHeight(0, 6)
)

func TestConnectionValidateBasic(t *testing.T) {
	testCases := []struct {
		name       string
		connection types.ConnectionEnd
		expPass    bool
	}{
		{
			"valid connection",
			types.ConnectionEnd{ClientId: clientID, Versions: []*types.Version{ibctesting.ConnectionVersion}, State: types.INIT, Counterparty: types.Counterparty{ClientId: clientID2, ConnectionId: connectionID2, Prefix: commitmenttypes.NewMerklePrefix([]byte("prefix"))}, DelayPeriod: 500},
			true,
		},
		{
			"invalid client id",
			types.ConnectionEnd{ClientId: "(clientID1)", Versions: []*types.Version{ibctesting.ConnectionVersion}, State: types.INIT, Counterparty: types.Counterparty{ClientId: clientID2, ConnectionId: connectionID2, Prefix: commitmenttypes.NewMerklePrefix([]byte("prefix"))}, DelayPeriod: 500},
			false,
		},
		{
			"empty versions",
			types.ConnectionEnd{ClientId: clientID, State: types.INIT, Counterparty: types.Counterparty{ClientId: clientID2, ConnectionId: connectionID2, Prefix: commitmenttypes.NewMerklePrefix([]byte("prefix"))}, DelayPeriod: 500},
			false,
		},
		{
			"invalid version",
			types.ConnectionEnd{ClientId: clientID, Versions: []*types.Version{{}}, State: types.INIT, Counterparty: types.Counterparty{ClientId: clientID2, ConnectionId: connectionID2, Prefix: commitmenttypes.NewMerklePrefix([]byte("prefix"))}, DelayPeriod: 500},
			false,
		},
		{
			"invalid counterparty",
			types.ConnectionEnd{ClientId: clientID, Versions: []*types.Version{ibctesting.ConnectionVersion}, State: types.INIT, Counterparty: types.Counterparty{ClientId: clientID2, ConnectionId: connectionID2, Prefix: emptyPrefix}, DelayPeriod: 500},
			false,
		},
	}

	for i, tc := range testCases {
		tc := tc

		err := tc.connection.ValidateBasic()
		if tc.expPass {
			require.NoError(t, err, "valid test case %d failed: %s", i, tc.name)
		} else {
			require.Error(t, err, "invalid test case %d passed: %s", i, tc.name)
		}
	}
}

func TestCounterpartyValidateBasic(t *testing.T) {
	testCases := []struct {
		name         string
		counterparty types.Counterparty
		expPass      bool
	}{
		{"valid counterparty", types.Counterparty{ClientId: clientID, ConnectionId: connectionID2, Prefix: commitmenttypes.NewMerklePrefix([]byte("prefix"))}, true},
		{"invalid client id", types.Counterparty{ClientId: "(InvalidClient)", ConnectionId: connectionID2, Prefix: commitmenttypes.NewMerklePrefix([]byte("prefix"))}, false},
		{"invalid connection id", types.Counterparty{ClientId: clientID, ConnectionId: "(InvalidConnection)", Prefix: commitmenttypes.NewMerklePrefix([]byte("prefix"))}, false},
		{"invalid prefix", types.Counterparty{ClientId: clientID, ConnectionId: connectionID2, Prefix: emptyPrefix}, false},
	}

	for i, tc := range testCases {
		tc := tc

		err := tc.counterparty.ValidateBasic()
		if tc.expPass {
			require.NoError(t, err, "valid test case %d failed: %s", i, tc.name)
		} else {
			require.Error(t, err, "invalid test case %d passed: %s", i, tc.name)
		}
	}
}

func TestIdentifiedConnectionValidateBasic(t *testing.T) {
	testCases := []struct {
		name       string
		connection types.IdentifiedConnection
		expPass    bool
	}{
		{
			"valid connection",
			types.NewIdentifiedConnection(clientID, types.ConnectionEnd{ClientId: clientID, Versions: []*types.Version{ibctesting.ConnectionVersion}, State: types.INIT, Counterparty: types.Counterparty{ClientId: clientID2, ConnectionId: connectionID2, Prefix: commitmenttypes.NewMerklePrefix([]byte("prefix"))}, DelayPeriod: 500}),
			true,
		},
		{
			"invalid connection id",
			types.NewIdentifiedConnection("(connectionIDONE)", types.ConnectionEnd{ClientId: clientID, Versions: []*types.Version{ibctesting.ConnectionVersion}, State: types.INIT, Counterparty: types.Counterparty{ClientId: clientID2, ConnectionId: connectionID2, Prefix: commitmenttypes.NewMerklePrefix([]byte("prefix"))}, DelayPeriod: 500}),
			false,
		},
	}

	for i, tc := range testCases {
		tc := tc

		err := tc.connection.ValidateBasic()
		if tc.expPass {
			require.NoError(t, err, "valid test case %d failed: %s", i, tc.name)
		} else {
			require.Error(t, err, "invalid test case %d passed: %s", i, tc.name)
		}
	}
}
