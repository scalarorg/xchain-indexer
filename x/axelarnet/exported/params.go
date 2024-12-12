package exported

import (
	"github.com/scalarorg/xchains-indexer/x/nexus/exported"
	tss "github.com/scalarorg/xchains-indexer/x/tss/exported"
)

const (
	// ModuleName exposes axelarnet module name
	ModuleName = "axelarnet"
)

var (
	// NativeAsset is the native asset on Axelarnet
	NativeAsset = "uaxl"

	// Axelarnet defines properties of the Axelar chain
	Axelarnet = exported.Chain{
		Name:                  "Axelarnet",
		SupportsForeignAssets: true,
		KeyType:               tss.None,
		Module:                ModuleName,
	}
)
