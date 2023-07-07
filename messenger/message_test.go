package messenger

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestSignature(t *testing.T) {
	timestamp := time.Now()
	message := Message{Id: uuid.New(), Content: "This is a test", SentAt: timestamp, DeliveredAt: timestamp}
	expected := fmt.Sprintf("Message ID#: %s; Sent At: %s; Delivered At: %s", message.Id, timestamp, timestamp)
	if message.Signature() != expected {
		t.Errorf("Expected signature to equal %q, got %q", expected, message.Signature())
	}
}
