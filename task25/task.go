package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	// After ждет заданное время и отправляет текущее значение времение в возвращаемый канал
	<-time.After(d)
	return
}

func main() {
	fmt.Println("Ждём")
	sleep(time.Second)
	fmt.Println("Готово")
}
