package message_broadcast_service

type BinaryMessage struct {
	Targets []int
	Lang    string
	Data    []byte
}
