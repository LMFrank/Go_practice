package main

import (
	"fmt"
	"strings"
)

func printA() {
	for i := 1; i < 101; i++ {
		fmt.Println(strings.Repeat("A", i))
	}
}

func main() {
	printA()
}
