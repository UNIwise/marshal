package marshal

import (
	"github.com/ThreeDotsLabs/watermill-nats/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/proto"
)

// Force implementation check.
var _ nats.MarshalerUnmarshaler = &Marshaler{}

type Marshaler struct{}

func NewMarshalerUnmarshaler() *Marshaler {
	return &Marshaler{}
}

// Marshal implements nats.MarshalerUnmarshaler
func (m *Marshaler) Marshal(topic string, msg *message.Message) ([]byte, error) {
	return msg.Payload, nil
}

// Unmarshal implements nats.MarshalerUnmarshaler
func (m *Marshaler) Unmarshal(stanMsg *stan.Msg) (*message.Message, error) {
	msg := message.NewMessage(uuid.NewString(), stanMsg.Data)
	return msg, nil
}
