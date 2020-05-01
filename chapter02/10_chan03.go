package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		fmt.Println("子goroutine...")
		//time.Sleep(3 * time.Second)
		data := <-ch1
		fmt.Println("data: ", data)
	}()

	time.Sleep(5 * time.Second)
	ch1 <- 10
	fmt.Println("main over...")
}
