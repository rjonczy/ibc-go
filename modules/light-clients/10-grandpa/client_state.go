package grandpa

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
)

var _ exported.ClientState = (*ClientState)(nil)

// The implementation of ClientState is needed to register the protobuf type, but none of the functions is actually
// used anywhere

func (m *ClientState) ClientType() string {
	panic("unimplemented")
}

func (m *ClientState) GetLatestHeight() exported.Height {
	panic("unimplemented")
}

func (m *ClientState) Validate() error {
	panic("unimplemented")
}

func (m *ClientState) Status(ctx types.Context, clientStore types.KVStore, cdc codec.BinaryCodec) exported.Status {
	panic("unimplemented")
}

func (m *ClientState) ExportMetadata(clientStore types.KVStore) []exported.GenesisMetadata {
	panic("unimplemented")
}

func (m *ClientState) ZeroCustomFields() exported.ClientState {
	panic("unimplemented")
}

func (m *ClientState) GetTimestampAtHeight(ctx types.Context, clientStore types.KVStore, cdc codec.BinaryCodec, height exported.Height) (uint64, error) {
	panic("unimplemented")
}

func (m *ClientState) Initialize(ctx types.Context, cdc codec.BinaryCodec, clientStore types.KVStore, consensusState exported.ConsensusState) error {
	panic("unimplemented")
}

func (m *ClientState) VerifyMembership(ctx types.Context, clientStore types.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, path exported.Path, value []byte) error {
	panic("unimplemented")
}

func (m *ClientState) VerifyNonMembership(ctx types.Context, clientStore types.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, proof []byte, path exported.Path) error {
	panic("unimplemented")
}

func (m *ClientState) VerifyClientMessage(ctx types.Context, cdc codec.BinaryCodec, clientStore types.KVStore, clientMsg exported.ClientMessage) error {
	panic("unimplemented")
}

func (m *ClientState) CheckForMisbehaviour(ctx types.Context, cdc codec.BinaryCodec, clientStore types.KVStore, clientMsg exported.ClientMessage) bool {
	panic("unimplemented")
}

func (m *ClientState) UpdateStateOnMisbehaviour(ctx types.Context, cdc codec.BinaryCodec, clientStore types.KVStore, clientMsg exported.ClientMessage) {
	panic("unimplemented")
}

func (m *ClientState) UpdateState(ctx types.Context, cdc codec.BinaryCodec, clientStore types.KVStore, clientMsg exported.ClientMessage) []exported.Height {
	panic("unimplemented")
}

func (m *ClientState) CheckSubstituteAndUpdateState(ctx types.Context, cdc codec.BinaryCodec, subjectClientStore, substituteClientStore types.KVStore, substituteClient exported.ClientState) error {
	panic("unimplemented")
}

func (m *ClientState) VerifyUpgradeAndUpdateState(ctx types.Context, cdc codec.BinaryCodec, store types.KVStore, newClient exported.ClientState, newConsState exported.ConsensusState, proofUpgradeClient, proofUpgradeConsState []byte) error {
	panic("unimplemented")
}
