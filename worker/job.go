package worker

type Job struct {
	UUID          string
	ReferenceUUID string
	QueueName     string
	Command       string
	Arguments     []string
	CreatedAt     string
	StartedAt     string
}
