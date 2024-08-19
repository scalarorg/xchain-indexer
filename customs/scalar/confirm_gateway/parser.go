package confirmgateway

import (
	"errors"
	"fmt"

	"github.com/DefiantLabs/cosmos-indexer/config"
	indexerTxTypes "github.com/DefiantLabs/cosmos-indexer/cosmos/modules/tx"
	"github.com/DefiantLabs/cosmos-indexer/db/models"
	"github.com/DefiantLabs/cosmos-indexer/parsers"
	stdTypes "github.com/cosmos/cosmos-sdk/types"
	"gorm.io/gorm"
)

// This defines the custom message parser for the call contract approve message type
// It implements the MessageParser interface
type ConfirmGatewayParser struct {
	Id string
}

func (c *ConfirmGatewayParser) Identifier() string {
	return c.Id
}

func (c *ConfirmGatewayParser) ParseMessage(cosmosMsg stdTypes.Msg, log *indexerTxTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	confirmGatewayTxRequest, ok := cosmosMsg.(*ConfirmGatewayTxRequest)
	if !ok {
		return nil, errors.New("not a confirm gateway tx request")
	}
	fmt.Println("confirmGatewayTxRequest %+v", confirmGatewayTxRequest)
	confirmGatewayTxEvent := ConfirmGatewayTxEvent{
		// Chain:  confirmGatewayTxRequest.Chain.String(),
		// Sender: confirmGatewayTxRequest.Sender.String(),
		// txID:   confirmGatewayTxRequest.TxID.Hex(),
	}
	storageVal := any(confirmGatewayTxEvent)
	return &storageVal, nil
}

// This method is called during database insertion. It is responsible for storing the parsed data in the database.
// The gorm db is wrapped in a transaction, so any errors will cause a rollback.
// Any errors returned will be saved as a parser error in the database as well for later debugging.
func (c *ConfirmGatewayParser) IndexMessage(dataset *any, db *gorm.DB, message models.Message, messageEvents []parsers.MessageEventWithAttributes, cfg config.IndexConfig) error {
	confirmGatewayTxEvent, ok := (*dataset).(ConfirmGatewayTxEvent)
	if !ok {
		return errors.New("failed to cast dataset to ConfirmGatewayTxEvent")
	}
	fmt.Printf("Event %v", confirmGatewayTxEvent)
	// confirmGatewayTxEventModel := ConfirmGatewayTxEvent{
	// 	Chain:  confirmGatewayTxEvent.Chain,
	// 	Sender: confirmGatewayTxEvent.Sender,
	// 	txID:   confirmGatewayTxEvent.txID,
	// }

	// err := db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "chain"}, {Name: "sender"}, {Name: "tx_id"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"chain", "sender", "tx_id"}),
	// }).Create(&confirmGatewayTxEventModel).Error

	return nil
}
