package main

import "fmt"

func main() {
	ch1 := make(chan int)
	//ch2 := make(chan <- int) //单向，只能写不能读
	//ch3 := make(<- chan int) //单向，只能读不能写

	go fun3(ch1)
	//go fun2(ch2)

	data := <-ch1
	fmt.Println("fun3中写出的数据是:", data)

	go fun4(ch1)
	ch1 <- 200
	fmt.Println("main over")

}

func fun3(ch chan<- int) {
	ch <- 100
	fmt.Println("fun3结束")
}

func fun4(ch <-chan int) {
	data := <-ch
	fmt.Println("fun2从ch中读取的数据是:", data)
}
