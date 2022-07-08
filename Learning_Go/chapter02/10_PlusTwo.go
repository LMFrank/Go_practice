package main

import "fmt"

func plusTwo() func(int) int {
	return func(x int) int {
		return x + 2
	}
}

func main() {
	p2 := plusTwo()
	fmt.Printf("%v", p2(2))
}
