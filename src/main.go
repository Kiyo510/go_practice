package main

import "fmt"

type Season int

const (
	Spring Season = iota + 1
	Summer
	Autumn
	Winter
)

func (s Season) String() string {
	seasons := [...]string{"string", "summer", "autumn", "winter"}
	if s < Spring || s > Winter {
		return fmt.Sprintf("Season(%d)", int(s))
	}
	return seasons[s-1]
}

func (s Season) IsValid() bool {
	switch s {
	case Spring, Summer, Autumn, Winter:
		return true
	}
	return false
}

func main() {
	mySeasons := []Season{Spring, Summer, Autumn, Winter, 6}
	for _, s := range mySeasons {
		fmt.Printf("season: %s, is valid: %t\n", s, s.IsValid())
	}
}
