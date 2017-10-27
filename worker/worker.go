package worker

var CurrentJob *Job

func IsBusy() bool {
	return CurrentJob != nil
}

func start(job *Job) (err error) {
	if IsBusy() {
		err = ErrWorkerBusy
		return
	}
	CurrentJob = job
	return
}

func stop() {
	CurrentJob = nil
	seconds, duration = 0, 0
}
