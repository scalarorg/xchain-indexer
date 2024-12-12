package messages

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/xchains-indexer/config"
	evmTypes "github.com/scalarorg/xchains-indexer/x/evm/types"
	github_com_axelarnetwork_axelar_core_x_nexus_exported "github.com/scalarorg/xchains-indexer/x/nexus/exported"
	"github.com/scalarorg/xchains-indexer/x/vote/exported"
)

type RefundMsgRequestValue struct {
	Type         string         `json:"@type,omitempty"`
	Sender       sdk.AccAddress `json:"sender,omitempty"`
	InnerMessage any            `json:"inner_message,omitempty"`
}

type VoteRequestValue struct {
	Type   string          `json:"@type,omitempty"`
	Sender sdk.AccAddress  `json:"sender,omitempty"`
	PollID exported.PollID `json:"poll_id,omitempty"`
	Vote   any             `json:"vote,omitempty"`
}

type VoteEvents struct {
	Type   string                                                          `json:"@type,omitempty"`
	Chain  github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName `json:"chain,omitempty"`
	Events []Event                                                         `json:"events,omitempty"`
}

type RegisterChainMaintainerRequestValue struct {
	Type   string                                                            `json:"@type,omitempty"`
	Sender sdk.AccAddress                                                    `json:"sender,omitempty"`
	Chains []github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName `json:"chains,omitempty"`
}

// type Event struct {
// 	Chain  github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/axelarnetwork/axelar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
// 	TxID   Hash                                                            `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3,customtype=Hash" json:"tx_id"`
// 	Index  uint64                                                          `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
// 	Status Event_Status                                                    `protobuf:"varint,4,opt,name=status,proto3,enum=axelar.evm.v1beta1.Event_Status" json:"status,omitempty"`
// 	// Types that are valid to be assigned to Event:
// 	//	*Event_TokenSent
// 	//	*Event_ContractCall
// 	//	*Event_ContractCallWithToken
// 	//	*Event_Transfer
// 	//	*Event_TokenDeployed
// 	//	*Event_MultisigOwnershipTransferred
// 	//	*Event_MultisigOperatorshipTransferred
// 	Event isEvent_Event `protobuf_oneof:"event"`
// }

type EventBase struct {
	Chain  github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName `json:"chain,omitempty"`
	Index  uint64                                                          `json:"index,omitempty"`
	Status string                                                          `json:"status,omitempty"`
	TxID   evmTypes.Hash                                                   `json:"tx_id,omitempty"`
}

type isEvent_Event interface {
	MarshalTo([]byte) (int, error)
	Size() int
}

type Event struct {
	EventBase
	Event isEvent_Event `json:"event,omitempty"`
}

func (p *Event) MarshalJSON() ([]byte, error) {
	if p.Event == nil {
		return nil, nil
	}
	switch v := (p.Event).(type) {
	case *evmTypes.Event_ContractCall:
		return json.Marshal(&struct {
			EventBase
			ContractCall *evmTypes.EventContractCall `json:"contract_call,omitempty"`
		}{
			EventBase:    p.EventBase,
			ContractCall: p.Event.(*evmTypes.Event_ContractCall).ContractCall,
		})
	case *evmTypes.Event_ContractCallWithToken:
		return json.Marshal(&struct {
			EventBase
			ContractCallWithToken *evmTypes.EventContractCallWithToken `json:"contract_call_with_token,omitempty"`
		}{
			EventBase:             p.EventBase,
			ContractCallWithToken: p.Event.(*evmTypes.Event_ContractCallWithToken).ContractCallWithToken,
		})
	case *evmTypes.Event_MultisigOperatorshipTransferred:
		return json.Marshal(&struct {
			EventBase
			MultisigOperatorshipTransferred *evmTypes.EventMultisigOperatorshipTransferred `json:"multisig_operatorship_transferred,omitempty"`
		}{
			EventBase:                       p.EventBase,
			MultisigOperatorshipTransferred: p.Event.(*evmTypes.Event_MultisigOperatorshipTransferred).MultisigOperatorshipTransferred,
		})
	case *evmTypes.Event_MultisigOwnershipTransferred:
		return json.Marshal(&struct {
			EventBase
			MultisigOwnershipTransferred *evmTypes.EventMultisigOwnershipTransferred `json:"multisig_ownership_transferred,omitempty"`
		}{
			EventBase:                    p.EventBase,
			MultisigOwnershipTransferred: p.Event.(*evmTypes.Event_MultisigOwnershipTransferred).MultisigOwnershipTransferred,
		})
	case *evmTypes.Event_TokenDeployed:
		return json.Marshal(&struct {
			EventBase
			TokenDeployed *evmTypes.EventTokenDeployed `json:"token_deployed,omitempty"`
		}{
			EventBase:     p.EventBase,
			TokenDeployed: p.Event.(*evmTypes.Event_TokenDeployed).TokenDeployed,
		})
	case *evmTypes.Event_TokenSent:
		return json.Marshal(&struct {
			EventBase
			TokenSent *evmTypes.EventTokenSent `json:"token_sent,omitempty"`
		}{
			EventBase: p.EventBase,
			TokenSent: p.Event.(*evmTypes.Event_TokenSent).TokenSent,
		})
	case *evmTypes.Event_Transfer:
		return json.Marshal(&struct {
			EventBase
			Transfer *evmTypes.EventTransfer `json:"transfer,omitempty"`
		}{
			EventBase: p.EventBase,
			Transfer:  p.Event.(*evmTypes.Event_Transfer).Transfer,
		})
	default:
		config.Log.Debugf("Event default %v, %v", p.Event, v)
		return nil, nil
	}
}
