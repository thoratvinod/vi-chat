package message

type Message struct {
	MsgType    MessageType `json:"msgType"`
	SenderID   uint        `json:"senderID"`
	ReceiverID uint        `json:"receiverID"` // this will be group id in case of group message
	MemberIDs  []uint      `json:"-"`          // members of group in case of group message
	Timestamp  int64       `json:"timeStamp"`
	Content    string      `json:"content"`
}

type MessageStatus uint32

const (
	MsgStatusNone MessageStatus = iota
	MsgStatusCreated
	MsgStatusSent
	MsgStatusSeen
)
