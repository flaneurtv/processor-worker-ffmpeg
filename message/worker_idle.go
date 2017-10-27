package message

import "github.com/242617/flaneurtv/process"

func NewWorkerIdle(proc *process.Process, request *Message) *Message {
	msg := NewMessage(proc.Namespace+"/worker_idle", proc)
	msg.TickReference = NewTickReference()
	msg.TickReference.UUID = request.Payload.UUID
	msg.TickReference.CreatedAt = request.CreatedAt
	return msg
}
