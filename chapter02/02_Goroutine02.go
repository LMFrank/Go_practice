package main

import (
	"fmt"
	"time"
)

func main() {
	go printNum1()
	go printAlp1()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main: Over")
}

func printNum1() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d\n", i)
	}
}

func printAlp1() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}
