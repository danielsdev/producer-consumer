package main

import (
	"fmt"
	"sync"
	"time"

)

var buffer = make(chan int, 10)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go producer(buffer)
	go consumer(buffer)

	wg.Wait()
}

func producer(buffer chan int) {
	defer wg.Done()
	for i:= 0; ; i++ {
		time.Sleep(1 * time.Second);
		fmt.Printf("Produzindo... %d\n", i);
		buffer <- i
	}
}

func consumer(buffer chan int){
	defer wg.Done()
	for{
		time.Sleep(2 * time.Second);
		fmt.Print("Consumindo... ");
		fmt.Print(<-buffer)
		fmt.Print("\n")
	}
}
