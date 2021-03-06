package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket1 = 10
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	wg.Add(4)
	go saleTickets1("售票口1")
	go saleTickets1("售票口2")
	go saleTickets1("售票口3")
	go saleTickets1("售票口4")
	wg.Wait()
	fmt.Println("main结束了...")
}

func saleTickets1(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		mutex.Lock()
		if ticket1 > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出", ticket1)
			ticket1--
		} else {
			mutex.Unlock()
			fmt.Println(name, "售罄")
			break
		}
		mutex.Unlock()
	}
}
