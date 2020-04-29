package main

import "fmt"

func main() {
	funcA()
	defer myPrint("defer main: 3...")
	funcB()
	defer myPrint("defer main: 4...")
	fmt.Println("main: over")
}

func myPrint(s string) {
	fmt.Println(s)
}

func funcA() {
	fmt.Println("funcA...")
}

func funcB() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "程序恢复了")
		}
	}()
	fmt.Println("funcB...")
	defer myPrint("defer funB: 1...")
	for i := 1; i <= 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			panic("funcB恐慌")
		}
	}
	defer myPrint("defer funB: 2...")
}
