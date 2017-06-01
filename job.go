package executors

type Job struct {
	status   bool
	statusCh chan bool
	result   []interface{}
	executor *Executor
}

func NewJob(e *Executor) *Job {
	j := new(Job)
	j.status = false
	j.statusCh = make(chan bool, 1)
	j.executor = e
	return j
}

func (j *Job) Finish() {
	if !j.status {
		j.status = true
		j.statusCh <- true
		<-j.executor.runningCh
	}
}

func (j *Job) Finished() bool {
	return j.status
}

func (j *Job) Join() {
	if !j.status {
		<-j.statusCh
	}
}

func (j *Job) SetResult(r ...interface{}) {
	j.result = r
}

func (j *Job) GetResult() []interface{} {
	return j.result
}
