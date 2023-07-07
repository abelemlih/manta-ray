package messenger

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	Content     string
	SentAt      time.Time
	DeliveredAt time.Time
}

func (m *Message) Signature() string {
	return fmt.Sprintf("Message ID#: %s; Sent At: %s; Delivered At: %s", m.Id, m.SentAt, m.DeliveredAt)
}
