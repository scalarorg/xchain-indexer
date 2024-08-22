package refundmsgrequest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type RefundMsgRequestValue struct {
	Type         string         `json:"@type,omitempty"`
	Sender       sdk.AccAddress `json:"sender,omitempty"`
	InnerMessage any            `json:"inner_message,omitempty"`
}
