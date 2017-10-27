package process

import "os"

func NewProcess() *Process {
	var proc Process
	proc.Namespace = os.Getenv("NAMESPACE")
	proc.ServiceUUID = os.Getenv("SERVICE_UUID")
	proc.ServiceName = os.Getenv("SERVICE_NAME")
	proc.ServiceHost = os.Getenv("SERVICE_HOST")
	return &proc
}

type Process struct {
	Namespace   string
	ServiceUUID string
	ServiceName string
	ServiceHost string
}

func (p *Process) Environ() []string {
	return []string{
		"NAMESPACE=" + p.Namespace,
		"SERVICE_UUID=" + p.ServiceUUID,
		"SERVICE_NAME=" + p.ServiceName,
		"SERVICE_HOST=" + p.ServiceHost,
	}
}
