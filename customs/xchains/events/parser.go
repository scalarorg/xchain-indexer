package events

import (
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
)

const (
	EVENT_TYPE_MESSAGE                string = "message"
	EVENT_TYPE_CONTRACT_CALL_APPROVED string = "axelar.evm.v1beta1.ContractCallApproved"
	EVENT_TYPE_MESSAGE_PROCESSING     string = "axelar.nexus.v1beta1.MessageProcessing"
	EVENT_TYPE_MESSAGE_EXECUTED       string = "axelar.nexus.v1beta1.MessageExecuted"
)

func ExtendEventsIndexer(indexer *indexer.Indexer) error {
	beginBlockEventParsers := []parsers.BlockEventParser{
		&XChainsMessageEventParser{
			Id:      EVENT_TYPE_MESSAGE,
			Indexer: indexer,
		},
	}
	endBlockEventParsers := []parsers.BlockEventParser{
		&XChainsContractCallApprovedEventParser{
			Id:      EVENT_TYPE_CONTRACT_CALL_APPROVED,
			Indexer: indexer,
		},
		&XChainsMessageProcessingEventParser{
			Id:      EVENT_TYPE_MESSAGE_PROCESSING,
			Indexer: indexer,
		},
		&XChainsMessageExecutedEventParser{
			Id:      EVENT_TYPE_MESSAGE_EXECUTED,
			Indexer: indexer,
		},
		&XChainsMessageEventParser{
			Id:      EVENT_TYPE_MESSAGE,
			Indexer: indexer,
		},
	}
	for _, parser := range beginBlockEventParsers {
		indexer.RegisterCustomBeginBlockEventParser(parser.Identifier(), parser)
	}
	for _, parser := range endBlockEventParsers {
		indexer.RegisterCustomEndBlockEventParser(parser.Identifier(), parser)
	}
	return nil
}
