package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count      int
	sync.Mutex //используется для блокировки одновременного доступа горутинами к структуре
}

func main() {
	c := counter{0, sync.Mutex{}}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			c.Lock()
			c.count++
			c.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.count)
}
