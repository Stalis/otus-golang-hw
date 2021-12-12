package hw05parallelexecution

import (
	"errors"
	"log"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(wg *sync.WaitGroup, num int, taskChan <-chan Task, errChan chan<- error, cancelToken *int32) {
	defer wg.Done()
	log.Printf("Worker %v started\n", num)

	for t := range taskChan {
		if err := t(); err != nil {
			log.Printf("Worker %v get error: %v", num, err)
			errChan <- err
		} else {
			log.Printf("Worker %v complete task", num)
		}

		token := atomic.LoadInt32(cancelToken)
		if token != 0 {
			break
		}
	}

	log.Printf("Worker %v done", num)
}

func errorScavenger(maxCount int32, errCh <-chan error, count *int32, token *int32) {
	for range errCh {
		cur := atomic.AddInt32(count, 1)
		if cur >= maxCount {
			atomic.StoreInt32(token, 1)
			return // My job here is done XD
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup

	taskChan := make(chan Task, len(tasks))
	errChan := make(chan error)
	var errCount int32
	var cancelToken int32

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(&wg, i, taskChan, errChan, &cancelToken)
	}

	go errorScavenger(int32(m), errChan, &errCount, &cancelToken)

	for _, t := range tasks {
		taskChan <- t
	}
	close(taskChan)

	wg.Wait()
	close(errChan)

	if cancelToken != 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
