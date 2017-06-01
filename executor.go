package executors

type Executor struct {
	jobs      []*Job
	runningCh chan bool
}

func NewExecutor(n int) *Executor {
	e := new(Executor)
	e.runningCh = make(chan bool, n)
	return e
}

func (e *Executor) New() *Job {
	e.runningCh <- true
	j := NewJob(e)
	e.jobs = append(e.jobs, j)
	return j
}

func (e *Executor) Join() {
	for _, job := range e.jobs {
		job.Join()
	}
}
