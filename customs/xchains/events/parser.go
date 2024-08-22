package events

import (
	"github.com/scalarorg/xchains-indexer/indexer"
)

const (
	EVENT_TYPE_MESSAGE string = "message"
)

func ExtendEventsIndexer(indexer *indexer.Indexer) error {
	messageParser := &XChainsMessageEventParser{
		Id:      "xchains-message",
		Indexer: indexer,
	}
	// pollCompletedParser := &XChainsPollCompletedEventParser{
	// 	Id:      "xchains-pollcompleted",
	// 	Indexer: indexer,
	// }
	// requestRegexMessageTypeFilter, err := filter.NewRegexMessageTypeFilter("^/" + rewardTypes.MSG_REWARD_REFUND_MSG_REQUEST + "$")
	// if err != nil {
	// 	log.Fatalf("Failed to create regex message type filter. Err: %v", err)
	// 	return err
	// }

	// indexer.RegisterMessageTypeFilter(requestRegexMessageTypeFilter)
	indexer.RegisterCustomBeginBlockEventParser(EVENT_TYPE_MESSAGE, messageParser)
	indexer.RegisterCustomEndBlockEventParser(EVENT_TYPE_MESSAGE, messageParser)
	// if err != nil {
	// 	log.Fatalf("Failed to extend message parser. Err: %v", err)
	// 	return err
	// }
	return nil
}
