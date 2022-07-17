package main

import (
	"fmt"
	"math/rand"
)

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Выбираем индекс опорного элемента в пределах длины массива
	pivotIndex := rand.Int() % len(a)

	// Ставим опорный элемент в правое положение
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Передвигаем все элементы меньшие чем опорный влево, увеличивая каждый раз индекс left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Ставим опорный элемент после последнего меньшего чем опорный элемента
	// т.о. элементы, которые меньше опорной точки оказались слева от нее,
	//а которые больше - справа
	a[left], a[right] = a[right], a[left]

	// повторяем для срезов до и после опорной точки пока
	// сортируемая часть среза не будет пустой или состоять из одного элемента
	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}

func main() {
	arr := []int{14, 12, 1, 156, 50, 5, 256, 0, -7, -22, 12}

	//sort.Ints(arr) //пакет сорт использует метод сортировки quicksort

	arr = quickSort(arr)
	fmt.Println(arr)
}
