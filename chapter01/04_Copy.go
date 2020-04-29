package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile := "D:\\Code\\Go\\src\\course\\chapter01\\test.txt"
	destFile := "chapter01/test2.txt"
	total, err := CopyFile1(srcFile, destFile)
	fmt.Println(total, err)
}

func CopyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	bs := make([]byte, 1024, 1024)
	total := 0
	for {
		n, err := file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕")
			break
		} else if err != nil {
			fmt.Println("报错了！")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}
