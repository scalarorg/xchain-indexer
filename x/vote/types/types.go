package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PollID uint64

var _ sdk.Msg = &VoteRequest{}

const MSG_VOTE_REQUEST = "axelar.vote.v1beta1.VoteRequest"

// GetSigners implements types.Msg.
func (m *VoteRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}

// ValidateBasic implements types.Msg.
func (*VoteRequest) ValidateBasic() error {
	return nil
}

// XXX_MessageName implements types.Msg.
// For codec compatibility.
func (*VoteRequest) XXX_MessageName() string {
	return MSG_VOTE_REQUEST
}
