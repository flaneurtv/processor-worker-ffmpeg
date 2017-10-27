package message

type Payload struct {
	UUID          string    `json:"uuid,omitempty"`
	ReferenceUUID string    `json:"reference_uuid,omitempty"`
	QueueName     string    `json:"queue_name,omitempty"`
	Command       string    `json:"command,omitempty"`
	Arguments     []string  `json:"args,omitempty"`
	CreatedAt     string    `json:"created_at,omitempty"`
	StartedAt     string    `json:"started_at,omitempty"`
	Progress      *Progress `json:"progress,omitempty"`
}
