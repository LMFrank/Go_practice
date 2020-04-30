package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go fun1()
	go fun2()

	fmt.Println("main进入阻塞状态...")
	wg.Wait()
	fmt.Println("main解除阻塞状态")
}

func fun1() {
	for i := 1; i < 10; i++ {
		fmt.Println("fun1:..", i)
	}
	wg.Done()
}

func fun2() {
	defer wg.Done()
	for j := 1; j < 10; j++ {
		fmt.Println("\tfun2:..", j)
	}
}
