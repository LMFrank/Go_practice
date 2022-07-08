package main

import "fmt"

func findMax(l []int) (max int) {
	max = l[0]
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	sl := []int{1, 4, 9, 13, 6, 7}
	fmt.Printf("%v", findMax(sl))
}
