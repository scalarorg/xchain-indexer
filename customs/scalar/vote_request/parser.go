package voterequest

import (
	"errors"
	"fmt"
	"log"

	"github.com/DefiantLabs/cosmos-indexer/config"
	txTypes "github.com/DefiantLabs/cosmos-indexer/cosmos/modules/tx"
	"github.com/DefiantLabs/cosmos-indexer/db/models"
	"github.com/DefiantLabs/cosmos-indexer/filter"
	"github.com/DefiantLabs/cosmos-indexer/indexer"
	"github.com/DefiantLabs/cosmos-indexer/parsers"
	"github.com/DefiantLabs/probe/client"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	stdTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/xchains-indexer/customs/scalar/common"
	evmTypes "github.com/scalarorg/xchains-indexer/x/evm/types"
	"github.com/scalarorg/xchains-indexer/x/vote/types"
	"gorm.io/gorm"
)

// This defines the custom message parser for the call contract approve message type
// It implements the MessageParser interface
type VoteRequestParser struct {
	Id      string
	Indexer *indexer.Indexer
}

func ExtendVoteRequestIndexer(indexer *indexer.Indexer) error {
	parser := &VoteRequestParser{
		Id:      "vote-request",
		Indexer: indexer,
	}
	messageTypeFilter, err := filter.NewRegexMessageTypeFilter("^/" + types.MSG_VOTE_REQUEST + "$")
	if err != nil {
		log.Fatalf("Failed to create regex message type filter. Err: %v", err)
		return err
	}
	// Register the custom database models. They will be migrated and included in the database when the indexer finishes setup.
	// indexer.RegisterCustomModels([]any{
	// 	&VoteRequestValue{},
	// })
	indexer.RegisterMessageTypeFilter(messageTypeFilter)
	indexer.RegisterCustomMessageParser("/"+types.MSG_VOTE_REQUEST, parser)
	if err != nil {
		log.Fatalf("Failed to extend message parser. Err: %v", err)
		return err
	}
	return nil
}

func (p *VoteRequestParser) Identifier() string {
	return p.Id
}

func (p *VoteRequestParser) ParseMessage(cosmosMsg stdTypes.Msg, logMsg *txTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	parsedMsg, ok := cosmosMsg.(*types.VoteRequest)
	if !ok {
		return nil, errors.New("not a Vote request")
	}
	// config.Log.Debugf("VoteRequest# PollId %d, Sender %s, Vote %+v", parsedMsg.PollID, parsedMsg.Sender, parsedMsg.Vote)
	parsedValue := VoteRequestValue{
		Type:   "/" + types.MSG_VOTE_REQUEST,
		PollID: parsedMsg.PollID,
		Sender: parsedMsg.Sender,
	}
	if parsedMsg.Vote != nil {
		msg, err := common.ParseInnerMessage(p.Indexer.ChainClient.Codec, parsedMsg.Vote, p.Indexer.CustomMessageParserRegistry, logMsg, cfg)
		if err == nil && msg != nil {
			config.Log.Debugf("VoteRequest# Success Parsed inner message")
			parsedValue.Vote = *msg
		} else {
			msg, err = p.ParseVoteEvents(p.Indexer.ChainClient.Codec, parsedMsg.Vote, logMsg, cfg)
			if err == nil {
				config.Log.Debugf("VoteRequest# Successfully parsed VoteEvents")
				parsedValue.Vote = *msg
			} else {
				config.Log.Debugf("VoteRequest# Failed to parse inner message")
			}
		}
	}
	storageVal := any(parsedValue)
	return &storageVal, nil
}
func (p *VoteRequestParser) ParseVoteEvents(codec client.Codec, vote *codecTypes.Any, logMsg *txTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	var events evmTypes.VoteEvents
	err := codec.InterfaceRegistry.UnpackAny(vote, &events)
	if err == nil {
		config.Log.Debugf("Successfully Unpacked message of type %s", vote.TypeUrl)
		config.Log.Debugf("VoteEvents# Events %v", events)
	} else {
		config.Log.Debug(fmt.Sprintf("ParseVoteEvents error: %v", err))
	}
	return nil, err
}

// This method is called during database insertion. It is responsible for storing the parsed data in the database.
// The gorm db is wrapped in a transaction, so any errors will cause a rollback.
// Any errors returned will be saved as a parser error in the database as well for later debugging.
func (p *VoteRequestParser) IndexMessage(dataset *any, db *gorm.DB, message models.Message, messageEvents []parsers.MessageEventWithAttributes, cfg config.IndexConfig) error {
	parsedEvent, ok := (*dataset).(VoteRequestValue)
	if !ok {
		return errors.New("failed to cast dataset to VoteRequestEvent")
	}
	fmt.Printf("Event %v", parsedEvent)

	// err := db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "chain"}, {Name: "sender"}, {Name: "tx_id"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"chain", "sender", "tx_id"}),
	// }).Create(&confirmGatewayTxEventModel).Error

	return nil
}
