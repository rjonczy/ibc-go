package main

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type GenericMsg struct {
	Signer string
}

// GetSigners Implements Msg.
func (msg GenericMsg) GetSigners() []sdk.AccAddress {

	return []sdk.AccAddress{}
}

func (m GenericMsg) Reset()                    {}
func (m GenericMsg) ProtoMessage()             {}
func (msg GenericMsg) String() string          { return "" }
func (msg GenericMsg) XXX_MessageName() string { return "sdsdafsdaf" }

// Route Implements Msg
func (msg GenericMsg) Route() string { return "ddd" }

// Type Implements Msg
func (msg GenericMsg) Type() string { return "" }

// ValidateBasic Implements Msg.
func (msg GenericMsg) ValidateBasic() error {
	// this just makes sure all the inputs and outputs are properly formatted,
	// not that they actually have the money inside
	return nil
}

func main() {
	testMsg := banktypes.MsgSend{FromAddress: "test", ToAddress: "test", Amount: sdk.NewCoins()}

	bz, _ := testMsg.XXX_Marshal([]byte{}, true)

	unmarshaled := &banktypes.MsgSend{}
	err := unmarshaled.XXX_Unmarshal(bz)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(unmarshaled)

	fmt.Println(unmarshaled.Route())
	fmt.Println(sdk.MsgTypeURL(unmarshaled))

}
