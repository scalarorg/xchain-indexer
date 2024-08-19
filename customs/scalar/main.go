package main

import (
	"errors"
	"log"

	"github.com/DefiantLabs/cosmos-indexer/cmd"
	"github.com/DefiantLabs/cosmos-indexer/config"
	indexerTxTypes "github.com/DefiantLabs/cosmos-indexer/cosmos/modules/tx"
	"github.com/DefiantLabs/cosmos-indexer/db/models"
	"github.com/DefiantLabs/cosmos-indexer/filter"
	"github.com/DefiantLabs/cosmos-indexer/indexer"
	"github.com/DefiantLabs/cosmos-indexer/parsers"
	confirmgw "github.com/scalarorg/xchains-indexer/customs/scalar/confirm_gateway"
	"github.com/scalarorg/xchains-indexer/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	// dbTypes "github.com/DefiantLabs/cosmos-indexer/db"
	stdTypes "github.com/cosmos/cosmos-sdk/types"
)

// This defines the custom message parser for the call contract approve message type
// It implements the MessageParser interface
type CallContractRequestParser struct {
	Id string
}

func (c *CallContractRequestParser) Identifier() string {
	return c.Id
}

func (c *CallContractRequestParser) ParseMessage(cosmosMsg stdTypes.Msg, log *indexerTxTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	callContractRequest, ok := cosmosMsg.(*types.CallContractRequest)
	if !ok {
		return nil, errors.New("not a call contract approved message")
	}

	// contractCallApprovedEvent := ConfirmGatewayTxRequestEvent{
	// 	Chain:  confirmGatewayTxRequest.Chain.String(),
	// 	Sender: confirmGatewayTxRequest.Sender.String(),
	// 	txID:   confirmGatewayTxRequest.TxID.Hex(),
	// }
	storageVal := any(callContractRequest)
	return &storageVal, nil
}

// This method is called during database insertion. It is responsible for storing the parsed data in the database.
// The gorm db is wrapped in a transaction, so any errors will cause a rollback.
// Any errors returned will be saved as a parser error in the database as well for later debugging.
func (c *CallContractRequestParser) IndexMessage(dataset *any, db *gorm.DB, message models.Message, messageEvents []parsers.MessageEventWithAttributes, cfg config.IndexConfig) error {
	confirmGatewayTxRequestEvent, ok := (*dataset).(ConfirmGatewayTxRequestEvent)
	if !ok {
		return errors.New("failed to cast dataset to ConfirmGatewayTxRequestEvent")
	}

	confirmGatewayTxRequestEventModel := ConfirmGatewayTxRequestEvent{
		Chain:  confirmGatewayTxRequestEvent.Chain,
		Sender: confirmGatewayTxRequestEvent.Sender,
		txID:   confirmGatewayTxRequestEvent.txID,
	}

	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chain"}, {Name: "sender"}, {Name: "tx_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"chain", "sender", "tx_id"}),
	}).Create(&confirmGatewayTxRequestEventModel).Error

	return err
}

// These are the indexer's custom models
// They are used to store the parsed data in the database

type ConfirmGatewayTxRequestEvent struct {
	ID     uint
	Chain  string
	Sender string
	txID   string
}

func extendMessageParser(indexer *indexer.Indexer, msgType string, msgParser parsers.MessageParser) error {
	// This indexer is only concerned with delegate and undelegate messages, so we create regex filters to only index those messages.
	// This significantly reduces the size of the indexed dataset, saving space and processing time.
	requestRegexMessageTypeFilter, err := filter.NewRegexMessageTypeFilter("^" + msgType + "$")
	if err != nil {
		log.Fatalf("Failed to create regex message type filter. Err: %v", err)
		return err
	}
	indexer.RegisterMessageTypeFilter(requestRegexMessageTypeFilter)
	indexer.RegisterCustomMessageParser(msgType, msgParser)
	return nil
}
func extendConfirmGatewayIndexer(indexer *indexer.Indexer) {
	parser := &confirmgw.ConfirmGatewayParser{
		Id: "confirm-gateway",
	}
	err := extendMessageParser(indexer, "/axelar.evm.v1beta1.ConfirmGatewayTxRequest", parser)
	if err != nil {
		log.Fatalf("Failed to extend message parser. Err: %v", err)
	}
}
func main() {
	// Register the custom database models. They will be migrated and included in the database when the indexer finishes setup.
	customModels := []any{
		&confirmgw.ConfirmGatewayTxEvent{},
	}

	indexer := cmd.GetBuiltinIndexer()

	// Register the custom types that will modify the behavior of the indexer
	indexer.RegisterCustomModels(customModels)
	extendConfirmGatewayIndexer(indexer)

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("Failed to execute. Err: %v", err)
	}
}
