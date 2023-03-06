package main

import (
	"fmt"
	"strconv"
)

func f[T fmt.Stringer](xs []T) []string {
	var result []string
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

func main() {
	fmt.Println(f([]MyInt{1, 2, 3}))
}
