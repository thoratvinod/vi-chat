package message

type Message struct {
	MsgType    MessageType `json:"msgType"`
	SenderID   uint        `json:"senderID"`
	ReceiverID uint        `json:"receiverID"`
	Timestamp  int64       `json:"timeStamp"`
	Content    string      `json:"content"`
}
