package main

import (
	"fmt"
	"strings"
)

func uniqueSymbols(s string) bool {
	s = strings.ToLower(s) //приводим всё к нижнему регистру

	//strings.Contains возвращает true при наличии искомой подстроки, при отсутствии - false
	//ищем каждую руну после её позиции в строке, при нахождении сразу возвращаем false

	for pos, r := range s {
		if strings.Contains(s[(pos+1):], string(r)) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(uniqueSymbols("abcd"))
	fmt.Println(uniqueSymbols("abCdefA"))
	fmt.Println(uniqueSymbols("aabcd"))
}
