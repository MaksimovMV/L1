package main

import (
	"fmt"
)

type human struct {
	velocity int
}

type action struct {
	human
}

func (h human) walk(hours int) (distance int) {
	return hours * h.velocity
}

func (h human) speak(m string) {
	fmt.Println(m)
}

func main() {
	a := action{human{velocity: 10}}
	//после встраивания стр. action может напрямую обращаться к методам стр. human
	d := a.walk(2)
	fmt.Printf("Скорость: %v, расстояние: %v\n", a.velocity, d) //а так же к полям
	a.speak("привет")
}
