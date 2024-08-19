package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &CallContractRequest{}

type CallContractRequest struct {
	Sender          sdk.AccAddress
	ChainName       string
	ContractAddress string
	Payload         []byte
	Fee             *sdk.Fee
}

// GetSigners implements types.Msg.
func (*CallContractRequest) GetSigners() []sdk.AccAddress {
	panic("unimplemented")
}

// ProtoMessage implements types.Msg.
func (*CallContractRequest) ProtoMessage() {
	panic("unimplemented")
}

// Reset implements types.Msg.
func (*CallContractRequest) Reset() {
	panic("unimplemented")
}

// String implements types.Msg.
func (*CallContractRequest) String() string {
	panic("unimplemented")
}

// ValidateBasic implements types.Msg.
func (*CallContractRequest) ValidateBasic() error {
	panic("unimplemented")
}
