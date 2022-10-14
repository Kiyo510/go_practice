package main

import "fmt"

type MyInt int

func (m MyInt) MyIntFunc(hoge string, huga int) int {
	fmt.Println("111111")
	return 1111
}

func main() {
	var myint MyInt
	myint = 2
	fmt.Printf("%T\n", myint)

	myint.MyIntFunc("hoge", 11)
}
