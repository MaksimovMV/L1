package main

import "fmt"

func makeSubset(arr []string) (result []string) {
	result = append(result, arr[0]) //сразу записываем первое значение исходного множества в подмножество

	//сравниваем значения исходного множества со элементами подмножества, при нахождении совпадений
	//заканчиваем итерирование и переходим к следующему значению, при отсутствии совпадений
	//добавляем элемент в подмножество
loop:
	for _, s := range arr {
		for _, s2 := range result {
			if s == s2 {
				continue loop
			}
		}
		result = append(result, s)

	}
	return
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(makeSubset(arr))
}
