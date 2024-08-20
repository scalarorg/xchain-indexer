package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &RefundMsgRequest{}

const MSG_REWARD_REFUND_MSG_REQUEST = "axelar.reward.v1beta1.RefundMsgRequest"

// GetSigners implements types.Msg.
func (m *RefundMsgRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}

// ValidateBasic implements types.Msg.
func (*RefundMsgRequest) ValidateBasic() error {
	return nil
}

// XXX_MessageName implements types.Msg.
// For codec compatibility.
func (*RefundMsgRequest) XXX_MessageName() string {
	return MSG_REWARD_REFUND_MSG_REQUEST
}
