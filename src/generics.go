package main

import "fmt"

type Number interface {
	~int | int32 | ~int64 | ~float32 | ~float64
}

type MyInt int

func Max[T Number](x, y T) T {
	if x >= y {
		return x
	}
	return y
}

func main() {
	var x, y MyInt = 1, 2
	fmt.Println(Max(x, y))
	fmt.Println(Max(1.2201, 1.220))
}
