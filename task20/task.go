package main

import (
	"fmt"
	"strings"
)

func reverseByWords(s string) string {
	arr := strings.Split(s, " ") //разбиваем строку по разделителю " " на массив строк
	// зеркально меняем значения местами
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return strings.Join(arr, " ")

}

func main() {
	s := "snow dog sun"
	fmt.Println(reverseByWords(s))
}
