package executors

import (
	"testing"
	"time"
)

func TestExecutor(t *testing.T) {
	executor := NewExecutor(5)
	for i := 0; i < 10; i++ {
		j := executor.New()
		go func(j *Job) {
			time.Sleep(1 * time.Millisecond)
			j.SetResult(i)
			j.Finish()
		}(j)
		j.Join()
		t.Logf("hello, %d", j.GetResult()[0])
	}
	executor.Join()
}
