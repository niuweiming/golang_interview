package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{name: "PHP", age: 1}
	//反射访问私有成员
	v := reflect.ValueOf(p)
	nameField := v.FieldByName("name")
	ageField := v.FieldByName("age")

	// 通过reflect 的安全Api来访问私有字段
	name := reflect.NewAt(nameField.Type(), unsafe.Pointer(nameField.UnsafeAddr())).Elem().Interface().(string)
	age := reflect.NewAt(ageField.Type(), unsafe.Pointer(nameField.UnsafeAddr())).Elem().Interface().(int)
	fmt.Println(name, age)
}
