package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func someFunc(ctx context.Context, wg *sync.WaitGroup) chan int {
	quit := make(chan int)
	wg.Add(1)

	go func() {
		for {
			select {
			// вариант с использованием time.After
			case <-time.After(time.Second * 3):
				fmt.Println("Goroutine stopped with time.After")
				wg.Done()
				return
			// вариант с использованием context
			case <-ctx.Done():
				fmt.Println("Goroutine stopped with context")
				wg.Done()
				return
			// вариант с закрытием канала извне
			case <-quit:
				fmt.Println("Goroutine stopped with cancel quit chan")
				wg.Done()
				return
			}
		}
	}()
	return quit
}

func main() {
	// первый вариант
	wg := sync.WaitGroup{}
	ctx := context.Background()
	ch := someFunc(ctx, &wg)
	time.Sleep(time.Second * 2)
	close(ch)
	wg.Wait()

	// второй вариант
	timeout, cancel := context.WithTimeout(ctx, time.Second*2)
	ch2 := someFunc(timeout, &wg)
	defer cancel()
	wg.Wait()
	close(ch2)

	// третий вариант
	ch3 := someFunc(ctx, &wg)
	wg.Wait()
	close(ch3)

}
