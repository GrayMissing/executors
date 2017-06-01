package executors

import (
	"testing"
	"time"
)

func TestExecutor(t *testing.T) {
	executor := NewExecutor(5)
	for i := 0; i < 10; i++ {
		go func(j *Job) {
			time.Sleep(500 * time.Millisecond)
			t.Log("hello")
			j.Finish()
		}(executor.New())
	}
	executor.Join()
}
