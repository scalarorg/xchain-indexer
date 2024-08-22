package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
	"github.com/scalarorg/xchains-indexer/indexer"
	"gorm.io/gorm"
)

type XChainsPollCompletedEventParser struct {
	Id      string
	Indexer *indexer.Indexer
}

func (p *XChainsPollCompletedEventParser) Identifier() string {
	return p.Id
}

func (p *XChainsPollCompletedEventParser) ParseBlockEvent(abci.Event, config.IndexConfig) (*any, error) {
	return nil, nil
}

func (p *XChainsPollCompletedEventParser) IndexBlockEvent(*any, *gorm.DB, models.Block, models.BlockEvent, []models.BlockEventAttribute, config.IndexConfig) error {
	return nil
}
