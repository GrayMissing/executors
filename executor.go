package executors

type Executor struct {
	jobs []*Job
}

func NewExecutor() *Executor {
	return new(Executor)
}

func (e *Executor) New() *Job {
	j := NewJob()
	e.jobs = append(e.jobs, j)
	return j
}

func (e *Executor) Join() {
	for _, job := range e.jobs {
		job.Join()
	}
}
