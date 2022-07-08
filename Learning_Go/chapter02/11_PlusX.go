package main

import "fmt"

func plusX(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	px := plusX(3)
	fmt.Printf("%v", px(4))

}
