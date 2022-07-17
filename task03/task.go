package main

import (
	"fmt"
	"sync"
)

const size = 5

type sum struct {
	sqSum      int
	sync.Mutex //используется для блокировки одновременного доступа горутинами к структуре
}

func squareSum(arr [size]int) int {
	wg := sync.WaitGroup{}
	wg.Add(size)
	s := sum{0, sync.Mutex{}}
	for _, n := range arr {
		n := n //избегаем захвата переменной
		go func() {
			sq := n * n //конкурентно вычисляем квадраты
			s.Lock()    //закрываем доступ к общей структуре для остальных горутин
			s.sqSum += sq
			fmt.Println(s.sqSum)
			s.Unlock() //открываем доступ
			wg.Done()
		}()
	}
	wg.Wait()
	return s.sqSum
}

func main() {
	arr := [size]int{2, 4, 6, 8, 10}
	fmt.Printf("Результат: %v", squareSum(arr))
}
