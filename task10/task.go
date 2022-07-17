package main

import "fmt"

func groupTemperatures(arr []float64) map[int][]float64 {
	m := map[int][]float64{}
	for _, t := range arr {
		group := int(t) / 10 * 10      //определяем группу для записи конкретной температуры
		m[group] = append(m[group], t) //записываем в мапе в соответствующий слайс по ключу группы
	}
	return m
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := groupTemperatures(temperatures)
	fmt.Println(m) //выводим данные из результирующей мапы
}
