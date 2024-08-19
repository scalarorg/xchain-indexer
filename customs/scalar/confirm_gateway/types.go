package confirmgateway

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

var _ sdk.Msg = &ConfirmGatewayTxRequest{}

// Hash wraps EVM Hash
type Hash common.Hash

// ZeroHash represents an empty 32-bytes hash
var ZeroHash = common.Hash{}

// IsZero returns true if the hash is empty; otherwise false
func (h Hash) IsZero() bool {
	return bytes.Equal(h.Bytes(), ZeroHash.Bytes())
}

// Bytes returns the actual byte array of the hash
func (h Hash) Bytes() []byte {
	return common.Hash(h).Bytes()
}

// Hex converts a hash to a hex string.
func (h Hash) Hex() string {
	return common.Hash(h).Hex()
}

// Marshal implements codec.ProtoMarshaler
func (h Hash) Marshal() ([]byte, error) {
	return h[:], nil
}

// MarshalTo implements codec.ProtoMarshaler
func (h Hash) MarshalTo(data []byte) (n int, err error) {
	bytesCopied := copy(data, h[:])
	if bytesCopied != common.HashLength {
		return 0, fmt.Errorf("expected data size to be %d, actual %d", common.HashLength, len(data))
	}

	return common.HashLength, nil
}

// Unmarshal implements codec.ProtoMarshaler
func (h *Hash) Unmarshal(data []byte) error {
	if len(data) != common.HashLength {
		return fmt.Errorf("expected data size to be %d, actual %d", common.HashLength, len(data))
	}

	*h = Hash(common.BytesToHash(data))

	return nil
}

// Size implements codec.ProtoMarshaler
func (h Hash) Size() int {
	return common.HashLength
}

// type ConfirmGatewayTxRequest struct {
// }

// GetSigners implements types.Msg.
func (*ConfirmGatewayTxRequest) GetSigners() []sdk.AccAddress {
	panic("unimplemented")
}

// // ProtoMessage implements types.Msg.
// func (*ConfirmGatewayTxRequest) ProtoMessage() {
// 	panic("unimplemented")
// }

// // Reset implements types.Msg.
// func (*ConfirmGatewayTxRequest) Reset() {
// 	panic("unimplemented")
// }

// // String implements types.Msg.
// func (*ConfirmGatewayTxRequest) String() string {
// 	panic("unimplemented")
// }

// ValidateBasic implements types.Msg.
func (*ConfirmGatewayTxRequest) ValidateBasic() error {
	panic("unimplemented")
}

// // These are the indexer's custom models
// // They are used to store the parsed data in the database

type ConfirmGatewayTxEvent struct {
	ID     uint
	Chain  string
	Sender string
	TxID   string
}
