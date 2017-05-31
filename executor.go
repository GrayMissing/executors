package executors


type Executor struct {
  jobs []*Job
}

type Job struct {
  status bool
  statusCh chan bool
  result interface{}
}

func NewJob() *Job {
  j := new(Job)
  j.status = false
  j.statusCh = make(chan bool)
  return j
}

func (j *Job) Finish() {
  j.status = true
  j.statusCh <- true
}

func (j *Job) Finished() bool {
  return j.status
}

func (j *Job) Join() {
  <- j.statusCh
}

func (j *Job) SetResult(r interface{}) {
  j.result = r
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
