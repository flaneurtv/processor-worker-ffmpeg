package message

import (
	"errors"
	"strings"

	"github.com/242617/flaneurtv/process"
	"github.com/242617/flaneurtv/utils"
)

var ErrUnknownType = errors.New("unknown message type")

const (
	TypeIdle = iota
	TypeJobAccepted
	TypeJobAssignment
	TypeJobFinished
	TypeTick
	TypeWorkerBusy
	TypeWorkerIdle
)

func NewMessage(topic string, proc *process.Process) *Message {
	var msg Message
	msg.Topic = topic
	msg.ServiceUUID = proc.ServiceUUID
	msg.ServiceName = proc.ServiceName
	msg.ServiceHost = proc.ServiceHost
	msg.CreatedAt = utils.Now()
	return &msg
}

type Message struct {
	Topic         string         `json:"topic"`
	ServiceUUID   string         `json:"service_uuid"`
	ServiceName   string         `json:"service_name"`
	ServiceHost   string         `json:"service_host"`
	CreatedAt     string         `json:"created_at"`
	TickReference *TickReference `json:"tick_reference,omitempty"`
	Payload       Payload        `json:"payload"`
}

func (m *Message) AddressedTo(p *process.Process) bool {
	return m.ServiceUUID == p.ServiceUUID
}

func (m *Message) Type() (t int, err error) {
	switch {
	case strings.Contains(m.Topic, "/job_accepted"):
		t = TypeJobAccepted
	case strings.Contains(m.Topic, "/job_assignment"):
		t = TypeJobAssignment
	case strings.Contains(m.Topic, "/job_finished"):
		t = TypeJobFinished
	case strings.Contains(m.Topic, "/tick"):
		t = TypeTick
	case strings.Contains(m.Topic, "/worker_busy"):
		t = TypeWorkerBusy
	case strings.Contains(m.Topic, "/worker_idle"):
		t = TypeWorkerIdle
	default:
		err = ErrUnknownType
	}
	return
}
