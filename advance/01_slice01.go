package main

import "fmt"

func main() {
	data := []string{"red", "", "black", "yellow", "", "white"}
	//newData := noEmpty1(data)
	newData := noEmpty2(data)
	fmt.Println(newData)
}

func noEmpty1(data []string) []string {
	res := data[:0]
	for _, k := range data {
		if k != "" {
			res = append(res, k)
		}
	}
	return res
}

func noEmpty2(data []string) []string {
	i := 0
	for _, k := range data {
		if k != "" {
			data[i] = k
			i++
		}
	}
	return data[:i]
}
