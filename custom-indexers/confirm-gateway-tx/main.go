package main

import (
	"log"

	"github.com/DefiantLabs/cosmos-indexer/cmd"
	"github.com/DefiantLabs/cosmos-indexer/filter"
)

func main() {
	// Register the custom database models. They will be migrated and included in the database when the indexer finishes setup.
	customModels := []any{}

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

	// // Register the custom message parser for the delegation message types. Our parser can handle both delegate and undelegate messages.
	// // However, they must be uniquely identified by the Identifier() method. This will make identifying any parser errors easier.
	// confirmGatewayTxRequestParser := &ConfirmGatewayTxRequestParser{Id: "contract-call"}
	// indexer.RegisterCustomMessageParser("/axelar.axelarnet.v1beta1.CallContractRequest", confirmGatewayTxRequestParser)

	err = cmd.Execute()
	if err != nil {
		log.Fatalf("Failed to execute. Err: %v", err)
	}
}
