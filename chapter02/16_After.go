package main

import (
	"fmt"
	"time"
)

func main() {
	ch3 := time.After(3 * time.Second)
	fmt.Printf("%T\n", ch3)
	fmt.Println(time.Now())

	time2 := <-ch3
	fmt.Println(time2)

}
