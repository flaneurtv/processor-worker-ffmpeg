package message

import (
	"github.com/242617/flaneurtv/process"
	"github.com/242617/flaneurtv/utils"
	"github.com/242617/flaneurtv/worker"
)

func NewJobAccepted(proc *process.Process) *Message {
	msg := NewMessage(proc.Namespace+"/job_accepted", proc)
	msg.Payload.UUID = worker.CurrentJob.UUID
	msg.Payload.ReferenceUUID = worker.CurrentJob.ReferenceUUID
	msg.Payload.QueueName = worker.CurrentJob.QueueName
	msg.Payload.Command = worker.CurrentJob.Command
	msg.Payload.Arguments = worker.CurrentJob.Arguments
	msg.Payload.CreatedAt = worker.CurrentJob.CreatedAt
	msg.Payload.StartedAt = utils.Now()
	return msg
}
