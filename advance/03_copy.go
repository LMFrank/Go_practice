package main

import "fmt"

func main() {
	data := []int{5, 6, 7, 8, 9}
	newData := remove(data, 2)
	fmt.Println(newData)
}

func remove(data []int, index int) []int {
	copy(data[index:], data[index+1:])
	return data[:index+2]
}
