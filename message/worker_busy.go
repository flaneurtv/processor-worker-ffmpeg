package message

import (
	"github.com/242617/flaneurtv/process"
	"github.com/242617/flaneurtv/worker"
)

func NewWorkerBusy(proc *process.Process) *Message {
	msg := NewMessage(proc.Namespace+"/worker_busy", proc)
	msg.Payload.UUID = worker.CurrentJob.UUID
	msg.Payload.ReferenceUUID = worker.CurrentJob.ReferenceUUID
	msg.Payload.QueueName = worker.CurrentJob.QueueName
	msg.Payload.Command = worker.CurrentJob.Command
	msg.Payload.Arguments = worker.CurrentJob.Arguments
	msg.Payload.CreatedAt = worker.CurrentJob.CreatedAt
	msg.Payload.StartedAt = worker.CurrentJob.StartedAt
	msg.Payload.Progress = NewProgress()
	msg.Payload.Progress.Progress = worker.Progress()
	return msg
}
