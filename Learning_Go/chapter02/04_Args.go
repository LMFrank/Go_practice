package main

import "fmt"

func foo(arg ...int) {
	for i, n := range arg {
		if i == len(arg)-1 {
			fmt.Printf("%d\n", n)
			break
		}
		fmt.Printf("%d ", n)
	}
}

func main() {
	foo(1, 2, 3, 4, 5, 6)
	foo(2, 3, 4, 5, 6)
}