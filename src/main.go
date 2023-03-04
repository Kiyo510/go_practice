package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return l[i].Name < l[j].Name
	} else {
		return l[i].Value < l[j].Value
	}
}

func main() {
	m := map[string]int{"S": 1, "A": 3, "I": 3, "K": 9, "O": 8}

	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}
	sort.Sort(lt)
	fmt.Println(lt)
}
