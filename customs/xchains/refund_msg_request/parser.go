package refundmsgrequest

import (
	"encoding/json"
	"errors"
	"log"

	stdTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/xchains-indexer/config"
	indexerTxTypes "github.com/scalarorg/xchains-indexer/cosmos/modules/tx"
	common "github.com/scalarorg/xchains-indexer/customs/xchains/common"
	voterequest "github.com/scalarorg/xchains-indexer/customs/xchains/vote_request"
	"github.com/scalarorg/xchains-indexer/db/models"
	"github.com/scalarorg/xchains-indexer/filter"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
	rewardTypes "github.com/scalarorg/xchains-indexer/x/reward/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// This defines the custom message parser for the call contract approve message type
// It implements the MessageParser interface
type RefundMsgRequestParser struct {
	Id      string
	Indexer *indexer.Indexer
}

func ExtendRefundMsgRequestIndexer(indexer *indexer.Indexer) error {
	parser := &RefundMsgRequestParser{
		Id:      "refund-msg-request",
		Indexer: indexer,
	}
	requestRegexMessageTypeFilter, err := filter.NewRegexMessageTypeFilter("^/" + rewardTypes.MSG_REWARD_REFUND_MSG_REQUEST + "$")
	if err != nil {
		log.Fatalf("Failed to create regex message type filter. Err: %v", err)
		return err
	}

	indexer.RegisterMessageTypeFilter(requestRegexMessageTypeFilter)
	indexer.RegisterCustomMessageParser("/"+rewardTypes.MSG_REWARD_REFUND_MSG_REQUEST, parser)
	if err != nil {
		log.Fatalf("Failed to extend message parser. Err: %v", err)
		return err
	}
	return nil
}

func (p *RefundMsgRequestParser) Identifier() string {
	return p.Id
}

func (p *RefundMsgRequestParser) ParseMessage(cosmosMsg stdTypes.Msg, messageLog *indexerTxTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	parsedMsg, ok := cosmosMsg.(*rewardTypes.RefundMsgRequest)
	if !ok {
		return nil, errors.New("not a refund message request")
	}
	parsedMessageValue := RefundMsgRequestValue{
		Type:   "/" + rewardTypes.MSG_REWARD_REFUND_MSG_REQUEST,
		Sender: parsedMsg.Sender,
	}
	innerMsg := parsedMsg.InnerMessage.GetCachedValue()
	if innerMsg != nil {
		voteRequest, ok := (innerMsg).(voterequest.VoteRequestValue)
		if ok {
			parsedMessageValue.InnerMessage = &voteRequest
			storageVal := any(parsedMessageValue)
			return &storageVal, nil
		} else {
			config.Log.Debug("RefundMsgRequestParser# Failed to cast inner message to VoteRequestValue")
		}
	} else {
		msg, err := common.ParseInnerMessage(p.Indexer.ChainClient.Codec, parsedMsg.InnerMessage, p.Indexer.CustomMessageParserRegistry, messageLog, cfg)
		if msg != nil && err == nil {
			voteRequest, ok := (*msg).(voterequest.VoteRequestValue)
			if ok {
				parsedMessageValue.InnerMessage = &voteRequest
				storageVal := any(parsedMessageValue)
				return &storageVal, nil
			} else {
				config.Log.Debug("RefundMsgRequestParser# Failed to cast inner message to VoteRequestValue")
			}

		} else {
			config.Log.Debug("RefundMsgRequestParser# Failed to parse inner message")
		}
	}

	return nil, nil
}

// This method is called during database insertion. It is responsible for storing the parsed data in the database.
// The gorm db is wrapped in a transaction, so any errors will cause a rollback.
// Any errors returned will be saved as a parser error in the database as well for later debugging.
func (p *RefundMsgRequestParser) IndexMessage(dataset *any, db *gorm.DB, message models.Message, messageEvents []parsers.MessageEventWithAttributes, cfg config.IndexConfig) error {
	parsedMessageEvent, ok := (*dataset).(RefundMsgRequestValue)
	if !ok {
		return errors.New("failed to cast dataset to RefundMsgRequestValue")
	}
	jsonValue, err := json.Marshal(parsedMessageEvent)
	if err == nil {
		txMessage := common.TxMessage{
			Tx:            message.Tx,
			MessageID:     message.ID,
			BlockId:       message.Tx.BlockID,
			MessageDetail: string(jsonValue),
		}

		err = db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "tx_id"}, {Name: "message_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"message_detail"}),
		}).Create(&txMessage).Error
		if err != nil {
			config.Log.Debugf("RefundMsgRequestParser# Failed to save message event %v", err)
		}
	} else {
		config.Log.Debugf("RefundMsgRequestParser# Failed to marshal message event %v", err)
		return err
	}

	return err
}
