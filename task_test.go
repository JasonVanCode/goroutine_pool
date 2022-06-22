package goroutine_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	task := NewTask(func() {
		fmt.Println(time.Now())
	})
	pool := NewPool(3)
	go func() {
		for {
			time.Sleep(time.Second)
			pool.ReceiveTask(task)
		}
	}()
	pool.Run()
}
