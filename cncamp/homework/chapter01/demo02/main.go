package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
将练习1.2中的生产者消费者模型修改成为多个生产者和多个消费者模式
*/

var cond sync.Cond

const productNum = 3

func producer(out chan<- int, nu int) {
	for {
		cond.L.Lock()
		for len(out) == 3 {
			cond.Wait()
		}
		num := rand.Intn(20)
		out <- num
		fmt.Printf("%dth ***producer produce***, num = %d, len(chan) = %d\n", nu, num, len(out))
		cond.L.Unlock()

		cond.Signal()
		time.Sleep(time.Second)
	}
}

func consumer(in <-chan int, nu int) {
	for {
		cond.L.Lock()
		for len(in) == 0 {
			cond.Wait()
		}
		num := <-in
		fmt.Printf("%dth ###consumer consume###, num = %d, len(chan) = %d\n", nu, num, len(in))
		cond.L.Unlock()
		cond.Signal()

		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	product := make(chan int, productNum)

	cond.L = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go producer(product, i)
	}

	for i := 0; i < 3; i++ {
		go consumer(product, i)
	}

	for {
		select {}
	}
}
