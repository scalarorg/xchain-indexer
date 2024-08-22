package parsers

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
	"gorm.io/gorm"
)

type BlockEventParser interface {
	Identifier() string
	ParseBlockEvent(abci.Event, config.IndexConfig) (*any, error)
	IndexBlockEvent(*any, *gorm.DB, models.Block, models.BlockEvent, []models.BlockEventAttribute, config.IndexConfig) error
}

type BlockEventParsedData struct {
	Data   *any
	Error  error
	Parser *BlockEventParser
}
