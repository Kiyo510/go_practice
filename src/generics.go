package main

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](xs ...T) Set[T] {
	s := make(Set[T])
	for _, v := range xs {
		s[v] = struct{}{}
	}
	return s
}

func main() {
	r := NewSet(2, 3, 4, 5, 6)
	fmt.Println(r)
}
