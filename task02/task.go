package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	arr := [5]int{2, 4, 6, 8, 10}
	for _, n := range arr {
		wg.Add(1) //увеличение счетчика wg на указанное число (1)
		go func(n int) { //запуск горутин
			fmt.Printf("Квадрат числа %v = %v\n", n, n*n)
			wg.Done() //снижение счетчика wg на 1
		}(n)
	}
	wg.Wait() //ждет выполнения всех горутин (wg.Done)
}
