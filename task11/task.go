package main

import "fmt"

func intersection(arr1 []int, arr2 []int) (result []int) {
	// перебираем значения первого среза и сравниваем со значениями второго
	// при совпадении значений записываем в результирующий срез
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				result = append(result, arr1[i])
			}
		}
	}
	return
}

func main() {
	arr1 := []int{1, 5, 7, 4, 135, 9, 12}
	arr2 := []int{120, 4, 5, 2, 135, 6, 9}
	fmt.Println(intersection(arr1, arr2))
}
