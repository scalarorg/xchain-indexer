package voteevents

import (
	"errors"
	"fmt"
	"log"

	"github.com/DefiantLabs/cosmos-indexer/config"
	indexerTxTypes "github.com/DefiantLabs/cosmos-indexer/cosmos/modules/tx"
	"github.com/DefiantLabs/cosmos-indexer/db/models"
	"github.com/DefiantLabs/cosmos-indexer/filter"
	"github.com/DefiantLabs/cosmos-indexer/indexer"
	"github.com/DefiantLabs/cosmos-indexer/parsers"
	"github.com/cosmos/cosmos-sdk/codec/types"
	stdTypes "github.com/cosmos/cosmos-sdk/types"
	evmTypes "github.com/scalarorg/xchains-indexer/x/evm/types"
	"gorm.io/gorm"
)

// This defines the custom message parser for the call contract approve message type
// It implements the MessageParser interface
type VoteEventsParser struct {
	Id      string
	Indexer *indexer.Indexer
}

func ExtendVoteEventsIndexer(indexer *indexer.Indexer) error {
	parser := &VoteEventsParser{
		Id:      "vote-events",
		Indexer: indexer,
	}
	messageTypeFilter, err := filter.NewRegexMessageTypeFilter("^/" + evmTypes.MSG_EVM_VOTE_EVENTS + "$")
	if err != nil {
		log.Fatalf("Failed to create regex message type filter. Err: %v", err)
		return err
	}
	// Register the custom database models. They will be migrated and included in the database when the indexer finishes setup.
	// indexer.RegisterCustomModels([]any{
	// 	&VoteEvents{},
	// })
	indexer.RegisterMessageTypeFilter(messageTypeFilter)
	indexer.RegisterCustomMessageParser("/"+evmTypes.MSG_EVM_VOTE_EVENTS, parser)
	if err != nil {
		log.Fatalf("Failed to extend message parser. Err: %v", err)
		return err
	}
	return nil
}

func (p *VoteEventsParser) Identifier() string {
	return p.Id
}

func (p *VoteEventsParser) ParseMessage(cosmosMsg stdTypes.Msg, logMsg *indexerTxTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	parsedMsg, ok := cosmosMsg.(*evmTypes.VoteEvents)
	if !ok {
		return nil, errors.New("not a VoteEvents")
	}
	parsedValue := VoteEvents{
		Type:  "/" + evmTypes.MSG_EVM_VOTE_EVENTS,
		Chain: parsedMsg.Chain,
	}
	if parsedMsg.Events != nil {
		parsedValue.Events = make([]Event, len(parsedMsg.Events))
		for index, event := range parsedMsg.Events {
			parsedValue.Events[index] = Event{
				EventBase: EventBase{
					Chain:  event.Chain,
					Index:  event.Index,
					Status: event.Status.String(),
					TxID:   event.TxID,
				},
				Event: event.Event,
			}
			//config.Log.Debugf("VoteEvent# Chain %s, Event %v, Index %d, Status %+v, txId %s", event.Chain, event.Event, event.Index, event.Status, event.TxID)
		}
	}
	storageVal := any(parsedValue)
	return &storageVal, nil
}

func (p *VoteEventsParser) ParseAnyMessage(cosmosMsg types.Any, log *indexerTxTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	return nil, errors.New("not a confirm gateway tx request")
}

// This method is called during database insertion. It is responsible for storing the parsed data in the database.
// The gorm db is wrapped in a transaction, so any errors will cause a rollback.
// Any errors returned will be saved as a parser error in the database as well for later debugging.
func (p *VoteEventsParser) IndexMessage(dataset *any, db *gorm.DB, message models.Message, messageEvents []parsers.MessageEventWithAttributes, cfg config.IndexConfig) error {
	parsedData, ok := (*dataset).(VoteEvents)
	if !ok {
		return errors.New("failed to cast dataset to ConfirmGatewayTxEvent")
	}
	fmt.Printf("Event %v", parsedData)

	// err := db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "chain"}, {Name: "sender"}, {Name: "tx_id"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"chain", "sender", "tx_id"}),
	// }).Create(&confirmGatewayTxEventModel).Error

	return nil
}
