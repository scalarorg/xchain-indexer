package exported

import (
	"github.com/scalarorg/xchains-indexer/x/nexus/exported"
	tss "github.com/scalarorg/xchains-indexer/x/tss/exported"
)

var (
	// Ethereum defines properties of the Ethereum chain
	Ethereum = exported.Chain{
		Name:                  "Ethereum",
		SupportsForeignAssets: true,
		KeyType:               tss.Multisig,
		Module:                "evm", // cannot use constant due to import cycle
	}
)
