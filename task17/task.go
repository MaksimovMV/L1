package main

import (
	"fmt"
	"sort"
)

const size int = 5

func leftBinarySearch(array [size]int, val int) int {
	left, right := 0, size-1
	for left <= right {
		middle := left + (right-left)/2
		if array[middle] < val {
			left = middle + 1
		} else if array[middle] >= val {
			right = middle - 1
		}
	}
	return left
}

func main() {
	arr := [size]int{14, 12, 1, 156, 50}
	slice := arr[:]
	sort.Ints(slice)
	fmt.Println(arr)
	num := 50
	i := sort.SearchInts(slice, num) //пакет сорт для поиска использует бинарный поиск
	index := leftBinarySearch(arr, num)
	fmt.Printf("Инедкс при использовании встроенного поиска %v\n", i)
	fmt.Printf("Инедкс при использовании собственной реализации поиска %v\n", index)
}
