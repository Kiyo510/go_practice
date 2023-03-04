package main

import "fmt"

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)

}

func main() {
	p := &Point{
		A: 100, B: "hoge",
	}

	fmt.Println(p)
}
