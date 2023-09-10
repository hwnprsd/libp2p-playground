package squad

type UpdateMessage interface {
	GetWireMessage() []byte
	GetIsBroadcast() bool
	GetPayload() []byte
}
