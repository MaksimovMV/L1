package main

import "fmt"

func reverse(s string) string {
	result := []rune(s) //приводим к массиву рун

	//итерируемся до середины (не включительно) массива и меняем зеркально значения местами
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return string(result) //возвращаем строку
}

func main() {
	s := "главрыба"
	fmt.Println(reverse(s))
}
