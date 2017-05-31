# executors
a simple goroutine utility

usage:
```go
executor := NewExecutor()
go func(j *Job) {
  time.Sleep(1 * time.Second)
	t.Log("hello")
	j.Finish()
}(executor.New())
executor.Join()
```
