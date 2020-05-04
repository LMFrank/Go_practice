package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.23
	fmt.Println("num的数值:", num)

	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	fmt.Println("类型：", newValue.Type())
	fmt.Println("是否可以修改数据:", newValue.CanSet())

	newValue.SetFloat(3.14)
	fmt.Println(num)
}
