package main

import "fmt"

func setBit(n int64, pos uint) int64 {
	n |= 1 << pos //побитово сдвигаем 1 на позицию pos и при помощи OR записываем в n
	return n
}

func clearBit(n int64, pos uint) int64 {
	n &^= 1 << pos //сдвигаем 1 на поз. pos, инвертируем полученное значение и при помощи & обнуляем нужный бит
	return n
}

func main() {
	var n int64
	var pos uint
	n, pos = 16, 1
	fmt.Println("Set bit")
	fmt.Printf("Number %v\nNumber in binary %b\nPosition %v\n", n, n, pos)
	n = setBit(n, pos)
	fmt.Printf("After shift\nNumber %v\nNumber in binary %b\n", n, n)

	fmt.Println("Clear bit")
	n, pos = 15, 2
	fmt.Printf("Number %v\nNumber in binary %b\nPosition %v\n", n, n, pos)
	n = clearBit(n, pos)
	fmt.Printf("After shift\nNumber %v\nNumber in binary %b\n", n, n)
}
