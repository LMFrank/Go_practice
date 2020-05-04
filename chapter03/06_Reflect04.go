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

func (p Person) Test(i, j int, s string) {
	fmt.Println(i, j, s)
}

func main() {
	p1 := Person{"小王", 20, "男"}
	value := reflect.ValueOf(p1)
	fmt.Printf("kind: %s, type: %s\n", value.Kind(), value.Type())

	methodValue1 := value.MethodByName("PrintInfo")
	fmt.Printf("kind: %s, type: %s\n", methodValue1.Kind(), methodValue1.Type())

	methodValue1.Call(nil)

	args := make([]reflect.Value, 0)
	methodValue1.Call(args)

	methodValue2 := value.MethodByName("Say")
	fmt.Printf("kind: %s, type: %s\n", methodValue2.Kind(), methodValue2.Type())

	args2 := []reflect.Value{reflect.ValueOf("反射机制")}
	methodValue2.Call(args2)

	methodValue3 := value.MethodByName("Test")
	fmt.Printf("kind: %s, type: %s\n", methodValue3.Kind(), methodValue3.Type())

	args3 := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200), reflect.ValueOf("aaa")}
	methodValue3.Call(args3)
}
