package main

import (
	"context"
	"fmt"
	"time"
)

func publisher(ctx context.Context) chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done(): //завершаем работу после заданного времени
				fmt.Println("Time's up")
				close(c)
				return
			case c <- i:
				i++
				time.Sleep(time.Second)
			}
		}
	}()
	return c
}

func main() {
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	ch := publisher(timeout)

	//читаем из канала
	for i := range ch {
		fmt.Println(i)
	}

}
