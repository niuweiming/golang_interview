package main

type MyStruct struct {
	field int
}

//func (ms MyStruct) ValueReceiverMethod() {
//	fmt.Println("Value receiver method")
//}
//func (ms *MyStruct) PointerReceiverMethod() {
//	fmt.Println("Pointer recevire method")
//}

type Feather struct {
	name string
}

type Son struct {
	Feather
	age int
}
