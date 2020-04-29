package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "D:\\Code\\Go\\src\\course\\chapter01\\aa.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	defer file.Close()
	for i := 0; i < 1000; i++ {
		bs := "abcd"
		n, err := file.WriteString(bs)
		fmt.Println(n)
		HandleErr(err)
	}
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
