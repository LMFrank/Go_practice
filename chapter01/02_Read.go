package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/Code/Go/src/course/chapter01/test.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer file.Close()
	bs := make([]byte, 4, 4)

	//n, err := file.Read(bs)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println(n)
	//fmt.Println(string(bs))

	for {
		n, err := file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读取完毕")
			return
		}
		fmt.Println(string(bs))
	}
}
