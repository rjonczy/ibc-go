/*
	This file contains the variables, constants, and default values
	used in the testing package and commonly defined in tests.
*/
package ibctesting

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	connectiontypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	commitmenttypes "github.com/cosmos/ibc-go/v3/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"
	"github.com/cosmos/ibc-go/v3/testing/mock"
)

const (
	FirstClientID     = "07-tendermint-0"
	FirstChannelID    = "channel-0"
	FirstConnectionID = "connection-0"

	// TrustingPeriod Default params constants used to create a TM client
	TrustingPeriod            = time.Hour * 24 * 7 * 2
	UnbondingPeriod           = time.Hour * 24 * 7 * 3
	MaxClockDrift             = time.Second * 10
	DefaultDelayPeriod uint64 = 0

	DefaultChannelVersion = mock.Version
	InvalidID             = "IDisInvalid"

	// TransferPort Application Ports
	TransferPort = ibctransfertypes.ModuleName
	MockPort     = mock.ModuleName

	// Title used for testing proposals
	Title       = "title"
	Description = "description"

	LongString = "LoremipsumdolorsitameconsecteturadipiscingeliseddoeiusmodtemporincididuntutlaboreetdoloremagnaaliquUtenimadminimveniamquisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequDuisauteiruredolorinreprehenderitinvoluptateelitsseillumoloreufugiatnullaariaturEcepteurintoccaectupidatatonroidentuntnulpauifficiaeseruntmollitanimidestlaborum"
)

var (
	DefaultOpenInitVersion *connectiontypes.Version

	// DefaultTrustLevel Default params variables used to create a TM client
	DefaultTrustLevel = ibctmtypes.DefaultTrustLevel
	TestCoin          = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100))

	UpgradePath = []string{"upgrade", "upgradedIBCState"}

	ConnectionVersion = connectiontypes.ExportedVersionsToProto(connectiontypes.GetCompatibleVersions())[0]

	MockAcknowledgement          = mock.Acknowledgement.Acknowledgement()
	MockPacketData               = mock.PacketData
	MockFailPacketData           = mock.FailPacketData
	MockRecvCanaryCapabilityName = mock.RecvCanaryCapabilityName

	prefix = commitmenttypes.NewMerklePrefix([]byte("ibc"))
)
