package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	School string
}

func main() {
	s1 := Student{"小明", 18, "pku"}
	fmt.Printf("%T\n", s1)
	p1 := &s1
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.Name)

	value := reflect.ValueOf(p1)
	if value.Kind() == reflect.Ptr {
		newValue := value.Elem()
		fmt.Println(newValue.CanSet())

		f1 := newValue.FieldByName("Name")
		f1.SetString("小王")
		f3 := newValue.FieldByName("School")
		f3.SetString("nju")
		fmt.Println(s1)
	}
}
