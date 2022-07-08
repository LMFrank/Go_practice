package main

import "fmt"

func fizzBuzz() {
	const (
		FIZZ = 3
		BUZZ = 5
	)
	for i := 1; i < 101; i++ {
		p := false
		if i % FIZZ == 0 {
			fmt.Println("FIZZ")
			p = true
		}
		if i % BUZZ == 0 {
			fmt.Println("BUZZ")
			p = true
		}
		if !p {
			fmt.Printf("%v\n", i)
		}
	}
}

func main() {
	fizzBuzz()
}