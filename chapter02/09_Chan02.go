package main

import "fmt"

func main() {
	var ch1 chan bool
	ch1 = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine...", i)
		}
		ch1 <- true
		fmt.Printf("结束")
	}()

	data := <-ch1
	fmt.Println("main...", data)
	fmt.Println("main, over")
}
