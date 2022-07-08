package main

import "fmt"

func fib(n int) []int {
	arr := make([]int, n)
    arr[0], arr[1] = 1, 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}

func main() {
	fmt.Println(fib(10))
}
