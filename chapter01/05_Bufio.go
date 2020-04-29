package main

import (
	"bufio"
	"fmt"
	"io"
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

	//1. Read()
	//p := make([]byte, 1024)
	//n1, err := b1.Read(p)
	//if err != nil {
	//	fmt.Println("error: ", err)
	//	return
	//}
	//fmt.Println(n1)
	//fmt.Println(string(p[:n1]))

	//2.Readline()
	//data, flag, err := b1.ReadLine()
	//fmt.Println(flag)
	//fmt.Println(err)
	//fmt.Println(data)
	//fmt.Println(string(data))

	//3.ReadString()
	//s1, err := b1.ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(s1)

	for {
		s1, err := b1.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		fmt.Println(s1)
	}
}
