package util

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
)

func ToNumeric(i *big.Int) decimal.Decimal {
	num := decimal.NewFromBigInt(i, 0)
	return num
}

// StrNotSet will return true if the string value provided is empty
func StrNotSet(value string) bool {
	return len(value) == 0
}

func RemoveDuplicatesFromUint64Slice(sliceList []uint64) []uint64 {
	allKeys := make(map[uint64]bool)
	list := []uint64{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// Encode a bytearray as hex string with 0x prefix.
func HexEncode(input []byte) string {
	return "0x" + hex.EncodeToString(input)
}

// Decode a hex string. Hex string can be optionally prefixed with 0x.
func HexDecode(input string) ([]byte, error) {
	return hex.DecodeString(strings.TrimPrefix(input, "0x"))
}
