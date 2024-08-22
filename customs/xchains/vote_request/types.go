package voterequest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/xchains-indexer/x/vote/exported"
)

type VoteRequestValue struct {
	Type   string          `json:"@type,omitempty"`
	Sender sdk.AccAddress  `json:"sender,omitempty"`
	PollID exported.PollID `json:"poll_id,omitempty"`
	Vote   any             `json:"vote,omitempty"`
}
