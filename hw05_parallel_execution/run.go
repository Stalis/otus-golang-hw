package hw05parallelexecution

import (
	"errors"
	"log"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(taskChan <-chan Task, wg *sync.WaitGroup, errCount *int32, num int) {
	defer wg.Done()
	log.Printf("Worker %v started\n", num)

	for t := range taskChan {
		if err := t(); err != nil {
			log.Printf("Worker %v get error: %v", num, err)
			atomic.AddInt32(errCount, 1)
		} else {
			log.Printf("Worker %v complete task", num)
		}
	}

	log.Printf("Worker %v done", num)
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup

	taskChan := make(chan Task)

	var errCount int32

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(taskChan, &wg, &errCount, i)
	}

	for _, t := range tasks {
		if m > 0 && atomic.LoadInt32(&errCount) >= int32(m) {
			break
		}
		taskChan <- t
	}
	close(taskChan)

	wg.Wait()

	if m > 0 && atomic.LoadInt32(&errCount) >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
