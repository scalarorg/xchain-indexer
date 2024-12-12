package core

import (
	"errors"
	"fmt"

	probeClient "github.com/DefiantLabs/probe/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

// Provides an in-app tx decoder.
// The primary use-case for this function is to allow fallback decoding if a TX fails to decode after RPC requests.
// This can happen in a number of scenarios, but mainly due to missing proto definitions.
// We can attempt a personal decode of the TX, and see if we can continue indexing based on in-app conditions (such as message type filters).
// This function skips a large chunk of decoding validations, and is not recommended for general use. Its main point is to skip errors that in
// default Cosmos TX decoders would cause the entire decode to fail.
func InAppTxDecoder(cdc probeClient.Codec) sdk.TxDecoder {
	return func(txBytes []byte) (sdk.Tx, error) {
		var raw tx.TxRaw
		var err error

		fmt.Println("--- InAppTxDecoder")
		err = cdc.Marshaler.Unmarshal(txBytes, &raw)
		if err != nil {
			return nil, err
		}

		var body tx.TxBody

		err = body.Unmarshal(raw.BodyBytes)
		if err != nil {
			return nil, errors.New("failed to unmarshal tx body")
		}

		for _, any := range body.Messages {
			if newAny, ok := convertMaintainerRequestAny(cdc, any); ok {
				// Find the index of the current message in body.Messages
				for i, msg := range body.Messages {
					if msg == any {
						body.Messages[i] = newAny
						break
					}
				}
				fmt.Println("--- newAny cached value", newAny.GetCachedValue())
			}
			var msg sdk.Msg
			// We deliberately ignore errors here to build up a
			// list of properly decoded messages for later analysis.
			cdc.Marshaler.UnpackAny(any, &msg) //nolint:errcheck
		}

		var authInfo tx.AuthInfo

		err = cdc.Marshaler.Unmarshal(raw.AuthInfoBytes, &authInfo)
		if err != nil {
			return nil, errors.New("failed to unmarshal auth info")
		}

		theTx := &tx.Tx{
			Body:       &body,
			AuthInfo:   &authInfo,
			Signatures: raw.Signatures,
		}

		return theTx, nil
	}
}

// convertMaintainerRequestAny attempts to convert a RegisterChainMaintainerRequest Any message
// to a new Any with the same type URL. Returns the new Any and a boolean indicating success.
func convertMaintainerRequestAny(cdc probeClient.Codec, any *types.Any) (*types.Any, bool) {
	maintainerReq, ok := GetMessageTypeFromURL(any.TypeUrl)
	if !ok {
		return nil, false
	}

	// Type assert maintainerReq to proto.Message which is what Unmarshal expects
	protoMsg, ok := maintainerReq.(codec.ProtoMarshaler)
	if !ok {
		return nil, false
	}

	err := cdc.Marshaler.Unmarshal(any.Value, protoMsg)
	if err != nil {
		return nil, false
	}

	newAny, err := types.NewAnyWithValue(protoMsg)
	if err != nil {
		fmt.Printf("Failed to create new Any: %v\n", err)
		return nil, false
	}

	newAny.TypeUrl = any.TypeUrl
	return newAny, true
}
