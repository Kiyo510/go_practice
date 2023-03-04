package main

import "fmt"

func main() {
	sl3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sl3)

	sl3 = append(sl3, 300)
	fmt.Println(sl3)

	sl2 := make([]int, 10)
	fmt.Println(sl2)

}
