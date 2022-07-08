package main

import "fmt"

func mapStr(f func(string) string, s []string) []string {
	j := make([]string, len(s))
	for k, v := range s {
		j[k] = f(v)
	}
	return j
}

func main() {
	m := []string{"a", "b", "c"}
	f := func(s string) string {
		return s + s
	}
	fmt.Printf("%v", mapStr(f, m))
}
