package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello, msg")
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n", p.Name, p.Age, p.Sex)
}

func main() {
	p1 := Person{"小明", 18, "男"}
	GetMessage(p1)
}

func GetMessage(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println(getType.Name())
	fmt.Println(getType.Kind())

	getValue := reflect.ValueOf(input)
	fmt.Println(getValue)

	//获取字段
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("字段名称：%s，字段类型：%s，字段数值：%v\n", field.Name, field.Type, value)
	}

	//获取方法
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称：%s，方法类型：%v\n", method.Name, method.Type)
	}
}
