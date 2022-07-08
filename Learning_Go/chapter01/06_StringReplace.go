package main

import "fmt"

func stringReplace() {
	str := "HelloWorld"
	r := []rune(str)
	copy(r[4:7], []rune("abc"))
	fmt.Printf("After replace: %s\n", string(r))
}

func main() {
	stringReplace()
}