package events

import (
	"encoding/base64"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
	"github.com/scalarorg/xchains-indexer/indexer"
	"gorm.io/gorm"
)

type XChainsMessageProcessingEventParser struct {
	Id      string
	Indexer *indexer.Indexer
}

func (p *XChainsMessageProcessingEventParser) Identifier() string {
	return p.Id
}

func (p *XChainsMessageProcessingEventParser) ParseBlockEvent(event abci.Event, conf config.IndexConfig) (*any, error) {
	config.Log.Debugf("XChainsMessageEventParser# ParseBlockEvent Type: %s", event.Type)
	for _, attribute := range event.Attributes {
		var value string
		var keyItem string
		if conf.Flags.BlockEventsBase64Encoded {
			// Should we even be decoding these from base64? What are the implications?
			valueBytes, err := base64.StdEncoding.DecodeString(attribute.Value)
			if err != nil {
				return nil, err
			}

			keyBytes, err := base64.StdEncoding.DecodeString(attribute.Key)
			if err != nil {
				return nil, err
			}

			value = string(valueBytes)
			keyItem = string(keyBytes)
		} else {
			value = attribute.Value
			keyItem = attribute.Key
		}
		config.Log.Debugf("Attributes# Key: %s, Value: %s, Index %t", keyItem, value, attribute.GetIndex())

	}
	return nil, nil
}

func (p *XChainsMessageProcessingEventParser) IndexBlockEvent(parsedData *any, db *gorm.DB, block models.Block, blockEvent models.BlockEvent, blockEventAttrs []models.BlockEventAttribute, cfg config.IndexConfig) error {
	return nil
}
