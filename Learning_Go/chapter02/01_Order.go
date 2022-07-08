package main

import "fmt"

func getOrder(a int, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func main() {
	res1, res2 := getOrder(5, 2)
	fmt.Println(res1, res2)
}
