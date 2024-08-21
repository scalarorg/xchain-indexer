package common

type MsgType int

const (
	MsgRecvPacket MsgType = iota
	MsgAcknowledgement
)

type TxMessage struct {
	ID            uint `gorm:"primaryKey"`
	TxID          uint `gorm:"uniqueIndex:txMessageIndex,priority:1"`
	MessageID     uint `gorm:"uniqueIndex:txMessageIndex,priority:2"`
	BlockId       uint
	MessageDetail string
}
