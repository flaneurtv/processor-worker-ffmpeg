package message

import "github.com/242617/flaneurtv/process"

func NewJobFinished(proc *process.Process) *Message {
	return NewMessage(proc.Namespace+"/job_finished", proc)
}
