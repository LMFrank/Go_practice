package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "D:\\Code\\Go\\src\\course\\chapter01\\aa.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	defer file.Close()
	b1 := bufio.NewReader(file)
	p := make([]byte, 1024)
	n1, err := b1.Read(p)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(n1)
	fmt.Println(string(p[:n1]))
}
