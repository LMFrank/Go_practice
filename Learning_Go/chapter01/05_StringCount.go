package main

import (
	"fmt"
	"unicode/utf8"
)

func stringCount(strA string) {
	b := []byte(strA)
	n1 := len(b)
	n2 := utf8.RuneCount(b)
	fmt.Printf("length: %d, runes: %d", n1, n2)
}

func main() {
	stringCount("Hello中文")
}