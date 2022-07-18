package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var implement = make(chan int)
var read = make(chan int)

// вариант с Mutex
type mutexCounter struct {
	count      int
	sync.Mutex //используется для блокировки одновременного доступа горутинами к структуре
}

// вариант с atomic
type atomicCounter struct {
	count int64
}

func (c *atomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.count)
}

// вариант с управляющей горутиной
func readCount() int {
	return <-read
}

func goroutineCounter() {
	var count int
	for {
		select {
		case <-implement:
			count++
		case read <- count:
		}
	}
}

func main() {
	mutCounter := mutexCounter{0, sync.Mutex{}}
	wg := sync.WaitGroup{}

	atomCounter := atomicCounter{}

	go goroutineCounter()

	for i := 0; i < 100; i++ {
		wg.Add(3)

		// инкрементируем счетчик 1 с Mutex
		go func() {
			mutCounter.Lock()
			defer mutCounter.Unlock()
			defer wg.Done()
			mutCounter.count++
		}()

		// инкрементируем счетчик 2 с atomic
		go func() {
			defer wg.Done()
			atomic.AddInt64(&atomCounter.count, 1)
		}()

		// инкрементируем счетчик 3 при помощи управляющей горутины
		go func() {
			defer wg.Done()
			implement <- 1
		}()
	}
	wg.Wait()
	fmt.Printf("Счетчик управляемый Mutex: %v\n", mutCounter.count)
	fmt.Printf("Счетчик управляемый Atomic: %v\n", atomCounter.Value())
	fmt.Printf("Счетчик управляемый управляющей горутиной: %v\n", readCount())
}
