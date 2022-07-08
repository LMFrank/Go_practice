package main

import (
	"fmt"
	"strconv"
)

type stack2 struct {
	i int
	data [10] int
}

func (s *stack2) push(k int ) {
	if s.i+1 > 9 {
		return
	}
	// println(k)
	// println(s.i)
	s.data[s.i] = k
	s.i++
}

func (s *stack2) String() string {
	var str string
	for j := 0; j < s.i; j++ {
		str = str + "[" + strconv.Itoa(j) + ":" + strconv.Itoa(s.data[j]) + "] "
	}
	return str
}

func main() {
	var s stack2
	s.push(25)
	s.push(14)
	s.push(88)
	fmt.Printf("%v\n",s)
	fmt.Printf("%v\n",s.String())
}

