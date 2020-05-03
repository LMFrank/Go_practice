package main

import (
	"fmt"
	"time"
)

func main() {
	//timer := time.NewTimer(3 * time.Second)
	//fmt.Printf("%T\n", timer)
	//fmt.Println(time.Now())
	//
	//ch2 := timer.C
	//fmt.Println(<-ch2)

	timer2 := time.NewTimer(5 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2结束了...开始...")
	}()

	time.Sleep(3 * time.Second)
	flag := timer2.Stop()
	if flag {
		fmt.Println("Timer 2停止了")
	}
}
