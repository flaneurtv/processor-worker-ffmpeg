package message

import (
	"github.com/242617/flaneurtv/process"
)

// Incoming!
func NewTick(recipient *process.Process, namespace string, payload *Payload) (msg *Message) {
	msg = NewMessage(namespace+"/tick", recipient)
	msg.Payload = *payload
	return msg
}
