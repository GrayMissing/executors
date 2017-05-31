package executors

import (
	"testing"
	"time"
)

func TestExecutor(t *testing.T) {
	executor := NewExecutor()
	for i := 0; i < 10; i++ {
		go func(j *Job) {
			time.Sleep(1 * time.Second)
			t.Log("hello")
			j.Finish()
		}(executor.New())
	}
	executor.Join()
}
