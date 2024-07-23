package main

import "fmt"

type MyStruct struct {
	field int
}

func (ms MyStruct) ValueReceiverMethod() {
	fmt.Println("Value receiver method")
}
func (ms *MyStruct) PointerReceiverMethod() {
	fmt.Println("Pointer recevire method")
}

func main() {
	var s MyStruct
	s.ValueReceiverMethod()
	s.PointerReceiverMethod()
	// 自动变成这样
	(&s).ValueReceiverMethod()

	ps := &s
	ps.ValueReceiverMethod()
	ps.PointerReceiverMethod()
	a, b := 1, 2
	{
		a, b := 3, 4
		fmt.Println(a, b)
	}
	fmt.Println(a, b)
}
