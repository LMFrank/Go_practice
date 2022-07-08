package main

import "fmt"

func stringReverse() {
	str := "Hello"
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Printf("After reverse: %s", string(r))
}

func main() {
	stringReverse()
}