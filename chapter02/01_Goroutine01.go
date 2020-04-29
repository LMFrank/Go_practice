package main

import (
	"fmt"
	"time"
)

func main() {
	go printNum()

	for i := 1; i < 100; i++ {
		fmt.Printf("\t主goroutine: A %d\n", i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("main: Over")
}

func printNum() {
	for i := 1; i < 100; i++ {
		fmt.Printf("子goroutine: %d\n", i)
	}
}
