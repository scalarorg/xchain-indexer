package exported

import (
	fmt "fmt"
	"strconv"
)

// PollID represents ID of polls
type PollID uint64

// String converts the given poll ID to string
func (id PollID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// Deprecated: String converts the given poll key to string
func (m PollKey) String() string {
	return fmt.Sprintf("%s_%s", m.Module, m.ID)
}
