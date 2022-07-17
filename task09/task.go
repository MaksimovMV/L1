package main

import (
	"fmt"
)

// функция перебора чисел из массива и записи в возвращаемый канал
func write(arr []int) chan int {
	out := make(chan int)
	go func() {
		for _, i := range arr {
			out <- i
		}
		close(out)
	}()
	return out
}

// функция умножения полученных в канале с чисел и записи в возвращаемый канал
func multiply(c chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range c {
			out <- i * 2
		}
		close(out)
	}()

	return out
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	c1 := write(arr)   //канал для записи чисел из массива
	c2 := multiply(c1) //канал для записи умноженных чисел

	//читаем с с2 и выводим stdout
	for i := range c2 {
		fmt.Println(i)
	}

}
