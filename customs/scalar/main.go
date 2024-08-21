package main

import (
	"log"

	"github.com/DefiantLabs/cosmos-indexer/cmd"
	"github.com/DefiantLabs/cosmos-indexer/indexer"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/xchains-indexer/customs/scalar/common"
	refundmsgrequest "github.com/scalarorg/xchains-indexer/customs/scalar/refund_msg_request"
	voteevents "github.com/scalarorg/xchains-indexer/customs/scalar/vote_events"
	voterequest "github.com/scalarorg/xchains-indexer/customs/scalar/vote_request"
	evmTypes "github.com/scalarorg/xchains-indexer/x/evm/types"
	rewardTypes "github.com/scalarorg/xchains-indexer/x/reward/types"
	voteTypes "github.com/scalarorg/xchains-indexer/x/vote/types"
)

func main() {
	// Get buildin indexer for extending with custom indexers
	buildinIndexer := cmd.GetBuiltinIndexer()
	// Register the custom database models. They will be migrated and included in the database when the indexer finishes setup.
	buildinIndexer.RegisterCustomModels([]any{
		&common.TxMessage{},
	})
	refundmsgrequest.ExtendRefundMsgRequestIndexer(buildinIndexer)
	voterequest.ExtendVoteRequestIndexer(buildinIndexer)
	voteevents.ExtendVoteEventsIndexer(buildinIndexer)
	buildinIndexer.PostSetupCustomFunction = func(dataset indexer.PostSetupCustomDataset) error {
		// Register msg types for the custom messages
		dataset.DB.AutoMigrate(&common.TxMessage{})
		if buildinIndexer.ChainClient != nil {
			buildinIndexer.ChainClient.Codec.InterfaceRegistry.RegisterInterface(rewardTypes.MSG_REWARD_REFUND_MSG_REQUEST, (*sdk.Msg)(nil), &rewardTypes.RefundMsgRequest{})
			buildinIndexer.ChainClient.Codec.InterfaceRegistry.RegisterInterface(voteTypes.MSG_VOTE_REQUEST, (*sdk.Msg)(nil), &voteTypes.VoteRequest{})
			buildinIndexer.ChainClient.Codec.InterfaceRegistry.RegisterInterface(evmTypes.MSG_EVM_VOTE_EVENTS, (*sdk.Msg)(nil), &evmTypes.VoteEvents{})
		}
		return nil
	}
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("Failed to execute. Err: %v", err)
	}
}
