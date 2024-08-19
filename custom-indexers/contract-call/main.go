package main

import (
	"errors"
	"log"

	"github.com/DefiantLabs/cosmos-indexer/cmd"
	"github.com/DefiantLabs/cosmos-indexer/config"
	indexerTxTypes "github.com/DefiantLabs/cosmos-indexer/cosmos/modules/tx"
	"github.com/DefiantLabs/cosmos-indexer/db/models"
	"github.com/DefiantLabs/cosmos-indexer/filter"
	"github.com/DefiantLabs/cosmos-indexer/parsers"
	"github.com/axelarnetwork/axelar-core/x/axelarnet/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	// dbTypes "github.com/DefiantLabs/cosmos-indexer/db"
	stdTypes "github.com/cosmos/cosmos-sdk/types"
)

// This defines the custom message parser for the call contract approve message type
// It implements the MessageParser interface
type ConfirmGatewayTxRequestParser struct {
	Id string
}

func (c *ConfirmGatewayTxRequestParser) Identifier() string {
	return c.Id
}

func (c *ConfirmGatewayTxRequestParser) ParseMessage(cosmosMsg stdTypes.Msg, log *indexerTxTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	confirmGatewayTxRequest, ok := cosmosMsg.(*types.CallContractRequest)
	if !ok {
		return nil, errors.New("not a call contract approved message")
	}

	contractCallApprovedEvent := ConfirmGatewayTxRequestEvent{
		Chain:  confirmGatewayTxRequest.Chain.String(),
		Sender: confirmGatewayTxRequest.Sender.String(),
		txID:   confirmGatewayTxRequest.TxID.Hex(),
	}
	storageVal := any(contractCallApprovedEvent)
	return &storageVal, nil
}

// This method is called during database insertion. It is responsible for storing the parsed data in the database.
// The gorm db is wrapped in a transaction, so any errors will cause a rollback.
// Any errors returned will be saved as a parser error in the database as well for later debugging.
func (c *ConfirmGatewayTxRequestParser) IndexMessage(dataset *any, db *gorm.DB, message models.Message, messageEvents []parsers.MessageEventWithAttributes, cfg config.IndexConfig) error {
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

func main() {
	// Register the custom database models. They will be migrated and included in the database when the indexer finishes setup.
	customModels := []any{
		&ConfirmGatewayTxRequestEvent{},
	}

	indexer := cmd.GetBuiltinIndexer()

	// Register the custom types that will modify the behavior of the indexer
	indexer.RegisterCustomModels(customModels)

	// This indexer is only concerned with delegate and undelegate messages, so we create regex filters to only index those messages.
	// This significantly reduces the size of the indexed dataset, saving space and processing time.
	confirmGatewayTxRequestRegexMessageTypeFilter, err := filter.NewRegexMessageTypeFilter("^axelar.axelarnet.v1beta1.CallContractRequest$")
	if err != nil {
		log.Fatalf("Failed to create regex message type filter. Err: %v", err)
	}

	indexer.RegisterMessageTypeFilter(confirmGatewayTxRequestRegexMessageTypeFilter)

	// Register the custom message parser for the delegation message types. Our parser can handle both delegate and undelegate messages.
	// However, they must be uniquely identified by the Identifier() method. This will make identifying any parser errors easier.
	confirmGatewayTxRequestParser := &ConfirmGatewayTxRequestParser{Id: "contract-call"}
	indexer.RegisterCustomMessageParser("/axelar.axelarnet.v1beta1.CallContractRequest", confirmGatewayTxRequestParser)

	err = cmd.Execute()
	if err != nil {
		log.Fatalf("Failed to execute. Err: %v", err)
	}
}
