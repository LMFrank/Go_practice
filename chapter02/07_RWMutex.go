package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup

func main() {
	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)

	//wg.Add(2)
	//go readData(1)
	//go readData(2)

	wg.Add(3)
	go writeData(1)
	go readData(2)
	go writeData(3)

	wg.Wait()
	fmt.Println("主函数结束")
}

func writeData(i int) {
	defer wg.Done()
	fmt.Println(i, "开始写")
	rwMutex.Lock()
	fmt.Println(i, "正在写")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock()
	fmt.Println(i, "写结束")
}

func readData(i int) {
	defer wg.Done()
	fmt.Println(i, "开始读")

	rwMutex.RLock()
	fmt.Println(i, "正在读取数据")
	time.Sleep(3 * time.Second)
	rwMutex.RUnlock()
	fmt.Println(i, "读结束")
}
