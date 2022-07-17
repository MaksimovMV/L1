package main

import "fmt"

func main() {
	a, b := 1, 2

	// первый и наилучший метод
	a, b = b, a
	fmt.Println(a, b)

	//второй метод (можно использовать * и /)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

	//третий метод с XOR
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)
}
