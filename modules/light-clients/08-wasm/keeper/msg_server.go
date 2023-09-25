package keeper

import (
	"context"
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v7/modules/light-clients/08-wasm/types"
)

var _ types.MsgServer = Keeper{}

// PushNewWasmCode defines a rpc handler method for MsgPushNewWasmCode
func (k Keeper) PushNewWasmCode(goCtx context.Context, msg *types.MsgPushNewWasmCode) (*types.MsgPushNewWasmCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Signer {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority: expected %s, got %s", k.authority, msg.Signer)
	}

	codeID, err := k.storeWasmCode(ctx, msg.Code)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "pushing new wasm code failed")
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			clienttypes.EventTypePushWasmCode,
			sdk.NewAttribute(clienttypes.AttributeKeyWasmCodeID, hex.EncodeToString(codeID)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, clienttypes.AttributeValueCategory),
		),
	})

	return &types.MsgPushNewWasmCodeResponse{
		CodeId: codeID,
	}, nil
}

// UpdateWasmCodeId defines a rpc handler method for MsgUpdateWasmCodeId
func (k Keeper) UpdateWasmCodeId(goCtx context.Context, msg *types.MsgUpdateWasmCodeId) (*types.MsgUpdateWasmCodeIdResponse, error) {
	if k.authority != msg.Signer {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority: expected %s, got %s", k.authority, msg.Signer)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)

	codeId := msg.CodeId
	if !store.Has(types.CodeID(codeId)) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCodeId, "code id %s does not exist", hex.EncodeToString(codeId))
	}

	clientId := msg.ClientId
	unknownClientState, found := k.clientKeeper.GetClientState(ctx, clientId)
	if !found {
		return nil, sdkerrors.Wrapf(clienttypes.ErrClientNotFound, "cannot update client with ID %s", clientId)
	}

	clientState, ok := unknownClientState.(*types.ClientState)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrInvalid, "client state type %T, expected %T", unknownClientState, (*types.ClientState)(nil))
	}

	clientState.CodeId = codeId

	k.clientKeeper.SetClientState(ctx, clientId, clientState)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			clienttypes.EventTypeUpdateWasmCodeId,
			sdk.NewAttribute(clienttypes.AttributeKeyClientID, clientId),
			sdk.NewAttribute(clienttypes.AttributeKeyWasmCodeID, hex.EncodeToString(codeId)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, clienttypes.AttributeValueCategory),
		),
	})

	return &types.MsgUpdateWasmCodeIdResponse{
		ClientId: clientId,
		CodeId:   codeId,
	}, nil
}
