package main

import "fmt"

func printCount1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printCount2()  {
	i := 0
Loop:
	if i < 10 {
		fmt.Printf("%d ", i)
		i ++
		goto Loop
	}
}

func printCount3()  {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
}

func main() {
	printCount1()
	printCount2()
	printCount3()
}
