package main

import "fmt"

func bubbleSort(l []int) []int {
	n := len(l)
	for i := 0; i < n-1; i++ {
		count := 0
		for j := 0; j < n-1; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
				count += 1
			}
		}
		if count == 0 {
			break
		}
	}
	return l
}

func main() {
	sl := []int{3, 1, 8, 5, 4, 6, 2}
	fmt.Printf("%v", bubbleSort(sl))
}
