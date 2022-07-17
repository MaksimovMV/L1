package main

import (
	"bytes"
	"fmt"
	"strings"
)

var justString string

func createHugeString(n int) string {
	var s bytes.Buffer
	for i := 0; i < n; i++ {
		s.WriteRune(rune(i))
	}
	return string(s.Bytes())
}

func someFunc() {
	v := createHugeString(1 << 10)
	//justString = strings.Clone(v[:100])

	//переводим в руны чтоб не обрезать двух или трех битные символы
	r := []rune(v)

	// Clone returns a fresh copy of s.
	// It guarantees to make a copy of s into a new allocation,
	// which can be important when retaining only a small substring
	// of a much larger string. Using Clone can help such programs
	// use less memory.
	justString = strings.Clone(string(r[:100]))
}

func main() {
	someFunc()
	fmt.Println(justString)
}
