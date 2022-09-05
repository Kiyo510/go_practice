package main

import "fmt"

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	map1 := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(map1, map1["apple"])
	for v, k := range map1 {
		fmt.Println(v, k)
	}
}
