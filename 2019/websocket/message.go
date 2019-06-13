package ws

type Message interface {
	Identify() int
	Body() []byte
}

type BaseMessage struct {
	identify int
	body     []byte
}

func (b *BaseMessage) New(identify int, body []byte) *BaseMessage {
	return &BaseMessage{
		identify: identify,
		body:     body,
	}
}

func (b *BaseMessage) Identify() int {
	return b.identify
}

func (b *BaseMessage) Body() []byte {
	return b.body
}
