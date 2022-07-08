package main

import "fmt"

func shower1(c chan int, quit chan bool) {
	for {
		select {
		case j := <-c:
			fmt.Printf("%d\n", j)
		case <-quit:
			break
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go shower1(ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- false
}
