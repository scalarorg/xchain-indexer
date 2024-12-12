package main

import (
	"log"

	"github.com/scalarorg/xchains-indexer/cmd"
	"github.com/scalarorg/xchains-indexer/customs/xchains/common"
	events "github.com/scalarorg/xchains-indexer/customs/xchains/events"
	"github.com/scalarorg/xchains-indexer/customs/xchains/messages"
)

func main() {
	// Get buildin indexer for extending with custom indexers
	indexer := cmd.GetBuiltinIndexer()
	// Register the custom database models. They will be migrated and included in the database when the indexer finishes setup.
	indexer.RegisterCustomModels([]any{
		&common.TxMessage{},
	})
	// indexer.RegisterCustomModuleBasics([]module.AppModuleBasic{
	// 	&nexus.AppModuleBasic{},
	// })
	messages.ExtendMessagesIndexer(indexer)
	events.ExtendEventsIndexer(indexer)

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("Failed to execute. Err: %v", err)
	}
}
