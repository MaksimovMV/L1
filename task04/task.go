package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var n int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&n) //считываем количество воркеров из stdin
	c := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)

	for i := 1; i <= n; i++ { //запускаем n воркеров
		i := i
		go func() {
			for {
				data, ok := <-c //считываем из канала с в data,  проверяем с помощью ok открыт ли канал
				if ok {
					time.Sleep(time.Second * 2)
					fmt.Printf("Воркер: %v, данные: %v\n", i, data)
				}

			}
		}()
	}

	sig := make(chan os.Signal, 1) //создаем канал прослушивания системных сигналов
	//отлавливаем SIGINT и при его получении отправляем сигнал в канал sig
	signal.Notify(sig, os.Interrupt)

	go func() {
		defer wg.Done()

		for i := 0; ; i++ {
			select {
			case <-sig: // при получении сигнала, завершаем все процессы
				fmt.Println("Received an interrupt")
				close(c)
				time.Sleep(time.Second * 3)
				return
			default:
				c <- i // по умолчанию отправляем сигналы в канал с
			}
		}
	}()
	wg.Wait()
}
