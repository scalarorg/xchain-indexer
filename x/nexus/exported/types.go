package exported

import (
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utils "github.com/scalarorg/xchains-indexer/util"
)

// ChainNameLengthMax bounds the max chain name length
const ChainNameLengthMax = 20

// ChainName ensures a correctly formatted EVM chain name
type ChainName string

// Validate returns an error, if the chain name is empty or too long
func (c ChainName) Validate() error {
	if err := utils.ValidateString(string(c)); err != nil {
		return sdkerrors.Wrap(err, "invalid chain name")
	}

	if len(c) > ChainNameLengthMax {
		return fmt.Errorf("chain name length %d is greater than %d", len(c), ChainNameLengthMax)
	}

	return nil
}

func (c ChainName) String() string {
	return string(c)
}

// Equals returns boolean for whether two chain names are case-insensitive equal
func (c ChainName) Equals(c2 ChainName) bool {
	return strings.EqualFold(c.String(), c2.String())
}
