package messages

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	stdTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/xchains-indexer/config"
	indexerTxTypes "github.com/scalarorg/xchains-indexer/cosmos/modules/tx"
	"github.com/scalarorg/xchains-indexer/db/models"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
	nexusTypes "github.com/scalarorg/xchains-indexer/x/nexus/types"
	"gorm.io/gorm"
)

type RegisterChainMaintainerRequestParser struct {
	Id      string
	Indexer *indexer.Indexer
}

func (p *RegisterChainMaintainerRequestParser) Identifier() string {
	return p.Id
}

func (p *RegisterChainMaintainerRequestParser) ParseMessage(cosmosMsg stdTypes.Msg, logMsg *indexerTxTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	parsedMsg, ok := cosmosMsg.(*nexusTypes.RegisterChainMaintainerRequest)
	if !ok {
		return nil, errors.New("not a RegisterChainMaintainerRequest")
	}

	parsedValue := RegisterChainMaintainerRequestValue{
		Type:   "/" + nexusTypes.MSG_NEXUS_REGISTER_CHAIN_MAINTAINER_REQUEST,
		Sender: parsedMsg.Sender,
		Chains: parsedMsg.Chains,
	}

	storageVal := any(parsedValue)
	return &storageVal, nil
}

func (p *RegisterChainMaintainerRequestParser) ParseAnyMessage(cosmosMsg types.Any, log *indexerTxTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	return nil, errors.New("not implemented")
}

func (p *RegisterChainMaintainerRequestParser) IndexMessage(dataset *any, db *gorm.DB, message models.Message, messageEvents []parsers.MessageEventWithAttributes, cfg config.IndexConfig) error {
	parsedData, ok := (*dataset).(RegisterChainMaintainerRequestValue)
	if !ok {
		return errors.New("failed to cast dataset to RegisterChainMaintainerRequestValue")
	}
	fmt.Printf("Event %v", parsedData)

	return nil
}
