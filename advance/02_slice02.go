package main

import "fmt"

func main() {
	data := []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	newData := noSame(data)
	fmt.Println(newData)
}

func noSame(data []string) []string {
	res := data[:1]
	for _, k := range data {
		i := 0
		for ; i < len(res); i++ {
			if k == res[i] {
				break
			}
		}
		if i == len(res) {
			res = append(res, k)
		}

	}
	return res
}
