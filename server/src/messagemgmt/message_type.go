package messagemgmt

type MessageType uint32

const (
	DirectMessage MessageType = iota
	GroupMessage
)
