package main

import "fmt"

func removePos(a []int, pos uint) []int {
	//добавляем к срезу значений до заданной позиции срез значений после позиции, т.о. исключая её из результата
	return append(a[:pos], a[pos+1:]...)
}

func main() {
	a := []int{10, 20, 32, 4, 52, 23, 642}
	fmt.Println(removePos(a, 2))
}
