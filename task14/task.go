package main

import (
	"fmt"
	"reflect"
)

// с помощью switch можно сравнить тип интерфейса и при положительном результате
// провести с этим типом необходимые операции
func defineType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		// определить тип можно с помощью пакета reflect
		var r = reflect.TypeOf(i)
		fmt.Printf("Other:%v\n", r)
	}

}

func main() {
	var a interface{} = 10
	defineType(a)
	// выяснить значение можно при помощи Printf и %T
	fmt.Printf("%v %T\n\n", a, a)

	var b interface{} = "some text"
	defineType(b)
	fmt.Printf("%v %T\n\n", b, b)

	var c interface{} = true
	defineType(c)
	fmt.Printf("%v %T\n\n", c, c)

	var d interface{} = make(chan int)
	defineType(d)
	fmt.Printf("%v %T\n\n", d, d)

	var e interface{} = 10.5
	defineType(e)
	fmt.Printf("%v %T\n\n", e, e)
}
