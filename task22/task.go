package main

import (
	"fmt"
	"math/big"
)

func sum(a, b *big.Int) *big.Int {
	result := new(big.Int)
	return result.Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	result := new(big.Int)
	return result.Sub(a, b)
}

func div(a, b *big.Int) *big.Int {
	result := new(big.Int)
	return result.Div(a, b)
}

func mul(a, b *big.Int) *big.Int {
	result := new(big.Int)
	return result.Mul(a, b)
}

func main() {
	num1 := "40000000000000000000"
	num2 := "20000000000000000000"

	a := new(big.Int)
	a.SetString(num1, 10) //получаем значение big.Int из строки
	b := new(big.Int)
	b.SetString(num2, 10) //получаем значение big.Int из строки

	// выводим необходимые результаты напрямую, либо через функции возвращающие нужные значения
	z := new(big.Int)
	fmt.Println(z.Add(a, b))
	fmt.Println(sum(a, b))
	fmt.Println(z.Sub(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(z.Div(a, b))
	fmt.Println(div(a, b))
	fmt.Println(z.Mul(a, b))
	fmt.Println(mul(a, b))
}
