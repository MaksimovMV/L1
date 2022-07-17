package main

import (
	"fmt"
	"strconv"
	"sync"
)

type someStruct struct {
	store      map[int]string
	sync.Mutex //используется для блокировки одновременного доступа горутинами к мапе
}

func main() {
	n := 10
	s := someStruct{make(map[int]string), sync.Mutex{}}
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		i := i //избегаем захвата переменной
		go func() {
			s.Lock()
			s.store[i] = "data" + strconv.Itoa(i) //запись значений в мапу
			s.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	for i, data := range s.store {
		fmt.Printf("key: %v, data: %v\n", i, data) //выводим значения из мапы
	}
}
