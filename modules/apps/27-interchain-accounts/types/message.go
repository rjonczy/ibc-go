package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SerializeCosmosTx(cdc codec.BinaryCodec, data interface{}) ([]byte, error) {

	msgs := make([]sdk.Msg, 0)
	switch data := data.(type) {
	case sdk.Msg:
		msgs = append(msgs, data)
	case []sdk.Msg:
		if len(data) == 0 {
			return []byte{}, nil
		}
		msgs = append(msgs, data...)
	default:
		return nil, ErrInvalidOutgoingData
	}

	msgAnys := make([]*codectypes.Any, len(msgs))

	for i, msg := range msgs {
		var err error
		msgAnys[i], err = codectypes.NewAnyWithValue(msg)
		if err != nil {
			return nil, err
		}
	}

	txBody := &IBCTxBody{
		Messages: msgAnys,
	}

	txRaw := &IBCTxRaw{
		BodyBytes: cdc.MustMarshal(txBody),
	}

	bz, err := cdc.Marshal(txRaw)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func DeserializeTx(cdc codec.BinaryCodec, txBytes []byte) ([]sdk.Msg, error) {
	if len(txBytes) == 0 {
		return []sdk.Msg{}, nil
	}
	var txRaw IBCTxRaw

	err := cdc.Unmarshal(txBytes, &txRaw)
	if err != nil {
		return nil, err
	}

	var txBody IBCTxBody

	err = cdc.Unmarshal(txRaw.BodyBytes, &txBody)
	if err != nil {
		return nil, err
	}

	anys := txBody.Messages
	res := make([]sdk.Msg, len(anys))
	for i, any := range anys {
		var msg sdk.Msg
		err := cdc.UnpackAny(any, &msg)
		if err != nil {
			return nil, err
		}
		res[i] = msg
	}

	return res, nil
}
