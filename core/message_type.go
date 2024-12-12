package core

import (
	nexusTypes "github.com/scalarorg/xchains-indexer/x/nexus/types"
)

const (
	// BatchRequestTypeURL          = "/axelar.auxiliary.v1beta1.BatchRequest"
	// BatchResponseTypeURL         = "/axelar.auxiliary.v1beta1.BatchResponse"
	// BatchResponseResponseTypeURL = "/axelar.auxiliary.v1beta1.BatchResponse.Response"

	// LinkRequest                     = "axelar.axelarnet.v1beta1.LinkRequest"
	// LinkResponse                    = "axelar.axelarnet.v1beta1.LinkResponse"
	// ConfirmDepositRequest           = "axelar.axelarnet.v1beta1.ConfirmDepositRequest"
	// ConfirmDepositResponse          = "axelar.axelarnet.v1beta1.ConfirmDepositResponse"
	// ExecutePendingTransfersRequest  = "axelar.axelarnet.v1beta1.ExecutePendingTransfersRequest"
	// ExecutePendingTransfersResponse = "axelar.axelarnet.v1beta1.ExecutePendingTransfersResponse"
	// RegisterIBCPathRequest          = "axelar.axelarnet.v1beta1.RegisterIBCPathRequest"
	// RegisterIBCPathResponse         = "axelar.axelarnet.v1beta1.RegisterIBCPathResponse"
	// AddCosmosBasedChainRequest      = "axelar.axelarnet.v1beta1.AddCosmosBasedChainRequest"
	// AddCosmosBasedChainResponse     = "axelar.axelarnet.v1beta1.AddCosmosBasedChainResponse"
	// RegisterAssetRequest            = "axelar.axelarnet.v1beta1.RegisterAssetRequest"
	// RegisterAssetResponse           = "axelar.axelarnet.v1beta1.RegisterAssetResponse"
	// RouteIBCTransfersRequest        = "axelar.axelarnet.v1beta1.RouteIBCTransfersRequest"
	// RouteIBCTransfersResponse       = "axelar.axelarnet.v1beta1.RouteIBCTransfersResponse"
	// RegisterFeeCollectorRequest     = "axelar.axelarnet.v1beta1.RegisterFeeCollectorRequest"
	// RegisterFeeCollectorResponse    = "axelar.axelarnet.v1beta1.RegisterFeeCollectorResponse"
	// RetryIBCTransferRequest         = "axelar.axelarnet.v1beta1.RetryIBCTransferRequest"
	// RetryIBCTransferResponse        = "axelar.axelarnet.v1beta1.RetryIBCTransferResponse"
	// RouteMessageRequest             = "axelar.axelarnet.v1beta1.RouteMessageRequest"
	// RouteMessageResponse            = "axelar.axelarnet.v1beta1.RouteMessageResponse"
	// CallContractRequest             = "axelar.axelarnet.v1beta1.CallContractRequest"
	// CallContractResponse            = "axelar.axelarnet.v1beta1.CallContractResponse"

	// SetGatewayRequest                  = "axelar.evm.v1beta1.SetGatewayRequest"
	// SetGatewayResponse                 = "axelar.evm.v1beta1.SetGatewayResponse"
	// ConfirmGatewayTxRequest            = "axelar.evm.v1beta1.ConfirmGatewayTxRequest"
	// ConfirmGatewayTxResponse           = "axelar.evm.v1beta1.ConfirmGatewayTxResponse"
	// ConfirmGatewayTxsRequest           = "axelar.evm.v1beta1.ConfirmGatewayTxsRequest"
	// ConfirmGatewayTxsResponse          = "axelar.evm.v1beta1.ConfirmGatewayTxsResponse"
	// ConfirmDepositRequest              = "axelar.evm.v1beta1.ConfirmDepositRequest"
	// ConfirmDepositResponse             = "axelar.evm.v1beta1.ConfirmDepositResponse"
	// ConfirmTokenRequest                = "axelar.evm.v1beta1.ConfirmTokenRequest"
	// ConfirmTokenResponse               = "axelar.evm.v1beta1.ConfirmTokenResponse"
	// ConfirmTransferKeyRequest          = "axelar.evm.v1beta1.ConfirmTransferKeyRequest"
	// ConfirmTransferKeyResponse         = "axelar.evm.v1beta1.ConfirmTransferKeyResponse"
	// LinkRequest                        = "axelar.evm.v1beta1.LinkRequest"
	// LinkResponse                       = "axelar.evm.v1beta1.LinkResponse"
	// CreateBurnTokensRequest            = "axelar.evm.v1beta1.CreateBurnTokensRequest"
	// CreateBurnTokensResponse           = "axelar.evm.v1beta1.CreateBurnTokensResponse"
	// CreateDeployTokenRequest           = "axelar.evm.v1beta1.CreateDeployTokenRequest"
	// CreateDeployTokenResponse          = "axelar.evm.v1beta1.CreateDeployTokenResponse"
	// CreatePendingTransfersRequest      = "axelar.evm.v1beta1.CreatePendingTransfersRequest"
	// CreatePendingTransfersResponse     = "axelar.evm.v1beta1.CreatePendingTransfersResponse"
	// CreateTransferOwnershipRequest     = "axelar.evm.v1beta1.CreateTransferOwnershipRequest"
	// CreateTransferOwnershipResponse    = "axelar.evm.v1beta1.CreateTransferOwnershipResponse"
	// CreateTransferOperatorshipRequest  = "axelar.evm.v1beta1.CreateTransferOperatorshipRequest"
	// CreateTransferOperatorshipResponse = "axelar.evm.v1beta1.CreateTransferOperatorshipResponse"
	// SignCommandsRequest                = "axelar.evm.v1beta1.SignCommandsRequest"
	// SignCommandsResponse               = "axelar.evm.v1beta1.SignCommandsResponse"
	// AddChainRequest                    = "axelar.evm.v1beta1.AddChainRequest"
	// AddChainResponse                   = "axelar.evm.v1beta1.AddChainResponse"
	// RetryFailedEventRequest            = "axelar.evm.v1beta1.RetryFailedEventRequest"
	// RetryFailedEventResponse           = "axelar.evm.v1beta1.RetryFailedEventResponse"

	// StartKeygenRequest      = "axelar.multisig.v1beta1.StartKeygenRequest"
	// StartKeygenResponse     = "axelar.multisig.v1beta1.StartKeygenResponse"
	// SubmitPubKeyRequest     = "axelar.multisig.v1beta1.SubmitPubKeyRequest"
	// SubmitPubKeyResponse    = "axelar.multisig.v1beta1.SubmitPubKeyResponse"
	// SubmitSignatureRequest  = "axelar.multisig.v1beta1.SubmitSignatureRequest"
	// SubmitSignatureResponse = "axelar.multisig.v1beta1.SubmitSignatureResponse"
	// RotateKeyRequest        = "axelar.multisig.v1beta1.RotateKeyRequest"
	// RotateKeyResponse       = "axelar.multisig.v1beta1.RotateKeyResponse"
	// KeygenOptOutRequest     = "axelar.multisig.v1beta1.KeygenOptOutRequest"
	// KeygenOptOutResponse    = "axelar.multisig.v1beta1.KeygenOptOutResponse"
	// KeygenOptInRequest      = "axelar.multisig.v1beta1.KeygenOptInRequest"
	// KeygenOptInResponse     = "axelar.multisig.v1beta1.KeygenOptInResponse"

	RegisterChainMaintainerRequest    = "axelar.nexus.v1beta1.RegisterChainMaintainerRequest"
	RegisterChainMaintainerResponse   = "axelar.nexus.v1beta1.RegisterChainMaintainerResponse"
	DeregisterChainMaintainerRequest  = "axelar.nexus.v1beta1.DeregisterChainMaintainerRequest"
	DeregisterChainMaintainerResponse = "axelar.nexus.v1beta1.DeregisterChainMaintainerResponse"
	ActivateChainRequest              = "axelar.nexus.v1beta1.ActivateChainRequest"
	ActivateChainResponse             = "axelar.nexus.v1beta1.ActivateChainResponse"
	DeactivateChainRequest            = "axelar.nexus.v1beta1.DeactivateChainRequest"
	DeactivateChainResponse           = "axelar.nexus.v1beta1.DeactivateChainResponse"
	RegisterAssetFeeRequest           = "axelar.nexus.v1beta1.RegisterAssetFeeRequest"
	RegisterAssetFeeResponse          = "axelar.nexus.v1beta1.RegisterAssetFeeResponse"
	SetTransferRateLimitRequest       = "axelar.nexus.v1beta1.SetTransferRateLimitRequest"
	SetTransferRateLimitResponse      = "axelar.nexus.v1beta1.SetTransferRateLimitResponse"
)

// GetMessageTypeFromURL returns the appropriate message type based on the TypeURL
func GetMessageTypeFromURL(typeURL string) (any, bool) {
	switch typeURL {
	// case BatchRequestTypeURL:
	// 	return &auxiliaryTypes.BatchRequest{}, true
	// case BatchResponseTypeURL:
	// 	return &auxiliaryTypes.BatchResponse{}, true
	// case BatchResponseResponseTypeURL:
	// 	return &auxiliaryTypes.BatchResponse_Response{}, true
	case RegisterChainMaintainerRequest:
		return &nexusTypes.RegisterChainMaintainerRequest{}, true
	default:
		return nil, false
	}
}
