package main

import (
	"time"
	"fmt"
)

func worker(workerId int, data chan int) {
	for x := range data {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished: %d\n", workerId, x)
	}
}

//Thread 1
func main() {
	channel := make(chan int) //cria um canal para comunicar entre as threads
	qtd := 100
	for i := 1; i <= qtd; i++ {
		fmt.Printf("Worker %d\n", i)

		go worker(i, channel) //Thread 2, 3, 4 etc...

		channel <- i // enche o canal
	}

	fmt.Println(<-channel) // esvazia o canal
}