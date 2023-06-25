package main

import (
	"time"
	"fmt"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second * 5)
	}
}

//Thread 1
func main() {
	channel := make(chan int) //cria um canal para comunicar entre as threads
	qtd := 1000000
	for i := 1; i <= qtd; i++ {
		go worker(i, channel) //Thread 2, 3, 4 etc...

		channel <- i
	}
}