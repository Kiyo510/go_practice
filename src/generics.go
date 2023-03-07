package main

import "fmt"

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Max[T Number](x, y T) T {
	if x >= y {
		return x
	}
	return y
}

func main() {
	fmt.Println(Max(1, 2))
	fmt.Println(Max(1.2201, 1.220))
}
