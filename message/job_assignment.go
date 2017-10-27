package message

import "github.com/242617/flaneurtv/process"

// Incoming!
func NewJobAssignment(proc *process.Process, namespace string, payload *Payload) (msg *Message) {
	msg = NewMessage(namespace+"/job_assignment", proc)
	msg.Payload = *payload
	return msg
}
