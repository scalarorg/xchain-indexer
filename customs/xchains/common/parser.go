package common

import (
	"fmt"

	"github.com/DefiantLabs/probe/client"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	stdTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/xchains-indexer/config"
	txTypes "github.com/scalarorg/xchains-indexer/cosmos/modules/tx"
	"github.com/scalarorg/xchains-indexer/parsers"
)

type IAnyMessageParser interface {
	ParseAnyMessage(types.Any, *txTypes.LogMessage, config.IndexConfig) (*any, error)
}

func ParseInnerMessage(codec client.Codec, msg *types.Any, customParsers map[string][]parsers.MessageParser, messageLog *txTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	var message stdTypes.Msg
	err := codec.InterfaceRegistry.UnpackAny(msg, &message)
	if err == nil && message != nil {
		config.Log.Debugf("Successfully Unpacked message of type %s", msg.TypeUrl)
		if customParsers != nil {
			if customMessageParsers, ok := customParsers[msg.TypeUrl]; ok {
				for _, customParser := range customMessageParsers {
					// We deliberately ignore the error here, as we want to continue processing the message even if a custom parser fails
					config.Log.Debugf("ParseInnerMessage# Parse message of type %s with customParser", msg.TypeUrl)
					parsedData, err := customParser.ParseMessage(message, messageLog, cfg)
					if err != nil && parsedData == nil {
						config.Log.Error(fmt.Sprintf("ParseInnerMessage# Error parsing message: %v", err))
						return nil, err
					} else if parsedData != nil {
						config.Log.Debug("ParseInnerMessage# Susscess Parsed inner message")
						return parsedData, nil
					}
				}
			}
		}
	}
	config.Log.Debugf("ParseInnerMessage# Cannot parse message of type %s", msg.TypeUrl)
	return nil, err
}
func MsgParse(codec client.Codec, msg *types.Any, customParsers map[string][]parsers.MessageParser) *stdTypes.Msg {
	var currMsgUnpack stdTypes.Msg
	err := codec.InterfaceRegistry.UnpackAny(msg, &currMsgUnpack)
	if err == nil && currMsgUnpack != nil {
		config.Log.Debugf("Successfully Unpacked message of type %s", msg.TypeUrl)
	} else {
		config.Log.Debug(fmt.Sprintf("Error unpacking message: %v", err))

	}
	return &currMsgUnpack
}
