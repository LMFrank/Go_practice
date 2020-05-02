package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch1 := make(chan int)
	fmt.Println(len(ch1), cap(ch1))
	//ch1 <- 100

	ch2 := make(chan int, 5)
	fmt.Println(len(ch2), cap(ch2))

	ch2 <- 100
	fmt.Println(len(ch2), cap(ch2))

	fmt.Println("=================================")
	ch3 := make(chan string, 4)
	go sendData2(ch3)
	for {
		v, ok := <-ch3
		if !ok {
			fmt.Println("读取完成", ok)
			break
		}
		fmt.Println("\t读取的数据是: ", v)
	}
	fmt.Println("main over")
}

func sendData2(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "数据" + strconv.Itoa(i)
		fmt.Printf("子goroutine第%d个数据\n", i)
	}
	close(ch)
}
